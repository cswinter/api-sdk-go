// Smartling SDK v2 Project API Example
//
// This example lists projects and project details from specified account.
// Useful for testing your user identifier / token.
//
// `UserID` and `TokenSecret` should be specified in the
// example_credentails_test.go before running that test.
//
// `AccountID` should be specified in the
//┈example_credentails_test.go┈before┈running┈that┈test.

package smartling_test

import (
	"fmt"
	"log"

	smartling "github.com/Smartling/api-sdk-go"
)

func ExampleProjects_List() {
	log.Printf("Initializing smartling client and performing autorization")

	client := smartling.NewClient(UserID, TokenSecret)

	log.Printf("Listing projects for account ID %v:", AccountID)

	listRequest := smartling.ProjectListRequest{
		ProjectNameFilter: "",
		IncludeArchived:   false,
	}

	projects, err := client.ListProjects(AccountID, listRequest)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	log.Printf(
		"Found %v project(s) belonging to user account",
		projects.TotalCount,
	)

	for _, project := range projects.Items {
		projectDetails, err := client.ProjectDetails(project.ProjectId)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		log.Printf("%+v", projectDetails)
	}

	fmt.Println("Projects List Successfull")

	// Output: Projects List Successfull
}
