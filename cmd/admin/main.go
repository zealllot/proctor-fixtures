// admin — minimal QOR admin panel used as a PRoctor end-to-end fixture.
//
// Two resources (User, Article) registered in configureAdmin(); mounted
// at /admin against a SQLite DB; healthz probe at /healthz so PRoctor's
// setup wait loop has something to ping.
package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
)

// User represents a person who can sign in to the admin.
type User struct {
	gorm.Model
	Name  string
	Email string
	Role  string
}

// Article represents a post in the admin panel.
type Article struct {
	gorm.Model
	Title  string
	Body   string
	Status string
	Author string
}

// configureAdmin registers all admin resources with their UI metadata.
// PRs that exercise PRoctor on this fixture should mostly modify this
// function and let the test plan focus on its visible effects.
func configureAdmin(db *gorm.DB) *admin.Admin {
	a := admin.New(&admin.AdminConfig{DB: db})

	user := a.AddResource(&User{})
	user.SearchAttrs("Name", "Email")
	user.IndexAttrs("ID", "Name", "Email", "Role", "CreatedAt")

	article := a.AddResource(&Article{})
	article.IndexAttrs("ID", "Title", "Status", "Author", "CreatedAt")
	article.EditAttrs("Title", "Body", "Status", "Author")

	return a
}

func main() {
	db, err := gorm.Open("sqlite3", "qor-fixture.db")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Article{})

	a := configureAdmin(db)

	mux := http.NewServeMux()
	a.MountTo("/admin", mux)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	const addr = "127.0.0.1:7000"
	log.Printf("admin listening on http://%s/admin", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
