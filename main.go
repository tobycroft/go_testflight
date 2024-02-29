package main

import (
	"fmt"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Redis"
)

func main() {
	var rs Redis.Stream
	a, b := rs.New("knet").XInfoGroups()
	fmt.Println(Jsong.Encode(a))
	fmt.Println(b)
}
