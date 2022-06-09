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

//this would be ripe for (issues on input) to the system in so many ways...
func GetRating(r int) Rating {
	if r > -3 && r < 3 {
		return Rating(r)
	}
	return MEH //zerivalues are what they are...
}

//pretty-print, gets string for the ~iota
func (r Rating) pp() string {
	switch r {
	case SUPER_GREAT:
		return "SUPER_GREAT"
	case KINDA_GREAT:
		return "KINDA_GREAT"
	case MEH:
		return "MEH"
	case KINDA_NEGATIVE:
		return "KINDA_NEGATIVE"
	case SUPER_NEGATIVE:
		return "SUPER_NEGATIVE"
	default:
		return "here were found dragons"
	}
}
