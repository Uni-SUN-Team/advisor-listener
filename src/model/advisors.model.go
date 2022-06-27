package model

type advisor struct {
	Data advisorData `json:"data"`
	Meta pagination  `json:"meta"`
}

type advisors struct {
	Data []advisorData `json:"data"`
	Meta pagination    `json:"meta"`
}

type advisorData struct {
	Id            int64            `json:"id"`
	FullName      string           `json:"full_name"`
	Email         string           `json:"email"`
	Telephone     string           `json:"telephone"`
	Dob           string           `json:"dob"`
	Content       string           `json:"content"`
	Short_content string           `json:"short_content"`
	Active        bool             `json:"active"`
	CreatedAt     string           `json:"createdAt"`
	UpdatedAt     string           `json:"updatedAt"`
	PublishedAt   string           `json:"publishedAt"`
	Locale        string           `json:"locale"`
	Slug          string           `json:"slug"`
	JobTitle      string           `json:"job_title"`
	Quip          string           `json:"quip"`
	Thumnail      thumbnailLarge   `json:"thumbnail"`
	Categories    []categories     `json:"categories"`
	Attachment    []thumbnailLarge `json:"attachment"`
	SEO           seo              `json:"SEO"`
}

type pagination struct {
	Pagination paginationContent `json:"pagination"`
}

type paginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
