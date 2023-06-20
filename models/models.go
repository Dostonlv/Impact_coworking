package models

import "github.com/jackc/pgtype"

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

// Resident ...
type Resident struct {
	Name string `json:"name"`
}

// BookingRequest ...
type BookingRequest struct {
	Resident Resident `json:"resident"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
}

// BookingResponse ...
type BookingResponse struct {
	Message string `json:"message"`
}

// AvailableRooms
type AvailableRooms struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Get Booking Rooms
type GetBookingRooms struct {
	ID       int            `json:"id"`
	Resident string         `json:"resident"`
	RoomID   int            `json:"room_id"`
	Period   pgtype.Tsrange `json:"period"`
}

// Get Booking Rooms Response
type GetBookingRoomsResponse struct {
	From any `json:"from"`
	To   any `json:"to"`
}

// Check
type Check struct {
	RoomID int    `json:"room_id"`
	Start  string `json:"start"`
	End    string `json:"end"`
}
