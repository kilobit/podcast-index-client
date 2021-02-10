/* Copyright 2021 Kilobit Labs Inc. */

package main

import "fmt"
import _ "errors"
import "context"
import "os"
import "strings"

import "kilobit.ca/go/podcastindex/client"

func main() {

	query := strings.Join(os.Args[1:], " ")

	key := os.Getenv("PODCAST_INDEX_API_KEY")
	secret := os.Getenv("PODCAST_INDEX_API_SECRET")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.PICAPIKey, key)
	ctx = context.WithValue(ctx, client.PICAPISecret, secret)

	pic := client.New(ctx)
	result, err := pic.Search(context.TODO(), query)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%d results returned for query, '%s'.\n\n", result.Count(), result.Query())

	for i, feed := range result.Feeds() {
		fmt.Printf("%2.d. %s\n", i+1, feed.Title())
	}

	fmt.Println()
}
