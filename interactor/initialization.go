package interactor

import (
	"errors"
	"fmt"
	"github.com/cjp2600/cutos/log"
	"github.com/cjp2600/cutos/wizard"
	"github.com/getkin/kin-openapi/openapi3"
	jsoniter "github.com/json-iterator/go"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// InitializationCmd
func InitializationCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("file name is required. Example: $ cutos init swagger.json [flags]")
	}
	fileName := strings.ToLower(strings.Trim(args[0], ""))

	// common wizard
	w := wizard.NewInfo(new(openapi3.Info))
	w.SetTitle(true).SetVersion(true).SetDescription().SetTermsOfService().
		SetAuthorName().SetAuthorEmail().SetAuthorURL().
		SetLicenseName().SetLicenseURL()

	swagger := BuildSwagger(w)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(swagger)
	if err != nil {
		return err
	}

	fo, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := fo.Write(b); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s Created file %s \n", promptui.Styler(promptui.FGGreen)("✔"), fileName)
	return nil
}

func BuildSwagger(w *wizard.Info) *openapi3.Swagger{
	return &openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: w.GetMeta(),
	}
}
