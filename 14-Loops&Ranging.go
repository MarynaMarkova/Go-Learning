// ===== 14. Loops and Ranging Over Data =====

package main

import "log"

// //==== The Simplest Loop ====

// func main() {
// 	for i := 0; i <= 5; i++ {
// 		log.Println(i)}
// }

// //===== Ranging the Slice =====

// func main() {
// 	animals := []string{"dog", "fish", "horse", "cow", "cat"}

// 	for i, animal := range animals{
// 	log.Println(i, animal)
// 	}
// }

// //==== Ranging with Blank Identifier ====
// func main() {
// 	animals := []string{"dog", "fish", "horse", "cow", "cat"}

// 	for _, animal := range animals{
// 	log.Println(animal)
// 	}
// }

// //==== Looping through the Map ====

// func main() {
// animals := make(map[string]string)
// animals["dog"] = "Fido"
// animals["cat"] = "Fluffy"

// for animalType, animal := range animals {
// log.Println(animalType, animal)
// }
// }

// //==== Ranging a String ====

// //In GO a string is a slice of bytes.

// func main() {
// var firstLine ="Once upon a midnight dreary"

// for i, l := range firstLine {
// log.Println(i, ";", l)
// }
// }

// ==== Ranging a TYPE ====

func main() {
type User struct{
	FirstName string
	LastName string
	Email string
	Age int
}

var users []User
users = append(users, User{"John", "Smith", "john@smith.com", 30})
users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

for _, l := range users {
	log.Println(l.FirstName, l.LastName, l.Email, l.Age)
}
}