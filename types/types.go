package types

type ImdbResponse struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
}

type Movie struct {
	Title  string `json:Title`
	Year   string `json:Year`
	ImdbID string `json:imdbID`
	Type   string `json:Type`
	Poster string `json:Poster`
}
