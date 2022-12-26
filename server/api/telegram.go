package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/cmd/server/telegram"
	"server/logger"
)

type SendMessageReq struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func SendTeleMessage(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &SendMessageReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.L.Err(err).Msg("Failed to encode message")
		return http.StatusBadRequest, err
	}

	message := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", req.Name, req.Email, req.Message)

	err = telegram.T.SendMessage(message)
	if err != nil {
		logger.L.Err(err).Msg("Failed to send message")
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
