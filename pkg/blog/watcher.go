package blog

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/fsnotify/fsnotify"
    "atp/pkg/config" // 导入配置包
)

// WatchContent 现在接收一个配置参数
func WatchContent(cfg *config.Config) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // 监视 content 目录及其子目录
    err = addRecursive(watcher, "content")
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }

            // 检查文件写入
            if event.Op&fsnotify.Write == fsnotify.Write {
                fmt.Println("Detected changes in", event.Name)
                time.Sleep(1 * time.Second) // 简单的防抖处理
                RenderContent(cfg) // 传递配置参数
            }

            // 处理文件重命名
            if event.Op&fsnotify.Rename == fsnotify.Rename {
                fmt.Println("Detected rename event:", event.Name)
                // 删除旧的 HTML 文件
                oldHTMLPath := strings.Replace(event.Name, "content/", "public/", 1)
                oldHTMLPath = strings.Replace(oldHTMLPath, ".md", ".html", 1)

                if err := os.Remove(oldHTMLPath); err != nil {
                    fmt.Printf("Failed to delete old HTML file: %s, error: %v\n", oldHTMLPath, err)
                } else {
                    fmt.Printf("Old HTML file deleted: %s\n", oldHTMLPath)
                }

                // 检查新文件是否存在，可能是重命名后的新文件
                if _, err := os.Stat(event.Name); !os.IsNotExist(err) {
                    time.Sleep(1 * time.Second) // 简单的防抖处理
                    RenderContent(cfg) // 重新渲染
                }
            }

            // 处理文件删除
            if event.Op&fsnotify.Remove == fsnotify.Remove {
                fmt.Println("Detected remove event:", event.Name)
                // 删除对应的 HTML 文件或文件夹
                htmlPath := strings.Replace(event.Name, "content/", "public/", 1)
                htmlPath = strings.Replace(htmlPath, ".md", ".html", 1)

                // 检查是否是目录删除
                if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
                    // 删除整个目录
                    if err := os.RemoveAll(htmlPath); err != nil {
                        fmt.Printf("Failed to delete directory: %s, error: %v\n", htmlPath, err)
                    } else {
                        fmt.Printf("Directory deleted: %s\n", htmlPath)
                    }
                } else {
                    // 删除单个文件
                    if err := os.Remove(htmlPath); err != nil {
                        fmt.Printf("Failed to delete HTML file: %s, error: %v\n", htmlPath, err)
                    } else {
                        fmt.Printf("HTML file deleted: %s\n", htmlPath)
                    }
                }
            }

            // 检查是否是目录创建事件
            if event.Op&fsnotify.Create == fsnotify.Create {
                // 检查新创建的路径是否是目录
                if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
                    fmt.Println("Detected new directory:", event.Name)
                    // 将新目录添加到 watcher
                    addRecursive(watcher, event.Name)
                }
            }
        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            fmt.Println("Error:", err)
        }
    }
}

// addRecursive 添加一个目录及其所有子目录到 watcher
func addRecursive(w *fsnotify.Watcher, path string) error {
    err := w.Add(path)
    if err != nil {
        return err
    }

    // 遍历目录，添加所有子目录
    return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && p != path { // 跳过根目录
            return w.Add(p) // 添加子目录
        }
        return nil
    })
}
