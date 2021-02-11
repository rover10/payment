package model

type Payment struct{}
type PaymentHistory struct {
	ID          int64          `json:"id"`
	UTR         string         `json:"utr"`
	Amount      float64        `json:"amount"`
	PaymentTime string         `json:"paymentTime"`
	Status      string         `json:"status"`
	Payment     string         `json:"payment"`
	FromBank    UserBankDetail `json:"fromBank"`
	ToBank      UserBankDetail `json:"toBank"`
}
