package helpers

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

type ResponsePaginate struct {
	Response
	Meta *MetaPaginate `json:"meta,omitempty"`
}

type MetaPaginate struct {
	Total       int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	Limit       int64 `json:"limit"`
	Pages       int64 `json:"pages"`
}
