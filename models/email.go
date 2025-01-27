package models

type EmailReservation struct {
	MailTo   string
	Code     string
	DateFrom string
	DateTo   string
	Total    float64
}
