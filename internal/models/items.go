package models

type Item struct {
	ChrtId      int    `json:"chrt_id" fake:"{number:1, 5000}" `
	TrackNumber string `json:"track_number" fake:"{regex:[A-Z]{14}}"`
	Price       int    `json:"price" fake:"{number:1, 5000000}" `
	Rid         string `json:"rid" fake:"{regex:[0-9a-z]{21}}"`
	Name        string `json:"name" fake:"{minecraftarmorpart}"`
	Sale        int    `json:"sale" fake:"{number:1, 100}"`
	Size        string `json:"size" fake:"{regex:[0-1]{3}}"`
	TotalPrice  int    `json:"total_price" fake:"{number:1, 5000000}"`
	NmId        int    `json:"nm_id" fake:"{number:1, 2389212}"`
	Brand       string `json:"brand" fake:"{firstname} {lastname} "`
	Status      int    `json:"status" fake:"{regex:[0-1]{3}}"`
}
