package types

// Great question! The specific date "2006-01-02 15:04:05" is not arbitrary.
// In Go, it's a reference layout that represents the format of date and time strings.
//The Go time package uses this exact date (January 2nd, 2006, 15:04:05) as the template for parsing and formatting dates.

const (
	MEAT        FoodType = "MEAT"
	PLANT       FoodType = "PLANT"
	OMNIVORE    FoodType = "OMNIVORE"
	defaultDate string   = "2006-01-02 15:04:05"
)
