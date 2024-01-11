package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ResquestBody struct {
	PhoneNumbers []string `json:"phone_numbers"`
	Message      string   `json:"message"`
}

type ResponseBody struct {
	WhatsappLinks []string `json:"whatsapp_links"`
}

func GenerateLinksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody ResquestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	whatsappLinks := make([]string, len(requestBody.PhoneNumbers))
	for i, phoneNumber := range requestBody.PhoneNumbers {
		whatsappLinks[i] = generateWhatsappLinks(phoneNumber, requestBody.Message)
	}

	responseBody := ResponseBody{WhatsappLinks: whatsappLinks}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateWhatsappLinks(phoneNumber, message string) string {
	encPhoneNumber := url.QueryEscape(phoneNumber)
	encMessage := url.QueryEscape(message)

	return fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", encPhoneNumber, encMessage)
}
