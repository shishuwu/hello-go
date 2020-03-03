package main

import "fmt"

// 引用包内的 go 文件
import "github.com/shishuwu/hello-go/common/math"

// 引用外部 包
import (
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World!")

	sum := math.Add(1, 2)

	fmt.Println(sum)
	
	log.Debug("this is debug log")
	log.Info("this is info log")
	log.Warn("this is warn log")
	log.Fatal("this is fatal log, it will exist the process")
}
