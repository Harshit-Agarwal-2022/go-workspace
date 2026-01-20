package main

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	cf := NewCmdFlags()
	cf.Execute(&todos)
	storage.save(todos)
}
