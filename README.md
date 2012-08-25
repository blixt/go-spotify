# Go Spotify

This is a simple module to communicate with the Spotify API.

**Note:** This is library is currently very simple and only supports a few
requests.

## Installation

    go get github.com/blixt/go-spotify/spotify

## Examples

### Perform parallel search queries

    package main
    
    import (
    	"fmt"
    	"github.com/blixt/go-spotify/spotify"
    )
    
    func main() {
    	api := spotify.GetApi()
    
    	queries := make(chan *spotify.SearchTrackQuery)
    
    	// Make two search queries in parallel.
    	go api.SearchTrack("never gonna give you up", queries)
    	go api.SearchTrack("trololo", queries)
    
    	// Print the first track of each query result.
    	for i := 0; i < 2; i++ {
    		query := <-queries
    		fmt.Printf("Got %d results for \"%s\". First result:\n", query.Info.NumResults, query.Query)
    
    		track := query.Tracks[0]
    		fmt.Println(track.Href, track.Name, "by", track.Artists[0].Name)
    	}
    }
