package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UsersList",
		"GET",
		"/users",
		UsersList,
	},
	Route{
		"UserDetail",
		"GET",
		"/users/{user_id}",
		UserDetail,
	},
	Route{
		"UserReviewList",
		"GET",
		"/users/{user_id}/reviews",
		UserReviewList,
	},
	Route{
		"UserReviewDetails",
		"GET",
		"/users/{user_id}/reviews/{review_id}",
		UserReviewDetails,
	},
}
