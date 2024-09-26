package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"htmgo-site/__htmgo"
	"htmgo-site/internal/markdown"
	"io/fs"
	"net/http"
)

func main() {
	locator := service.NewLocator()
	staticAssets := GetStaticAssets()
	markdownAssets := GetMarkdownAssets()

	service.Set(locator, service.Singleton, markdown.NewRenderer)

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {

			app.UseWithContext(func(w http.ResponseWriter, r *http.Request, context map[string]any) {
				context["embeddedMarkdown"] = markdownAssets
			})

			sub, err := fs.Sub(staticAssets, "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))

			__htmgo.Register(app.Router)
		},
	})
}
