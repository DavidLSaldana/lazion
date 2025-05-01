package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
// menu fields go here, will be iota'd
)

func main() {
	fmt.Println("Hello World!")
}

type model struct {
	inputs  []textinput.Model
	focused int
}
