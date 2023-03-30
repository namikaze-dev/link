package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func Parse(r io.Reader) ([]Link, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return traverse(root), nil
}

func traverse(node *html.Node) []Link {
	var links []Link
	if isAnchorNode(node) {
		links = append(links, Link{
			Href: getHref(node),
			Text: strings.TrimSpace(getText(node.FirstChild)),
		})
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, traverse(child)...)
	}
	return links
}

func isAnchorNode(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

func getHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func getText(node *html.Node) string {
	var res string
	for next := node; next != nil; next = next.NextSibling {
		if next.Type == html.TextNode {
			res += cleanText(next.Data)
		}

		if next.Type == html.ElementNode {
			res += getText(next.FirstChild)
		}
	}
	return res
}

func cleanText(text string) string {
	text = strings.ReplaceAll(text, "\n", "")
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return ""
	}

	if text[0] == ' ' {
		trimmed = " " + trimmed
	}
	if text[len(text)-1] == ' ' {
		trimmed += " "
	}
	return trimmed
}