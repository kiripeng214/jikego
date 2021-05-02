package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"kiripeng214/jikego/week3/xnet"
)

func main() {
	group := &errgroup.Group{}
	group.Go(xnet.InitServer)
	if err := group.Wait(); err != nil {
		fmt.Println("Get errors:", err)
	}
}
