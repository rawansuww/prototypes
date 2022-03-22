package main

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"xorm.io/core"
	//_ "github.com/go-sql-driver/mysql"
)

// The tablename of the corresponding database must be student
// When executing mysql, the corresponding fields are xxx, yyy, zzz; you can also omit the default MySQL fields are id, username, address
type Person struct {
	//	Id       int    `xorm:"INTEGER GENERATED ALWAYS AS IDENTITY 'id'"`
	FirstName string `xorm:"VARCHAR(64) 'firstname'"`
	LastName  string `xorm:"VARCHAR(256) 'lastname'"`
	Email     string `xorm:"VARCHAR(64) 'email'"`
}

func main() {
	engine, err := xorm.NewEngine("postgres", "user=postgres dbname=xormdb2 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := engine.Sync2(new(Person))
	if err2 != nil {
		fmt.Println("whyyy")
		return
	}
	engine.ShowSQL(true) // Displays the execution of SQL for easy debugging and analysis
	engine.SetTableMapper(core.SnakeMapper{})
	st1 := new(Person)
	st1.FirstName = "taoge"
	st1.LastName = "taogee"
	st1.Email = "taoge@china.com"
	affected, err := engine.Insert(st1)
	fmt.Println(affected)
	st2 := new(Person)
	result, err := engine.Where("id=?", 1).Get(st2)
	fmt.Println(result)
}
