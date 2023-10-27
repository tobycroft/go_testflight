package main

import (
	"fmt"
	"main.go/tuuz"
)

func main() {
	tuuz.Db()
	db := tuuz.Db().Table("aaa")
	db.Where("1=1")
	//fmt.Println(db.BuildSql("select"))
	//db.Query()
	fmt.Println(db.Get())
}
