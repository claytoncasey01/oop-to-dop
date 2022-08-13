package main

import (
	"fmt"
	"github.com/claytoncasey01/go-oop-to-dod/dod"
	"github.com/claytoncasey01/go-oop-to-dod/oop"
)

func main() {
	posts := oop.MakePosts(10)
	found := oop.FindPostsByAuthorName("Casey", posts)
	fmt.Printf("Length of Oop Posts: %d\n", len(posts))
	fmt.Printf("Found %d posts\n", len(found))

	dodAuthors := dod.MakeAuthors(1)
	dodPosts := dod.MakePosts(10, dodAuthors)
	fmt.Printf("Length of Dod Posts: %d\n", len(dodPosts.Ids))
	//dodSurveys := dod.MakeSurveys(10, 0)
	//dodSurveyInstances := dod.MakeSurveyInstances(1000, dodSurveys.IDs)
	//fmt.Printf("DodSurveys Length: %d\n", len(dodSurveyInstances.IDs))

}
