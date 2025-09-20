package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Prefix("/api/v1").Group(func(router route.Router) {
		AuthServiceRoutes(router)
	})
}
