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
	Target			string
}

func GetUrl(url string) ([]UrlModel, error) {
	data := new([]UrlModel)

	err := database.PGConnection.NewSelect().
		Model(data).
		Where("url = ?", url).
		Limit(1).
		Scan(ctx)
	if err != nil {
		log.Fatal("Error GetUrl:", url, err)
		return *data, err
	}

	return *data, nil
}

func InsertUrl(value UrlModel) (UrlModel, error) {
	data := value

	_, err := database.PGConnection.NewInsert().
		Model(&data).
		Column("url_id", "url", "target").
		Exec(ctx)
	if err != nil {
		log.Fatal("Error InsertUrl:", err)
		return UrlModel{}, err
	}

	return data, nil
}
