package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{FirstName: "Chris", LastName: "Elten"},
		User{FirstName: "Riley", LastName: "Draward"},
		User{FirstName: "Saim", LastName: "Khan"},
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseUint(vars["user_id"], 0, 64)
	user := User{
		FirstName: "Saim",
		LastName:  "Khan",
		Id:        userId,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func UserReviewList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseUint(vars["user_id"], 0, 64)
	reviews := []UserReview{
		UserReview{Title: "I cant believe my ears", UserId: userId},
		UserReview{Title: "This sucks", UserId: userId},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		panic(err)
	}
}

func UserReviewDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseUint(vars["user_id"], 0, 64)
	reviewId, _ := strconv.ParseUint(vars["review_id"], 0, 64)
	review := UserReview{
		Title:         "Its adorable that they try",
		Description:   "Listening to this is like giving a baby a guitar. it looks cute, but you know that baby isnt playing anything worthwhile",
		UserId:        userId,
		ReviewEventId: 1,
		Id:            reviewId,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(review); err != nil {
		panic(err)
	}
}
