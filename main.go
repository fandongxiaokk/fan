package main

import (
	"fan/ss/ss"
	"fmt"
	"time"
	// "github.com/eddycjy/go-retract-demo"
)

func main() {
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println(ss.Yy())
}
