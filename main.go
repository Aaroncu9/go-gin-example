package main

import (
	"fmt"
)

func main() {
	// 默认路由
	rt := SetupRouter()

	// 默认0.0.0.1：8080 run
	if err := rt.Run(); err != nil {
		fmt.Printf("startup service faild, err: %v\n", err)
	}

}
