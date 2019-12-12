package models

// Verifizierte E-Mail ist Pflicht für 'Minuskonten'
// Die Pin des Nutzers.
type User struct {
	Model
	Name      string    `json:"name,omitempty" gorm:"unique"`
	ImageUrl  string    `json:"imageUrl,omitempty"`
	Email     string    `json:"email,omitempty"`
	Verified  bool      `json:"verified,omitempty"`
	Gravatar  bool      `json:"gravatar,omitempty"`
	Balance   int32     `json:"balance,omitempty"`
	Pin       string    `json:"pin,omitempty"`
}
