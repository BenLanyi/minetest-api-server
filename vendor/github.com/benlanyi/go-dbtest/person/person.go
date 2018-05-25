package person

// Person : represents a person with name and age
type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
	ID  int    `db:"id"`
}
