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

// BasicCmd
func BasicCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("file name is required. Example: $ cutos edit basic swagger.json [flags]")
	}
	fileName := strings.ToLower(strings.Trim(args[0], ""))

	loader := openapi3.NewSwaggerLoader()
	sw, err := loader.LoadSwaggerFromFile(fileName)
	if err != nil {
		return err
	}

	w := wizard.NewInfo(sw.Info, nil)
	w.SetTitle(true).
		SetVersion(true).
		SetDescription().
		SetTermsOfService().
		SetAuthorName().SetAuthorEmail().SetAuthorURL().
		SetLicenseName().SetLicenseURL()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(sw)
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
	fmt.Printf("%s Information updated \n", promptui.Styler(promptui.FGGreen)("✔"))
	return nil
}
