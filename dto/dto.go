package dto

type Response struct {
	Name 	string `json:"Имя"`
	Slug 	string `json:"Транскрипция"`
	Preview string `json:"Превью"`
	Video 	string `json:"Видео"`
}
