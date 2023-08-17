package controller

type AuthModels struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

var BASE_URL = "https://sandbox.partner.api.bri.co.id/oauth/client_credential/accesstoken"
