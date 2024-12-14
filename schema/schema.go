package schema

type RouteSchema struct {
	Route  string
	Params []string
}

var Schema []RouteSchema

func Init() {
	Schema = []RouteSchema{
		{
			Route:  "/register",
			Params: []string{"name", "email", "password"},
		},
		{
			Route:  "/login",
			Params: []string{"email", "password"},
		},
	}
}
