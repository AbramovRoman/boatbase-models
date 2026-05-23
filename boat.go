package models

type Boat struct {
	ID            int    `json:"id"`
	StateNumber   string `json:"state_number"`    //гос номер лодки
	TrailerNumber string `json:"trailer_number"`  //номер прицепа
	OwnerFullName string `json:"owner_full_name"` //ФИО владельца
}
