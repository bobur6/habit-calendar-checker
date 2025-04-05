package models

type List struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Lists = []List{
	{ID: "1", Title: "Спорт", Description: "Тренировки"},
}
