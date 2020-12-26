package main

import (
	"flag"
	"fmt"
	link "pkg/link"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/xml"
	"os"
)

func IsExternalUrl(href string) bool{
	if(strings.HasPrefix(href, "#") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "mailto") || strings.HasPrefix(href, "http")){
		return true
	}
	return false
}

func Bfs(links []link.Link, vis map[string]int, parentUrl string)  {
	// fmt.Println("Parent = ", parentUrl)
	for _,l := range links {
		var href string = l.Href
		if(strings.HasPrefix(href, "/")){
			href = parentUrl + href
		}else if(IsExternalUrl(href)){
			continue
		}
		_, ok := vis[href]
		// fmt.Println(href, " exists ", ok)
		if(ok == false){
			// not visited
			vis[href] = 1
			// fmt.Println("Visiting = ",href)
			pageLinks := GetPageLinks(href)
			Bfs(pageLinks, vis, parentUrl)
		}
	}
}


func GetPageLinks(url string) []link.Link {
	resp, err := getPage(url)
	HandleError(err)
	var r,_ = ioutil.ReadAll(resp.Body)
	rs := string(r)
	return link.Parse(rs)
}

func getPage(url string) (*http.Response, error){
	resp, err := http.Get(url) 
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func HandleError(err error){
	if err != nil{
		panic(err)
	}
}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Loc struct {
	Value string `xml:"loc"`
}

type Urlset struct {
	Urls []Loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func GenerateXML(vis map[string]int){
	fmt.Print(xml.Header)

	urlset := Urlset{
		Xmlns: xmlns,
	}
	for url, _ := range vis{
		// fmt.Println(url)
		urlset.Urls = append(urlset.Urls, Loc{url})
		// fmt.Println(urlset.Urls)
	}

	enc := xml.NewEncoder(os.Stdout)
	err := enc.Encode(urlset)
	HandleError(err)
	fmt.Println()
}

func main()  {
	websiteUrl := flag.String("-url", "https://gophercises.com", "url for which sitemap is to be built")
	flag.Parse()
	fmt.Print(*websiteUrl)
	links := GetPageLinks(*websiteUrl)
	vis := make(map[string]int)
	vis[*websiteUrl + "/"] = 1;
	Bfs(links, vis, *websiteUrl)
	GenerateXML(vis)
}