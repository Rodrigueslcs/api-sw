package namespace

import (
	"fmt"
	"strings"
)

type Namespace struct {
	Root      string
	Component string
}

func New(root string) *Namespace {
	return &Namespace{
		Root: root,
	}
}

func (n *Namespace) AddComponet(name string) {
	n.Component = fmt.Sprintf("%s.%s", n.Root, name)
}

func (n *Namespace) Concat(paths ...string) string {
	if n.Component != "" {
		return strings.Join(append([]string{n.Component}, paths...), ".")
	}
	return strings.Join(append([]string{n.Root}, paths...), ".")
}
