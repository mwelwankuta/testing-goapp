package routes

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mwelwankuta/goweb/api"
	"github.com/mwelwankuta/goweb/components"
)

func (h *Home) OnMount(ctx app.Context) {
	response, err := api.GetPokimon()
	h.Error = err
	h.Data = response

	h.Update()
}

type Home struct {
	app.Compo
	Loading bool
	Data    api.Pokimen
	Error   string
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.If(
			h.Loading,
			&components.Loading{},
		).Else(
			app.Range(h.Data).Slice(func(idx int) app.UI {
				item := h.Data[idx]
				return app.Div().Body(
					app.P().Text(item.Name),
					app.A().Text(item.Url).Href(item.Url),
				)
			}),
		),
	)
}
