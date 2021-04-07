package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.ListUsers,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.FindUserById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUserById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUserById,
		RequireAuthentication: false,
	},
}
