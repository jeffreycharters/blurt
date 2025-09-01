package database

import (
	"github.com/jeffreycharters/blurt/.gen/model"

	. "github.com/jeffreycharters/blurt/.gen/table"
)

func (v *HandlerDB) LikBlurt(username string, blurtID string) error {
	user, err := v.SetUser(username)
	if err != nil {
		return err
	}

	insert := Lik.INSERT(Lik.Username, Lik.BlurtID).
		MODEL(model.Lik{
			Username: user.Username,
			BlurtID:  blurtID,
		})

	if _, err := insert.Exec(v.db); err != nil {
		return err

	}

	return nil
}
