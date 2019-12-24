package cms

import (
	"database/sql"

	// Use the Postgres SQL driver
	_ "github.com/lib/pq"
)

var (
	store = newDB()
)

// PgStore is a reference to the password store.
type PgStore struct {
	DB *sql.DB
}

// CreatePage creates a new page in our DB
func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}

// CreatePost creates a new post in our DB
func CreatePost(p *Post) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO posts(title, content, date_created) VALUES($1, $2, $3) RETURNING id", p.Title, p.Content, p.DatePublished).Scan(&id)
	return id, err
}

// GetPage gets the page by it's ID.
func GetPage(id string) (*Page, error) {
	var p Page
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content)
	return &p, err
}

// GetPages returns every page from our DB.
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=goprojects host=127.0.0.1 port=32768 dbname=goprojects sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}
