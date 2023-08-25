package example

import "github.com/caleb-sideras/gox/utils"

func TodoCsr() (interface{}, []string, string) {
	return utils.Render{}.DefaultRender(nil, []string{"pages/example/_components/todo_csr.html"}, "todo-csr")
}
func TodoSsr() (interface{}, []string, string) {
	return utils.Render{}.DefaultRender(nil, []string{"pages/example/_components/todo_ssr.html"}, "todo-ssr")
}
