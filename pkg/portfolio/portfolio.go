package portfolio

import "errors"

func GetPortfolio(id int) (string, error) {
	switch id {
	case 1:
		return "My Portfolio", nil
	case 2:
		return "My Managed Portfolio", nil
	default:
		return "", errors.New("id not found")
	}
}
