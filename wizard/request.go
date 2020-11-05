package wizard

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/cjp2600/cutos/log"
	"github.com/manifoldco/promptui"
	"strings"
)

type Request struct {
	Description       string
	Tag               string
	Request           string
	RequestValidation func(input string) error
	Response          string
}

func NewRequest() *Request {
	return &Request{}
}

func (w *Request) SetRequestValidation(rv func(input string) error) {
	w.RequestValidation = rv
}

// SetTag
func (w *Request) SetTag() *Request {
	title := "Tag"
	if len(w.Tag) > 0 {
		title = title + " (...)"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.Tag = strings.ToLower(result)
	}
	return w
}

// SetDescription
func (w *Request) SetDescription() *Request {
	title := "Description"
	if len(w.Description) > 0 {
		title = title + " (...)"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.Description = result
	}
	return w
}

// SetRequest
func (w *Request) SetRequest(p string) *Request {
	fmt.Printf("%s Сopy data on your clipboard ... \n", promptui.Styler(promptui.FGYellow)("REQUEST ▶"))
	var text string
	for {
		if strings.Contains(text, p) {
			break
		}
		text, _ = clipboard.ReadAll()
	}
	w.Request = text
	fmt.Printf("%s copied \n", promptui.Styler(promptui.FGGreen)("✔"))
	return w
}

func (w *Request) isJSON(s string) bool {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return false
	}
	return true
}

// SetResponse
func (w *Request) SetResponse() *Request {
	fmt.Printf("%s Сopy json response on your clipboard ... \n", promptui.Styler(promptui.FGYellow)("RESPONSE ▶"))
	var text string
	for {
		if w.isJSON(text) {
			break
		}
		text, _ = clipboard.ReadAll()
	}
	w.Response = text
	fmt.Printf("%s copied \n", promptui.Styler(promptui.FGGreen)("✔"))
	return w
}
