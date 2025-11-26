/*
Notes:

There are two ways you can implement env variables injection
1. using os.LookupEnv(key)
2. using .env file with godotenv package

Below method is to read from os directly using os package

*/

package env

import (
	"os"
	"strconv"
)

/*
This is simple and easy way to implement env variables runtime fetching
*/
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

/*
Above function by default gives values in string type,
In case if we want values to be int or boolean we need to convert them explicitly
*/
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
