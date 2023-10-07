package flags

import (
	"flag"
	"log"
	"os"

	"github.com/yzbtdiy/BlockIpHelper/models"
	"github.com/yzbtdiy/BlockIpHelper/utils"
)

// 文件导入flag选项
var importFlag = flag.String("imp", "",
	`white   导入白名单(./data/whitelist.txt)`)

// black   导入黑名单(./data/blacklist.txt)`)

// 文件生成flag选项
var generateFlag = flag.String("gen", "",
	`merge   纯真IP解压文件(./data/qqwry.txt)生成的xdb源文件(./data/ip_merge.txt)
xdb     xdb源文件(./data/ip_merge.txt)生成ip2region的xdb文件(./data/ip2region.xdb)`)

func UseFlags(config *models.Config) {
	flag.Parse()
	if *importFlag == "white" {
		log.Println("开始导入白名单 ...")
		utils.InsertWhite(config.WhiteFile)
		os.Exit(0)
	} else if *generateFlag == "merge" {
		log.Println("转化UTF-8编码, 删除空行, 生成source.txt文件 ...")
		utils.ConventUtf8(config.Ip2Region.CzTxt)
		log.Println("将source.txt转化为ip_merge.txt ...")
		utils.GenIp2RegionMergeFile("./data/source.txt", config.Ip2Region.MergeFile, config.Ip2Region.CnKeys)
		os.Exit(0)
	} else if *generateFlag == "xdb" {
		log.Println("使用ip_merge.txt生成ip2region.xdb文件 ...")
		utils.GenIp2RegionXdbFile(config.Ip2Region.MergeFile, config.Ip2Region.XdbFile)
		os.Exit(0)
	}
}
