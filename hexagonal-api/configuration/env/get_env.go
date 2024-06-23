package env

import "os"

func GetNewsTokenAPI() string {
	return os.Getenv("NEWS_API_TOKEN")
}

func GetNewsBaseURLAPI() string {
	return os.Getenv("NEWS_API_BASE_URL")
}
