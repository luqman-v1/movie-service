package entity

import "encoding/json"

type List struct {
	Search       []Search `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type Search struct {
	Title  string `json:"Title" db:"title"`
	Year   string `json:"Year" db:"year"`
	ImdbID string `json:"imdbID" db:"imdb_id"`
	Type   string `json:"Type" db:"type"`
	Poster string `json:"Poster" db:"poster"`
}

func (l *List) Decode(data []byte) error {
	return json.Unmarshal(data, &l)
}

type Detail struct {
	Response string `json:"Response"`
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
}

func (l *Detail) Decode(data []byte) error {
	return json.Unmarshal(data, &l)
}

func (l *Detail) Encode() ([]byte, error) {
	return json.Marshal(l)
}
