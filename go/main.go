package main

import (
	"fmt"
	"github.com/claytoncasey01/oop-to-dod/dod"
	"github.com/claytoncasey01/oop-to-dod/oop"
	"strconv"
)

func main() {
	authors := oop.MakeAuthors(100)
	posts := oop.MakePosts(10000, authors)

	dodAuthors := dod.MakeAuthors(100)
	dodPosts := dod.MakePosts(10000, dodAuthors)

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
		case "DodFindById":
			for i := 0; i < parsedN; i++ {
				dodPosts.FindById(dodPosts.Ids[50])
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

	//dodSurveys := dod.MakeSurveys(10, 0)
	//dodSurveyInstances := dod.MakeSurveyInstances(1000, dodSurveys.IDs)
	//fmt.Printf("DodSurveys Length: %d\n", len(dodSurveyInstances.IDs))

}
