package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Trick struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Progress    string `json:"progress"`
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./tricks.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable()
	router := mux.NewRouter()
	router.HandleFunc("/tricks/{id}", handleTricksPut).Methods("PUT")
	router.HandleFunc("/tricks", handleTricks).Methods("GET", "POST")
	http.ListenAndServe(":8080", router)
}

func createTable() {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS tricks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        description TEXT,
		difficulty TEXT,
		progress TEXT
    )`)
	if err != nil {
		panic(err)
	}
}

func handleTricks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleTricksGet(w, r)
	case http.MethodPost:
		handleTricksPost(w, r)
	case http.MethodPut:
		handleTricksPut(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func handleTricksGet(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM tricks")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tricks := make([]Trick, 0)
	for rows.Next() {
		var t Trick
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Difficulty, &t.Progress); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		tricks = append(tricks, t)
	}

	if err := json.NewEncoder(w).Encode(tricks); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func handleTricksPost(w http.ResponseWriter, r *http.Request) {
	var t Trick
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO tricks (name,description,difficulty,progress) VALUES (?,?,?,?)", t.Name, t.Description, t.Difficulty, t.Progress)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	t.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(t); err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func handleTricksPut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid Trick ID", http.StatusBadRequest)
		return
	}

	var t Trick
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t.ID = id

	dbreq, err := db.Exec("UPDATE tricks SET name='" + t.Name + "', description='" + t.Description + "', difficulty='" + t.Difficulty + "', progress='" + t.Progress + "' WHERE id=" + strconv.Itoa(id))

	fmt.Println(dbreq)

	fmt.Println("SQL Statement: ", "UPDATE tricks SET name='"+t.Name+"', description='"+t.Description+"', difficulty='"+t.Difficulty+"', progress='"+t.Progress+"' WHERE id="+strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
