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

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func textNodes(n *html.Node) (ret string) {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += textNodes(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func buildLink(n *html.Node) link {
	return link{href: n.Attr[0].Val, text: textNodes(n)}
}

func parse(htmlCont string) (linkArr []link, err error) {
	doc, err := html.Parse(strings.NewReader(htmlCont))
	if err != nil {
		return nil, err
	}

	linkNodes := linkNodes(doc)
	for _, node := range linkNodes {
		linkArr = append(linkArr, buildLink(node))
	}
	return linkArr, nil
}

func main() {
	htmlCont, err := readHtmlFromFile(EX4)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parse(htmlCont)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", links)
}
