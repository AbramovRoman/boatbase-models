package models

type LoginRequest struct { //Вход на сайт
	StateNumber   string `json:"state_number"`   //госномер
	TrailerNumber string `json:"trailer_number"` //номер прицепа
}

type Session struct {
	Token         string `json:"token"`           //Уникальный токен
	StateNumber   string `json:"state_number"`    //госномер
	OwnerFullName string `json:"owner_full_name"` //Имя
	Role          string `json:"role"`
}
