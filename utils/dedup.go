package utils

import (
	"fmt"
)

// 读取键盘输入
func GetChoice() bool {
	var input string
	fmt.Println("是否执行XXX? 输入'y'生成")
	_, err := fmt.Scanf("%s", &input)

	if err != nil {
		fmt.Println(err)
	}
	if input == "y" {
		fmt.Println("执行成功")
		return true
	} else {
		return false
	}
}

// 数组去重
func RemoveDuplicates(arr []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, addr := range arr {
		if !seen[addr] {
			result = append(result, addr)
			seen[addr] = true
		}
	}
	return result
}
