package main

import (
	"fmt"
	"main.go/tuuz/Redis"
	"time"
)

func main() {
	var rs Redis.Stream
	group := rs.New("knet")
	fmt.Println(group.XInfoGroups())
	fmt.Println(group.XGroupCreate("group1", "0"))
	fmt.Println(group.XGroupCreateConsumer("com1"))

	var rs2 Redis.Stream
	group2 := rs2.New("knet")
	fmt.Println(group2.XInfoGroups())
	fmt.Println(group2.XGroupCreate("group1", "0"))
	fmt.Println(group2.XGroupCreateConsumer("com2"))
	//go fmt.Println(group.XReadGroup())
	//go fmt.Println(group2.XReadGroup())
	fmt.Println(group.XPending())
	fmt.Println(group2.XPending())
	fmt.Println(group.XInfoGroups())
	time.Sleep(1 * time.Second)

}
