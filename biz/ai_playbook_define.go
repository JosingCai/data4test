package biz

type AiPlaybook struct {
	PlaybookBase
	Product string `gorm:"column:product" json:"product"`
	CreateBase
}

type InputPlaybook struct {
	AiTemplate string `gorm:"column:ai_template" json:"ai_template"`
	CommonExtend
}

type PlaybookBase struct {
	PlaybookDesc string `gorm:"column:playbook_desc" json:"playbook_desc"`
	DataFileList string `gorm:"column:data_file_list" json:"data_file_list"`
	LastFile     string `gorm:"column:last_file" json:"last_file"`
	PlaybookType string `gorm:"column:playbook_type" json:"playbook_type"`
	Priority     string `gorm:"column:priority" json:"priority"`
	Result       string `gorm:"column:result" json:"result"`
	FailReason   string `gorm:"column:fail_reason" json:"fail_reason"`
}
