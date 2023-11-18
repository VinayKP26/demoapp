package router

import (
	"context"

	"github.com/VinayKP26/demoapp/controllers"
)

type Handlers struct {
	handlers controllers.Controllers
}

func NewHandlers(handlers controllers.Controllers) Handlers {
	return Handlers{handlers}
}

func StartServer(port string, ctx context.Context) {
	c := controllers.NewControllers(ctx)
	h := NewHandlers(c)
	h.Server(port)
}
