package main

/*
package关键字代表的是当前go文件属于哪一个包，启动文件通常是main包，启动函数是main函数，在自定义包和函数时命名应当尽量避免与之重复。

import是导入关键字，后面跟着的是被导入的包名。

func是函数声明关键字，用于声明一个函数。

fmt.Println("Hello 世界!")是一个语句，调用了fmt包下的Println函数进行控制台输出。

*/

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
