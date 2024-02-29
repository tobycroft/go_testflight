package main

import (
	"fmt"
	"main.go/tuuz/Redis"
)

func main() {
	var rs Redis.Stream
	group := rs.New("knet")
	fmt.Println(group.XInfoGroups())
	fmt.Println(group.XGroupCreate("group1", "0"))
	fmt.Println(group.XGroupCreateConsumer("com1"))
	fmt.Println(group.XReadGroup())
}
