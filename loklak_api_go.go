// Package loklak_api_go is a library for interacting with http://loklak.org/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/hokaccha/go-prettyjson"
)

// The Loklak Object structure.
type Loklak struct {
	baseUrl    string
	name       string
	followers  string
	following  string
	query      string
	since      string
	until      string
	source     string
	count      string
	fields     string
	from_user  string
	limit      string
}

// Initiation of the loklak object
func (l *Loklak) Connect(urlString string) {
	u, err := url.Parse(urlString)
	if (err != nil) {
		fmt.Println(u)
		fatal(err)
	} else {
		l.baseUrl = urlString
	}
}

// A generic query URL request and fetch JSON response
// This should be suitable for a majority of the JSON based responses
// Plain text and CSV format responses need another custom control function.
// Function name: getJSON()
// Scope        : globally accessible
// Parameters   : string              , Variable => route
// Return Types : JSON Response       , Variable => string
//              : Error Response      , Variable => error
//
// Makes a request to the given URL and returns the JSON response obtained
// and error if any.

func getJSON(route string) (string, error) {
	r, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	var b interface{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return "", err
	}
	out, err := prettyjson.Marshal(b)
	return string(out), err
}


// The API function for the /api/hello.json api call.
func (l *Loklak) hello() (string) {
	apiQuery := l.baseUrl + "api/hello.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API function for the /api/peers.json api call.
func (l *Loklak) peers() (string) {
	apiQuery := l.baseUrl + "api/peers.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API function for the /api/status.json api call.
func (l *Loklak) status() (string) {
	apiQuery := l.baseUrl + "api/status.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for the /api/apps.json api call.
func (l *Loklak) apps() (string) {
	apiQuery := l.baseUrl + "api/apps.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for /api/settings.json api call.
// This is only a localhost query
func (l *Loklak) settings() (string) {
	apiQuery := "http://localhost:9000/api/settings.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// Helper function to return the error responses to stderr
// Function name: fatal()
// Scope        : globally accessible
// Parameters   : Error               , Variable => err
// Exits the program as soon as a fatal error is obtained.

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// Everything below this is purely for testing only
// The library ends above and needs to be packaged once all the API
// components in the library are ready.

func main() {
	loklakObject := new(Loklak)
	loklakObject.Connect("http://192.168.0.3:9000/")
	helloResponse := loklakObject.hello()
	fmt.Println(helloResponse)
	peersResponse := loklakObject.peers()
	fmt.Println(peersResponse)
	statusResponse := loklakObject.status()
	fmt.Println(statusResponse)
	appsResponse := loklakObject.apps()
	fmt.Println(appsResponse)
	settingsResponse := loklakObject.settings()
	fmt.Println(settingsResponse)
}
