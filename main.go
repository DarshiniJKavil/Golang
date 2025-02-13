#1
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

#2
package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

#3
package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to  %v  booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
}


#4
package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to  %v  booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}


#5
package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to  %v  booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}


#6
package main

import "fmt"

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v,%v,%v", grade1, grade2, grade3)
}

#7
package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v", grades)
}


#8
package main

import "fmt"

func main() {
	var students [3]string
	fmt.Printf("students : %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("students : %v\n", students)
}


#9
package main

import "fmt"

func main() {
	var students [3]string
	fmt.Printf("students : %v\n", students)
	students[0] = "Lisa"
	students[1] = "Rose"
	students[2] = "Jeny"
	fmt.Printf("students : %v\n", students)
}


#10
package main

import "fmt"

func main() {
	var students [3]string
	fmt.Printf("students : %v\n", students)
	students[0] = "Lisa"
	students[1] = "Rose"
	students[2] = "Jeny"
	fmt.Printf("students #1: %v\n", students[1])
	fmt.Printf("Number of students : %v\n", len(students))
}
