package main

type Rating int

const (
	SUPER_NEGATIVE Rating = iota - 2
	KINDA_NEGATIVE
	MEH
	KINDA_GREAT
	SUPER_GREAT
)

// func checkIfRating(s string) Rating {
// 	switch s {
// 	case "SUPER_NEGATIVE":
// 		return SUPER_NEGATIVE
// 	//// 	...
// 	//// default:
// 	}
// }
// can generate nice string output etc w fmters+gogenerate.... meh
