package main

import (
	"fmt"
	"github.com/apcera/termtables"
)

func main() {
	// create table object
	t := termtables.CreateTable()
	// add headers
	t.AddHeaders("User", "Age")
	// add rows
	t.AddRow("", 18)
	t.AddRow("darjun", 30)
	// print rendered content
	// default mode is terminal
	t.SetModeTerminal()
	fmt.Println("Terminal: \n" + t.Render())

	/*// mode is HTML
	t.SetModeHTML()
	fmt.Println("HTMl: \n" + t.Render())

	// mode is Markdown
	t.SetModeMarkdown()
	fmt.Println("Markdown: \n" + t.Render())*/
}
