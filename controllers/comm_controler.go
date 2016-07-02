package controllers

type CommentController struct {
	BaseController
}

var comment_rules = map[string]int{
	"View":   0,
	//"CreateJump": identify.Login,
}

func (this *CommentController) getRules(action string) int {
	return comment_rules[action]
}

type Comment struct {
	Content string
	ID uint
}
type Comments struct {
	Size    uint
	Comments []Comment
}
/*show 20 comment ordered by submit time of one post
id post id;  start:comment offset  */
func (this *CommentController) Comment() {
	comments :=  Comments{Size:2,Comments:[]Comment{Comment{Content:"hello",ID:1},Comment{Content:"hello",ID:2}}}
	this.Data["json"] = &comments
	this.ServeJSON()
}