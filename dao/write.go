package dao

import "github.com/yzbtdiy/BlockIpHelper/models"

// 添加地址到白名单
func AddWhiteListToDb(rowData *models.WhiteListTable) bool {
	result := db.Create(&rowData)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

// 添加地址到黑名单
// func AddBlackListToDb(rowData *models.BlackListTable) bool {
// 	result := db.Create(&rowData)
// 	if result.RowsAffected != 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }