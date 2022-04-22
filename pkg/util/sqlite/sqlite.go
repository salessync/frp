package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/salessync/frp/pkg/util/log"
)

type PastRequest struct {
	Id            int    `json:"id"`
	Original_host string `json:"original_host"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Query_params  string `json:"query_params"`
	Body          string `json:"body"`
	Headers       string `json:"headers"`
	Created       string `json:"created"`
}

type PastRequests struct {
	Data []PastRequest `json:"data"`
}

func InitDB() {
	db := getDB()
	createRequestTable(db)
}

func RunQuery(query string) bool {
	db := getDB()

	_, err := db.Exec(query)
	if err != nil {
		log.Warn(err.Error())
		return false
	}

	return true
}

func FetchRequests() PastRequests {
	db := getDB()

	rows, err := db.Query("SELECT * FROM request")
	if err != nil {
		log.Info(err.Error())
	}
	defer rows.Close()

	allPastRequests := PastRequests{}

	for rows.Next() {
		var id int
		var original_host string
		var url string
		var method string
		var query_params string
		var body string
		var headers string
		var created string

		err := rows.Scan(&id, &original_host, &url, &method, &query_params, &body, &headers, &created)
		if err != nil {
			log.Warn(err.Error())
		}

		pastReq := PastRequest{
			Id:            id,
			Original_host: original_host,
			Url:           url,
			Method:        method,
			Query_params:  query_params,
			Body:          body,
			Headers:       headers,
			Created:       created,
		}

		allPastRequests.Data = append(allPastRequests.Data, pastReq)
	}

	return allPastRequests
}

func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./frps.db")
	if err != nil {
		log.Warn(err.Error())
	}

	return db
}

func createRequestTable(db *sql.DB) {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS request (id integer not null primary key, original_host text, url text, method text, query_params text, body text, headers text, created TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Warn("%q: %s\n", err, sqlStmt)
	}
}
