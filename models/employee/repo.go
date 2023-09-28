package employee

type Employee struct {
	Id        int
	LastName  string `json:"last_name"`  // Change to match the SQL column name
	FirstName string `json:"first_name"` // Change to match the SQL column name
	Address   string `json:"address"`
	City      string `json:"city"`
}
