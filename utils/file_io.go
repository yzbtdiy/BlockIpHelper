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
	// 循环读取
	var lineArr []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := strings.TrimSpace(buf.Text()) //获取每一行
		lineArr = append(lineArr, line)
	}
	return lineArr, nil
}

// 保存IP和区域信息到文件
func WriteIpAndMaskFile(dataArr []models.IpAndRegion, path string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range dataArr {
		fmt.Fprintln(w, v.Ip+"/32"+" "+v.Region)
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
