package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type student struct {
	roll int
	name string
	marks map[string]int
	age int
	next *student
}

func recoverFun() {
	r := recover()
	if r != nil {
		fmt.Println(r)
	}
}

func setName(st *student,roll int) {
	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	inputReader := bufio.NewReader(os.Stdin)
	for st != nil {

		if(st.roll == roll) {
			fmt.Println("Enter name to update")
			name,_ := inputReader.ReadString('\n')
			name = strings.TrimSpace(name)
			st.name = name
			return
		}
		st = st.next
	}
	if(st == nil) {
		fmt.Println("Enter valid roll number")
	}
}

func setAge(st *student,roll int) {

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}
	var age int = 0
	for st != nil {
		if(st.roll == roll) {
			fmt.Println("Enter age to update")
			fmt.Scanln(&age)
			st.age = age
			return
		}
		st = st.next
	}
	if(st == nil) {
		fmt.Println("Enter valid roll number")
	}
}

func setMarks(st *student,roll int) {

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	inputReader := bufio.NewReader(os.Stdin)
	var mark int = 0
	for st != nil {
		if(st.roll == roll) {
			fmt.Println("Enter name of subject to update mark")
			sub,err := inputReader.ReadString('\n')

			if err != nil {
				fmt.Println(err)
			}

			sub = strings.TrimSpace(sub)
			fmt.Println("Enter marks of subject ")
			fmt.Scanln(&mark)
			for mark < 0 || mark > 101 {
				fmt.Println("Enter valid marks")
				fmt.Scanln(&mark)
			}
			st.marks[sub] = mark
			return
		}
		st = st.next
	}
	if(st == nil) {
		fmt.Println("Enter valid roll number")
	}
}

func getName(st *student,roll int) string{

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	for st != nil {
		if(st.roll == roll) {
			break
		}
		st = st.next
	}
	if st == nil {
		return " "
	}
	
	return st.name
}

func getAge(st *student,roll int) int{

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	for st != nil {
		if(st.roll == roll) {
			break
		}
		st = st.next
	}
	if(st == nil) {
		return -1
	}
	return st.age
}

func getMarks(st *student,roll int) map[string]int{

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	for st != nil {
		if(st.roll == roll) {
			break
		}
		st = st.next
	}
	if(st == nil) {
		return nil
	}
	return st.marks
}

func percentage(st *student,roll int) (per float32){

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	var total int = 0

	for st != nil {
		if(st.roll == roll) {

			for _,mark := range(st.marks) {
				total = total + mark
			}

			per = float32(total) / float32(len(st.marks))
			return
		}
		st = st.next
	}
	if(st == nil) {

		fmt.Println("Enter valid roll number")
		return -1

	} else {
		return 
	}
}

func display(s *student) {
	
	defer recoverFun()

	if(s == nil) {
		panic("Error Invalid student")
	}

	fmt.Println("--------------------------------------")
	for s != nil {
		fmt.Printf("Roll Number ->%v\n",s.roll)
		fmt.Printf("Name ->%v\n",s.name)
		fmt.Printf("Marks\t")

		for sub,mark := range(s.marks) {
			fmt.Printf("%v->%v\n",sub,mark)
		}

		fmt.Printf("Age->%v\n",s.age)
		fmt.Println()
		s = s.next
	}
}

func deleteDetail(st **student,roll int) *student{
	
	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	var ret *student = *st
	if((*st).roll == roll) {
		(*st) = (*st).next
		ret = *st

	} else {

		var temp **student = st
		ret = *st
		for (*temp).next != nil {
			if((*temp).next.roll == roll) {

				(*temp).next = (*temp).next.next
				break
			}
			*temp = (*temp).next
		}
		if((*temp).roll == roll) {
			*temp = nil
		}
	}
	return ret
}

