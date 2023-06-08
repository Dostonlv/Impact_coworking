package models

import "time"

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
}

// Rooms Response ...
type RoomsResponse struct {
	Page     int    `json:"page"`
	Count    int    `json:"count"`
	PageSize int    `json:"page_size"`
	Results  []Room `json:"results"`
}

// RoomsRequest ...
type RoomsRequest struct {
	Search   string `json:"search"`
	Type     string `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

// CreateRoomRequest ...
type CreateRoomRequest struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
}

// Resident ...
type Resident struct {
	Name string `json:"name"`
}

// CreateBookingRequest ...
type CreateBookingRequest struct {
	Resident Resident  `json:"resident"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}
