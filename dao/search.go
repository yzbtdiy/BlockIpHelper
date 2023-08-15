package dao

import (
	"log"
	"strings"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 查找是否白名单中有此地址
func IsExistInWhiteList(Ip string) bool {
	result := db.First(&models.WhiteListTable{}, "white_ip = ?", Ip)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

// 查找是否黑名单中有此地址
// func IsExistInBlackList(Ip string) bool {
// 	result := db.First(&models.BlackListTable{}, "black_addr = ?", Ip)
// 	if result.RowsAffected != 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func GetWhiteList() (subnetsList []string, addrList []string) {
	var whiteList []models.WhiteListTable
	result := db.Select("white_addr").Find(&whiteList)
	if result.RowsAffected != 0 {
		for _, addr := range whiteList {
			if IsSubnet(addr.WhiteAddr) {
				subnetsList = append(subnetsList, addr.WhiteAddr)
			} else {
				addrList = append(addrList, addr.WhiteAddr)
			}
		}
		return subnetsList, addrList
	} else {
		log.Println("读取白名单失败或者白名单为空")
		return nil, nil
	}
}

func IsSubnet(addr string) bool {
	if strings.Contains(addr, "/") {
		return true
	} else {
		return false
	}
}
