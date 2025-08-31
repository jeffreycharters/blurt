package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jeffreycharters/blurt/.gen/model"
	. "github.com/jeffreycharters/blurt/.gen/table"

	. "github.com/go-jet/jet/v2/sqlite"
)

type LoadedBlurt struct {
	ID      string `sql:"primary_key"`
	Content string
	Created string
	Author  BlurtAuthor `alias:"author.*"`
}

type BlurtAuthor struct {
	ID       string `sql:"primary_key"`
	Username string
}

func (db *HandlerDB) SetBlurt(user_id string, blurt string) (*model.Blurt, error) {
	insert := Blurt.INSERT(Blurt.ID, Blurt.Content, Blurt.AuthorID, Blurt.Created).
		MODEL(model.Blurt{
			ID:       uuid.NewString(),
			Content:  blurt,
			AuthorID: user_id,
			Created:  time.Now().Format(time.RFC3339),
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

func (db *HandlerDB) GetBlurts(offset int) ([]LoadedBlurt, error) {
	stmt := SELECT(Blurt.AllColumns.Except(Blurt.AuthorID).As("loaded_blurt"),
		User.AllColumns.Except(User.Created).As("author")).
		FROM(Blurt.INNER_JOIN(User, Blurt.AuthorID.EQ(User.ID))).
		ORDER_BY(Blurt.Created.DESC()).
		OFFSET(int64(offset)).
		LIMIT(25)

	var blurts []LoadedBlurt
	err := stmt.Query(db.db, &blurts)
	if err != nil {
		return nil, err
	}

	return blurts, nil
}
