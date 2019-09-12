package main

import (
    "fmt"

    "github.com/ShiinaOrez/ginny/router"
    "github.com/ShiinaOrez/ginny/model"
)

func main() {
    model.DB.Init()        // 初始化数据库
    defer model.DB.Close() // 记得关闭数据库

    router.InitRouter()    // 初始化路由
    router.Router.Run()    // 运行
    fmt.Println("Running... Successful!")
}
