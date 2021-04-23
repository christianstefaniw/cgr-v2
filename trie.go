package cgr

import (
	"net/http"
	"strings"
)

type tree struct {
	method map[string]*node
}

type node struct {
	route    *route
	children map[string]*node
}

const (
	pathDelimiter  string = "/"
	paramDelimiter string = ":"
)

func newTree() *tree {
	return &tree{
		method: map[string]*node{
			http.MethodGet: {
				route:    nil,
				children: make(map[string]*node),
			},
			http.MethodPost: {
				route:    nil,
				children: make(map[string]*node),
			},
			http.MethodPut: {
				route:    nil,
				children: make(map[string]*node),
			},
			http.MethodDelete: {
				route:    nil,
				children: make(map[string]*node),
			},
			http.MethodPatch: {
				route:    nil,
				children: make(map[string]*node),
			},
			http.MethodOptions: {
				route:    nil,
				children: make(map[string]*node),
			},
		},
	}
}

func (t *tree) insert(r route) {
	currNode := t.method[r.method]

	for _, s := range deleteEmpty(strings.Split(r.path, pathDelimiter)) {
		if nextNode, ok := currNode.children[s]; ok {
			currNode = nextNode
		} else {
			currNode.children[s] = &node{
				route:    &r,
				children: make(map[string]*node),
			}
			currNode = currNode.children[s]
		}
	}
}

func (t *tree) search(path, method string) *route {
	currNode := t.method[method]

	for _, s := range deleteEmpty(strings.Split(path, pathDelimiter)) {
		if nextNode, ok := currNode.children[s]; ok {
			currNode = nextNode
		}
	}
	return currNode.route
}
