package blog

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func HandleNewCommand() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: atp new [blogname]")
        os.Exit(1)
    }

    blogname := os.Args[2]
    err := createBlog(blogname)
    if err != nil {
        log.Fatalf("could not create blog: %v", err)
    }

    fmt.Printf("Blog '%s' created successfully.\n", blogname)
}

func createBlog(blogname string) error {
    dirs := []string{
        "content",
        "layouts",
        "static",
        "public",
        "themes",
    }

    err := os.Mkdir(blogname, 0755)
    if err != nil {
        return err
    }

    for _, dir := range dirs {
        path := filepath.Join(blogname, dir)
        err := os.MkdirAll(path, 0755)
        if err != nil {
            return err
        }
    }

    // 创建 config.yml 文件
    configPath := filepath.Join(blogname, "config.yml")
    configContent := `
# Blog configuration
title: "My Blog"
description: "This is my blog"

# header
menu:
  - name: "Home"
    url: "/"
  - name: "Post"
    url: "/post"
  - name: "About"
    url: "/about"
  - name: "Link"
    url: "/link"

# footer
contact:
  - name: "Github"
    url: "https://github.com"
  - name: "Gitlab"
    url: "https://gitlab.com"
`
    err = os.WriteFile(configPath, []byte(configContent), 0644)
    if err != nil {
        return err
    }

    // 创建 index.md 文件
    indexPath := filepath.Join(blogname, "content", "index.md")
    indexContent := `
# Welcome to My Blog
This is the homepage of your new blog. Edit this file to update your content.
`
    err = os.WriteFile(indexPath, []byte(indexContent), 0644)
    if err != nil {
        return err
    }

    // 创建 layouts/base.html 文件
    basePath := filepath.Join(blogname, "layouts", "base.html")
    baseContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Site.Title }}</title>
    <style>
        body {
            background-color: black; /* 背景颜色 */
            color: white;            /* 字体颜色 */
            font-family: Arial, sans-serif; /* 字体 */
            margin: 0;
            padding: 20px;          /* 内边距 */
        }

        h1, h2, h3, h4, h5, h6 {
            color: white;           /* 标题字体颜色 */
        }

        p {
            color: white;           /* 段落字体颜色 */
        }

        a {
            color: gray;            /* 链接颜色 */
        }

        a:hover {
            color: purple;          /* 悬停时链接颜色 */
        }

        blockquote {
            border-left: 5px solid gray; /* 引用左侧边框 */
            padding-left: 10px;          /* 引用内边距 */
            margin: 10px 0;              /* 引用外边距 */
            color: lightgray;             /* 引用字体颜色 */
        }
    </style>
</head>
<body >
    <header>
        <nav>
            <ul>
                {{ range .Site.Menu }}
                <li><a href="{{ .URL }}">{{ .Name }}</a></li>
                {{ end }}
            </ul>
        </nav>
    </header>

    <main>
        {{ .Content }}
    </main>

    <footer>
        <nav>
            <ul>
                {{ range .Site.Contact }}
                <li><a href="{{ .URL }}">{{ .Name }}</a></li>
                {{ end }}
            </ul>
        </nav>
    </footer>
</body>
</html>`

    err = os.WriteFile(basePath, []byte(baseContent), 0644)
    if err != nil {
        return err
    }

    // 创建 CSS 文件
    cssDir := filepath.Join(blogname, "static")
    cssPath := filepath.Join(cssDir, "style.css")
    cssContent := `body {
    background-color: black; /* 背景颜色 */
    color: white;            /* 字体颜色 */
    font-family: Arial, sans-serif; /* 字体 */
    margin: 0;
    padding: 20px;          /* 内边距 */
}

h1, h2, h3, h4, h5, h6 {
    color: white;           /* 标题字体颜色 */
}

p {
    color: white;           /* 段落字体颜色 */
}

a {
    color: gray;            /* 链接颜色 */
}

a:hover {
    color: purple;          /* 悬停时链接颜色 */
}`
    err = os.WriteFile(cssPath, []byte(cssContent), 0644)
    if err != nil {
        return err
    }

    return nil
}
