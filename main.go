package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkoukk/tiktoken-go"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/teris-io/shortid"
)

var model = "gpt-4-0314"
var cnt = 5

func main() {
	// Initialize reusable components
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatal(err)
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	generators := []func(){
		func() { gen("shortid", cnt, func() string { id, _ := sid.Generate(); return id }) },
		func() { gen("nanoid", cnt, func() string { id, _ := gonanoid.Generate(alphabet, 10); return id }) },
		func() { gen("snowflake", cnt, func() string { return node.Generate().String() }) },
		func() { gen("xid", cnt, func() string { return xid.New().String() }) },
		func() { gen("ksuid", cnt, func() string { return ksuid.New().String() }) },
		func() { gen("uuid", cnt, func() string { return uuid.New().String() }) },
	}

	for _, generate := range generators {
		generate()
	}
}

func gen(name string, count int, generateFunc func() string) {
	totalTokens := 0
	for i := 0; i < count; i++ {
		id := generateFunc()
		tokens := measureTokens(id)
		totalTokens += tokens
		print(id, tokens)
	}
	averageTokens := int(math.Ceil(float64(totalTokens) / float64(count)))
	fmt.Printf("-- %s (avg token: %d)\n\n", name, averageTokens)
}


func print(id string, tokens int) {
	fmt.Printf("%s (len=%d) (t=%d)\n", id, len(id), tokens)
}

func measureTokens(id string) (numTokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		log.Printf("encoding for model: %v", err)
		return
	}
	return len(tkm.Encode(id, nil, nil))
}
