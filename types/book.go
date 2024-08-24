package types

type Categories struct{}
type Book struct {
	Title            string   `bson:"title"`
	Author           string   `bson:"author"`
	Translator       string   `bson:"translator"`
	Publication      string   `bson:"publication"`
	Categories       []string `bson:"categories"`
	Rate             string   `bson:"rate"`
	TotalRate        string   `bson:"total_rate"`
	CoverImage       string   `bson:"cover_image"`
	Description      string   `bson:"description"`
	ShortDescription string   `json:"short_description"`
	URL              string   `json:"url"`
	PublishDate      string   `json:"publish_date"`
	BookID           string   `json:"book_id"`
}
