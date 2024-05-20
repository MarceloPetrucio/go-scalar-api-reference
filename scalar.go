package scalar

import (
	"encoding/json"
	"fmt"
	"strings"
)

func safeJSONConfiguration(options *Options) string {
	// Serializes the options to JSON
	jsonData, _ := json.Marshal(options)
	// Escapes double quotes into HTML entities
	escapedJSON := strings.ReplaceAll(string(jsonData), `"`, `&quot;`)
	return escapedJSON
}

func specContentHandler(specContent interface{}) string {
	switch spec := specContent.(type) {
	case func() map[string]interface{}:
		// If specContent is a function, it calls the function and serializes the return
		result := spec()
		jsonData, _ := json.Marshal(result)
		return string(jsonData)
	case map[string]interface{}:
		// If specContent is a map, it serializes it directly
		jsonData, _ := json.Marshal(spec)
		return string(jsonData)
	case string:
		// If it is a string, it returns directly
		return spec
	default:
		// Otherwise, returns empty
		return ""
	}
}

func ApiReferenceHTML(optionsInput *Options) (string, error) {
	options := DefaultOptions(*optionsInput)

	if options.SpecURL == "" && options.SpecContent == nil {
		return "", fmt.Errorf("specURL or specContent must be provided")
	}

	if options.SpecContent == nil && options.SpecURL != "" {

		if strings.HasPrefix(options.SpecURL, "http") {
			content, err := fetchContentFromURL(options.SpecURL)
			if err != nil {
				return "", err
			}
			options.SpecContent = content
		} else {
			urlPath, err := ensureFileURL(options.SpecURL)
			if err != nil {
				return "", err
			}

			content, err := readFileFromURL(urlPath)
			if err != nil {
				return "", err
			}

			options.SpecContent = string(content)
		}
	}

	dataConfig := safeJSONConfiguration(options)
	specContentHTML := specContentHandler(options.SpecContent)

	var pageTitle string

	if options.CustomOptions.PageTitle != "" {
		pageTitle = options.CustomOptions.PageTitle
	} else {
		pageTitle = "Scalar API Reference"
	}

	customThemeCss := CustomThemeCSS

	if options.Theme != "" {
		customThemeCss = ""
	}

	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
      <head>
        <title>%s</title>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <style>%s</style>
      </head>
      <body>
        <script id="api-reference" type="application/json" data-configuration="%s">%s</script>
        <script src="%s"></script>
      </body>
    </html>
  `, pageTitle, customThemeCss, dataConfig, specContentHTML, options.CDN), nil
}
