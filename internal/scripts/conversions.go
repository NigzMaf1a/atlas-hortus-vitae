package scripts

import (
	"log"
	"strconv"
)

func ConvertToInteger(value string) (int64, error) {
	num, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		log.Fatal("Error occurred while converting to integer")
		return 0, err
	}

	log.Println("String converted successfully")

	return num, nil
}
