package model

import "time"


type ShareBasic struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	ExpiredTime        int
	ClickNum           int
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`  // 可以更新时间戳的 tag 
 }

func (table UserBasic) ShareBasic() string {
	return "share_basic"
}