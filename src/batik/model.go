package batik

import "time"

type (
	Pajak struct {
		ID   uint    `json:"id"`
		Nama string  `json:"nama"`
		Rate float64 `json:"rate"`
	}

	Item struct {
		ID         uint      `json:"id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		DeletedAt  time.Time `json:"deleted_at"`
		Nama       string    `json:"nama" gorm:"nama"`
		ChangeName string    `json:"change_nama"`
	}

	ResponseItem struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    Item   `json:"data"`
	}

	ResponseItemDetail struct {
		Success bool           `json:"success"`
		Message string         `json:"message"`
		Data    []DataResponse `json:"data"`
	}
	DataResponse struct {
		Id    uint       `json:"id"`
		Nama  string     `json:"nama"`
		Pajak []ResPajak `json:"pajak" gorm:"-"`
	}

	ResPajak struct {
		ID   uint   `json:"id"`
		Nama string `json:"nama"`
		Rate string `json:"rate"`
	}
)
