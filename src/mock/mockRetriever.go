package mock

import "fmt"

type Retriever struct {
	Contents string
}

//重写string方法
func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever:{Contents=%s}",r.Contents,
	)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
