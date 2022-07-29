package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

const EX1 = "ex1.html"
const EX2 = "ex2.html"
const EX3 = "ex3.html"
const EX4 = "ex4.html"

type link struct {
	href string
	text string
}

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func parse(htmlCont string) ([]link, error) {
	doc, err := html.Parse(strings.NewReader(htmlCont))
	if err != nil {
		log.Fatal(err)

	}
	linksArr := []link{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					text := ""
					text = n.FirstChild.Data
					if n.FirstChild != n.LastChild && n.LastChild.Type != html.CommentNode {
						text += n.LastChild.Data
					}
					l := link{href: a.Val, text: text}
					linksArr = append(linksArr, l)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return linksArr, nil
}

func main() {
	htmlCont, err := readHtmlFromFile(EX2)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parse(htmlCont)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(links)
}
