package domain

type Event struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Sold     int    `json:"sold"`
}
