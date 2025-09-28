package database

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Predefined list of people
var Peoples = []Person{
	{ID: 1, Name: "Liam", Age: 24, Email: "liam@example.com", Password: "pass123"},
	{ID: 2, Name: "Olivia", Age: 29, Email: "olivia@example.com", Password: "pass234"},
	{ID: 3, Name: "Noah", Age: 31, Email: "noah@example.com", Password: "pass345"},
	{ID: 4, Name: "Emma", Age: 27, Email: "emma@example.com", Password: "pass456"},
	{ID: 5, Name: "Oliver", Age: 35, Email: "oliver@example.com", Password: "pass567"},
	{ID: 6, Name: "Ava", Age: 28, Email: "ava@example.com", Password: "pass678"},
	{ID: 7, Name: "Elijah", Age: 33, Email: "elijah@example.com", Password: "pass789"},
	{ID: 8, Name: "Sophia", Age: 26, Email: "sophia@example.com", Password: "pass890"},
	{ID: 9, Name: "William", Age: 30, Email: "william@example.com", Password: "pass901"},
	{ID: 10, Name: "Isabella", Age: 25, Email: "isabella@example.com", Password: "pass012"},
}


var People = []Person{}

func HandleInit() {
	People = Peoples
}
