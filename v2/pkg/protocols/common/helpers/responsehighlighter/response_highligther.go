package responsehighlighter

import (
	"strings"

	"github.com/logrusorgru/aurora"

	"github.com/projectdiscovery/nuclei/v2/pkg/operators"
)

func Highlight(operatorResult *operators.Result, response string, noColor bool) string {
	result := response
	if operatorResult != nil && !noColor {
		for _, matches := range operatorResult.Matches {
			if len(matches) > 0 {
				colorizer := aurora.NewAurora(true)
				for _, currentMatch := range matches {
					result = strings.ReplaceAll(result, currentMatch, colorizer.Green(currentMatch).String())
				}
			}
		}
	}

	return result
}