package main

import "fmt"

type instructor interface {
	teach()
}

type Teacher struct {
	name    string
	subject string
}

type Professor struct {
	name    string
	subject string
}

func (p Professor) teach() {
	fmt.Println("Teaching " + p.subject)
}

func (t Teacher) teach() {
	fmt.Println("Teaching " + t.subject)
}

func find(t instructor) {
	switch t.(type) {
	case Professor:
		fmt.Println("Professor Executed "+t.(Professor).subject)
		break
	case Teacher:
		fmt.Println("Teacher executed "+t.(Teacher).subject)
		break
	}
}

func main() {
	t := Teacher{"Bapan", "Physics"}
	p := Professor{"H", "Magnet"}
	faculty := []instructor{t, p}
	for _, t := range faculty {
		t.teach()
	}
	fmt.Println("Displaying tutor finder")
	find(t)
	find(p)
}
