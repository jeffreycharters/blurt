package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jeffreycharters/blurt/.gen/model"
	. "github.com/jeffreycharters/blurt/.gen/table"

	. "github.com/go-jet/jet/v2/sqlite"
)

type LoadedBlurt struct {
	ID       string `sql:"primary_key"`
	Content  string
	Created  string
	Author   string
	Likd     bool
	LikCount int
}

func (db *HandlerDB) SetBlurt(username string, blurt string) (*model.Blurt, error) {
	insert := Blurt.INSERT(Blurt.ID, Blurt.Content, Blurt.Author, Blurt.Created).
		MODEL(model.Blurt{
			ID:      uuid.NewString(),
			Content: blurt,
			Author:  username,
			Created: time.Now().Format(time.RFC3339),
		}).RETURNING(Blurt.AllColumns)

	var dest struct {
		Blurt model.Blurt
	}
	err := insert.Query(db.db, &dest)
	if err != nil {
		return nil, err
	}

	return &dest.Blurt, nil
}

func (db *HandlerDB) GetBlurts(username string, offset int) ([]LoadedBlurt, error) {
	stmt := SELECT(
		Blurt.ID.AS("loaded_blurt.id"),
		Blurt.Content.AS("loaded_blurt.content"),
		Blurt.Created.AS("loaded_blurt.created"),
		Blurt.Author.AS("loaded_blurt.author"),
		EXISTS(
			Lik.
				SELECT(Lik.Username).
				WHERE(
					Lik.Username.EQ(String(username)).
						AND(Lik.BlurtID.EQ(Blurt.ID)))).AS("loaded_blurt.likd"),
		SELECT(
			COUNT(Lik.BlurtID)).
			FROM(Lik).
			WHERE(Lik.BlurtID.EQ(Blurt.ID)).AS("loaded_blurt.lik_count"),
	).
		FROM(Blurt.
			LEFT_JOIN(Lik, Lik.BlurtID.EQ(Blurt.ID)),
		).
		ORDER_BY(Blurt.Created.DESC()).
		OFFSET(int64(offset)).
		LIMIT(25)

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
