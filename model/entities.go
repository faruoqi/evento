package model

type Event struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Duration  int      `json:"duration"`
	StartDate int64    `json:"start_date"`
	EndDate   int64    `json:"end_date"`
	Location  Location `json:"location"`
}

type Location struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	OpenTime  int    `json:"open_time"`
	CloseTime int    `json:"close_time"`
	Halls     []Hall `json:"halls"`
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
