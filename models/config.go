package models

type Config struct {
	TargetFile string         `yaml:"target_file"`
	WhiteFile  string         `yaml:"in_white_list"`
	ExportFile ExportFileConf `yaml:"export_file"`
	Template   []TemplateConf `yaml:"template"`
	Ip2Region  Ip2RegionConf  `yaml:"ip2region"`
}

type ExportFileConf struct {
	InWhitelist string `yaml:"in_whitelist"`
	InChina     string `yaml:"in_china"`
	OutCina     string `yaml:"out_china"`
}

type TemplateConf struct {
	Name       string `yaml:"name"`
	Enable     bool   `yaml:"enable"`
	ExportPath string `yaml:"export_path"`
}
type Ip2RegionConf struct {
	CzSource  string   `yaml:"cz_source"`
	MergeFile string   `yaml:"merge_file"`
	XdbFile   string   `yaml:"xdb_file"`
	CnKeys    []string `yaml:"cn_keys"`
}
