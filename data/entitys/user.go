package entitys

import "time"

// User 用户实体
type User struct {
	BaseEntity                `xorm:"extends"`
	UserName       string     `xorm:"varchar(64)"` // 用户名
	PassWord       string     `xorm:"varchar(40)"`  // 密码(sha1(md5(明文))加密)
	Email          string     `xorm:"varchar(255)"`    // 邮箱
	Phone          string     `xorm:"varchar(20)"`     // 手机号
	LoginCount     int32      `xorm:"int"`            // 状态(1:启用 2:停用)
	LastIp         string     `xorm:"varchar(36)"`         // 最后登录ip
	LastLoginTime  time.Time  `xorm:"detetime"`      //最后登录时间
	AuthKey        string     `xorm:"varchar(500)"`  //登录key
	Active         bool       `xorm:"bit"`  //是否激活
}
