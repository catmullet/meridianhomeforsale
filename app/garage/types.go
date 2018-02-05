package garage

type Validate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ValidateResponse struct {
	SecurityToken string `json:"SecurityToken"`
	ReturnCode string `json:"ReturnCode"`
	ErrorMessage string `json:"ErrorMessage"`
	CorrelationId string `json:"CorrelationId"`
}