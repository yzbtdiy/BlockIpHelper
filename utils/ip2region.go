package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/lionsoul2014/ip2region/maker/golang/xdb"
)

// 简单处理纯真地址source.txt转为ip_merge.txt
// 仅仅通过包含省市名称来判断国内国外地址, 然后将区域信息整体存入
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
}

// 生成ip2region xdb文件
// 代码片段来自https://github.com/lionsoul2014/ip2region/blob/master/maker/golang/main.go
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
}
