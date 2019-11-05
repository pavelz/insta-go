package main

import (
	"database/sql"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"context"
	"github.com/pavelz/insta-go/ent"
	"github.com/pavelz/insta-go/ent/user"
	"log"
	"os"
	"github.com/pavelz/insta-go/src/location"
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
		fmt.Printf("file: %s saved of %d bytes\n", header.Filename, header.Size)

		client := doConnection()
		_, err = client.User.
			Create().
			SetAge(30).
			SetName("whot").
			SetNick("gimmiphotos").
			SetImage(data).
			Save(context.Background())

		if err != nil {
			fmt.Printf("failed creating user: %v", err)
			os.Exit(-1)
		}


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
			fmt.Printf("File: %s served %d bytes\n", f, stat.Size())
		} else {
			// check if there is a record with this name:
			client := doConnection()
			res := client.User.Query().Where(user.Name(f))
			c, err := res.Count(context.Background())
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(-1)
			}

			if c > 0 {
				d, err := res.First(context.Background())
				if err != nil {
					fmt.Printf("error!: %s\n", err)
				}

				w.Write(d.Image)
			}
		}
	}
}


// Ent example
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SetNick("Womp").
		SetImage([]byte("Code")).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}

	log.Println("user was created: ", u)
	return u, nil
}

func doConnection()(*ent.Client){
	db, err := sql.Open("postgres","user=pavel dbname=insta-go sslmode=disable")
	if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(-1)
	}
	drv := entsql.OpenDB("postgres", db)

	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	log.Fatalf("failed opening connection to sqlite: %v", err)
	//}
	//defer client.Close()

	client := ent.NewClient(ent.Driver(drv))

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func entTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("ent test")
	client := doConnection()
	CreateUser(context.Background(),client)
}

func doLocation(w http.ResponseWriter, r *http.Request){
	l := new(location.Location)
	l.Save()
	fmt.Printf("Locations!\n")
}

func main() {
	fmt.Println("insta-go 0.01a")

	http.HandleFunc("/ent", entTest)
	http.HandleFunc("/locations", doLocation)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
