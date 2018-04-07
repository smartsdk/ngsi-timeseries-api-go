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
		"/v2/",
		Index,
	},

	Route{
		"Retrieve API Resources",
		"GET",
		"/v2/",
		Retrieve API Resources,
	},

	Route{
		"Get attribute value",
		"GET",
		"/v2/entities/{entityId}/attrs/{attrName}/value",
		Get attribute value,
	},

	Route{
		"Update attribute value",
		"PUT",
		"/v2/entities/{entityId}/attrs/{attrName}/value",
		Update attribute value,
	},

	Route{
		"Get attribute data",
		"GET",
		"/v2/entities/{entityId}/attrs/{attrName}",
		Get attribute data,
	},

	Route{
		"Remove a single attribute",
		"DELETE",
		"/v2/entities/{entityId}/attrs/{attrName}",
		Remove a single attribute,
	},

	Route{
		"Update attribute data",
		"PUT",
		"/v2/entities/{entityId}/attrs/{attrName}",
		Update attribute data,
	},

	Route{
		"Query",
		"POST",
		"/v2/op/query",
		Query,
	},

	Route{
		"Update",
		"POST",
		"/v2/op/update",
		Update,
	},

	Route{
		"Create entity",
		"POST",
		"/v2/entities",
		Create entity,
	},

	Route{
		"List entities",
		"GET",
		"/v2/entities",
		List entities,
	},

	Route{
		"Remove entity",
		"DELETE",
		"/v2/entities/{entityId}",
		Remove entity,
	},

	Route{
		"Replace all entity attributes",
		"PUT",
		"/v2/entities/{entityId}/attrs",
		Replace all entity attributes,
	},

	Route{
		"Retrieve entity",
		"GET",
		"/v2/entities/{entityId}",
		Retrieve entity,
	},

	Route{
		"Retrieve entity attributes",
		"GET",
		"/v2/entities/{entityId}/attrs",
		Retrieve entity attributes,
	},

	Route{
		"Update existing entity attributes",
		"PATCH",
		"/v2/entities/{entityId}/attrs",
		Update existing entity attributes,
	},

	Route{
		"Update or append entity attributes",
		"POST",
		"/v2/entities/{entityId}/attrs",
		Update or append entity attributes,
	},

	Route{
		"Create a new subscription",
		"POST",
		"/v2/subscriptions",
		Create a new subscription,
	},

	Route{
		"Delete subscription",
		"DELETE",
		"/v2/subscriptions/{subscriptionId}",
		Delete subscription,
	},

	Route{
		"Retrieve subscription",
		"GET",
		"/v2/subscriptions/{subscriptionId}",
		Retrieve subscription,
	},

	Route{
		"Retrieve subscriptions",
		"GET",
		"/v2/subscriptions",
		Retrieve subscriptions,
	},

	Route{
		"Update subscription",
		"PATCH",
		"/v2/subscriptions/{subscriptionId}",
		Update subscription,
	},

	Route{
		"Retrieve entity type",
		"GET",
		"/v2/types/{entityType}",
		Retrieve entity type,
	},

	Route{
		"Retrieve entity types",
		"GET",
		"/v2/types/",
		Retrieve entity types,
	},

}
