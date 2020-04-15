package main

import "time"

type ReviewItem struct {
	Id      uint64   `json:"id"`
	Title   string   `json:"title"`
	Length  uint32   `json:"length"`
	Artists []string `json:"artists"`
}

type ReviewEvent struct {
	Id             uint64    `json:"id"`
	Description    string    `json:"description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	SpotifyId      string    `json:"spotify_id"`
	ParticipantIds []uint64  `json:"participant_ids"`
	ReviewItemIds  []uint64  `json:"review_item_ids"`
}

type UserReview struct {
	Id            uint64    `json:"id"`
	UserId        uint64    `json:"user_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	DateSubmitted time.Time `json:"date_submitted"`
	ReviewEventId uint64    `json:"review_event_id"`
	ReviewItemIds []uint64  `json:"review_item_ids"`
}
