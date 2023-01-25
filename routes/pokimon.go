package routes

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mwelwankuta/goweb/api"
	"github.com/mwelwankuta/goweb/components"
)

func (p *Pokimon) OnNav(ctx app.Context) {

	p.Loading = true

	url := app.Window().URL()
	id := strings.Split(string(url.Path), "/pokimon/")[1]

	response, err := api.GetCurrentPokimon(id)

	if err != nil {
		p.Error = err.Error()
	}

	p.Data = response
	ctx.Page().SetTitle("Pokimon App | " + p.Data.Name)

	p.Loading = false
	p.Update()
}

type Pokimon struct {
	app.Compo
	Loading bool
	Data    api.CurrentPokimon
	Error   string
}

func (p *Pokimon) Render() app.UI {
	return app.Main().Body(
		app.If(
			p.Loading, &components.Loading{},
		).Else(app.Div().Body(
			app.Div().Body(
				app.P().Text(
					string(p.Data.Name),
				).Class("text-2xl font-semibold text-center"),
				app.Div().Body(
					app.Img().Src(p.Data.Sprites.FrontDefault).Class("w-48"),
					app.Img().Src(p.Data.Sprites.BackDefault).Class("w-48"),
				).Class("flex items-center"),
			),
		)),
	).Class("container mx-auto")
}
