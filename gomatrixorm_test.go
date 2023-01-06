package GoMatrixORM

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// create table tab1( id int auto_increment primary key, name varchar(128) not null );

type Tab1 struct {
	Id   int
	Name string
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("mysql", "root:root@tcp(localhost:33066)/gee?charset=utf8")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

func TestInsert(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	session := engine.NewSession()
	res, err := session.Raw("insert into tab1 values(0,'godhearing')").Exec()
	if err != nil {
		t.Fatal("filed to insert", err)
	}
	fmt.Println(res)
}

func TestQuery(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	session := engine.NewSession()
	res, err := session.Raw("select * from tab1").QueryRows()
	if err != nil {
		t.Fatal("query data error ", err)
	}
	result := make([]*Tab1, 0)
	for res.Next() {
		scan := &Tab1{}
		err = res.Scan(&scan.Id, &scan.Name)
		if err != nil {
			t.Fatal("scan error ", err)
		}
		result = append(result, scan)
	}
	for i := range result {
		fmt.Println(result[i])
	}
}
