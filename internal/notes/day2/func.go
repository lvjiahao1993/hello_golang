package main

import "fmt"

// 函数定义
// func 函数名(参数列表) (返回值列表) {
// 函数体
// }
// 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
// 参数列表：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
// 返回值列表：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
// 函数体：实现指定功能的代码块。

// 函数的返回值可以命名，也可以不命名。
// 命名的返回值就相当于在函数中声明一个变量，可以被随意修改，并在函数中使用return直接返回。
// 不命名返回值就相当于在函数中声明一个变量，但是不对该变量进行赋值，函数中使用return返回该变量，由于没有对该变量赋值，因此返回该类型的默认值。
// 命名返回值的函数示例：
func add1(x int, y int) (sum int) {
	sum = x + y
	return sum
}

// 不命名返回值的函数示例：
func add2(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add3(x, y int) (sum int) {
	sum = x + y
	return sum
}

// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
// 直接返回的函数示例：
func add4(x, y int) (sum int) {
	sum = x + y
	return
}

// func main() {
// 	x := 1
// 	y := 2
// 	sum := add4(x, y)
// 	fmt.Println(sum)
// }

// 返回多个值的函数示例：
func add5(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// func main() {
// 	x := 1
// 	y := 2
// 	sum, sub := add5(x, y)
// 	fmt.Println(sum, sub)
// }

// 函数的参数列表和返回值列表的数据类型可以是值类型和引用类型。
// 值类型包括：基础数据类型int、float、bool、string等正常使用的类型，以及数组和结构体。
// 引用类型包括：指针、slice切片、map、管道chan、interface等复杂数据类型。
// 值类型的参数传递是值传递，函数接收的是参数值的副本，对副本的修改不会影响原值。
// 引用类型的参数传递是引用传递，函数接收的是指针，对指针指向的值的修改会影响原值。
// 函数的参数传递示例：
// func main() {
// 	x := 1
// 	y := 2
// 	sum := add(x, y)
// 	fmt.Println(sum)
// }

// 函数的参数传递示例：
// func main() {
// 	x := 1
// 	y := 2
// 	sum := add(&x, &y)
// 	fmt.Println(sum)
// }

//   func add(x *int, y *int) (sum int) {
// 	*x = 3
// 	*y = 4
// 	sum = *x + *y
// 	return sum
// }
