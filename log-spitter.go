package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		result, _ := strconv.Atoi(value)
		return result
	}
	return fallback
}

func main() {

	count := 0
	log_count := getEnv("LOG_COUNT", 20000)
	period := getEnv("PERIOD", 0)
	max_count := getEnv("MAX_COUNT", 1000000)

	start := time.Now()

	for count < max_count {
		for i := 0; i < log_count; i++ {
			count += 1
			elapsed := time.Since(start)
			log.Printf("Time elapsed: %s, total log count(s): %d, %f log/s", elapsed, count, float64(count)/elapsed.Seconds())
		}

		time.Sleep(time.Duration(period) * time.Millisecond)
	}

}
