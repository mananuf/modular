package utils

import (
	"time"

	"github.com/mananuf/library-manager/types"
)

func CalculateLateFees(b types.BorrowedBook) float64 {
	totalHours := time.Since(b.ReturnDate).Hours()
	days := totalHours / 24

	return days * types.LateFee
}