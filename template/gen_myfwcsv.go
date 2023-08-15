package template

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 生成明御防火墙黑名单模板
func GenerageMyFwCSVFile(inCns, notInCns []models.IpAndRegion, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败")
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)
	writer.Write([]string{"# 源Addr:合法addr格式(IP/Domain/MAC)", "", ""})
	writer.Write([]string{"# 开关:(enable:启用，disable:禁用)", "", ""})
	writer.Write([]string{"# 生命周期:以秒为单位计算", "", ""})
	writer.Write([]string{"IP/Domain/MAC", "Enable", "Age"})
	for _, cnIp := range inCns {
		writer.Write([]string{cnIp.Ip, "enable", "2592000"})
	}
	for _, notCnIp := range notInCns {
		writer.Write([]string{notCnIp.Ip, "enable", "permanent"})
	}
	writer.Flush()
}
