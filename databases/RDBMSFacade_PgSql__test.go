package databases

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

const constrPgsql = "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"

func TestZenqDB_PgSqlConnection(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {
		t.Fatal(err)
	} else {
		err = conn.Ping()
		defer conn.Close()
	}

}

func TestZenqDB_PgSql_ExecuteQuery(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		limit := 4

		result, err := Query[Users](conn, "select * from public.users LIMIT $1", limit)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}

}

func TestZenqDB_PgSql_ExecuteSingleQuery(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		id := 0

		result, err := SingleQuery[Users](conn, "select * from users where Id>$1", id)

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			t.Error("Execute Single Failed")
		}

		fmt.Println(result)

		id = 1

		result, err = SingleQuery[Users](conn, "select * from users where Id=$1", id)

		if err != nil {
			t.Error("Execute Single Failed")
		}
		fmt.Println(result)
	}

}

func TestSqlInjection1_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "mohammad';Drop Table users;"

		result, err := Query[Users](conn, "select * from users where Name = $1", name)

		if err != nil {
			fmt.Println(err)
		}

		name = "mohammad"

		result, err = Query[Users](conn, "select * from users where Name = $1", name)

		if err != nil {
			t.Error("sql injected")
		}

		fmt.Println(result)
	}

}

func TestSqlInjection2_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "'mohammad--' or 1=1"

		result, err := Query[Users](conn, "select * from users where Name = $1", name)

		if err != nil {
			fmt.Println(err)
		}

		if len(result) > 0 {
			t.Error("sql injected!")
		}

		fmt.Println(result)
	}

}

func TestExecMySqlCommand_update_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {
		t.Fatal(err)
	} else {
		age := 65
		id := 1
		result := Exec(conn, "update users set Age = $1 where Id = $2", age, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_delete_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {
		t.Fatal(err)
	} else {
		id := 4
		result := Exec(conn, "delete from users where id = $1", id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_insert_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {
		t.Fatal(err)
	} else {
		name := "javid"
		age := 65
		result := Exec(conn, "INSERT INTO users (name,age) values($1,$2)", name, age)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}
func Test_StreamFromMySql_PgSql(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	if conn, err := Connect("postgres", constrPgsql); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0
		stream :=
			FromSqlRows[Users](ctx, conn,
				"select * from users where id>$1", func(rows *sql.Rows) (Users, error) {
					var id, age int
					var name string
					var err error

					err = rows.Scan(&id, &name, &age)
					model := Users{
						ID:   id,
						Age:  age,
						Name: name,
					}
					return model, err
				}, id)

		if stream.Initiated {
			for v := range stream.FilterStream(func(model Users) bool {
				return model.Age > 25
			}).Throttle(time.Millisecond * 5000).Channel {

				/// business logic

				business_logic_satisfied := true

				if business_logic_satisfied {

					result := Exec(conn, "update users set Name = $1 where Id = $2", v.Name+" - old ", v.ID)
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

func Test_Transaction_PgSql(t *testing.T) {

	if conn, err := Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	} else {
		conn.Begin()
		res := Exec(conn, "DELETE FROM users WHERE Id =$1", 2)

		if res.Err == nil {
			name := "asghar"
			age := 65

			cmd2 := Exec(conn, "INSERT INTO users (name,age) values($1,$2)", name, age)

			if cmd2.Err != nil {

				conn.Rollback()

				res3 := Exec(conn, "DELETE FROM users WHERE Id =$1", 100)

				if res3.Err == nil {
					fmt.Println("success")
				} else {
					fmt.Println(res3.Err)
				}

			} else {
				conn.Rollback()
				res4 := Exec(conn, "DELETE FROM users WHERE Id =$1", 100)
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
