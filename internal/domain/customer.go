package domain

type Customer struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Orders   []*Order `json:"orders"`
}
