package routes

import (
	"fmt"

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

func (h *Home) OnNav(ctx app.Context) {
	ctx.Page().SetTitle("Pokimon App | Home")
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
			app.Div().Body(
				app.H1().Text("Pokimon + Go Lang").Class("text-2xl font-semibold text-center"),
				app.Range(h.Data).Slice(func(idx int) app.UI {
					item := h.Data[idx]
					pokimonUrl := fmt.Sprintf("%v/%v", "/pokimon", item.Url)

					return app.Ul().Body(
						app.Li().Body(
							app.P().Text(idx+1).Class("mr-2"),
							app.A().Text(item.Name).Href(pokimonUrl).Class("w-full"),
						).Class("py-[20px] border-y border-gray-400 flex items-center"),
					)
				}),
			),
		),
	).Class("container mx-auto")
}
