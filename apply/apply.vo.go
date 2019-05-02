package apply

// Applicant struct
type Applicant struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Number     int    `json:"number"`
	Vegetarian int    `json:"vegetarian"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Line       string `json:"line"`
	Relation   string `json:"relation"`
	Card       bool   `json:"card"`
	Address    string `json:"address"`
	Other      string `json:"other"`
}
