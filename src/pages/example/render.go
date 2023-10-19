package example

import "github.com/caleb-sideras/goxstack/gox/render"

func TodoCsr_() render.StaticF {
	return render.StaticF{[]string{"pages/example/_components/todo_csr.html"}, nil, "todo-csr"}
}
func TodoSsr_() render.StaticF {
	return render.StaticF{[]string{"pages/example/_components/todo_ssr.html"}, nil, "todo-ssr"}
}
