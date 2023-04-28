package main

type Trick struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Description    string `json:"description"`
	Difficulty     string `json:"difficulty"`
	Progress       string `json:"progress"`
}

func createTable() {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS tricks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
		translatedName TEXT,
        description TEXT,
		difficulty TEXT,
		progress TEXT
    )`)
	if err != nil {
		panic(err)
	}
}
