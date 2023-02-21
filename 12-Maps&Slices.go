// package main

// import (
// 	"log"
// )

// // type User struct {
// // 	FirstName string
// // 	LastName  string
// // }

// // ===== MAPS =====

// // func main() {

// // === myMap as a string ===

// // myMap := make(map[string]string)

// // myMap["dog"] = "Samson"
// // myMap["other-dog"] = "Cassie"

// // myMap["dog"] = "Fido"

// // log.Println(myMap["dog"])
// // log.Println(myMap["other-dog"])

// // === myMap as an int ===

// // myMap := make(map[string]int)

// // myMap["First"] = 1
// // myMap["Second"] = 2

// // log.Println(myMap["First"], myMap["Second"])

// // === myMap with a struct ===

// // myMap := make(map[string]User)

// // me := User{
// // 	FirstName: "Maryna",
// // 	LastName:  "Markova"}

// // myMap["me"] = me
// // log.Println(myMap["me"].FirstName)

// // //  === myMap with any type of data (NOT recommended)===
// // myMap := make(map[string]interface{})

// // Maps are immutable. They stay the same no matter where they go. You never have to bother pointing a passing a pointer to a map. You can just pass the map itself and that map will remain constant no matter where in the program it is accessed.

// // Maps are programmatically built into the system not sorted. So if you put things into a map in a certain order and then you try to access them later on, you cannot assume that they are going to be in the order you originally added. You always must look up by key. They may be in RANDOM order.
// // }

// // ===== SLICES =====

// func main() {
// // var mySlice []string

// // mySlice = append(mySlice, "Maryna")
// // mySlice = append(mySlice, "John")
// // mySlice = append(mySlice, "Olha")
// //add to the end of slice

// // var mySlice []int

// // mySlice = append(mySlice, 2)
// // mySlice = append(mySlice, 1)
// // mySlice = append(mySlice, 3)

// // sort.Ints(mySlice)
// // //sorts slice

// // numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// // log.Println(numbers[0:2])
// //prints part of slice (counting starts from 0 like in arrays)

// names := []string{"one", "seven", "fish", "cat"}

// log.Println(names)

// }