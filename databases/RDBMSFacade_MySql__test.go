package databases

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

type UserModel struct {
	UserId   int    `zdb:"Id"`
	UserName string `zdb:"Name"`
	Age      int    `zdb:"Age"`
}

const mysqlconstr = "root:1245Sa@tcp(localhost:30306)/Test?parseTime=true&charset=utf8mb4"

func TestZenqDB_NewMySqlConnection(t *testing.T) {

	if conn, err := Connect("mysql", mysqlconstr); err != nil {
		fmt.Println(err)
		t.Fatal(err)
	} else {
		err = conn.Ping()
		defer conn.Close()
	}

}

func TestZenqDB_ExecuteQuery(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		limit := 4

		result, err := Query[UserModel](conn, "select * from Test.users  limit ?", limit)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}

}

func TestZenqDB_ExecuteSingleQuery(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		id := 0
		result, err := SingleQuery[UserModel](conn, "select * from Test.users where Id>?", id)

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			t.Error("Execute Single Failed")
		}

		fmt.Println(result)

		id = 1

		result, err = SingleQuery[UserModel](conn, "select * from Test.users where Id=?", id)

		if err != nil {
			t.Error("Execute Single Failed")
		}
		fmt.Println(result)
	}

}

func TestSqlInjection1(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "mohammad';Drop Table users;"

		result, err := Query[UserModel](conn, "select * from Test.users where Name = ?", name)

		if err != nil {
			fmt.Println(err)
		}

		name = "mohammad"

		result, err = Query[UserModel](conn, "select * from Test.users where Name = ?", name)

		if err != nil {
			t.Error("sql injected")
		}

		fmt.Println(result)
	}

}

func TestSqlInjection2(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		type user struct {
			UserId   int    `zdb:"Id"`
			UserName string `zdb:"Name"`
			Age      int    `zdb:"Age"`
		}

		name := "'mohammad--' or 1=1"

		result, err := Query[user](conn, "select * from Test.users where Name = ?", name)

		if err != nil {
			fmt.Println(err)
		}

		if len(result) > 0 {
			t.Error("sql injected!")
		}

		fmt.Println(result)
	}

}

func TestExecMySqlCommand_update(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {
		t.Fatal(err)
	} else {
		age := 65
		id := 1
		result := Exec(conn, "update Test.users set Age = ? where Id =?", age, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_delete(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"
	if conn, err := Connect("mysql", constr); err != nil {
		t.Fatal(err)
	} else {
		id := 4
		result := Exec(conn, "delete from Test.users where id = ?", id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_insert(t *testing.T) {

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {
		t.Fatal(err)
	} else {
		id := 7
		name := "javid"
		age := 65
		result := Exec(conn, "INSERT INTO Test.users values(?,?,?)", id, name, age)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}
func Test_StreamFromMySql(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0
		stream :=
			FromSqlRows[UserModel](ctx, conn,
				"select * from Test.users where id>?", func(rows *sql.Rows) (UserModel, error) {
					var id, age int
					var name string
					var err error

					err = rows.Scan(&id, &name, &age)
					model := UserModel{
						UserId:   id,
						Age:      age,
						UserName: name,
					}
					return model, err
				}, id)

		if stream.Initiated {
			for v := range stream.FilterStream(func(model UserModel) bool {
				return model.Age > 25
			}).Throttle(time.Millisecond * 5000).Channel {

				/// business logic

				business_logic_satisfied := true

				if business_logic_satisfied {

					result := Exec(conn, "update Test.users set Name = ? where Id =?", v.UserName+" - old ", v.UserId)
					if result.Err != nil {
						t.Error(result.Err)
					} else {
						fmt.Println(v, " - updated. ", result.RowsAffected)
					}
				}

			}
		} else {
			fmt.Println("stream not initiated")
		}

	}
}

func Test_Transaction(t *testing.T) {
	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {
		conn.Begin()
		res := Exec(conn, "DELETE FROM Test.users WHERE Id =?", 1)
		fmt.Println(res.Err)
		if res.Err == nil {
			id := 11
			name := "asghar"
			age := 65

			cmd2 := Exec(conn, "INSERT INTO Test.users values(?,?,?)", id, name, age)

			if cmd2.Err != nil {

				fmt.Println(cmd2.Err)

				fmt.Println("rolling back")

				conn.Rollback()

				res3 := Exec(conn, "DELETE FROM Test.users WHERE Id =?", 100)

				if res3.Err == nil {
					fmt.Println("success")
				} else {
					fmt.Println(res3.Err)
				}

			} else {
				conn.Commit()
				res4 := Exec(conn, "DELETE FROM Test.users WHERE Id =?", 100)
				if res4.Err == nil {
					fmt.Println(res4)
					fmt.Println("success")
				} else {
					fmt.Println(res4.Err)
				}
			}
		}
	}
}
