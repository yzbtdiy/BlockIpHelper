package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

// 读取配置文件, 返回结构体
func GetConfig(path string) (config models.Config) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Read file error: ", err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Yaml unmarshal error:  ", err)
	}
	return config
}

// 生成默认配置文件
func GenerateConfig() {
	config := models.Config{
		TargetFile: "./target.txt",
		WhiteFile:  "./data/whitelist.txt",
		ExportFile: models.ExportFileConf{
			InWhitelist: "./白名单IP.txt",
			InChina:     "./国内攻击IP.txt",
			OutCina:     "./国外攻击IP.txt",
		},
		Template: []models.TemplateConf{{Name: "k01", Enable: true, ExportPath: "./k01Block.csv"}, {Name: "myfw", Enable: true, ExportPath: "./myFwBlock.csv"}},
		Ip2Region: models.Ip2RegionConf{
			CzTxt:  "./data/qqwry.txt",
			MergeFile: "./data/ip_merge.txt",
			XdbFile:   "./data/ip2region.xdb",
			CnKeys: []string{"中国",
				"北京",
				"广东",
				"山东",
				"江苏",
				"河南",
				"上海",
				"河北",
				"浙江",
				"香港",
				"陕西",
				"湖南",
				"重庆",
				"福建",
				"天津",
				"云南",
				"四川",
				"广西",
				"安徽",
				"海南",
				"江西",
				"湖北",
				"山西",
				"辽宁",
				"台湾",
				"黑龙江",
				"内蒙古",
				"澳门",
				"贵州",
				"甘肃",
				"青海",
				"新疆",
				"西藏",
				"吉林",
				"宁夏"},
		},
	}
	conf, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./data/config.yaml", conf, 0644)
}
