package database


 type Person struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
 }



 // Predefined list of people
	var Peoples = []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 28},
		{Name: "Ethan", Age: 32},
		{Name: "Fiona", Age: 27},
		{Name: "George", Age: 29},
		{Name: "Hannah", Age: 31},
		{Name: "Ian", Age: 26},
		{Name: "Jane", Age: 33},
	}

  var People = []Person{}

  func HandleInit(){
	People = Peoples
  }