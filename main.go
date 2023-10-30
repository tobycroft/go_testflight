package main

import (
	"fmt"
	"main.go/tuuz"
)

func main() {
	tuuz.Db()
	db := tuuz.Db().Table("aaa")
	db.Fields()
	//db.Where("id")
	//db.Where("\"NUMBER\"")
	//db.Where("NUMBER")
	db.Data(map[string]any{
		"id":     5,
		"val":    "ggg",
		"val2":   "sss",
		"NUMBER": 7,
	})
	//fmt.Println(db.BuildSql("replace"))
	fmt.Println(db.Replace())
	//db.Query()
	//fmt.Println(db.Get())
}
