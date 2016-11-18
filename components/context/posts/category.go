package posts

import (
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
)

type Category struct {
	Tags   []m.Tag
	Topics []m.Topic
}

func (this *Category)NewInstant() {
	database.O.QueryTable("topic").Filter("visible", true).All(&this.Topics)
	database.O.QueryTable("tag").Filter("visible", true).All(&this.Tags)
	//RW database.DB.Where("visible = ?", true).Find(&this.Topics)
	//RW database.DB.Where("visible = ?", true).Find(&this.Tags)
}

