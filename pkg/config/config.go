package config

import (
    "io/ioutil"
//    "log"

    "gopkg.in/yaml.v3"
)

// Config 结构体定义
type Config struct {
    Title       string `yaml:"title"`
    Description string `yaml:"description"`
    Menu        []struct {
        Name string `yaml:"name"`
        URL  string `yaml:"url"`
    } `yaml:"menu"`
    Contact []struct {
        Name string `yaml:"name"`
        URL  string `yaml:"url"`
    } `yaml:"contact"`
}

// LoadConfig 加载配置文件
func LoadConfig(file string) (*Config, error) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
