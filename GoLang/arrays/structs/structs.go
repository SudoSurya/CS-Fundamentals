package structs

import "fmt"

type Employee struct {
	Id       int
	Name     string
	Address  string
	Salary   int
	Position string
}

func main() {
}
func Swapable() {

	var dilbert Employee
	dilbert.Id = 1
	dilbert.Name = "Dilbert"
	dilbert.Address = "Bangalore"
	dilbert.Salary = 10000
	dilbert.Position = "Software Engineer"
	fmt.Println(dilbert)
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(dilbert)
}
