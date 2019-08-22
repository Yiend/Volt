package entitys

import "time"

type TagArticle struct {
	BaseEntity            `xorm:"extends"`
	TagId      int      `xorm:"int"`          //标签id
	ArticleId     int64      `xorm:"int"`          //文章id
	ArticleStatus int64      `xorm:"int"`          //内容状态
	ArticleTime   time.Time  `xorm:"datetime"`     //发布时间
}
