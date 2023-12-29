package domain

type HttpResponse struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