func insertDetails(st **student) *student{

	defer recoverFun()

	if(st == nil) {
		panic("Error Invalid student")
	}

	var roll,age,mark,ch int
	var marks = make(map[string]int)
	var ret *student = *st
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter roll number")
	fmt.Scanln(&roll)
	fmt.Println("Enter name of student")
	name,err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	name = strings.TrimSpace(name)

	for(true) {

		fmt.Println("Enter subject name")
		sub,err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		sub = strings.TrimSpace(sub)
		fmt.Println("Enter marks ")
		fmt.Scanln(&mark)

		for mark < 0 || mark > 101 {
			fmt.Println("Enter valid marks")
			fmt.Scanln(&mark)
		}

		marks[sub] = mark
		fmt.Println("Do you want to add another subject(1/0)")
		fmt.Scanln(&ch) 
		if(ch == 0) {
			break
		}
	}

	fmt.Println("Enter age of student")
	fmt.Scanln(&age)

	if(*st == nil) {

		*st = new(student)
		(*st).roll = roll
		(*st).name = name
		(*st).marks = marks
		(*st).age = age
		(*st).next = nil
		ret = *st

	} else {

		var temp **student = st
		ret = *st

		if((*st).roll == roll) {
			fmt.Println("Roll number already exist")
			return ret
		}

		for((*temp).next != nil){
			if((*temp).roll == roll) {
				break
			}
			*temp = (*temp).next
		}

		if((*temp).roll == roll) {
			fmt.Println("Roll number already exist")
			return ret
		}

		(*temp).next = new(student)
		(*temp).next.roll = roll
		(*temp).next.name = name
		(*temp).next.marks = marks
		(*temp).next.age = age
		(*temp).next.next = nil
	}
	return ret
}

func main() {

	var ch int = 0
	var st *student = nil

	for ch != 11 {

		fmt.Println("--------------------------------------")
		fmt.Println("1.Insert details of student")
		fmt.Println("2.Delete details of student")
		fmt.Println("3.Update name of student")
		fmt.Println("4.Update marks of student")
		fmt.Println("5.Update age of student")
		fmt.Println("6.Calculate percentage of student")
		fmt.Println("7.Display name of student")
		fmt.Println("8.Display marks of student")
		fmt.Println("9.Display age of student")
		fmt.Println("10.Display details of student")
		fmt.Println("11.exit")
		fmt.Println("--------------------------------------")

		fmt.Scanln(&ch)

		switch ch {
		case 1:

			st = insertDetails(&st)
		case 2:

			var roll int = 0
			fmt.Println("Enter roll number to delete details")
			fmt.Scanln(&roll)
			st = deleteDetail(&st,roll)
		case 3:

			var roll int = 0
			fmt.Println("Enter roll number to update name")
			fmt.Scanln(&roll)
			setName(st,roll)
		case 4:

			var roll int = 0
			fmt.Println("Enter roll number to update marks")
			fmt.Scanln(&roll)
			setMarks(st,roll)
		case 5:

			var roll int = 0
			fmt.Println("Enter roll number to update age")
			fmt.Scanln(&roll)
			setAge(st,roll)
		case 6:

			var roll int = 0
			fmt.Println("Enter roll number to Calculate percentage")
			fmt.Scanln(&roll)
			per := percentage(st,roll)
			if per != -1{
				fmt.Println(per)
			}
		case 7:

			var roll int = 0
			fmt.Println("Enter roll number to display name")
			fmt.Scanln(&roll)
			name := getName(st,roll)
			fmt.Println(name)

		case 8:

			var roll int = 0
			fmt.Println("Enter roll number to marks")
			fmt.Scanln(&roll)
			marks := getMarks(st,roll)
			for sub,mark := range(marks) {
				fmt.Println(sub,"->",mark)
			}

		case 9:

			var roll int = 0
			fmt.Println("Enter roll number to age")
			fmt.Scanln(&roll)
			age := getAge(st,roll)
			if age != -1 {
				fmt.Println(age)
			}

		case 10:

			display(st)
			
		case 11:
			break
			
		default :
			fmt.Println("Enter valid choice")
		}
	}	
}
