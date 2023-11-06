package tests

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"testing"
)

func TestQueryAllWords(t *testing.T) {
	c, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	status, body, err := c.Get(context.Background(), nil, "http://localhost:8888/buzzwords")
	if err != nil {
		return
	}
	fmt.Println(status)
	fmt.Println(string(body))

}
func TestAddBuzzword(t *testing.T) {
	c, err := client.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	args := protocol.Args{}
	args.Set("word", "嫦娥一号")
	args.Set("paraphrase", "")
	args.Set("tags", "[{id:123,type_id:123,value:\"中性\"}]")
	status, body, err := c.Post(context.Background(), nil, "http://localhost:8888/buzzword", &args)
	if err != nil {
		return
	}
	fmt.Println(status)
	fmt.Println(string(body))
}
