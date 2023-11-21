package main

import (
	"fmt"
	"main.go/tuuz"
)

func main() {
	tuuz.Db()
	db := tuuz.Db().Table("aaa")
	//db.Fields()
	//db.Where("id")
	//db.Where("\"NUMBER\"")
	//db.Where("NUMBER")
	//db.Data(map[string]any{
	//	"id":     5,
	//	"val":    "ggg",
	//	"val2":   "sss",
	//	"NUMBER": 7,
	//})
	//fmt.Println(db.BuildSql("replace"))
	db.Where("val", "ggg")
	db.Data(map[string]any{
		"val": nil,
	})
	//db.OrderBy("val desc, val2 desc")
	//db.Limit(1)
	//fmt.Println(db.Get())
	fmt.Println(db.PaginatorWG())
	//db.Query()
	//fmt.Println(db.Get())
}
