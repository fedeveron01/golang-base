package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type WhatsappRepository struct {
	client *whatsmeow.Client
}

func NewWhatsappRepository(client *whatsmeow.Client) *WhatsappRepository {
	connectClient(client)

	return &WhatsappRepository{client: client}
}

func (r *WhatsappRepository) SendText(to string, text string) error {

	fmt.Println("Sending message")
	recipient := types.JID{
		User:   to,
		Server: types.DefaultUserServer,
	}
	res, err := r.client.SendMessage(context.Background(), recipient, &waProto.Message{
		Conversation: proto.String(text),
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)

	return nil
}

func connectClient(client *whatsmeow.Client) {
	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)

			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}

}
