package main

import (
	"log"
	"time"

	"github.com/yzbtdiy/BlockIpHelper/dao"
	"github.com/yzbtdiy/BlockIpHelper/flags"
	"github.com/yzbtdiy/BlockIpHelper/template"
	"github.com/yzbtdiy/BlockIpHelper/utils"
)

func init() {
	// 检查./data目录是否存在,不存在自动创建
	utils.CheckPath("./data")
	// 初始化数据库
	dao.InitDb()
}

func main() {
	config, err := utils.CheckConfig("./data/config.yaml")
	if err != nil {
		log.Fatal("读取配置文件异常, 请检查")
	}
	// 检查命令是否携带参数
	flags.UseFlags(config)
	// 记录开始时间
	var tStart = time.Now()
	// 读取白名单, 拆分携带掩码的和不携带掩码的
	whiteSubnets, whiteAddrs := dao.GetWhiteList()
	// 读取需要判断的地址列表
	targets, err := utils.ReadTxt(config.TargetFile)
	if err != nil {
		log.Fatal(err)
	}
	// 去除重复地址
	targetAddrs := utils.RemoveDuplicates(targets)
	// 创建数组存放匹配到白名单的地址
	var inWhiteList []string
	// 创建数组存放未匹配到白名单的地址
	var needBlocks []string
	// 遍历需要判断的数组, 筛选出白名单地址和非白名单地址
	for _, addr := range targetAddrs {
		// 带掩码的判断地址是否在子网内
		InWhiteSubnets := utils.InWhiteSubnets(addr, whiteSubnets)
		// 不带掩码的地址直接比较字符串
		InWhiteAddr := utils.InWhiteAddrs(addr, whiteAddrs)
		if InWhiteAddr || InWhiteSubnets {
			// fmt.Println(addr + "白名单地址, 跳过")
			inWhiteList = append(inWhiteList, addr)
		} else {
			// fmt.Println(addr + "非白名单的地址, 进行分类")
			needBlocks = append(needBlocks, addr)
		}
	}
	// 判断存放白名单匹配结果得数组是否为空
	if len(inWhiteList) != 0 {
		log.Println("#############################################################")
		log.Println("目标地址列表存在白名单地址, 保存到 " + config.ExportFile.InWhitelist)
		utils.WriteFile(inWhiteList, config.ExportFile.InWhitelist)
	} else {
		log.Println("#############################################################")
		log.Println("目标地址列表未发现白名单地址")
	}
	// 判断存放非白名单地址的数组是否为空
	if len(needBlocks) != 0 {
		// 调用ip2region获取IP归属地
		inCns, notInCns := utils.FilterCn(needBlocks, config.Ip2Region.XdbFile)
		// 国内地址和国外地址拆分存储
		utils.WriteIpAndMaskFile(inCns, config.ExportFile.InChina)
		utils.WriteIpAndMaskFile(notInCns, config.ExportFile.OutCina)
		log.Println("#############################################################")
		log.Println("白名单IP , 国内攻击IP, 国外攻击IP 分类完成, 请查看对应txt文件")
		log.Println("#############################################################")
		// 生成设备黑名单导入文件
		for _, temp := range config.Template {
			if temp.Enable == true {
				template.GenerateCSV(inCns, notInCns, temp)
			}
		}
	} else {
		log.Println("#############################################################")
		log.Println("未发现需要加入黑名单的地址, 请确认 " + config.TargetFile)
		log.Println("#############################################################")
	}
	// 执行结束, 输出结果
	log.Printf("本次处理%v个地址, 耗时%v", len(targets), time.Since(tStart))
}
