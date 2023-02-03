package model

type InfoBpSeason struct {
	Sid       int32 `gorm:"column:sid;" json:"sid"` // 赛季id
	StartTime int32 `gorm:"column:start_time;" json:"start_time"`
	EndTime   int32 `gorm:"column:end_time;" json:"end_time"`
}

func (i *InfoBpSeason) TableName() string {
	return "info_bp_season"
}
