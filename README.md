Podcast Index Client
====================

A client for using the Podcast Index API in Golang.

Status: In-Development

```
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

```

Features
--------

- Currently implements the search/byterm endpoint.
- Pass API keys etc via Context objects.
- Parses Result and Feed responses.
- Includes a simple search cmd.

Installation
------------

```
go get kilobit.ca/go/podcastindex
go test -v ./...
```

Building
--------

```
go get kilobit.ca/go/podcastindex
cd cmd/search
go build
PODCAST_INDEX_API_KEY="MyAPIKey" PODCAST_INDEX_API_SECRET="MySecretSecret" ./search no agenda
12 results returned for query, 'no agenda'.

 1. No Agenda
 2. No Agenda
 3. No Agenda Pre-Show Out-Takes (NAPSOT) podcast
 4. No Hidden Agenda Podcast
 5. No Agenda No Apologies
 6. No Agenda BACK_UP
 7. No Agenda Podcast
 8. No Agenda Music's Podcast
 9. No Fixed Agenda
10. No Agenda, Just Vibes
11. NO AGENDA WEEKLY
12. No Agenda

```

Contribute
----------

Please help!  Submit pull requests through
[Github](https://github.com/kilobit/podcast-index-client).

Support
-------

Please submit issues through
[Github](https://github.com/kilobit/podcast-index-client).

License
-------

See LICENSE.

--  
Created: Feb 10, 2021  
By: Christian Saunders <cps@kilobit.ca>  
Copyright 2021 Kilobit Labs Inc.  
