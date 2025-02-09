package main

import (
	"fmt"
	"time"

	"github.com/mananuf/library-manager/types"
	"github.com/mananuf/library-manager/utils"
)

func main() {
	thingsFallApartMedia := types.NewMedia(
		"Things Fall Apart",
		"2002",
		true,
	)
	thingsFallApart := types.NewBook(
		"Chinua Achebe",
		"QWErt1234",
		thingsFallApartMedia,
	)

	bc := types.NewBookCollection()

	bc.AddBookToCollection(thingsFallApart)

	fmt.Println(thingsFallApart.MediaAvailability())

	borrowDate := time.Date(2025, 1, 1, 8, 30, 0, 0, time.UTC)
	returnDate := time.Date(2025, 1, 14, 8, 30, 0, 0, time.UTC)

	bb := types.NewBorrowedBook(
		borrowDate,
		returnDate,
		5,
		"Badman Newls",
		types.Borrowed,
		thingsFallApart,
	)

	bbc := types.NewBorrowedBooksCollection()
	bbc.AddBorrowedBookToCollection(bb)

	fmt.Printf("%v\n", bb)

	lateFees := utils.CalculateLateFees(*bb)
	fmt.Printf("late Fees: $%.2f\n", lateFees)
}