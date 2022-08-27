package entities

type News struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Keyword string `json:"keyword"`
	Source  string `json:"source"`
}
