package entitys

type Option struct {
	BaseEntity            `xorm:"extends"`
	Name       string     `xorm:"varchar(64)"`
	Value      string     `xorm:"text"`
}
