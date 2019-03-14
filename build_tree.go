package fields

import (
	"sort"
	"strings"
)

/*
BuildTree receives an input string with a GraphQL request
and returns a map with requested fields for each field.

	This request:
		query {
		  users {
		    id
		    name
		  }
		}

	Will generates:
		:= map[string][]string{
			"":      []string{"users"},
			"users": []string{"id", "name"},
		}
*/
func BuildTree(request string) map[string][]string {
	tree := make(map[string][]string)

	if schemaRegex.MatchString(request) {
		return tree
	}

	request = applyFragments(request, extractAndGroupFragments(request))
	request = removeFragments(request, extractAndGroupFragments(request))
	request = removeParams(request)
	request = removeCommas(request)
	request = normalizeSpaces(request)
	request = removeAlias(request)
	request = removeQuery(request)

	currentLevel := 0

	var currentTree []string

	var lastField string

	for _, line := range strings.Split(strings.TrimSuffix(request, "\n"), "\n") {
		field := levelUpRegex.ReplaceAllString(line, "")
		field = levelDownRegex.ReplaceAllString(field, "")
		field = strings.TrimSpace(field)

		if field != "" {
			lastField = field

			path := strings.Join(currentTree, ".")

			tree[path] = appendIfMissing(tree[path], field)
		}

		if levelUpRegex.MatchString(line) {
			currentTree = append(currentTree, lastField)
			currentLevel++
		} else if levelDownRegex.MatchString(line) {
			currentTree = currentTree[:len(currentTree)-1]
			currentLevel--
		}
	}

	return tree
}

func removeParams(request string) string {
	return paramsRegex.ReplaceAllString(request, "")
}

func removeCommas(request string) string {
	return commasRegex.ReplaceAllString(request, "\n")
}

func removeAlias(request string) string {
	return aliasRegex.ReplaceAllString(request, "\n")
}

func normalizeSpaces(request string) string {
	return spacesRegex.ReplaceAllString(request, "\n")
}

func extractFragmentName(fragment string) string {
	fragmentName := fragmentNameStartRegex.ReplaceAllString(fragment, "")

	return "..." + fragmentNameEndRegex.ReplaceAllString(fragmentName, "")
}

func extractFragmentBody(fragment string) string {
	fragmentBody := fragmentBodyStartRegex.ReplaceAllString(fragment, "")

	fragmentBody = fragmentBodyEndRegex.ReplaceAllString(fragmentBody, "")

	return fragmentBody
}

func extractAndGroupFragments(request string) map[string]string {
	var fragments = make(map[string]string)

	for _, fragmentResult := range fragmentsStartRegex.FindAllStringIndex(request, -1) {
		openedLevels := 1
		closedLevels := 0

		currentIndex := fragmentResult[1]

		for openedLevels > closedLevels {
			currentIndex++

			if string(request[currentIndex]) == "}" {
				closedLevels++
			} else if string(request[currentIndex]) == "{" {
				openedLevels++
			}
		}

		fragmetText := request[fragmentResult[0] : currentIndex+1]

		fragmentName := extractFragmentName(fragmetText)

		fragments[fragmentName] = fragmetText
	}

	return fragments
}

const fragmentsLimit int = 1000

func applyFragments(request string, fragments map[string]string) string {
	var sortedFragments []string

	for name := range fragments {
		sortedFragments = append(sortedFragments, name)
	}

	sort.Sort(byLength(sortedFragments))

	continueApplying := true

	replacedFragments := 0

	for continueApplying {
		if partialFragmentRegex.MatchString(request) {
			for _, name := range sortedFragments {
				request = strings.Replace(
					request, name, extractFragmentBody(fragments[name]), -1)
			}

			replacedFragments++

			if replacedFragments > fragmentsLimit {
				continueApplying = false
			}
		} else {
			continueApplying = false
		}
	}

	return request
}

func removeFragments(request string, fragments map[string]string) string {
	for _, fragmentText := range fragments {
		request = strings.Replace(request, fragmentText, "", -1)
	}

	return request
}

func removeQuery(request string) string {
	request = queryStartRegex.ReplaceAllString(request, "")
	return queryEndRegex.ReplaceAllString(request, "")
}

func appendIfMissing(tree []string, fieldToAppend string) []string {
	for _, field := range tree {
		if field == fieldToAppend {
			return tree
		}
	}
	return append(tree, fieldToAppend)
}
