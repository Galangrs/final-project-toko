package helper

import (
	"fmt"
	"strings"
)

func FormatRupiah(amount int) string {
	amountStr := fmt.Sprintf("%d", amount)

	var groups []string
	for i := len(amountStr); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		groups = append([]string{amountStr[start:i]}, groups...)
	}

	result := "Rp " + strings.Join(groups, ".")

	return result
}
