package interactor

import (
	"fmt"
	"github.com/cjp2600/cutos/log"
	"github.com/cjp2600/cutos/parser"
	"github.com/cjp2600/cutos/wizard"
	"github.com/getkin/kin-openapi/openapi3"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// ListeningCmd Listen to the buffer to process the content
func ListeningCmd(cmd *cobra.Command, args []string, parserType parser.Type) error {
	if len(args) == 0 {
		return fmt.Errorf("file name is required. Example: $ cutos add %s swagger.json [flags]", parserType)
	}
	fileName := strings.ToLower(strings.Trim(args[0], ""))

	loader := openapi3.NewSwaggerLoader()
	fromFile, err := loader.LoadSwaggerFromFile(fileName)
	if err != nil {
		return err
	}

	// get parser
	newParser, err := parser.NewParser(parserType)
	if err != nil {
		return err
	}

	// input data set wizardRequest
	wizardRequest := wizard.NewRequest()
	wizardRequest.SetRequestValidation(newParser.RequiredField())
	wizardRequest.SetRequest(string(parserType))
	wizardRequest.SetResponse()
	wizardRequest.SetDescription()
	wizardRequest.SetTag()

	// set source string
	newParser.SetSource(wizardRequest)
	// set swagger
	newParser.SetSwagger(fromFile)
	// update and parse
	updatedSw := newParser.BuildPathMethod()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(updatedSw)
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
		return err
	}
	return ListeningCmd(cmd, args, parserType)
}
