package main

import do "gopkg.in/godo.v2"

func tasks(p *do.Project) {
	p.Task("server", nil, func(c *do.Context) {
		c.Start("server.go", nil)
	}).Src("**/*.{go, gohtml}").Debounce(3 * 1000)
}

func main() {
	do.Godo(tasks)
}
