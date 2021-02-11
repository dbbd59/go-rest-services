package hn

import (
	"fmt"
	"time"

	"github.com/PaulRosset/go-hacknews"
	"github.com/gin-gonic/gin"
)

type HackerNewsResponse struct {
	Posts []hackerNewsPostResponse
}

type hackerNewsPostResponse struct {
	By    string
	Id    int
	Score int
	Time  time.Time
	Title string
	Type  string
	Url   string
}

func GetHackerNews(c *gin.Context, s string) HackerNewsResponse {
	//topstories/newstories/beststories/askstories/showstories/jobstories
	init := hacknews.Initializer{Story: s, NbPosts: 10}

	codes, err := init.GetCodesStory()

	if err != nil {
		fmt.Println(err)
	}
	posts, err := init.GetPostStory(codes)
	if err != nil {
		fmt.Println(err)
	}
	var postsResponse []hackerNewsPostResponse
	for _, post := range posts {
		postsResponse = append(postsResponse, hackerNewsPostResponse{By: post.By, Id: post.Id, Score: post.Score, Time: time.Unix(int64(post.Time), 0), Title: post.Title, Type: post.Type, Url: post.Url})
	}

	return HackerNewsResponse{Posts: postsResponse}

}
