package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/malikhan-dev/zenql/streams/v2"
)

const mysqlconstr_init = "root:1245Sa@tcp(localhost:30306)/?parseTime=true&charset=utf8mb4"
const constrPgsql_init = "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"

type Users struct {
	ID        int       `zql:"id"`
	Name      string    `zql:"name"`
	Age       int       `zql:"age"`
	CreatedAt time.Time `zql:"created_at"`
}

var dbNameOfTestRun string
var dbNameOfTestRun_postgres string

func setup_db() {

	var wg sync.WaitGroup
	wg.Add(2)
	go MySqlDbSetup(&wg)
	go PgsqlDbSetup(&wg)

	wg.Wait()

}

func PgsqlDbSetup(wg *sync.WaitGroup) {

	dbNameOfTestRun_postgres = fmt.Sprintf("test_zenql_%d", time.Now().UnixNano())

	create_sql := "CREATE DATABASE " + dbNameOfTestRun_postgres + ";"

	if conn, err := Connect("postgres", constrPgsql_init); err != nil {
		panic(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		Exec(conn, create_sql)

		if dbcon, error := Connect("postgres", GetRelevantPostgresConstr()); error == nil {

			table_statement := "CREATE TABLE IF NOT EXISTS users (id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,name VARCHAR(100) NOT NULL,age INT NOT NULL,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"

			res := Exec(dbcon, table_statement)

			fmt.Println(table_statement)

			fmt.Println(res.Err)
			fmt.Println(res.RowsAffected)
			seed_staement := "INSERT INTO users (name, age)\nSELECT v.name, v.age\nFROM (\n    VALUES \n        ('mohammad', 25),\n        ('sara', 30),\n        ('ali', 28)\n) AS v(name, age)\nWHERE NOT EXISTS (\n    SELECT 1 FROM users\n);"

			seedres := Exec(dbcon, seed_staement)

			fmt.Println(seedres.RowsAffected)
			fmt.Println(seedres.Err)

		} else {
			panic(err)
		}
		wg.Done()
	}

}

func MySqlDbSetup(wg *sync.WaitGroup) {

	dbNameOfTestRun = fmt.Sprintf("test_zenql_%d", time.Now().UnixNano())

	create_sql := "CREATE DATABASE IF NOT EXISTS " + dbNameOfTestRun + ";"

	if conn, err := Connect("mysql", mysqlconstr_init); err != nil {

		panic(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		Exec(conn, create_sql)

		Exec(conn, "USE "+dbNameOfTestRun)

		table_statement := "CREATE TABLE IF NOT EXISTS users (\n    id INT PRIMARY KEY AUTO_INCREMENT,\n    name VARCHAR(100) NOT NULL,\n    age INT NOT NULL,\n    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\n);"

		Exec(conn, table_statement)

		seed_staement := "INSERT INTO users (name, age)\nSELECT * FROM (\n    SELECT 'mohammad', 25\n    UNION ALL\n    SELECT 'sara', 30\n    UNION ALL\n    SELECT 'ali', 28\n) AS tmp\nWHERE NOT EXISTS (SELECT 1 FROM users LIMIT 1);"

		Exec(conn, seed_staement)

		wg.Done()
	}
}

func init() {

	setup_db()
}

func TestSetup(t *testing.T) {
	fmt.Println("Test_setup")
}

func GetRelevantConstr() string {
	constr := "root:1245Sa@tcp(localhost:30306)/" + dbNameOfTestRun + "?parseTime=true&charset=utf8mb4"
	return constr
}

func GetRelevantPostgresConstr() string {

	constr := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=" + GetRelevantDbName_postgres() + " sslmode=disable"
	return constr
}
func GetRelevantDbName() string {
	return dbNameOfTestRun
}
func GetRelevantDbName_postgres() string {
	return dbNameOfTestRun_postgres
}
func TestZenqDB_NewMySqlConnection(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {
		fmt.Println(err)
		t.Fatal(err)
	} else {
		err = conn.Ping()
		defer conn.Close()
	}

}

func TestZenqDB_ExecuteQuery(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		limit := 4

		query := fmt.Sprintf(
			`SELECT id, name, age, created_at
	 				FROM %s.users
	 				limit ?`,
			GetRelevantDbName(),
		)

		result, err := Query[Users](conn, query, limit)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(result)
	}

}

func TestZenqDB_ExecuteSingleQuery(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		id := 0

		query := fmt.Sprintf(
			`SELECT id, name, age, created_at
	 				FROM %s.users
	 				WHERE Id > ?`,
			GetRelevantDbName(),
		)

		result, err := SingleQuery[Users](conn, query, id)

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			t.Error("Execute Single Failed")
		}

		fmt.Println(result)

		id = 1

		query = fmt.Sprintf(
			`SELECT id, name, age, created_at
	 				FROM %s.users
	 				WHERE Id = ?`,
			GetRelevantDbName(),
		)

		result, err = SingleQuery[Users](conn, query, id)

		if err != nil {
			t.Error("Execute Single Failed")
		}
		fmt.Println(result)
	}

}

