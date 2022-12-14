package main

import (
	"fmt"
	"strconv"

	"github.com/claytoncasey01/oop-to-dop/dop"
	"github.com/claytoncasey01/oop-to-dop/oop"
)

func main() {
	authors := oop.MakeAuthors(100)
	posts := oop.MakePosts(10000, authors)

	dopAuthors := dop.MakeAuthors(100)
	dopPosts := dop.MakePosts(10000, dopAuthors)

  // TODO: Possibly display a menu here with the commands or pass by arg?
	var command string
	fmt.Print("Enter a command to run: ")
	_, err := fmt.Scanln(&command)
	fmt.Print("\n Enter the amount of times: ")
	var n string
	_, err = fmt.Scanln(&n)
	parsedN, _ := strconv.Atoi(n)

	if err == nil {
		fmt.Printf("Running command %s %d times\n", command, parsedN)
		switch command {
		case "OopFindById":
			for i := 0; i < parsedN; i++ {
				oop.FindPostById(posts[0].Id, posts)
			}
		case "dopFindById":
			input := dop.FindPostByIdInput{Ids: dopPosts.Ids, Id: dopPosts.Ids[50]}
			for i := 0; i < parsedN; i++ {
				dop.FindPostById(input)
				//fmt.Printf("Post Index: %d\n", postId)
			}
		case "OopFindByTitle":
			for i := 0; i < parsedN; i++ {
				oop.FindPostByTitle(posts[0].Title, posts)
			}
		case "OopFindByAuthorName":
			for i := 0; i < parsedN; i++ {
				oop.FindPostsByAuthorName("Author 0", posts)
			}
		default:
			fmt.Println("Unknown argument")
		}
	}

	//dopSurveys := dop.MakeSurveys(10, 0)
	//dopSurveyInstances := dop.MakeSurveyInstances(1000, dopSurveys.IDs)
	//fmt.Printf("dopSurveys Length: %d\n", len(dopSurveyInstances.IDs))

}
