package batik

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func itemGetService(db *gorm.DB) (res ResponseItemDetail, err error) {

	var idItems []uint
	var nameItems string
	var idPajak []uint

	dataItem := []DataResponse{}
	db.Table("items").Select("id").Where("deleted_at IS NULL").Scan(&idItems)
	for i := range idItems {
		db.Table("items").Select("nama").Where("id = ?", idItems[i]).Find(&nameItems)
		db.Table("pajak_items").Select("id_pajak").Where("id_item = ?", idItems[i]).Find(&idPajak)
		arrayPajak := []ResPajak{}
		// fmt.Println(idPajak, idItems[i])
		if idPajak != nil {
			for p := range idPajak {
				pajak := Pajak{}
				db.Table("pajaks").Select("id, nama, rate").Where("id = ?", idPajak[p]).Find(&pajak)
				rate := fmt.Sprintf("%v", pajak.Rate)
				arrayPajak = append(arrayPajak, ResPajak{
					ID:   pajak.ID,
					Nama: pajak.Nama,
					Rate: rate + "%",
				})
			}
		}

		dataItem = append(dataItem, DataResponse{
			Id:    idItems[i],
			Nama:  nameItems,
			Pajak: arrayPajak,
		})
	}
	res = ResponseItemDetail{
		Data: dataItem,
	}
	return
}

func itemPostService(db *gorm.DB, req Item) (res ResponseItem, err error) {
	itemReq := Item{
		Nama:      req.Nama,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Select("nama").Create(&itemReq)

	res = ResponseItem{
		Success: true,
		Data: Item{
			ID:        itemReq.ID,
			CreatedAt: itemReq.CreatedAt,
			UpdatedAt: itemReq.UpdatedAt,
			Nama:      itemReq.Nama,
		},
	}

	return
}

func itemPatchService(db *gorm.DB, req Item) (res ResponseItem, err error) {
	itemReq := Item{
		Nama:       req.Nama,
		UpdatedAt:  time.Now(),
		ChangeName: req.ChangeName,
	}
	tx := db.Model(&Item{}).Where("nama=?", itemReq.Nama)
	var itemCheck Item

	check := tx
	check.Find(&itemCheck)
	if itemCheck.ID == 0 {
		res = ResponseItem{
			Success: false,
			Message: "Nama Tidak Ditemukan",
		}
		return
	}

	update := tx
	err = update.Updates(Item{Nama: itemReq.ChangeName, UpdatedAt: itemReq.UpdatedAt}).Error
	if err != nil {
		res = ResponseItem{
			Success: false,
			Message: err.Error(),
		}
		return
	}
	db.Table("items").Select("id, created_at, updated_at, nama").Where("nama = ?", itemReq.ChangeName).Scan(&itemCheck)
	if itemCheck.ID == 0 {
		res = ResponseItem{
			Success: false,
			Message: err.Error(),
		}
	}
	res = ResponseItem{
		Success: true,
		Data:    itemCheck,
	}

	return
}

func itemDeleteService(db *gorm.DB, req Item) (res ResponseItem, err error) {
	itemReq := Item{
		Nama:      req.Nama,
		DeletedAt: time.Now(),
	}
	tx := db.Model(&Item{}).Where("nama=?", itemReq.Nama)
	var itemCheck Item

	check := tx
	check.Find(&itemCheck)
	if itemCheck.ID == 0 {
		res = ResponseItem{
			Success: false,
			Message: "Nama Tidak Ditemukan",
		}
		return
	}

	update := tx
	err = update.Updates(Item{DeletedAt: itemReq.DeletedAt}).Error
	if err != nil {
		res = ResponseItem{
			Success: false,
			Message: err.Error(),
		}
		return
	}
	db.Table("items").Select("id, created_at, updated_at, deleted_at, nama").Where("nama = ?", itemReq.Nama).Scan(&itemCheck)
	if itemCheck.ID == 0 {
		res = ResponseItem{
			Success: false,
			Message: err.Error(),
		}
	}
	res = ResponseItem{
		Success: true,
		Data:    itemCheck,
	}

	return
}
