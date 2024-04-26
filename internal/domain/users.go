package domain

type Users struct {
	ID       uint   `storm:"id,increment" json:"id"`
	Username string `storm:"unique" json:"username"`
	Password string `json:"_"`
}
