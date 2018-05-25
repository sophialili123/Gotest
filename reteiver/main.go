package main

import (
	"fmt"
	"mock"
	"real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.zhenlihan.com")
}
func post(poster Poster) {
	poster.Post("http://www.zhenlihan.com", map[string]string{
		"name":   "ccmouse",
		"course": "goLang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake zhenlihan.com"}
	//fmt.Printf("%T %v\n", r, r)
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Printf("%T %v\n", r, r)
	inspect(r)
	//fmt.Println(download(r))

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))

	fmt.Println(`dafsdaf
dsafdsafdaf
dsaf`)

}

func inspect(r Retriever) {
	fmt.Println("Inspecting,", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Printf("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("Contents", v.UserAgent)
	}

	fmt.Println()
}
