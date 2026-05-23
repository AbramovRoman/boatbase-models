package models

type LaunchRequest struct { //Создание заявки на спуск
	ReturnTime    string `json:"return_time"`    //Время возвращения
	ReturnDate    string `json:"return_date"`    //Дата возвращения
	SwimmingArea  string `json:"swimming_area"`  //Район плавания
	TrailerNumber string `json:"trailer_number"` //Номер прицепа
	StateNumber   string `json:"state_number"`   //Госномер лодки
}

type QueueStatus string //Статус очереди для админки

const (
	StatusWaiting  QueueStatus = "Ожидает спуска"
	StatusOnWater  QueueStatus = "На воде"
	StatusReturned QueueStatus = "Ожидает подьем"
	StatusOnBase   QueueStatus = "На базе"
)

type LiftRequest struct { //Запрос на подьем
	StateNumber string `json:"state_number"`
}

type QueueItem struct { //Для админки
	ID            int         `json:"id"`
	StateNumber   string      `json:"state_number"`   //Госномер
	TrailerNumber string      `json:"trailer_number"` //Номер прицепа
	ReturnTime    string      `json:"return_time"`    //Время возвращения
	ReturnDate    string      `json:"return_date"`    //Дата возвращения
	SwimmingArea  string      `json:"swimming_area"`  //Район плавания
	Status        QueueStatus `json:"status"`         //Текущий статус лодки
	CreatedAt     string      `json:"created_at"`     //Время создания заявки
}
