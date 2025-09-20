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

func (db *HandlerDB) GetBlurts(offset int, count int) ([]LoadedBlurt, error) {
	stmt := SELECT(
		Blurt.ID.AS("loaded_blurt.id"),
		Blurt.Content.AS("loaded_blurt.content"),
		Blurt.Created.AS("loaded_blurt.created"),
		Blurt.Author.AS("loaded_blurt.author"),
		Lik.Username.AS("loaded_blurt.likkers")).
		FROM(Blurt.
			LEFT_JOIN(Lik, Lik.BlurtID.EQ(Blurt.ID)),
		).
		ORDER_BY(Blurt.Created.DESC()).
		LIMIT(int64(count)).
		OFFSET(int64(offset))

	var blurts []LoadedBlurt
	err := stmt.Query(db.db, &blurts)
	if err != nil {
		return nil, err
	}

	if len(blurts) == 0 {
		return []LoadedBlurt{}, nil
	}

	return blurts, nil
}
