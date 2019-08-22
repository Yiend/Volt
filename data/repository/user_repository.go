package repository

import (
	"github.com/go-xorm/xorm"
	"volt/data/entitys"
)

type UserRepository struct {
	db *xorm.Engine
}

func NewUserRepository(db *xorm.Engine) *UserRepository {
	return &UserRepository{db}
}

func (this *UserRepository) GetById(Id int32) (*entitys.User,error) {
	user := new(entitys.User)
	_, err := this.db.Id(Id).Get(user)
	return user,err
}

func (this *UserRepository) GetByName(name string) (*entitys.User,error)  {
	user := new(entitys.User)
	_, err := this.db.Alias("o").Where("o.username = ?", name).Get(user)
	return user,err
}


func (this *UserRepository) SearchByPage() ([]entitys.Article, int, error) {
	total, err := this.db.Count(this)
	var posts []entitys.Article

	session := this.db.NewSession()
	defer session.Close()
	//if len(this.GetSortProperties()) > 0 {
	//	session = session.OrderBy(this.GetSortProperties()[0].Column + " " + self.GetSortProperties()[0].Direction)
	//}
	//err = session.Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&posts, this)
	return posts, int(total), err
}

