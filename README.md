# 🍬bingo🍬

A practice project designed to create a minimalist static blog generator. Built with Go, it generates a static site from markdown files with minimalist configuration.

## Features
- **Markdown-based content:** Write your blog posts using Markdown syntax.
- **Configurable layout:** Customize your site's title, menu, footer, and theme colors using a `config.yml` file.
- **Automatic file watching:** Automatically detect changes to your content and update the site.

## Usage

### Configuring the Blog

Edit the `config.yml` file to set your site's title, menu items, contact information, and theme colors. Example:

```yaml
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
```
##  Have a try

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.

To generate and serve your site locally, use the following command:

```bash
git clone https://github.com/antipeth/bingo.git
cd bingo
go mod tidy
go build -o bingo .
./bingo new yourBlogName
cd yourBlogName
../bingo start
```
Visit localhost:3333 in your browser.




