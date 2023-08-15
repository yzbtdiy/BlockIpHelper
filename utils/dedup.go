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
	// 创建一个 map 来记录元素是否已经出现过
	seen := make(map[string]bool)
	result := []string{}
	// 遍历数组，将未出现过的元素添加到结果切片中
	for _, addr := range arr {
		if !seen[addr] {
			result = append(result, addr)
			seen[addr] = true
		}
	}
	return result
}
