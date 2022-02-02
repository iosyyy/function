package iosyyy

type Person struct {
	sex bool
}

type Student struct {
	Person
	id       int
	name     string
	birthday int64
}

func (per Person) getSex() (sex bool) {
	sex = per.sex
	sex = true
	return
}

func (per *Person) setSex(sex bool) {
	per.sex = sex
}

func testStudent(student *Student) {
	student.id++
}

func test() {
	/*student := Student{Person{sex: false}, 1, "test", 1000}

	setSex := (*Student).setSex
	setSex(&student, true)
	fmt.Println(student)*/
	//setSex := student.setSex
	//setSex(true)
	//fmt.Println(student)
	/*fmt.Printf("student: %+v\n", student)
	testStudent(&student)
	fmt.Printf("student: %+v\n", student)*/

	/*p1 := &student
	println(p1)
	*/
}
