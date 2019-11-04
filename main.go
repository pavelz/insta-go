package main

import (
	"log"
	"os"
)
import "fmt"
import "net/http"
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func handler(w http.ResponseWriter, r *http.Request){
	f := r.URL.Path[1:]
	if(r.Method == "POST"){
		// read photo and store it
		fmt.Printf("a POST request!\n")
		a_file, header, _ := r.FormFile("file")
		write_to,_ := os.OpenFile(header.Filename ,os.O_RDWR|os.O_CREATE, os.FileMode(0666))
		fmt.Printf("filename upload: %s\nfile size:%d\n", header.Filename, header.Size)
		data := make([]byte, header.Size)
		n, err := a_file.Read(data)
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
		fmt.Printf("Read: %d bytes\n", n)
		_, err = write_to.Write(data)
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
		_ = write_to.Close()
	} else {

		if (Exists(f)) {
			// serve the file
			file, err := os.OpenFile(f, 0, os.FileMode(0777))

			if (err != nil) {
				os.Exit(-1)
			}
			stat, _ := file.Stat()
			data := make([]byte, stat.Size())
			file.Read(data)
			w.Write(data)
		}
	}
}

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
