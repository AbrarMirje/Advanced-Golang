package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	// urlPath := "https://example.com:8080/path?query=param#fragment"
	urlPath := "https://example.com:8080/path/to/resource?user=john&id=42#section2"
	parsedUrl, err := url.Parse(urlPath) //Parse mean extracting data from something
	if err != nil {
		fmt.Println("Error while parsing URL", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	fmt.Println("______________________")
	rawUrl := "https://example.com/path?name=John&age=30"
	parsedUrl1, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error while parsing URL", err)
		return
	}
	queryParams := parsedUrl1.Query()
	fmt.Printf("%T\n", queryParams)
	fmt.Println(queryParams)
	// Method .Has() check whether the given param is present or not
	if queryParams.Has("name") && queryParams.Has("age") {
		fmt.Println("Name:", queryParams.Get("name"))
		fmt.Println("Age:", queryParams.Get("age"))
	} else {
		fmt.Println("param[name] and param[age] not found❌")
	}

	fmt.Println("--------------Building URLs------------------")
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/path/to/resource",
	}
	query := baseUrl.Query()
	query.Set("name", "abrar")
	query.Set("age", "23")
	query.Add("country", "india")
	baseUrl.RawQuery = query.Encode() // Convert / Encode into a URL format
	fmt.Println(baseUrl)

	values := url.Values{}
	values.Add("name", "David")
	values.Add("age", "21")
	values.Add("city", "London")
	values.Add("country", "UK")
	encodeQueries := values.Encode()
	fmt.Println(encodeQueries)

	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodeQueries
	fmt.Println(fullUrl)
}
