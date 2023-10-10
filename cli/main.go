package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Por favor, informe pelo menos um número de telefone!")
		return
	}

	phoneNumbers := os.Args[1:]
	whatsappLinks := make(chan string)

	for _, phoneNumber := range phoneNumbers {
		go func(phoneNumber string) {
			message := "Olá, segue o link da promoção!"
			whatsappLinks <- generateWhatsappLink(phoneNumber, message)

			whatsappLinks <- <-whatsappLinks
		}(phoneNumber)
	}

	for range phoneNumbers {
		fmt.Println(<-whatsappLinks)
	}
}

func generateWhatsappLink(phoneNumber, message string) string {
	encPhoneNumner := url.QueryEscape(phoneNumber)
	encMessage := url.QueryEscape(message)

	return fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", encPhoneNumner, encMessage)
}
