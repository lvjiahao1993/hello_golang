package main

// // 推荐 （）的形式导入包
// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

// 导入包的另一种形式
// import "fmt"
// import "math/rand"
//
// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }
//

// 你可以导入一个包而不使用它，这是不允许的，编译器会报错。
// import "fmt"
//
// func main() {
// }
//

// 你可以使用.来导入包，这样就可以省略包名。
// import (
// 	. "fmt"
// )
//
// func main() {
// 	Println("Hello, World!")
// }

// 你可以给导入的包起一个别名。
// import (
// 	f "fmt"
// )
//
// func main() {
// 	f.Println("Hello, World!")
// }

// 你也可以导入包，但是不使用包中的任何函数，这个时候需要使用_来重命名这个包。
// import (
// 	_ "fmt"
// )
//
// func main() {
// }

// 你可以导入多个包。
// import (
// 	"fmt"
// 	"math/rand"
// )
//
// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

// 导出名字
// 在 Go 中，导出和访问控制是通过命名来进行实现的，首字母大写的名称是被导出的。
// 在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。
// Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
