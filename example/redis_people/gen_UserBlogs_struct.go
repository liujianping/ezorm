package test

import "time"

var _ time.Time

type UserBlogs struct {
	UserId int32 `json:"user_id"`
	BlogId int32 `json:"blog_id"`
	isNew  bool
}

func (p *UserBlogs) GetNameSpace() string {
	return "people"
}

func (p *UserBlogs) GetClassName() string {
	return "UserBlogs"
}
func (p *UserBlogs) GetStoreType() string {
	return "set"
}

func (p *UserBlogs) GetPrimaryKey() string {
	return "UserId"
}

func (p *UserBlogs) GetIndexes() []string {
	idx := []string{}
	return idx
}

type _UserBlogsMgr struct {
}

var UserBlogsMgr *_UserBlogsMgr

func (m *_UserBlogsMgr) NewUserBlogs() *UserBlogs {
	rval := new(UserBlogs)
	return rval
}
