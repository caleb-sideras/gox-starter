package example

import "github.com/caleb-sideras/goxstack/gox/utils"

func TodoCsr_() utils.RenderFilesType {
	return utils.RenderFilesType{nil, []string{"pages/example/_components/todo_csr.html"}, "todo-csr"}
}
func TodoSsr_() utils.RenderFilesType {
	return utils.RenderFilesType{nil, []string{"pages/example/_components/todo_ssr.html"}, "todo-ssr"}
}
