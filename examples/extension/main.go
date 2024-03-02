package main

import (
	"github.com/westcoastcode-se/gohtml/a"
	"github.com/westcoastcode-se/gohtml/h"
	"log"
	"net/http"
)

const InputText = "text"
const InputButton = "button"

// Input represents an input field
type Input struct {
	Id      string
	Name    string
	Type    string
	Heading string
}

// Label creates a HTML Label node and makes sure that it has the correct for attribute so that
// it points to the input field
func (i Input) Label() h.Node {
	return h.Label(
		a.For(i.Id),
		h.Text(i.Heading),
	)
}

// Node returns the actual input html node
func (i Input) Node() h.Node {
	return h.Input(
		a.ID(i.Id),
		a.Type(i.Type),
		a.Name(i.Name),
	)
}

// Form represents a builder pattern that will, eventually, create a form with input fields
type Form struct {
	Method     string
	URL        string
	Fields     []Input
	SubmitText string
}

// Build the entire form
func (f Form) Build() h.Node {
	return h.Form(
		a.Method(f.Method),
		a.Action(f.URL),
		h.Table(
			h.EmitArray(f.Fields, func(i Input) h.Node {
				return h.Tr(
					h.Td(
						i.Label(),
					),
					h.Td(
						i.Node(),
					),
				)
			}),
			h.Tr(
				h.Td(
					a.Colspan(2),
					h.Input(
						a.Type("submit"),
						a.Value("Login"),
					),
				),
			),
		),
	)
}

func index(w http.ResponseWriter, r *http.Request) {
	// generate the actual html
	_, _ = h.Html(a.Lang("en"),
		h.Head(
			h.Meta(a.Charset("UTF-8")),
			h.Title("Example: Extensions"),
		),
		h.Body(
			h.H1(
				h.Text("Example: Extensions"),
			),
			// Create a form using a custom builder
			Form{
				Method: http.MethodPost,
				URL:    "/",
				Fields: []Input{
					{
						Id:      "user",
						Name:    "username",
						Type:    "text",
						Heading: "Username",
					},
					{
						Id:      "pass",
						Name:    "password",
						Type:    "password",
						Heading: "Password",
					},
				},
			}.Build(),
		),
	)(w)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
