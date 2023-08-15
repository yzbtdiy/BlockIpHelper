package utils

import (
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 使用ip2region判断是否为国内地址
func InChinaOffline(targetIp, xdbFile string) (bool, string) {
	searcher, err := xdb.NewWithFileOnly(xdbFile)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return false, ""
	}
	defer searcher.Close()
	region, err := searcher.SearchByStr(targetIp)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", targetIp, err)
		return false, ""
	}
	zone := strings.Split(region, "|")
	if zone[0] == "中国" {
		// fmt.Println(targetIp + " 国内地址")
		return true, zone[2]
	} else {
		// fmt.Println(targetIp + " 国外地址")
		return false, zone[2]
	}
}

// 非白名单地址分类为不同数组, 包含ip地址和区域信息
func FilterCn(needBlocks []string, xdbFile string) (inCns, notInCns []models.IpAndRegion) {
	for _, v := range needBlocks {
		inChina, region := InChinaOffline(v, xdbFile)
		if inChina {
			inCns = append(inCns, models.IpAndRegion{Ip: v, Region: region})
		} else {
			notInCns = append(notInCns, models.IpAndRegion{Ip: v, Region: region})
		}
	}
	return inCns, notInCns
}
