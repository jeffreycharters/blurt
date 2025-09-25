package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jeffreycharters/blurt/.gen/model"
	. "github.com/jeffreycharters/blurt/.gen/table"

	. "github.com/go-jet/jet/v2/sqlite"
)

type LoadedBlurt struct {
	ID      string   `sql:"primary_key" json:"id"`
	Content string   `json:"content"`
	Created string   `json:"created"`
	Author  string   `json:"author"`
	Likkers []string `json:"likkers" alias:"loaded_blurt.likkers"`
}

func (db *HandlerDB) SetBlurt(username string, blurt string) (*LoadedBlurt, error) {
	insert := Blurt.INSERT(Blurt.ID, Blurt.Content, Blurt.Author, Blurt.Created).
		MODEL(model.Blurt{
			ID:      uuid.NewString(),
			Content: blurt,
			Author:  username,
			Created: time.Now().Format(time.RFC3339),
		}).RETURNING(Blurt.AllColumns.As("loaded_blurt"))

	var dest struct {
		Blurt LoadedBlurt
	}
	err := insert.Query(db.db, &dest)
	if err != nil {
		return nil, err
	}

	return &dest.Blurt, nil
}

func (db *HandlerDB) GetBlurts(offset int64, count int64) ([]LoadedBlurt, error) {
	selected_blurts := SELECT(Blurt.AllColumns).
		FROM(Blurt).
		ORDER_BY(Blurt.Created.DESC()).
		OFFSET(offset).
		LIMIT(count).
		AsTable("selected_blurts")

	blurt_id := Blurt.ID.From(selected_blurts)

	stmt := SELECT(
		selected_blurts.AllColumns().As("loaded_blurt"),
		Lik.Username.AS("loaded_blurt.likkers")).
		FROM(selected_blurts.
			LEFT_JOIN(Lik, Lik.BlurtID.EQ(blurt_id)))

	var blurts []LoadedBlurt
	err := stmt.Query(db.db, &blurts)
	if err != nil {
		return nil, err
	}

	return blurts, nil
}
