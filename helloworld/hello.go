package main

import "fmt"

//引用包内的 go 文件
import "github.com/shishuwu/hello-go/common/math"

func main() {
	fmt.Println("Hello World!")

	sum := math.Add(1,2)

	fmt.Println(sum)
}
