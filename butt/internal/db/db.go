package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jeffreycharters/blurt/ent"
	"github.com/jeffreycharters/blurt/ent/blurt"
	_ "github.com/mattn/go-sqlite3"
)

var client *ent.Client

func Client() *ent.Client {
	if client != nil {
		return client
	}

	client, err := ent.Open("sqlite3", "file:/home/jeffrey/blurtDB/blurts.db?mode=rw&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func addBlurt(content string, userID uuid.UUID) (*ent.Blurt, error) {
	newBlurt, err := Client().Blurt.Create().
		SetContent(content).
		SetAuthorID(userID).
		Save(context.Background())

	if err != nil {
		return nil, errors.New("failed to add blurt")
	}

	return Client().Blurt.Query().WithAuthor().WithLiks().Where(blurt.ID(newBlurt.ID)).Only(context.Background())
}

type MessageType string

const (
	Blurt MessageType = "blurt"
	Lik   MessageType = "lik"
)

type incomingMessage struct {
	MessageType MessageType            `json:"type"`
	Payload     map[string]interface{} `json:"payload"`
}

type BlurtPayload struct {
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userID"`
}

type DBResponse struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

func ParseMessage(message []byte) ([]byte, error) {
	var msg = incomingMessage{}
	err := json.Unmarshal([]byte(message), &msg)

	if err != nil {
		fmt.Println("error unmarshaling")
		return nil, err
	}

	var payload interface{}

	switch msg.MessageType {
	case Blurt:
		userID, err := uuid.Parse(msg.Payload["userID"].(string))
		if err != nil {
			return nil, errors.New("invalid blurt userID")
		}

		blurt, err := addBlurt(msg.Payload["content"].(string), userID)
		if err != nil {
			return nil, errors.New("failed to add blurt")
		}

		payload = map[string]interface{}{
			"blurt": blurt,
		}

	case Lik:
		userID, err := uuid.Parse(msg.Payload["userID"].(string))
		if err != nil {
			return nil, errors.New("invalid lik userID")
		}
		blurtID, err := uuid.Parse(msg.Payload["blurtID"].(string))
		if err != nil {
			return nil, errors.New("invalid lik blurtID")
		}
		lik, err := Client().Lik.Create().
			SetUserID(userID).
			SetBlurtID(blurtID).
			Save(context.Background())

		if err != nil {
			return nil, errors.New("failed to add like")
		}

		payload = map[string]interface{}{
			"lik": lik,
		}

	}

	return json.Marshal(&DBResponse{Success: true, Payload: payload})

}

type wrappedError struct {
	Success bool
	Error   string
}

func WrapError(err error) []byte {
	ret, _ := json.Marshal(wrappedError{Success: false, Error: err.Error()})
	return ret
}
