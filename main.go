package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/inhies/go-bytesize"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	maxSize := readFileSize(reader)
	fileName := readFileName(reader)

	header := readHeader(reader)
	content := readContenFile(reader)
	header = append(header, "created_at")

	fmt.Println("begin gen file csv...")
	time.Sleep(1 * time.Second)

	file, err := os.Create(fmt.Sprintf("%v.csv", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSize := float64(0)
	fistRow := true
	for {
		writer := bufio.NewWriter(file)

		rowContent := append(content, time.Now().UTC().Format(time.RFC3339Nano))
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

func readHeader(reader *bufio.Reader) (header []string) {
	fmt.Print("input header csv [format with Comma]: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(text, ",") {
		colName := strings.TrimSpace(v)
		header = append(header, colName)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	return header
}

func readFileName(reader *bufio.Reader) string {
	fmt.Print("input file csv name (not included file extension): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fileName := strings.Replace(text, "\n", "", -1)
	if fileName == "" {
		log.Fatal("file name is empty")
	}
	return fileName
}

func readFileSize(reader *bufio.Reader) float64 {
	fmt.Println(`Valid byte units are "B", "KB", "MB", "GB", "TB", "PB" and "EB"]: `)
	fmt.Print(`file size wanted: `)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fileSizeWanted := strings.Replace(text, "\n", "", -1)
	b, err := bytesize.Parse(fileSizeWanted)
	if err != nil {
		log.Fatal(err)
	}

	return float64(b)
}

func readContenFile(reader *bufio.Reader) []string {
	fmt.Print(`content file [format with Comma]: `)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	content := []string{}
	for _, v := range strings.Split(text, ",") {
		colName := strings.TrimSpace(v)
		content = append(content, colName)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	if len(content) == 0 {
		log.Fatal("content file is empty")
	}
	return content
}
