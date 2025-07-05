package utils

const (
	BDT = "BDT"
	AFG = "AFG"
	USD = "USD"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case BDT, AFG, USD:
		return true
	}
	return false
}