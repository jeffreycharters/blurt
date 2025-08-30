package database

import (
	"time"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/jeffreycharters/blurt/.gen/model"
	. "github.com/jeffreycharters/blurt/.gen/table"
)

func (db *HandlerDB) GetUsers() []model.User {
	return []model.User{}
}

func (db *HandlerDB) SetUser(username string) (*model.User, error) {
	stmt := SELECT(User.AllColumns).
		FROM(User).
		WHERE(User.Username.EQ(String(username))).
		LIMIT(1)

	var dest struct {
		User model.User
	}

	err := stmt.Query(db.db, &dest)
	if err == nil {
		return &dest.User, nil
	}

	if err != qrm.ErrNoRows {
		return nil, err
	}

	user := model.User{
		Username: username,
		Created:  time.Now().Format(time.RFC3339),
	}

	insert := User.INSERT(User.Username, User.Created).
		MODEL(user)

	if _, err := insert.Exec(db.db); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *HandlerDB) GetUserCount() (int, error) {
	stmt := SELECT(COUNT(User.Username).AS("Count")).FROM(User)

	var dest struct {
		Count int
	}

	err := stmt.Query(db.db, &dest)
	if err != nil {
		return 0, err
	}

	return dest.Count, nil
}
