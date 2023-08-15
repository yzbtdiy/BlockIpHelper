package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yzbtdiy/BlockIpHelper/dao"
	"github.com/yzbtdiy/BlockIpHelper/flags"
	"github.com/yzbtdiy/BlockIpHelper/template"
	"github.com/yzbtdiy/BlockIpHelper/utils"
)

func init() {
	utils.CheckPath("./data")
	dao.InitDb()
}

func main() {
	flags.UseFlags()
	var tStart = time.Now()
	whiteSubnets, whiteAddrs := dao.GetWhiteList()
	targets, err := utils.ReadTxt("./target.txt")
	if err != nil {
		log.Fatal(err)
	}
	targetAddrs := utils.RemoveDuplicates(targets)
	var inWhiteList []string
	var needBlocks []string
	for _, addr := range targetAddrs {
		InWhiteSubnets := utils.InWhiteSubnets(addr, whiteSubnets)
		InWhiteAddr := utils.InWhiteAddrs(addr, whiteAddrs)
		if InWhiteAddr || InWhiteSubnets {
			fmt.Println(addr + "白名单地址, 跳过")
			inWhiteList = append(inWhiteList, addr)
		} else {
			fmt.Println(addr + "非白名单的地址, 进行分类")
			needBlocks = append(needBlocks, addr)
		}
	}
	if len(inWhiteList) != 0 {
		log.Println("#############################################################")
		log.Println("目标地址列表存在白名单地址, 保存到 白名单IP.txt 文件")
		utils.WriteFile(inWhiteList, "./白名单IP.txt")
	} else {
		log.Println("#############################################################")
		log.Println("目标地址列表未发现白名单地址")
	}
	if len(needBlocks) != 0 {
		inCns, notInCns := utils.FilterCn(needBlocks, "./data/ip2region.xdb")
		utils.WriteIpAndMaskFile(inCns, "./国内攻击IP.txt")
		utils.WriteIpAndMaskFile(notInCns, "./国外攻击IP.txt")
		log.Println("#############################################################")
		log.Println("已完成地址分类: 白名单IP.txt, 国内攻击IP.txt, 国外攻击IP.txt")
		log.Println("#############################################################")
		log.Println("尝试生成k01黑名单导入文件 k01Block.csv")
		log.Println("#############################################################")
		template.GenerageK01CSVFile(inCns, notInCns, "./k01Block.csv")
		log.Println("尝试生成明防御防火墙黑名单导入文件 myFwBlock.csv")
		log.Println("#############################################################")
		template.GenerageMyFwCSVFile(inCns, notInCns, "./myFwBlock.csv")
	} else {
		log.Println("#############################################################")
		log.Println("未发现需要加入黑名单的地址, 请确认 target.txt 文件")
		log.Println("#############################################################")
	}
	log.Printf("本次处理%v个地址, 耗时%v", len(targets), time.Since(tStart))
}
