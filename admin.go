package models

type ManualBoatAdd struct {
	StateNumber  string `json:"state_number"`
	ReturnTime   string `json:"return_time"`
	ReturnDate   string `json:"return_date"`
	SwimmingArea string `json:"swimming_area"`
}
