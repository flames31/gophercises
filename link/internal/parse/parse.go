package parse

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHTML(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return []Link{}, err
	}

	nodes := getNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func getNodes(root *html.Node) []*html.Node {
	if root.Type == html.ElementNode && root.Data == "a" {
		return []*html.Node{root}
	}
	var nodes []*html.Node

	for node := root.FirstChild; node != nil; node = node.NextSibling {
		nodes = append(nodes, getNodes(node)...)
	}

	return nodes
}

func buildLink(node *html.Node) Link {
	var link Link
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	link.Text = text(node)

	return link
}

func text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	var ret string
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		ret += text(n)
	}

	return strings.Join(strings.Fields(ret), " ")
}
