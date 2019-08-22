package entitys

import "time"

// Model base model
type BaseEntity struct {
	ID         int32       `xorm:"pk autoincr notnull unique"`
	CreateTime time.Time  `xorm:"created"`
	UpdateTime time.Time  `xorm:"updated"`
}
