package blog

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "os"
    "path/filepath"
    "reflect"
    "strings"

    "github.com/russross/blackfriday/v2"
    "gopkg.in/yaml.v3" // 导入yaml解析库
    "atp/pkg/config"   // 导入配置
)

// FrontMatter 定义 frontmatter 的结构
type FrontMatter struct {
    Title string `yaml:"title"`
    Date  string `yaml:"date"`
    Tags  []string `yaml:"tags"`
}

func RenderContent(cfg *config.Config) error {
    // 检查 cfg.Menu 的类型和数据
    fmt.Printf("cfg.Menu Type: %s\n", reflect.TypeOf(cfg.Menu))
    fmt.Printf("cfg.Menu Data: %+v\n", cfg.Menu)

    // 读取 base.html 文件内容
    basePath := filepath.Join("layouts", "base.html")
    baseTemplate, err := ioutil.ReadFile(basePath)
    if err != nil {
        return err
    }

    // 编译模板
    tmpl, err := template.New("main").Parse(string(baseTemplate))
    if err != nil {
        return err
    }

    // 遍历 content 目录及其子目录
    err = filepath.Walk("content", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 检查是否是目录
        if info.IsDir() {
            return nil
        }

        // 只处理 .md 文件
        if filepath.Ext(path) == ".md" {
            content, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }

            // 将 frontmatter 和 Markdown 内容分开
            contentStr := string(content)
            var frontMatter FrontMatter
            parts := strings.SplitN(contentStr, "---", 3)
            if len(parts) >= 3 {
                err = yaml.Unmarshal([]byte(parts[1]), &frontMatter) // 解析 frontmatter
                if err != nil {
                    return fmt.Errorf("error parsing frontmatter in %s: %v", path, err)
                }
                contentStr = parts[2] // 保留实际 Markdown 内容
            }

            // 打印调试信息以确认 frontmatter 和 Markdown 内容
            fmt.Printf("Rendering file: %s\n", path)
            fmt.Printf("FrontMatter: %+v\n", frontMatter)
            fmt.Printf("Markdown content: %s\n", contentStr)

            // 将 Markdown 转换为 HTML
            htmlContent := blackfriday.Run([]byte(contentStr))

            // 打印转换后的 HTML 内容
            fmt.Printf("HTML content: %s\n", string(htmlContent))

            // 准备模板数据，包含 frontmatter 数据
            data := struct {
                Site struct {
                    Title   string `yaml:"title"`
                    Menu    []struct {
                        Name string `yaml:"name"`
                        URL  string `yaml:"url"`
                    } `yaml:"menu"`
                    Contact []struct {
                        Name string `yaml:"name"`
                        URL  string `yaml:"url"`
                    } `yaml:"contact"`
                }
                Content     template.HTML
                FrontMatter FrontMatter
            }{
                Content:     template.HTML(htmlContent),
                FrontMatter: frontMatter,
            }

            data.Site.Title = cfg.Title
            data.Site.Menu = cfg.Menu
            data.Site.Contact = cfg.Contact

            // 生成输出文件路径，将 .md 替换为 .html
            outputFile := strings.Replace(path, "content/", "public/", 1)
            outputFile = strings.Replace(outputFile, ".md", ".html", 1)

            // 确保输出目录存在
            outputDir := filepath.Dir(outputFile)
            if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
                return err
            }

            // 创建输出文件
            outFile, err := os.Create(outputFile)
            if err != nil {
                return err
            }
            defer outFile.Close()

            // 打印输出文件路径
            fmt.Printf("Output file path: %s\n", outputFile)

            // 执行模板，将数据写入输出文件
            err = tmpl.Execute(outFile, data)
            if err != nil {
                return err
            }
        }

        return nil
    })

    return err
}

