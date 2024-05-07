package go_scalar

import (
	"encoding/json"
	"fmt"
	"strings"
)

func safeJSONConfiguration(options *Options) string {
	// Serializa as opções para JSON
	jsonData, _ := json.Marshal(options)
	// Escapa as aspas duplas para entidades HTML
	escapedJSON := strings.ReplaceAll(string(jsonData), `"`, `&quot;`)
	return escapedJSON
}

func specContentHandler(specContent interface{}) string {
	switch spec := specContent.(type) {
	case func() map[string]interface{}:
		// Se specContent é uma função, chama a função e serializa o retorno
		result := spec()
		jsonData, _ := json.Marshal(result)
		return string(jsonData)
	case map[string]interface{}:
		// Se specContent é um mapa, serializa diretamente
		jsonData, _ := json.Marshal(spec)
		return string(jsonData)
	case string:
		// Se é uma string, retorna diretamente
		return spec
	default:
		// Caso contrário, retorna vazio
		return ""
	}
}

func ApiReferenceHTML(optionsInput *Options) (string, error) {
	options := DefaultOptions(*optionsInput)

	if options.SpecURL == "" && options.SpecContent == nil {
		return "", fmt.Errorf("specURL or specContent must be provided")
	}

	if options.SpecContent == nil && options.SpecURL != "" {
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
