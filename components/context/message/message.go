package message

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
)

type Notification struct {
	ID          uint
	RelatedID   uint
	SubjectType int
	Data        string
	IsRead      bool
}

type PostMessage struct {
	ID              uint
	RelatedID       uint
	SubjectType     int
	RelatedUsername string
	PostID          uint
	PostTitle       string
	IsRead          bool
}

func FindLatestNotifications(uid uint, types []int) ([]Notification) { //todo 指针
	//todo user id
	notices := []m.Notification{}
	database.O.QueryTable("notification").Filter("subject_type__in", types).All(&notices)
	//database.DB.Where("subject_type in (?)", types).Find(&notices)
	n := make([]Notification, 0, len(notices))
	for _, no := range notices {
		n = append(n, Notification{ID:no.Id, RelatedID:no.Related.Id,
			SubjectType:no.SubjectType, Data:no.Data, IsRead:no.IsRead});
	}
	return n
}

func FindLatestPostMessages(uid uint, types []int) ([]PostMessage) {
	//todo user id
	db_messages := []m.PostMessage{}
	database.O.QueryTable("post_message").Filter("subject_type__in", types).All(&db_messages)
	messages := make([]PostMessage, 0, len(db_messages))
	for _, msg := range db_messages {
		messages = append(messages, PostMessage{ID:msg.Id, RelatedID:msg.Related.Id, SubjectType:msg.SubjectType,
			RelatedUsername:msg.RelatedUsername, PostID:msg.PostId, PostTitle:msg.PostTitle, IsRead:msg.IsRead});
	}
	return messages
}
