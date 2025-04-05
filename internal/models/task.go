package models

type Task struct {
	ID     string `json:"id"`
	ListID string `json:"list_id"`
	Title  string `json:"title"`
	Emoji  string `json:"emoji"`
}

var Tasks = []Task{
	{ID: "1", ListID: "1", Title: "100 отжиманий", Emoji: "💪"},
}
