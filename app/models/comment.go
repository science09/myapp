package models

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"time"
)

type Comment struct {
	BlogId  bson.ObjectId
	Email   string
	CDate   time.Time
	Content string
}

func (comment *Comment) Validate(v *revel.Validation) {
	v.Check(comment.Email,
		revel.Required{},
		revel.MaxSize{50})
	v.Email(comment.Email)
	v.Check(comment.Content,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{10000})
}

func (dao *Dao) InsertComment(comment *Comment) error {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comment.CDate = time.Now()
	err := commentCollection.Insert(comment)
	if err != nil {
		revel.WARN.Panicf("Unable to save Comment: %v error %v", comment, err)
	}
	return err
}

func (dao *Dao) FindCommentsByBlogId(id bson.ObjectId) []Comment {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comms := []Comment{}
	query := commentCollection.Find(bson.M{"blogid": id}).Sort("CDate")
	query.All(&comms)
	return comms
}
