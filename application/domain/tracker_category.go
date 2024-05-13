package domain

type TrackerCategory struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
	Now string `json:"now" gorm:"column:now"`
	End string `json:"end" gorm:"column:end"`
}


func (c *TrackerCategory) TableName() string {
	return "tracker_category"
}