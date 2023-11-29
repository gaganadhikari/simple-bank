package util

const (
	NPR = "NPR"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, NPR, EUR:
		return true
	}
	return false
}
