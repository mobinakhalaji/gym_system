package model

type Member struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Phone  *string `json:"phone"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
}
type Trainer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	Phone     string `json:"phone"`
}
type Class struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TrainerID   int    `json:"trainer_id"`
	TrainerName string `json:"trainer_name,omitempty"`
	Capacity    int    `json:"capacity"`
	Price       int    `json:"price"`
}
type Register struct {
	ID       int `json:"id"`
	MemberID int `json:"member_id"`
	ClassID  int `json:"class_id"`
}
