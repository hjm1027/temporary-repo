package model

// 必须要写一个TableName函数，返回table的名字，否则gorm读取不到表。
func (u *MoodModel) TableName() string {
	return "mood"
}

// 添加心情
func (mood *MoodModel) New() error {
	d := DB.Self.Create(mood)
	return d.Error
}