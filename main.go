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
	fmt.Println(group.XGroupCreate("group3", "0"))
	fmt.Println(group.XGroupCreateConsumer("com1"))

	fmt.Println(group.XReadGroup())
	fmt.Println(group.XDelete())
	fmt.Println(group.XPending())
	time.Sleep(1 * time.Second)

}
