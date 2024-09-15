package template

import "github.com/toolsascode/gomodeler/internal/template/render"

func renderFile() RenderInterface {
	return &render.Commands{}
}

func RunRender() {
	renderFile().Run()
}
