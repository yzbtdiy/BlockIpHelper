package template

import "github.com/yzbtdiy/BlockIpHelper/models"

func GenerateCSV(inCns, notInCns []models.IpAndRegion, temp models.TemplateConf) {
	if temp.Enable == true {
		if temp.Name == "k01" {
			GenerateK01CSVFile(inCns, notInCns, temp.ExportPath)
		} else if temp.Name == "myfw" {
			GenerateMyFwCSVFile(inCns, notInCns, temp.ExportPath)
		}
	}
}
