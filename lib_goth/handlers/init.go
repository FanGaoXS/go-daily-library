package handlers

import "html/template"

var (
	ptTemplate *template.Template
)

func Init() {
	// 读取templates下的模板文件
	ptTemplate = template.Must(template.New("").ParseGlob("lib_goth/templates/*.tpl"))
}
