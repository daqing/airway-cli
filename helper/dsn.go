package helper

import "os"

func GetDSN() string {
	return os.Getenv("AIRWAY_PG_URL")
}
