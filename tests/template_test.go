package test

import (
	"testing"
	"os"
	"text/template"
)

const (
	view_temp = "../view_temp/"
	view_out = "../views/"
)

type Data struct {
	StaticPath string
	Title      string
}

func TestLogin(t *testing.T) {
	d := Data{"../", "登录"}
	generate(view_temp + "account/signin.html",view_out+"account/signin.html",d)
}

func TestRegister(t *testing.T) {
	d := Data{"../", "注册"}
	generate(view_temp + "account/signup.html",view_out+"account/signup.html",d)
}

func generate(container string,out string,data Data){
	s1, _ := template.ParseFiles(view_temp + "layouts/header.html",view_temp + "layouts/navbar.html",
		container, view_temp + "layouts/center_script.html",
		view_temp + "layouts/footer.html")
	file, _ := os.Create(out);
	s1.ExecuteTemplate(file, "content", data)
}