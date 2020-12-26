package main

import (
	"strings"
	"golang.org/x/net/html"
	"fmt"
	"io/ioutil"
)

type Link struct{
	Href string
	Text string
}

func getText(n *html.Node) string{
	str := ""
	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		if(c.Type == html.TextNode){
			str += strings.TrimSpace(c.Data);
		}else if(c.Type == html.ElementNode){
			str += getText(c);
		}
	}
	return str
}
func Dfs(n *html.Node, links []Link) ([]Link){
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := Link{}
		for _, a := range n.Attr {
			fmt.Println(a)
			if a.Key == "href"{
				newLink.Href = a.Val
				newLink.Text = getText(n);
				// fmt.Println(newLink)
				links = append(links, newLink)
				break		
			}
		}
		return links
	}


	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		links = Dfs(c, links)
	}
	return links
}

func getData(filename string) string{
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(file);
}

func main(){
	// define a static html string
	doc := getData("examples/ex4.html")

	// create a new IO reader
	r := strings.NewReader(doc)
	// parse HTML, which returns doc and err
	// package : https://godoc.org/golang.org/x/net/html#Node
	nodes, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	// create a custom function dfs
	// create a struct type
	/**
		Link{
		Href: "/dog",
		Text: "Something in a span Text not in a span Bold text!",
		}
	*/
	links := make([]Link, 0)
	// call a custom function which runs DFS on the elements in doc
	links = Dfs(nodes, links)

	fmt.Println("final >>>>>", links)	
}