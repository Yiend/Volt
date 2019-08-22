package entitys

import "time"

type Article struct {
	BaseEntity            `xorm:"extends"`
	UserId      int32     `xorm:"int"`             //用户id
	Author      string    `xorm:"varchar(64)"`     //作者
	Title       string    `xorm:"varchar(100)"`    //标题
	Color       string    `xorm:"varchar(50)"`      //标题颜色
	UrlName		string    `xorm:"varchar(100)"`    //url名
	UrlType     int32       `xorm:"int"`             //url访问形式
	Content     string    `xorm:"text"`            //内容
	Tags        string    `xorm:"varchar(100)"`    //标签
	Views       int32       `xorm:"int"`             //查看次数
	Status      int32       `xorm:"int"`             //状态{0:正常,1:草稿,2:回收站}
	PublishTime time.Time `xorm:"datetime"`        //发布时间
	IsTop       bool      `xorm:"bit"`             //是否置顶
	IsHome      bool      `xorm:"bit"`             //是否首页显示
}
