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
	ShortDescription string   `bson:"short_description"`
	URL              string   `bson:"url"`
	PublishDate      string   `bson:"publish_date"`
	BookID           string   `bson:"book_id"`
}
