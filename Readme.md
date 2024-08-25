# Persian Book Crawler 📚🕵️‍♂️
Persian Book Crawler is a Go-based tool designed to scrape book information from taaghche.com, a popular Persian book website. It extracts and saves book details into a MongoDB database. Perfect for building a collection of Persian literature!

### Features 🌟
- Targeted Crawling: Specifically built to navigate and scrape Persian book details from Taaghche.
- MongoDB Integration: Stores data with a unique index on BookID to avoid duplicates.
Configurable Depth and Threads: Set to crawl up to 2 levels deep and supports up to 3 concurrent threads for efficient data collection.
- User Agent Randomization: Uses random user agents to help avoid detection and blocking by the target website.
- Referer Header Setting: Includes a referer header to mimic real browser behavior during crawling.

### Installation 🚀
Clone the repository:

~~~bash
git clone https://github.com/yourusername/book_crawler.git
cd book_crawler
~~~
Install dependencies: Ensure Go is installed, then run:

~~~bash
go mod tidy
~~~
Set up MongoDB: Make sure MongoDB is running and update your .env file with the connection details.

Create a `.env` file: Add the following content to configure your MongoDB connection:

~~~bash
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=your_database_name
COLLECTION_NAME=books
~~~
Run the Crawler: Start the crawler with:


`go run main.go`
### Usage 📖
The crawler visits taaghche.com, extracts book information, and saves it to your MongoDB database.
It checks if a book is already in the database using BookID and only inserts new entries.
### Contributing 🤝
Contributions are welcome! If you have ideas or improvements:

- Fork the repository
- Create a new branch 
- Commit your changes 
- Push to the branch 
- Open a pull request

