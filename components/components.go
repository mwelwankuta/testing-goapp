package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Loading struct {
	app.Compo
}

func (l *Loading) Render() app.UI {
	return app.H1().Text("Loading...")
}
