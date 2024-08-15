package types

type Book struct {
	Name             string   `json:"book"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	Categories       []string `json:"categories"`
	CoverImage       string   `json:"cover_image"`
	AuthorName       string   `json:"author_name"`
	Publication      string   `json:"publication"`
	PageCount        int32    `json:"page_count"`
	Rate             int8     `json:"rate"`
	TotalRate        int      `json:"total_rate"`
	URL              string   `json:"url"`
}
