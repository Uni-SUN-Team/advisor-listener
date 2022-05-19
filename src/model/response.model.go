package model

type ResponseAdvisorsAll struct {
	Error      error             `json:"error"`
	Result     []AdvisorsData    `json:"result"`
	Pagination PaginationContent `json:"pagination"`
}

type ResponseAdvisor struct {
	Error  error       `json:"error"`
	Result AdvisorData `json:"result"`
}
