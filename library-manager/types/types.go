package types

import "time"

type Books []Book
type Magazines []Magazine
type Status uint

const (
	Borrowed Status = iota
	Returned
	LateFee = 1
)

type Media struct {
	Id              uint
	Title           string
	PublicationYear string
	IsAvailable     bool
}

type Book struct {
	Media
	Author string
	ISBN   string
}

type Magazine struct {
	Media
	IssueNumber uint
	Publisher   string
}

type BookCollection struct {
	AllBooks []Book
}

type MagazineCollection struct {
	AllMagazines []Magazine
}

type BorrowedBook struct {
	Name         string
	BorrowDate   time.Time
	ReturnDate   time.Time
	Fees         uint
	BorrowStatus Status
	Book
}

type BorrowedBooksCollection struct {
	AllBorrowedBooks map[uint]BorrowedBook
}
