package main

import (
	"flag"
	"fmt"

	"github.com/mochi-yu/kns23-catch-up/app"
)

var env = flag.String("env", "prod", "Execute environment")

func main() {
	// 実行環境をフラグから取得する
	flag.Parse()
	fmt.Println("env: ", *env)

	server := app.NewServer()
	server.Engine.Run(":8080")
}
