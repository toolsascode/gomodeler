package template

import "github.com/toolsascode/gomodeler/internal/template/render"

type RenderInterface interface {
	Run()
	RenderFile(data *render.Data, filename string, outputFile *string)
}
