package boatbase_models

type Boat struct {
	ID            int    `json:"id"`
	StateNumber   int    `json:"state_number"`    //гос номер лодки
	TrailerNumber int    `json:"trailer_number"`  //номер прицепа
	OwnerFullName string `json:"owner_full_name"` //ФИО владельца
}
