package cgr

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

type Tree struct {
	Method map[string]*Node
}

type Node struct {
	Route    *Route
	Children map[string]*Node
}

type SearchResult struct {
	Route  *Route
	Params Params
}

type Param struct {
	Key, Value string
}

type Params []*Param

const (
	pathDelimiter  string = "/"
	paramDelimiter string = ":"
	regexWildCard  string = "(.+)"
)

func NewTree() *Tree {
	return &Tree{
		Method: map[string]*Node{
			http.MethodGet: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
			http.MethodPost: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
			http.MethodPut: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
			http.MethodDelete: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
			http.MethodPatch: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
			http.MethodOptions: {
				Route:    nil,
				Children: make(map[string]*Node),
			},
		},
	}
}

func (t *Tree) Insert(r *Route) {
	for _, method := range r.Methods {
		currNode := t.Method[method]

		if r.Path == pathDelimiter {
			currNode.Route = r
			return
		}

		for _, s := range deleteEmpty(strings.Split(r.Path, pathDelimiter)) {
			if nextNode, ok := currNode.Children[s]; ok {
				currNode = nextNode
				currNode.Route = r
			} else {
				currNode.Children[s] = &Node{
					Route:    r,
					Children: make(map[string]*Node),
				}
				currNode = currNode.Children[s]
			}
		}
	}
}

func (t *Tree) Search(path, method string) (*SearchResult, error) {
	var params Params
	currNode := t.Method[method]
	for _, s := range deleteEmpty(strings.Split(path, pathDelimiter)) {
		if nextNode, ok := currNode.Children[s]; ok {
			currNode = nextNode
		} else {
			if len(currNode.Children) == 0 {
				return nil, errors.New("handler is not registered")
			}
			for c := range currNode.Children {
				ptn := regexWildCard

				reg := regexp.MustCompile(ptn)

				if reg.Match([]byte(s)) {
					param := getParam(c)
					params = append(params, &Param{Key: param, Value: s})
					currNode = currNode.Children[c]
				}
			}
		}
	}
	return &SearchResult{Route: currNode.Route, Params: params}, nil
}

func getParam(section string) string {
	return section[1:]
}
