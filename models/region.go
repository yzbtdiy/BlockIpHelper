package models

type IpAndRegion struct {
	Ip     string `gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
	Region string `gorm:"column:black_addrs;type:TEXT NOT NULL UNIQUE;"`
}
