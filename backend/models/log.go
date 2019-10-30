package models

type Log struct {
	Model
	User    string    `json:"user,omitempty"`
	Product string    `json:"product,omitempty"`
	Action  string    `json:"action,omitempty"`
	Amount  int32     `json:"amount,omitempty"`
}
