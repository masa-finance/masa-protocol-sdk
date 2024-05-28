package db

import "oracledata"

type Work struct {
	// Define the structure fields here
}

func FireData(uid string, payload []byte, response []byte) error {
	// Implement your function here
	return nil
}

func FireEvent(uid string, value []byte) error {
	// Implement your function here
	return nil
}

func GetData(uid string) ([]oracledata.Work, error) {
	// Implement your function here
	return nil, nil
}
