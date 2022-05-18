package model

type Advisor struct {
	Data AdvisorData `json:"data"`
	Meta Pagination  `json:"meta"`
}

type AdvisorsSlug struct {
	Data []AdvisorData `json:"data"`
	Meta Pagination    `json:"meta"`
}

type Advisors struct {
	Data []AdvisorsData `json:"data"`
	Meta Pagination     `json:"meta"`
}

type AdvisorsData struct {
	Id            int64    `json:"id"`
	FullName      string   `json:"full_name"`
	Email         string   `json:"email"`
	Telephone     string   `json:"telephone"`
	Dob           string   `json:"dob"`
	Content       string   `json:"content"`
	Short_content string   `json:"short_content"`
	Active        bool     `json:"active"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
	PublishedAt   string   `json:"publishedAt"`
	Locale        string   `json:"locale"`
	Slug          string   `json:"slug"`
	JobTitle      string   `json:"job_title"`
	Quip          string   `json:"quip"`
	Thumnail      Thumnail `json:"thumnail"`
}

type AdvisorData struct {
	Id            int64        `json:"id"`
	FullName      string       `json:"full_name"`
	Email         string       `json:"email"`
	Telephone     string       `json:"telephone"`
	Dob           string       `json:"dob"`
	Content       string       `json:"content"`
	Short_content string       `json:"short_content"`
	Active        bool         `json:"active"`
	CreatedAt     string       `json:"createdAt"`
	UpdatedAt     string       `json:"updatedAt"`
	PublishedAt   string       `json:"publishedAt"`
	Locale        string       `json:"locale"`
	Slug          string       `json:"slug"`
	JobTitle      string       `json:"job_title"`
	Quip          string       `json:"quip"`
	Thumnail      Thumnail     `json:"thumnail"`
	Categories    []Categories `json:"categories"`
	Attachment    []Thumnail   `json:"attachment"`
}

type Pagination struct {
	Pagination PaginationContent `json:"pagination"`
}

type PaginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
