package template

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 生成k01黑名单导入文件
func GenerateK01CSVFile(inCns, notInCns []models.IpAndRegion, path string) {
	log.Println("开始生成k01黑名单导入文件, 保存到 " + path)
	log.Println("#############################################################")
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败")
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)
	writer.Write([]string{"IP/IP段(0.0.0.0不可输入)", "类型(固定格式：源地址/目的地址)", "过期时长(固定格式：x月x天x时x分x秒)", "备注(不超过20个字符)", "创建时间(导入时生成新的创建时间)"})
	for _, cnIp := range inCns {
		writer.Write([]string{cnIp.Ip + "/32", "源地址", "30天0时0分0秒", K01Len20(cnIp.Region)})
	}
	for _, notCnIp := range notInCns {
		writer.Write([]string{notCnIp.Ip + "/32", "源地址", "", K01Len20(notCnIp.Region)})
	}
	writer.Flush()
}

// k01模板限制20个字符, 对区域信息截取20个字符
func K01Len20(region string) (regSlice string) {
	regRune := []rune(region)
	if len(regRune) > 20 {
		regSlice = string(regRune[0:20])
		return regSlice
	} else {
		return string(regRune)
	}
}
