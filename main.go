package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// create a file
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if err := w.Write([]string{
		"key",
		"input_channel_id",
		"input_video_id",
	}); err != nil {
		log.Fatalln("error writing header record to csv:", err)
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	for i := 0; i < 100099993939339; i++ {
		fmt.Println(i)
		if err := w.Write([]string{
			time.Now().UTC().String(),
			"UCPW580o8ITQJFLCQlV6uuAw",
			"LO3ZdbYK8yI",
		}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		// Write any buffered data to the underlying writer (standard output).
		w.Flush()
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
