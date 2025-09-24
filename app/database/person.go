package database


 type Person struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
	ID   int    `json:"id"`
 }



 // Predefined list of people
	var Peoples = []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
		{ID: 4, Name: "Diana", Age: 28},
		{ID: 5, Name: "Ethan", Age: 32},
		{ID: 6, Name: "Fiona", Age: 27},
		{ID: 7, Name: "George", Age: 29},
		{ID: 8, Name: "Hannah", Age: 31},
		{ID: 9, Name: "Ian", Age: 26},
		{ID: 10, Name: "Jane", Age: 33},
	}

  var People = []Person{}

  func HandleInit(){
	People = Peoples
  }