package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lionsoul2014/ip2region/maker/golang/xdb"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// 纯真IP数据替换部分字符串, 删除空行
func qqwryRepalce(srcCon []byte) string {
	utf8Str := strings.Replace(string(srcCon), "  CZ88.NET", "", -1)
	reSumInfo := regexp.MustCompile(`IP数据库共有数据\s\S\s\d+\s条`)
	reBlankRow := regexp.MustCompile(`^\s*$[\r\n]*|[\r\n]+\s+\r*$`)
	reSumInfo.FindAllStringIndex(utf8Str, -1)
	noSumInfo := reSumInfo.ReplaceAllString(utf8Str, "")
	noBlankRow := reBlankRow.ReplaceAllString(noSumInfo, "")
	return noBlankRow
}

// 纯真IP导出的TXT编码格式为GB2312, 使用GB18030解码, UTF8重新编码
func ConventUtf8(qqwryTxt string) {
	fileBytes, err := os.ReadFile(qqwryTxt)
	if err != nil {
		fmt.Println("文件读取失败: ", err)
		os.Exit(1)
	}
	utf8Bytes, conErr := simplifiedchinese.GB18030.NewDecoder().Bytes(fileBytes)
	if err != nil {
		fmt.Println("编码转化失败: ", conErr)
		os.Exit(1)
	}
	utf8Strings := qqwryRepalce(utf8Bytes)
	wErr := os.WriteFile("./data/source.txt", []byte(utf8Strings), os.ModePerm)
	if wErr != nil {
		fmt.Println("文件保存失败: ", wErr)
		os.Exit(1)
	}
	fmt.Println("文件转化完成")
}

// 简单处理纯真地址source.txt转为ip_merge.txt
// 仅通过是否包含省市名称来区分国内地址与国外地址, 然后将区域信息整体保存
func GenIp2RegionMergeFile(srcFile, dstFile string, cnProvince []string) {
	IpLibs, _ := ReadTxt(srcFile)
	var mergeArry []string
	for _, row := range IpLibs {
		reg := regexp.MustCompile(`\s+`)
		infoArr := reg.Split(row, -1)
		regionInfo := strings.Join(infoArr[2:], "")
		var regionArry [5]string
		for _, provi := range cnProvince {
			if strings.Contains(regionInfo, provi) {
				regionArry[0] = "中国"
				break
			} else {
				regionArry[0] = "外国"
			}
		}
		regionArry[1] = "0"
		regionArry[2] = regionInfo
		regionArry[3] = "0"
		regionArry[4] = "0"
		mergeRow := strings.Join(infoArr[:2], "|") + "|" + strings.Join(regionArry[0:5], "|")
		mergeArry = append(mergeArry, mergeRow)
	}
	err := WriteFile(mergeArry, dstFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("已生成ip_merge.txt")
}

// 生成ip2region xdb文件
// 代码来自https://github.com/lionsoul2014/ip2region/blob/master/maker/golang/main.go
func GenIp2RegionXdbFile(srcFile, dstFile string) {
	tStart := time.Now()
	var err error
	var indexPolicy = xdb.VectorIndexPolicy
	maker, err := xdb.NewMaker(indexPolicy, srcFile, dstFile)
	if err != nil {
		fmt.Printf("failed to create %s\n", err)
		return
	}

	err = maker.Init()
	if err != nil {
		fmt.Printf("failed Init: %s\n", err)
		return
	}

	err = maker.Start()
	if err != nil {
		fmt.Printf("failed Start: %s\n", err)
		return
	}

	err = maker.End()
	if err != nil {
		fmt.Printf("failed End: %s\n", err)
	}
	log.Printf("Done, elapsed: %s\n", time.Since(tStart))
	fmt.Println("已生成ip2region.xdb")
}
