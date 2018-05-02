package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/context/v2/",
		Index,
	},

	Route{
		"Retrieve API Resources",
		"GET",
		"/context/v2/",
		Retrieve API Resources,
	},

	Route{
		"Get attribute value",
		"GET",
		"/context/v2/entities/{entityId}/attrs/{attrName}/value",
		Get attribute value,
	},

	Route{
		"Update attribute value",
		"PUT",
		"/context/v2/entities/{entityId}/attrs/{attrName}/value",
		Update attribute value,
	},

	Route{
		"Get attribute data",
		"GET",
		"/context/v2/entities/{entityId}/attrs/{attrName}",
		Get attribute data,
	},

	Route{
		"Remove a single attribute",
		"DELETE",
		"/context/v2/entities/{entityId}/attrs/{attrName}",
		Remove a single attribute,
	},

	Route{
		"Update attribute data",
		"PUT",
		"/context/v2/entities/{entityId}/attrs/{attrName}",
		Update attribute data,
	},

	Route{
		"Query",
		"POST",
		"/context/v2/op/query",
		Query,
	},

	Route{
		"Update",
		"POST",
		"/context/v2/op/update",
		Update,
	},

	Route{
		"Create entity",
		"POST",
		"/context/v2/entities",
		Create entity,
	},

	Route{
		"List entities",
		"GET",
		"/context/v2/entities",
		List entities,
	},

	Route{
		"Remove entity",
		"DELETE",
		"/context/v2/entities/{entityId}",
		Remove entity,
	},

	Route{
		"Replace all entity attributes",
		"PUT",
		"/context/v2/entities/{entityId}/attrs",
		Replace all entity attributes,
	},

	Route{
		"Retrieve entity",
		"GET",
		"/context/v2/entities/{entityId}",
		Retrieve entity,
	},

	Route{
		"Retrieve entity attributes",
		"GET",
		"/context/v2/entities/{entityId}/attrs",
		Retrieve entity attributes,
	},

	Route{
		"Update existing entity attributes",
		"PATCH",
		"/context/v2/entities/{entityId}/attrs",
		Update existing entity attributes,
	},

	Route{
		"Update or append entity attributes",
		"POST",
		"/context/v2/entities/{entityId}/attrs",
		Update or append entity attributes,
	},

	Route{
		"Create a new subscription",
		"POST",
		"/context/v2/subscriptions",
		Create a new subscription,
	},

	Route{
		"Delete subscription",
		"DELETE",
		"/context/v2/subscriptions/{subscriptionId}",
		Delete subscription,
	},

	Route{
		"Retrieve subscription",
		"GET",
		"/context/v2/subscriptions/{subscriptionId}",
		Retrieve subscription,
	},

	Route{
		"Retrieve subscriptions",
		"GET",
		"/context/v2/subscriptions",
		Retrieve subscriptions,
	},

	Route{
		"Update subscription",
		"PATCH",
		"/context/v2/subscriptions/{subscriptionId}",
		Update subscription,
	},

	Route{
		"Retrieve entity type",
		"GET",
		"/context/v2/types/{entityType}",
		Retrieve entity type,
	},

	Route{
		"Retrieve entity types",
		"GET",
		"/context/v2/types/",
		Retrieve entity types,
	},

}
