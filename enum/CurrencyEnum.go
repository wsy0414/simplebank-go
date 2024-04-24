package enum

const (
	TWD string = "TWD"
	USD string = "USD"
	JPN string = "JPN"
)

func IsCurrencyValid(currency string) bool {
	switch currency {
	case TWD:
		return true
	case USD:
		return true
	case JPN:
		return true
	default:
		return false
	}
}
