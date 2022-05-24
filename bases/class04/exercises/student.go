package exercises

type Student struct {
	Name      string
	Surname   string
	Document  string
	CreatedAt uint16
}

func Build(name, surname, document string, createdAt uint16) Student {
	return Student{name, surname, document, createdAt}
}
