package pointers

type Person struct {
	Name string
	Age  int
}

var p *Person

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func ChangeName(person *Person, newName string) *Person {
	person.Name = newName
	return person
}

func Exec() {
	p = NewPerson("Bruno", 27)
	ChangeName(p, "Brabo")
	print(p.Name)
}