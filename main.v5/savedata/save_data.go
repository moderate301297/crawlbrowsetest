package savedata

import (
	"fmt"

	"../dbconnection"
)

type DataWeb struct {
	Body string
}

//
func SaveLink(url string) {
	_, err := dbconnection.Connect.Exec("insert all_links set link = ?", url)
	if err != nil {
		fmt.Println("Error mysql: ", err)
	}
}

// save url error
func SaveUrlError(url string) {
	_, err := dbconnection.Connect.Exec("insert urls_error set url = ?", url)
	if err != nil {
		fmt.Println("Error mysql: ", err)
	}
}
