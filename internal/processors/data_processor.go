package processors

// ProcessAuctionData processes raw auction data and returns structured items
func ProcessAuctionData(key string, value string) (interface{}, bool) {
	switch key {
	case "Lõpeb":
		return value, true
	case "Kogupindala":
		return value, true
	default:
		return nil, false
	}
}
