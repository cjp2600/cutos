package wizard

import (
	"errors"
	"github.com/cjp2600/cutos/log"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/manifoldco/promptui"
)

type Info struct {
	meta   *openapi3.Info
	server openapi3.Servers
}

func NewInfo(meta *openapi3.Info, server openapi3.Servers) *Info {
	return &Info{meta: meta, server: server}
}

// requiredField
func (w *Info) requiredField() func(input string) error {
	return func(input string) error {
		if len(input) == 0 {
			err := errors.New("required field")
			return err
		}
		return nil
	}
}

// SetTitle
func (w *Info) SetTitle(isRequired bool) *Info {
	title := "Enter project name"
	if len(w.meta.Title) > 0 {
		title = title + " (" + w.meta.Title + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	if isRequired {
		prompt.Validate = w.requiredField()
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.meta.Title = result
	}
	return w
}

// SetAuthorName
func (w *Info) SetAuthorName() *Info {
	if w.meta.Contact == nil {
		w.meta.Contact = new(openapi3.Contact)
	}
	title := "Author Name"
	if len(w.meta.Contact.Name) > 0 {
		title = title + " (" + w.meta.Contact.Name + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if len(result) > 0 {
		w.meta.Contact.Name = result
	}
	return w
}

// SetAuthorEmail
func (w *Info) SetAuthorEmail() *Info {
	if w.meta.Contact == nil {
		w.meta.Contact = new(openapi3.Contact)
	}
	title := "Author Email"
	if len(w.meta.Contact.Email) > 0 {
		title = title + " (" + w.meta.Contact.Email + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if len(result) > 0 {
		w.meta.Contact.Email = result
	}
	return w
}

// SetAuthorEmail
func (w *Info) SetAuthorURL() *Info {
	if w.meta.Contact == nil {
		w.meta.Contact = new(openapi3.Contact)
	}
	title := "Author URL"
	if len(w.meta.Contact.URL) > 0 {
		title = title + " (" + w.meta.Contact.URL + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.meta.Contact.URL = result
	}
	return w
}

// SetLicenseName
func (w *Info) SetLicenseName() *Info {
	if w.meta.License == nil {
		w.meta.License = new(openapi3.License)
	}
	title := "License Name"
	if len(w.meta.License.Name) > 0 {
		title = title + " (" + w.meta.License.Name + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.meta.License.Name = result
	}
	return w
}

// SetLicenseURL
func (w *Info) SetLicenseURL() *Info {
	if w.meta.License == nil {
		w.meta.License = new(openapi3.License)
	}
	title := "License URL"
	if len(w.meta.License.URL) > 0 {
		title = title + " (" + w.meta.License.URL + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if len(result) > 0 {
		w.meta.License.URL = result
	}
	return w
}

// SetVersion
func (w *Info) SetBaseURL() *Info {
	prompt := promptui.Prompt{
		Label: "Set API URL",
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.server = openapi3.Servers{{URL: result}}
	}
	return w
}

// SetVersion
func (w *Info) SetVersion(isRequired bool) *Info {
	title := "Version"
	if len(w.meta.Version) > 0 {
		title = title + " (" + w.meta.Version + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	if isRequired {
		prompt.Validate = w.requiredField()
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.meta.Version = result
	}
	return w
}

// SetVersion
func (w *Info) SetDescription() *Info {
	title := "Description"
	if len(w.meta.Description) > 0 {
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
		w.meta.Description = result
	}
	return w
}

// SetTermsOfServices
func (w *Info) SetTermsOfService() *Info {
	title := "Terms Of Service"
	if len(w.meta.TermsOfService) > 0 {
		title = title + " (" + w.meta.TermsOfService + ")"
	}
	prompt := promptui.Prompt{
		Label: title,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		w.meta.TermsOfService = result
	}
	return w
}

func (w *Info) GetMeta() *openapi3.Info {
	return w.meta
}

func (w *Info) GetServers() openapi3.Servers {
	return w.server
}
