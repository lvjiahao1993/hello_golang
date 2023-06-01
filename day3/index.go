// 数据类型

package main

import (
	"fmt"
)

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
// 就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。
// 除了 var 语句之外，还可以通过简短赋值语句 := 来声明并初始化局部变量。
// 例如：在函数中 var f string = "short" 可以简写为 f := "short" 。
// 注意：函数外的每个语句都必须以关键字开始（var、func 等等），因此 := 结构不能在函数外使用。
var a string = "G"

// 变量声明可以包含初始值，每个变量对应一个。
// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
var b = "G"

// 省略 var 的声明方式 := 函数内的简洁写法，不需要使用 var 关键字声明变量，但是它有一些限制：
// 1. 定义变量
// 2. 不能提供数据类型
// 3. 只能用在函数内部
// 4. 不能用于声明全局变量
// c := "G"

// 常量的声明和变量类似，只不过是使用 const 关键字。
// 常量可以是字符、字符串、布尔值或数值。
// 常量不能用 := 语法声明。
const Pi = 3.14

// 数值常量是高精度的 _值_。
// 一个未指定类型的常量由上下文来决定其类型。
// 也尝试一下输出 needInt(Big) 吧。
const (
	Big   = 1 << 100
	Small = Big >> 99
)

// Go 的基本类型有
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // uint8 的别名
// rune // int32 的别名 代表一个Unicode码
// float32 float64
// complex64 complex128
// int、uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位。
// 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。

// int、uint 和 uintptr 在使用时具有特定的位宽，例如，int 是 32 位宽，在所有平台都是相同的大小，除非程序需要和特定硬件平台保持一致。
// 但是，int 并不是固定的大小或者说它并不是一个具体的类型，它的大小取决于编译器的实现。
// 一个 int 可以是 32 位或者 64 位的，但是在同一个平台上所有的 int 都是相同的大小。
// 由于历史原因，有些函数接受一个 int 类型的数值，而有些则接受 int32 或 int64 类型的数值，这取决于该函数的使用场景。
// 例如，math.MaxInt32 可以返回一个 int32 类型的最大值，而 math.MaxInt64 则返回一个 int64 类型的最大值。
// 但是，如果你需要一个整数值，你应该总是使用 int 类型，因为它可以被编译器很好地处理，而且在 32 位系统上它是 32 位，在 64 位系统上是 64 位。

// 没有明确初始值的变量声明会被赋予它们的 零值。

// 零值是：
// 数值类型为 0，
// 布尔类型为 false，
// 字符串为 ""（空字符串）。

// 类型转换
// 表达式 T(v) 将值 v 转换为类型 T。

// 一些关于数值的转换：
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// 或者，更加简单的形式：
// i := 42
// f := float64(i)
// u := uint(f)

// 与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换。

// 派生类型
// 指针类型（Pointer）
// 数组类型
// 结构化类型(struct)
// Channel 类型
// 函数类型
// 切片类型
// 接口类型（interface）
// Map 类型

// 指针
// 指针保存了值的内存地址。
// 类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。
// var p *int
// & 符号会生成一个指向其作用对象的指针。
// i := 42
// p = &i
// * 符号表示指针指向的底层的值。
// fmt.Println(*p) // 通过指针 p 读取 i
// *p = 21         // 通过指针 p 设置 i
// 这也就是通常所说的“间接引用”或“非直接引用”。

// 结构体
// 一个结构体（struct）就是一个字段的集合。
// （而 type 的含义跟其字面意思相符。）
// type Vertex struct {
// 	X int
// 	Y int
// }

// 数组
// 类型 [n]T 表示拥有 n 个 T 类型的值的数组。
// var a [10]int
// 数组的长度是其类型的一部分，因此数组不能改变大小。
// 这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
// 切片。

// 切片
// 切片（slice）代表变长的序列，序列中每个元素都有相同的类型。
// 一个切片的类型一般写作 []T，其中 T 代表切片中元素的类型；切片像数组一样，但切片的长度是可变的，可以随时增加或删除元素。
// 切片的零值是 nil。
// var s []int
// 切片的长度和容量
// 切片拥有 长度 和 容量。
// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
// 切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
// 你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的边界(extend its boundaries)，
// 看看会发生什么。

// 切片的默认行为
// 在进行切片时，你可以利用它的默认行为来忽略上下界。
// 切片下界的默认值为 0，上界则是该切片的长度。

// 切片的切片
// 切片可包含任何类型，甚至包括其它的切片。
// var board = [][]string{
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// }

// 切片的零值
// 切片的零值是 nil。
// nil 切片的长度和容量为 0 且没有底层数组。
// var s []int
// if s == nil {
// 	fmt.Println("nil!")
// }

// 使用 make 创建切片
// 切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
// a := make([]int, 5)  // len(a)=5
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4

// Channel 类型
// Channel 可以被声明为带有一个方向的类型，只发送或者只接收。
// var c chan int
// c = make(chan int)
// c <- 1    // 发送
// <-c       // 接收

// Channel 缓冲区
// Channel 可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
// ch := make(chan int, 100)

// Channel 的关闭
// 发送者可以 close 一个 channel 来表示再没有值会被发送了。
// 接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：
// v, ok := <-ch
// 如果 ok 是 false，那么说明 channel 已经没有任何数据并且已经被关闭。

// 函数类型
// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。

// 函数闭包
// Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
// 该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
// 例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }
// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }

// 方法
// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
// 在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

// 接口
// 接口类型 是由一组方法签名定义的集合。
// 接口类型的变量可以保存任何实现了这些方法的值。
// 注意: 示例代码的 22 行存在一个错误。由于 Abs 只定义在 *Vertex（指针类型）上，所以 Vertex（值类型）并未实现 Abser。
// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}
// 	a = f  // a MyFloat 实现了 Abser
// 	a = &v // a *Vertex 实现了 Abser
// 	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
// 	// 所以没有实现 Abser。
// 	a = v
// 	fmt.Println(a.Abs())
// }

// 错误
// Go 程序使用 error 值来表示错误状态。
// 与 fmt.Stringer 类似，error 类型是一个内建接口：
// type error interface {
// 	Error() string
// }
// 通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。
// i, err := strconv.Atoi("42")
// if err != nil {
// 	fmt.Printf("couldn't convert number: %v\n", err)
// 	return
// }
// fmt.Println("Converted integer:", i)

// 映射 (map)
// 映射将键映射到值。
// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
// make 函数会返回给定类型的映射，并将其初始化备用。
// var m map[string]Vertex
// m = make(map[string]Vertex)
// m["Bell Labs"] = Vertex{
// 	40.68433, -74.39967,
// }
// fmt.Println(m["Bell Labs"])

// 映射的文法
// 映射的文法与结构体相似，不过必须有键名。
// type Vertex struct {
// 	Lat, Long float64
// }
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }
// fmt.Println(m)

// 映射的文法续
// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
// type Vertex struct {
// 	Lat, Long float64
// }
// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }
// fmt.Println(m)

// 修改映射
// 在映射 m 中插入或修改一个元素：
// m[key] = elem
// 获得元素：
// elem = m[key]
// 删除元素：
// delete(m, key)
// 通过双赋值检测某个键存在：
// elem, ok = m[key]
// 若 key 在 m 中，ok 为 true ；否则，ok 为 false。

// 函数值
// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。
// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }
// func main() {
// 	hypot := func(x, y float64) float64 {
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	fmt.Println(hypot(5, 12))
// 	fmt.Println(compute(hypot))
// 	fmt.Println(compute(math.Pow))
// }

// 函数的闭包
func main() {
	fmt.Println("Hello, World!")
}