func TestSqlInjection1(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "mohammad';Drop Table users;"

		query := fmt.Sprintf(
			`SELECT id, name, age, created_at
	 				FROM %s.users
	 				WHERE name = ?`,
			GetRelevantDbName(),
		)

		result, err := Query[Users](conn, query, name)

		if err != nil {
			fmt.Println(err)
		}
		if len(result) > 0 {
			t.Error(result)
		}

		name = "mohammad"

		result, err = Query[Users](conn, query, name)

		if err != nil {
			t.Error("sql injected")
		}
	}

}

func TestSqlInjection2(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "'mohammad--' or 1=1"

		query := fmt.Sprintf(
			`SELECT id, name, age, created_at
	 				FROM %s.users
	 				WHERE name = ?`,
			GetRelevantDbName(),
		)

		result, err := Query[Users](conn, query, name)

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

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {
		t.Fatal(err)
	} else {

		age := 65

		id := 1

		cmd := fmt.Sprintf(`update %s.users set age = ? where id = ?`, GetRelevantDbName())

		result := Exec(conn, cmd, age, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_delete(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {
		t.Fatal(err)
	} else {

		id := 4

		cmd := fmt.Sprintf(`delete from %s.users where id = ?`, GetRelevantDbName())

		result := Exec(conn, cmd, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_insert(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {
		t.Fatal(err)
	} else {
		name := "javid"
		age := 65
		cmd := fmt.Sprintf(`INSERT INTO %s.users (name,age) values(?,?)`, GetRelevantDbName())
		result := Exec(conn, cmd, name, age)
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

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0

		query := fmt.Sprintf(`select * from %s.users where id>?`, GetRelevantDbName())

		stream := streams.FromSqlRows[Users](ctx, conn, query, id)

		fmt.Println(stream.Err)

		if stream.Initiated {
			for v := range stream.FilterStream(func(model Users) bool {
				return model.Age > 0
			}).Throttle(time.Millisecond * 200).Channel {

				/// business logic

				fmt.Println(v)
				business_logic_satisfied := true

				if business_logic_satisfied {

					cmd := fmt.Sprintf(`update %s.users set Name = ? where Id =?`, GetRelevantDbName())

					result := Exec(conn, cmd, v.Name+" - old ", v.ID)
					if result.Err != nil {
						t.Error(result.Err)
					} else {
						fmt.Println(v, " - updated. ", result.RowsAffected)
					}
				}

			}
		} else {
			t.Error("stream not initiated")
		}

	}
}

func Test_Transaction_Fail(t *testing.T) {

	if conn, err := Connect("mysql", GetRelevantConstr()); err != nil {

		t.Fatal(err)

	} else {
		conn.Begin()

		q := fmt.Sprintf(`DELETE FROM %s.users WHERE Id =?`, GetRelevantDbName())

		res := Exec(conn, q, 1)

		if res.Err == nil {

			cmd2 := Exec(conn, q, 985000)

			if cmd2.Err != nil {
				t.Error("transaction fail")

			} else if cmd2.RowsAffected > 0 {
				t.Error("transaction fail")
			} else {
				conn.Rollback()
			}
		} else {
			t.Error("transaction fail")
		}
	}
}

func TestZenqDB_PgSqlConnection(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {
		t.Fatal(err)
	} else {
		err = conn.Ping()
		defer conn.Close()
	}

}

func TestZenqDB_ExecuteQuery_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		limit := 4

		query := "SELECT id, name, age, created_at FROM users limit $1"

		result, err := Query[Users](conn, query, limit)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(result)
	}

}

func TestZenqDB_ExecuteSingleQuery_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		id := 0

		query := "SELECT id, name, age, created_at FROM users WHERE Id > $1"

		result, err := SingleQuery[Users](conn, query, id)

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			t.Error("Execute Single Failed")
		}

		fmt.Println(result)

		id = 1

		query = "SELECT id, name, age, created_at FROM users WHERE Id = $1"

		result, err = SingleQuery[Users](conn, query, id)

		if err != nil {
			t.Error("Execute Single Failed")
		}
		fmt.Println(result)
	}

}

func TestSqlInjection1_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "mohammad';Drop Table users;"

		query := "SELECT id, name, age, created_at FROM users WHERE name = $1"

		result, err := Query[Users](conn, query, name)

		if err != nil {
			fmt.Println(err)
		}
		if len(result) > 0 {
			t.Error(result)
		}

		name = "mohammad"

		result, err = Query[Users](conn, query, name)

		if err != nil {
			t.Error("sql injected")
		}
	}

}

