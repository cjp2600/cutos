package parser

import (
	"errors"
	"mvdan.cc/xurls/v2"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/google/shlex"
	jsoniter "github.com/json-iterator/go"
)

// Curl structure
type Curl struct {
	text string
}

func (c *Curl) Text() string {
	return c.text
}

// NewCurl processing constructor
func NewCurl() *Curl {
	return &Curl{}
}

// SetText
func (c *Curl) SetText(text string) {
	c.text = text
}

// Split divide the transmitted messages into commands.
func (c *Curl) Split() ([]string, error) {
	split, err := shlex.Split(c.text)
	if err != nil {
		return nil, err
	}
	return split, nil
}

// RequiredField validation field
func (c *Curl) RequiredField() func(input string) error {
	return func(input string) error {
		if len(input) == 0 {
			err := errors.New("required field")
			return err
		}
		if !strings.Contains(input, "curl") {
			err := errors.New("its not curl scheme")
			return err
		}
		return nil
	}
}

// isValidUUID
func (c *Curl) isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// isObjectID
func (c *Curl) isObjectID(objectId string) bool {
	r := regexp.MustCompile(`(?m)(\{){0,1}[a-f\d]{24}(\}){0,1}`)
	return r.MatchString(objectId)
}

// isNumeric
func (c *Curl) isNumeric(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// SetUniqueName
func (c *Curl) SetUniqueName(path *Path) {
	var b strings.Builder
	items := strings.Split(path.TemplatePath, "/")
	for _, item := range items {
		if len(item) > 0 {
			if string(item[0]) != "{" {
				b.WriteString(strings.Replace(strings.Title(item), "-", "", 50))
			}
		}
	}
	b.WriteString(strings.Title(path.Method))
	path.UniqueName = b.String()
}

// Unmarshal
func (c *Curl) Unmarshal(b string) (map[string]interface{}, error) {
	var j map[string]interface{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal([]byte(b), &j); err != nil {
		return nil, err
	}
	return j, nil
}

// CreateTemplatePath - trying to identify variables on the path
func (c *Curl) CreateTemplatePath(path *Path) string {
	var tp []string
	stp := strings.Split(path.URL.Path, "/")
	for i, item := range stp {
		if c.isValidUUID(item) || c.isNumeric(item) || c.isObjectID(item) {
			uName := "{ID}"
			if i > 0 {
				p := pluralize.NewClient()
				variableName := strings.Title(p.Singular(stp[i-1]))
				uName = "{" + variableName + "ID}"
				var currentType = "string"
				if c.isValidUUID(item) {
					currentType = "uuid"
				}
				if c.isNumeric(item) {
					currentType = "integer"
				}
				path.PathVariables = append(path.PathVariables, &Variable{
					name:        variableName + "ID",
					varType:     currentType,
					description: variableName + " ID",
				})
			}
			tp = append(tp, uName)
		} else {
			tp = append(tp, item)
		}
	}
	return strings.Join(tp, "/")
}

func (c *Curl) FindType(s string) string {
	if _, err := strconv.Atoi(s); err == nil {
		return "integer"
	}
	return "string"
}

// Parse process the transferred data and return the path.
func (c *Curl) Parse() (*Path, error) {
	var response = new(Path)
	var setUrl bool
	var setMethod bool

	response.Headers = make(map[string]string)
	response.Method = "GET"

	splitItems, err := c.Split()
	if err != nil {
		return nil, err
	}

	setURL := func(resp *Path, nextItem string) error {
		resp.SourceURL = nextItem
		recognized, err := url.Parse(resp.SourceURL)
		if err != nil {
			return err
		}
		resp.URL = recognized

		if len(recognized.Query()) > 0 {
			for k, v := range recognized.Query() {
				var vl string
				if len(v) > 0 {
					vl = v[0]
				}
				resp.QueryVariables = append(resp.QueryVariables, &Variable{
					name:    k,
					varType: c.FindType(vl),
					Example: vl,
				})
			}
		}
		return nil
	}

	for i, item := range splitItems {

		var currentItem = strings.ToLower(c.clean(item))

		if c.isDataItem(currentItem) {
			var nextItem = splitItems[i+1]
			response.SourceRequest = strings.Trim(c.clean(nextItem), "$")
			parserMap, err := c.Unmarshal(response.SourceRequest)
			if err != nil {
				return nil, err
			}
			response.ParseRequest = parserMap
			continue
		}

		if c.isUrlItem(currentItem) {
			var nextItem = splitItems[i+1]
			setUrl = true
			err := setURL(response, nextItem)
			if err != nil {
				return response, err
			}
			response.TemplatePath = c.CreateTemplatePath(response)
			continue
		}

		if c.isRequestItem(currentItem) {
			var nextItem = splitItems[i+1]
			response.Method = strings.ToUpper(nextItem)
			setMethod = true
			continue
		}

		if c.isHeaderItem(currentItem) {
			var nextItem = splitItems[i+1]
			headerString := strings.ToLower(c.clean(nextItem))
			headerValues := strings.Split(headerString, ":")
			if len(headerValues) > 1 {
				response.Headers[headerValues[0]] = strings.ToLower(c.clean(headerValues[1]))
			}
			continue
		}
	}

	if !setUrl {
		for i, item := range splitItems {
			var nextItem = splitItems[i+1]
			bItem := strings.ToLower(c.clean(item))
			if c.isCurlItem(bItem) {
				rxRelaxed := xurls.Relaxed()
				findUrl := rxRelaxed.FindString(c.Text())
				if findUrl != "" {
					nextItem = findUrl
				}

				setUrl = true
				if err := setURL(response, nextItem); err != nil {
					return response, err
				}
				response.TemplatePath = c.CreateTemplatePath(response)
				break
			}
		}
	}

	if !setMethod {
		for _, item := range splitItems {
			bItem := strings.ToLower(c.clean(item))
			if c.isDataItem(bItem) {
				response.Method = "POST"
				break
			}
		}
	}

	c.SetUniqueName(response)
	return response, nil
}

// isCurlItem
func (c *Curl) isCurlItem(bItem string) bool {
	return strings.Contains(bItem, "curl")
}

func (c *Curl) isLocation(currentItem string) bool {
	return strings.Contains(currentItem, "--location")
}

// isHeaderItem
func (c *Curl) isHeaderItem(currentItem string) bool {
	return strings.Contains(currentItem, "--header") || strings.Contains(currentItem, "-h")
}

// isRequestItem
func (c *Curl) isRequestItem(currentItem string) bool {
	return strings.Contains(currentItem, "--request") || strings.Contains(currentItem, "-x")
}

// isUrlItem
func (c *Curl) isUrlItem(bItem string) bool {
	return strings.Contains(bItem, "--url")
}

// clean
func (c *Curl) clean(item string) string {
	return strings.Trim(item, " ")
}

// isDataItem - validate curl data
func (c *Curl) isDataItem(item string) bool {
	var dataArg = []string{"--data-binary", "--data", "--d", "--data-raw"}
	var isData bool
	for i := 0; i < len(dataArg); i++ {
		if strings.Contains(item, dataArg[i]) {
			isData = true
		}
	}
	return isData
}
