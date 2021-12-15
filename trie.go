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
	PathDelimiter  string = "/"
	ParamDelimiter string = ":"
	RegexWildCard  string = "(.+)"
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

		if r.Path == PathDelimiter {
			currNode.Children[PathDelimiter] = &Node{
				Route:    r,
				Children: make(map[string]*Node),
			}
			return
		}

		for _, s := range deleteEmpty(strings.Split(r.Path, PathDelimiter)) {
			if nextNode, ok := currNode.Children[s]; ok {
				currNode = nextNode
			} else {
				currNode.Children[s] = &Node{
					Route:    r,
					Children: make(map[string]*Node),
				}
				currNode = currNode.Children[s]
			}
		}
		currNode.Route = r
	}
}

func (t *Tree) Search(path, method string) (*SearchResult, error) {
	var params Params
	var count int
	currNode := t.Method[method]

	if path == PathDelimiter {
		return pathIsPathDelim(currNode)
	}

	for _, s := range deleteEmpty(strings.Split(path, PathDelimiter)) {
		if nextNode, ok := currNode.Children[s]; ok {
			currNode = nextNode
		} else {
			if len(currNode.Children) == 0 {
				return HandlerNotRegisted()
			}
			children := currNode.Children
			for section := range children {
				if string(section[0]) == ParamDelimiter {
					ptn := RegexWildCard
					reg := regexp.MustCompile(ptn)
					if reg.Match([]byte(s)) {
						param := getParam(section)
						params = append(params, &Param{Key: param, Value: s})
						currNode = children[section]
						count++
					} else if count == len(children)-1 {
						return HandlerNotRegisted()
					}
				} else {
					return HandlerNotRegisted()
				}
			}
		}
	}

	if currNode.Route.Path != path {
		if params == nil {
			return HandlerNotRegisted()
		}
	}
	return &SearchResult{Route: currNode.Route, Params: params}, nil
}

func pathIsPathDelim(currNode *Node) (*SearchResult, error) {
	if node, ok := currNode.Children[PathDelimiter]; ok {
		return &SearchResult{Route: node.Route, Params: nil}, nil
	}
	return nil, errors.New("handler is not registered")
}

func getParam(section string) string {
	return section[1:]
}
