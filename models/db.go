package models

type WhiteListTable struct {
	Id        int    `gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
	WhiteAddr string `gorm:"column:white_addr;type:TEXT NOT NULL UNIQUE;"`
}

func (WhiteListTable) TableName() string {
	return "white_list"
}

// type BlackListTable struct {
// 	Id        int    `gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
// 	BlackAddr string `gorm:"column:black_addr;type:TEXT NOT NULL UNIQUE;"`
// }

// func (BlackListTable) TableName() string {
// 	return "black_list"
// }
