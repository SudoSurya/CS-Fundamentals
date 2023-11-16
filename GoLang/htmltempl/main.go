package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct{
    Type NodeType
    Data string
    Attr []Attribute
    FirstChild, NextSibling *Node
}

type NodeType int32

const (
    ErrorNode NodeType = iota
    TextNode
    DocumentNode
    ElementNode
    CommentNode
    DoctypeNode
)

type Attribute struct{
    Key, Val string
}
func Parse(r io.Reader) (*Node, error){
    return nil, nil
}
func main(){
    doc, err := Parse(os.Stdin)
    if err != nil{
        fmt.Fprintf(os.Stderr, "parse error: %s", err)
        os.Exit(1)
    }
    for _,link := range visit(nil, doc){
        fmt.Println(link)
    }
}

func visit(links []string, n *Node) []string{
    if n.Type == ElementNode && n.Data == "a"{
        for _, a := range n.Attr{
            if a.Key == "href"{
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling{
        links = visit(links, c)
    }
    return links
}


