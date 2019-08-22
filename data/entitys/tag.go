package entitys

type Tag struct {
	BaseEntity            `xorm:"extends"`
	Name       string     `xorm:"varchar(64)"`
	Count      int32      `xorm:"int"`          //使用次数
}
