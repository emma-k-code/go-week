package kernel

type Weather struct {
	City    string `gorm:"column:city;type:varcher(20);"`
	Content string `gorm:"column:content;type:text;"`
}

// 設定 Table 名稱
func (Weather) TableName() string {
	return "weather"
}
