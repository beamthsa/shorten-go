package model

import (
	"github.com/opn-ooo/gin-boilerplate/config/database"
	"github.com/uptrace/bun"
	"log"
)

type UrlModel struct {
	bun.BaseModel   `bun:"url,alias:u"`
	UrlID      		uint64 `bun:",pk"`
	Url     		string
}

func GetUrl(id uint64) ([]UrlModel, error) {
	data := new([]UrlModel)

	err := database.PGConnection.NewSelect().
		Model(data).
		Where("url_id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		log.Fatal("Error GetUrl:", id, err)
		return *data, err
	}

	return *data, nil
}

func InsertUrl(value UrlModel) (UrlModel, error) {
	data := value

	_, err := database.PGConnection.NewInsert().
		Model(&data).
		Column("url_id", "url").
		Exec(ctx)
	if err != nil {
		log.Fatal("Error InsertUrl:", err)
		return UrlModel{}, err
	}

	return data, nil
}
