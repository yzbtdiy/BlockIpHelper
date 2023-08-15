package utils

import (
	"log"

	"github.com/yzbtdiy/BlockIpHelper/dao"
	"github.com/yzbtdiy/BlockIpHelper/models"
)

// sqlite插入白名单地址
func InsertWhite(path string) {
	whiteList, _ := ReadTxt(path)
	for _, whiteAddr := range whiteList {
		row := models.WhiteListTable{WhiteAddr: whiteAddr}
		result := dao.AddWhiteListToDb(&row)
		if result {
			log.Println(whiteAddr + " 已添加到数据库白名单地址表中")
		}
	}
}

// sqlite插入黑名单
// func InsertBlack(path string) {
// 	blackList, _ := ReadTxt(path)
// 	for _, blackAddr := range blackList {
// 		row := models.BlackListTable{BlackAddr: blackAddr}
// 		result := dao.AddBlackListToDb(&row)
// 		if result {
// 			log.Println(blackAddr + " 已添加到本地数据库黑名单列表中")
// 		}
// 	}
// }

// 黑名单添加地址
// func AppendBlack(blackList []string) {
// 	for _, blackAddr := range blackList {
// 		row := models.BlackListTable{BlackAddr: blackAddr}
// 		result := dao.AddBlackListToDb(&row)
// 		if result {
// 			log.Println(blackAddr + " 已添加到本地数据库黑名单地址表中")
// 		}
// 	}
// }
