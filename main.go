package main

import "main.go/tuuz/Redis"

func main() {
	var rs Redis.Stream
	rs.New("knet").XGroupCreateConsumer("group1", "com1")
}
