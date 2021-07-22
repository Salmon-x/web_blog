package utils

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)


// 渲染html
func Render(data string) string {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	md := []byte(data)
	content := markdown.ToHTML(md, nil, renderer)
	return string(content)
}