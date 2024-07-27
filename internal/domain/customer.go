package domain

type Customer struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Balance  float64  `json:"balance"`
	Orders   []*Order `json:"orders"`
}
