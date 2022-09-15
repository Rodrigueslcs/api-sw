package entities

type Planet struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"terrain"`
	Apparitions int    `json:"apparitions"`
}
