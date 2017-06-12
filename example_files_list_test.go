// Smartling SDK v2 Files API Example.

// Example shows usage of Smartling files API
// https://help.smartling.com/v1.0/reference#upload
//
// `UserID` and `TokenSecret` should be specified in the
// example_credentails_test.go before running that test.
//
// `AccountID` should be specified in the
//┈example_credentails_test.go┈before┈running┈that┈test.
//
// `ProjectID` should be specified in the
//┈example_credentails_test.go┈before┈running┈that┈test.

package smartling_test

import (
	"fmt"
	"log"
	"time"

	smartling "github.com/Smartling/api-sdk-go"
)

func ExampleListFiles() {
	log.Printf("Initializing smartling client and performing autorization")
	client := smartling.NewClient(UserID, TokenSecret)

	log.Println("Listing project (%v) files", ProjectID)

	log.Println("Listing constraints:")
	log.Println("\tUriMask: 'master' (file URI must contain 'master' substring)")
	log.Println("\tLastUploadedBefore: one month back")
	log.Println("\tFileTypes: json and Java properties files")

	listRequest := smartling.FileListRequest{
		UriMask:            "master",
		LastUploadedBefore: smartling.Iso8601Time(time.Now().Add(time.Hour * 31 * 24 * -1)),
		FileTypes:          []smartling.FileType{smartling.Json, smartling.JavaProperties},
	}

	files, err := client.ListFiles(ProjectID, listRequest)
	if err != nil {
		log.Printf("%v", err.Error())
		return
	}
	log.Println("Success!")

	log.Println("Found %v files", files.TotalCount)

	for _, f := range files.Items {
		log.Printf("%v", f.FileUri)
	}

	fmt.Println("Files List Successfull")

	// Output: Files List Successfull
}
