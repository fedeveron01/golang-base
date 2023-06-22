package handler_person

import (
	"fmt"
	"net/http"

	"context"

	"github.com/fedeveron01/golang-base/cmd/core/usecases/calculate_age"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
)

//get all

type PersonGetAllHandler struct {
	// use cases
	CalculateAge calculate_age.CalculateAgeUseCase
}

func NewPersonGetAllHandler(calculateAge calculate_age.CalculateAgeUseCase) PersonGetAllHandler {
	return PersonGetAllHandler{
		CalculateAge: calculateAge,
	}
}

func SendMessage(client *whatsmeow.Client, receptor types.JID, message string) {
	fmt.Println("Sending message")
	client.SendMessage(context.Background(), receptor, &waProto.Message{
		Conversation: proto.String(message),
	})
}

func (p PersonGetAllHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Person handler")
	//p.CalculateAge.CalculateAge()

}
