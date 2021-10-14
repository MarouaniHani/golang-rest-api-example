package model

type Product struct {
	ID       string `json:"id" db:"id"`
	Category string `json:"category" db:"category"`
	NameLen  int    `json:"name_len" db:"name_len"`
	DescLen  int    `json:"desc_len" db:"desc_len"`
	Photos   int    `json:"photos" db:"photos"`
	Weight   int    `json:"weight" db:"weight"`
	Length   int    `json:"length" db:"length"`
	Height   int    `json:"height" db:"height"`
	Width    int    `json:"width" db:"width"`
}
