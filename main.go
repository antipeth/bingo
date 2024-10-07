package main

import (
    "fmt"
    "os"

    "atp/pkg/blog"
    "atp/pkg/config"
    "atp/pkg/server"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("expected 'new' or 'start' command")
        os.Exit(1)
    }


    switch os.Args[1] {
    case "new":
        blog.HandleNewCommand()
    case "start":
    // 加载配置
        cfg, err := config.LoadConfig("config.yml")
        if err != nil {
            fmt.Println("Error loading config:", err)
            os.Exit(1)
        }

        blog.RenderContent(cfg) // 传递配置
        go blog.WatchContent(cfg) // 在这里传递配置
        server.StartServer() // 传递配置
    default:
        fmt.Println("unknown command:", os.Args[1])
        os.Exit(1)
    }
}
