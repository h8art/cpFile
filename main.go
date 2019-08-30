package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var from string
var to string
var offset int64
var limit int

func init() {
	flag.StringVar(&from, "from", "", "file copy from")
	flag.StringVar(&to, "to", "", "file copy to")
	flag.Int64Var(&offset, "offset", 0, "offset in copy file")
	flag.IntVar(&limit, "limit", 0, "limit of buffer")
}
func main() {
	flag.Parse()
	var file *os.File
	file, err := os.OpenFile(from, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File", from, " does not exist", err)
		}
	}
	defer file.Close()
	if offset == 0 {
		if limit != 0 {
			b := make([]byte, limit)
			_, err := io.ReadFull(file, b)
			if err != nil {
				log.Fatal(err)
			}
			createdfile, _ := os.Create(to)
			written, err := createdfile.Write(b)
			if err != nil {
				log.Panicf("failed to write: %v", err, written)
			}
		}
		if limit == 0 {
			read, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			createdfile, _ := os.Create(to)
			written, err := createdfile.Write(read)
			if err != nil {
				log.Panicf("failed to write: %v", err, written)
			}

		}
	}
	if offset > 0 {
		file.Seek(offset, 0)
		if limit != 0 {
			b := make([]byte, limit)
			_, err := io.ReadFull(file, b)
			if err != nil {
				log.Fatal(err)
			}
			createdfile, _ := os.Create(to)
			written, err := createdfile.Write(b)
			if err != nil {
				log.Panicf("failed to write: %v", err, written)
			}
		}
		if limit == 0 {
			read, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			createdfile, _ := os.Create(to)
			written, err := createdfile.Write(read)
			if err != nil {
				log.Panicf("failed to write: %v", err, written)
			}

		}
	}
}
