package main

/*
package关键字代表的是当前go文件属于哪一个包，启动文件通常是main包，启动函数是main函数，在自定义包和函数时命名应当尽量避免与之重复。

import是导入关键字，后面跟着的是被导入的包名。

func是函数声明关键字，用于声明一个函数。

fmt.Println("Hello 世界!")是一个语句，调用了fmt包下的Println函数进行控制台输出。

*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello_golang/config"
	"hello_golang/internal/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	// Load application configuration from config.yaml
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// 设置 MySQL 数据库连接信息
	username := "root"
	password := "1993729"
	dbname := "mydatabase"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", username, password, dbname))
	if err != nil {
		fmt.Println("数据库连接失败！")
		panic(err)
	}
	fmt.Println("数据库连接成功！")
	defer db.Close()

	// 尝试连接数据库
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Create a new router and set up routes
	router := routers.NewRouter(db)

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Server listening on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}

// 自定义包
// 除了使用标准库中的包，我们还可以自定义包，自定义包的格式如下：
// package 包名
//
// import "包名"
//
// // 函数名首字母小写，表示该函数只能在本包内使用
// func 函数名(参数列表) (返回值列表) {
// 函数体
// }
//
// // 函数名首字母大写，表示该函数能在本包以外的包中使用
// func 函数名(参数列表) (返回值列表) {
// 函数体
// }

// 自定义包的使用
// 自定义包的使用与标准库中的包使用方法相同，只需要在导入包时使用自定义包的路径即可。
// 例如，我们在cmd/main.go中导入自定义包internal/day2，然后调用其中的函数：
// package main
//
// import (
// 	"fmt"
// 	"learn-go/internal/day2"
// )
//
// func main() {
// 	fmt.Println(day2.Add(1, 2))
// }

// 自定义包的导入路径
// 自定义包的导入路径是相对于src目录的相对路径，例如，我们在src目录下创建一个internal/day2包，那么它的导入路径就是internal/day2。
// 但是，如果我们在src目录下创建一个internal/day2/day2包，那么它的导入路径就是internal/day2/day2。
// 也就是说，导入路径是相对于src目录的相对路径，而不是相对于当前文件的相对路径。

// internal
// internal是一个特殊的包名，它的作用是限制包的使用范围，internal包下的内容只能被和internal同级的包所使用。
