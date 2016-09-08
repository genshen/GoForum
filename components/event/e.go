package event

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
	u "gensh.me/goforum/components/utils"
)

func OnAccountCreated(email string, username string, uid uint) {
//todo none
}

//todo: none
func OnPostCreated() {

}

//Posts:id,title,comment_count
func OnCommentSubmitted(post *m.Posts, comment *m.Comment, username string) {
	var message = m.PostMessage{RelatedUsername:username, PostId:post.Id, PostTitle:post.Title, Quote:post.Summary,
		Summary:comment.Content, SubjectType:u.POST_COMMENT,
		User:&m.Profile{Id:post.Author.Id}, Related:&m.Profile{Id:comment.Author}}
	database.O.Insert(&message)
	//多次create,message 可以复用
}

/*id userId;id_r:related_id;name:*/
func OnFollowed(id uint, id_r uint, name string) {
	var notify = m.Notification{Data:"{\"username\":\"" + name + "\"}",
		User:&m.Profile{Id:id}, Related:&m.Profile{Id:id_r}, SubjectType:u.FOLLOW_ADD}
	database.O.Insert(&notify)
}

func OnUnFollowed() {

}