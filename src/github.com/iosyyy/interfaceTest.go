package iosyyy

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Tester interface {
	sayHi()
}

type Humaner interface {
	sayHello()
	Tester
}

type Persons struct {
	word string
}

func (per Persons) sayHi() {
	fmt.Println("hi ", per.word)

}

func (per Persons) sayHello() {
	fmt.Println("hello ", per.word)
}

func testErr(a, b int) (c int, err error) {
	if b == 0 {
		err = errors.New("err is happened")
	} else {
		c = a / b
	}
	return
}

func testRecover() (err error) {
	defer func() {
		if rev := recover(); rev != nil {
			err = fmt.Errorf("%s", rev)
		}
	}()
	panic("test panic")
}

func init() {
	float := strconv.FormatFloat(1.65, 'f', -1, 64)
	println(float)
	split := strings.Split("hello,,,", ",")
	for _, value := range split {
		if value != "" {
			fmt.Println(value)
		}
	}
	/*	err := testRecover()
		if err != nil {
			fmt.Println(err)
		}*/
	/*value, err := testErr(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}*/
	/*any := []interface{}{1, "", Persons{word: "test"}}

	for _, value := range any {
		switch val := value.(type) {
		case Persons:
			fmt.Println(val)
		}

		if val, ok := value.(int); ok {
			fmt.Println("int: ", val)
		}
	}
	*/
	// hello := Humaner.sayHello
	// persons := Persons{word: "world"}

	// hello(persons)
	// var inter Humaner = persons

	// var test Tester = inter
	// inter.sayHello()
	// inter.sayHi()
}
