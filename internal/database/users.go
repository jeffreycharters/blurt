package database

import (
	"log/slog"

	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/google/uuid"
	"github.com/jeffreycharters/blurt/.gen/model"
	. "github.com/jeffreycharters/blurt/.gen/table"
)

func (db *HandlerDB) GetUsers() []model.User {
	return []model.User{}
}

func (db *HandlerDB) SetUser(username string) (*model.User, error) {
	user := model.User{
		ID:       uuid.NewString(),
		Username: username,
	}

	stmt := User.INSERT(User.ID, User.Username).
		MODEL(user).
		ON_CONFLICT(User.Username).
		DO_NOTHING()

	if _, err := stmt.Exec(db.db); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *HandlerDB) GetUserCount() (int, error) {
	stmt := SELECT(COUNT(User.ID).AS("Count")).FROM(User)

	var dest struct {
		Count int
	}

	err := stmt.Query(db.db, &dest)
	if err != nil {
		return 0, err
	}

	slog.Info("got user count", "count", dest)
	return dest.Count, nil
}