func TestSqlInjection2_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {

		t.Fatal(err)

	} else {

		err = conn.Ping()

		defer conn.Close()

		name := "'mohammad--' or 1=1"

		query := "SELECT id, name, age, created_at FROM users WHERE name = $1"

		result, err := Query[Users](conn, query, name)

		if err != nil {
			fmt.Println(err)
		}

		if len(result) > 0 {
			t.Error("sql injected!")
		}

		fmt.Println(result)
	}

}

func TestExecMySqlCommand_update_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {
		t.Fatal(err)
	} else {

		age := 65

		id := 1

		cmd := "update users set age = $1 where id = $2"

		result := Exec(conn, cmd, age, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_delete_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {
		t.Fatal(err)
	} else {

		id := 4

		cmd := "delete from users where id = $1"

		result := Exec(conn, cmd, id)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func TestExecMySqlCommand_insert_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {
		t.Fatal(err)
	} else {
		name := "javid"
		age := 65
		cmd := "INSERT INTO users (name,age) values($1,$2)"
		result := Exec(conn, cmd, name, age)
		if result.Err != nil {
			t.Error(result.Err)
		} else {
			fmt.Println("command executed, rows affected: ", result.RowsAffected)
		}
	}
}

func Test_StreamFromPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0

		query := "select * from users where id>$1"

		stream :=
			streams.FromSqlRows[Users](ctx, conn, query, id)

		if stream.Initiated {
			for v := range stream.FilterStream(func(model Users) bool {
				return model.Age > 0
			}).Throttle(time.Millisecond * 200).Channel {

				/// business logic

				business_logic_satisfied := true

				if business_logic_satisfied {

					cmd := "update users set Name = $1 where Id =$2"

					result := Exec(conn, cmd, v.Name+" - old ", v.ID)
					if result.Err != nil {
						t.Error(result.Err)
					} else {
						fmt.Println(v, " - updated. ", result.RowsAffected)
					}
				}

			}
		} else {
			t.Error("stream not initiated")
		}

	}
}

func Test_Transaction_Fail_postgres(t *testing.T) {

	if conn, err := Connect("postgres", GetRelevantPostgresConstr()); err != nil {

		t.Fatal(err)

	} else {

		conn.Begin()

		q := "DELETE FROM users WHERE Id =$1"

		res := Exec(conn, q, 1)

		if res.Err == nil {

			cmd2 := Exec(conn, q, 985000)

			if cmd2.Err != nil {
				t.Error("transaction fail")

			} else if cmd2.RowsAffected > 0 {
				t.Error("transaction fail")
			} else {
				conn.Rollback()
			}
		} else {
			t.Error("transaction fail")
		}
	}
}
