package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/inhies/go-bytesize"
)

var (
	maxSize float64 = 100
	header          = []string{
		"key",
		"input_channel_id",
		"input_video_id",
	}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input file name: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fileName := strings.Replace(text, "\n", "", -1)
	if fileName == "" {
		log.Fatal("file name is empty")
	}

	fmt.Println("")
	fmt.Println(`Valid byte units are "B", "KB", "MB", "GB", "TB", "PB" and "EB"]: `)
	fmt.Print(`file size wanted: `)
	text, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fileSizeWanted := strings.Replace(text, "\n", "", -1)
	if fileName == "" {
		log.Fatal("file size is empty")
	}
	b, err := bytesize.Parse(fileSizeWanted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
	maxSize = float64(b)

	fmt.Println("begin gen file csv...")
	time.Sleep(2 * time.Second)

	file, err := os.Create(fmt.Sprintf("%v.csv", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSize := float64(0)
	fistRow := true
	for {
		writer := bufio.NewWriter(file)
		rowContent := []string{
			time.Now().UTC().String(),
			"UCPW580o8ITQJFLCQlV6uuAw",
			"LO3ZdbYK8yI",
		}

		if fistRow {
			rowContent = header
			fistRow = false
		}

		len, err := writer.WriteString(strings.Join(rowContent, ",") + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fileSize += float64(len)

		b1 := bytesize.New(fileSize)
		fmt.Printf("wrote %s \n", b1)
		writer.Flush()

		if fileSize >= maxSize {
			break
		}
	}

	bt := bytesize.New(fileSize)
	fmt.Printf("gen %s file [%s] csv success.\n", fileName, bt)
}
