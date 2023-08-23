package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 读取文件, 返回数组
func ReadTxt(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	var lineArr []string
	for {
		if !buf.Scan() {
			break
		}
		line := strings.TrimSpace(buf.Text())
		lineArr = append(lineArr, line)
	}
	return lineArr, nil
}

// 保存IP和区域信息到文件
func WriteIpToFile(dataArr []models.IpAndRegion, path string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range dataArr {
		fmt.Fprintln(w, v.Ip+" "+v.Region)
	}
	return w.Flush()
}

// 将数组写入文件, 一行一个元素
func WriteFile(dataArr []string, path string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range dataArr {
		fmt.Fprintln(w, v)
	}
	// w.WriteString(line)
	return w.Flush()
}

// 检查路径目录是否存在
func CheckPath(pathStr string) bool {
	s, err := os.Stat(pathStr)
	if err != nil {
		log.Println("data 目录不存在尝试创建")
		err := os.MkdirAll(pathStr, os.ModePerm)
		if err != nil {
			log.Println("创建 data 目录失败")
			return false
		}
		return true
	} else {
		if s.IsDir() {
			return true
		} else {
			log.Println("当前 data 不是目录, 请检查!!!")
			return false
		}
	}
}

// 检查配置文件, 不存在自动生成配配置并退出
func CheckConfig(pathStr string) (*models.Config, error) {
	s, err := os.Stat(pathStr)
	if err != nil {
		log.Println("配置文件不存在, 生成默认配置, 请检查配置后重新运行")
		GenerateConfig()
		os.Exit(0)
		return nil, nil
	} else {
		if !s.IsDir() {
			conf := GetConfig(pathStr)
			return &conf, nil
		} else {
			return nil, nil
		}
	}
}
