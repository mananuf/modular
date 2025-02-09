package types

import "time"

var mediaId uint

// ---------------------------------------------------------------------------
// ----------------------------- Receiver Methods ----------------------------
// ---------------------------------------------------------------------------

func (bc *BookCollection) AddBookToCollection(book *Book) {
	bc.AllBooks = append(bc.AllBooks, *book)
}

func (m *Media) MediaAvailability() bool {
	return m.IsAvailable
}

// --------------------------------------------------------------------------
// ---------------------------- Constructor Methods -------------------------
// ---------------------------------------------------------------------------

func NewMedia(title, publicationYear string, isAvailable bool) *Media {
	return &Media{
		Id:              mediaId + 1,
		Title:           title,
		PublicationYear: publicationYear,
		IsAvailable:     isAvailable,
	}
}

func NewBook(author, isbn string, media *Media) *Book {
	return &Book{
		Media:  *media,
		Author: author,
		ISBN:   isbn,
	}
}

func NewMagazine(issueNumber uint, publisher string, media *Media) *Magazine {
	return &Magazine{
		Media:       *media,
		IssueNumber: issueNumber,
		Publisher:   publisher,
	}
}

func NewBookCollection() *BookCollection {
	return &BookCollection{
		AllBooks: make(Books, 0),
	}
}

func NewMagazineCollection() *MagazineCollection {
	return &MagazineCollection{
		AllMagazines: make(Magazines, 0),
	}
}

func NewBorrowedBook(collected, toBeRreturned time.Time, fees uint, name string, status Status, book *Book) *BorrowedBook {
	return &BorrowedBook{
		Name:         name,
		BorrowDate:   collected,
		ReturnDate:   toBeRreturned,
		Fees:         fees,
		BorrowStatus: status,
		Book:         *book,
	}
}

func NewBorrowedBooksCollection() *BorrowedBooksCollection {
	return &BorrowedBooksCollection{
		AllBorrowedBooks: make(map[uint]BorrowedBook),
	}
}

func (bbc *BorrowedBooksCollection) AddBorrowedBookToCollection(borrowedBook *BorrowedBook) {
	bbc.AllBorrowedBooks[borrowedBook.Book.Id] = *borrowedBook
}