package handler_whatsapp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fedeveron01/golang-base/cmd/repositories"
)

//get all

type WhatsappHandler struct {
	// use cases
	WhatsappRepository repositories.WhatsappRepository
}

func NewWhatsappHandlerHandler(whatsappRepository repositories.WhatsappRepository) WhatsappHandler {
	return WhatsappHandler{
		WhatsappRepository: whatsappRepository,
	}
}

func (wh WhatsappHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting whatsapp server")
	fmt.Fprint(w, "whatsapp server")
	c := make(chan os.Signal)
	wh.WhatsappRepository.ConnectClient(c)
	//p.CalculateAge.CalculateAge()

}
