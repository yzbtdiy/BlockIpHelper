package flags

import (
	"flag"
	"log"
	"os"

	"github.com/yzbtdiy/BlockIpHelper/models"
	"github.com/yzbtdiy/BlockIpHelper/utils"
)

var importFlag = flag.String("imp", "",
	`white   导入白名单(./data/whitelist.txt)`)

// black   导入黑名单(./data/blacklist.txt)`)

var generateFlag = flag.String("gen", "",
	`merge   纯真IP源文件(./data/source.txt)生成的xdb源文件(./data/ip_merge.txt)
xdb     xdb源文件(./data/ip_merge.txt)生成ip2region的xdb文件(./data/ip2region.xdb)`)

func UseFlags(config *models.Config) {
	flag.Parse()
	if *importFlag == "white" {
		log.Println("开始导入白名单 ...")
		utils.InsertWhite("./data/whitelist.txt")
		os.Exit(0)
	} else if *generateFlag == "merge" {
		log.Println("开始将纯真IP导出的source.txt转化为ip_merge.txt ...")
		utils.GenIp2RegionMergeFile("./data/source.txt", "./data/ip_merge.txt", config.Ip2Region.CnKeys)
		os.Exit(0)
	} else if *generateFlag == "xdb" {
		log.Println("使用ip_merge.txt生成ip2region.xdb文件 ...")
		utils.GenIp2RegionXdbFile("./data/ip_merge.txt", "./data/ip2region.xdb")
		os.Exit(0)
	}
}
