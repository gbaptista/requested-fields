package fields

import (
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

	request = applyFragments(request, extractFragments(request))
	request = removeFragments(request)
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

			tree[path] = append(tree[path], field)
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

	return fragmentBodyEndRegex.ReplaceAllString(fragmentBody, "")
}

func extractFragments(request string) map[string]string {
	var fragments = make(map[string]string)

	for _, fragment := range fragmentsRegex.FindAllString(request, -1) {
		fragments[extractFragmentName(fragment)] = extractFragmentBody(fragment)
	}

	return fragments
}

func applyFragments(request string, fragments map[string]string) string {
	// TODO: while ... exists, with a max of loops.
	for i := 5; i > 0; i-- {
		for name, body := range fragments {
			request = strings.Replace(request, name, body, -1)
		}
	}

	return request
}

func removeFragments(request string) string {
	return fragmentsRegex.ReplaceAllString(request, "")
}

func removeQuery(request string) string {
	request = queryStartRegex.ReplaceAllString(request, "")
	return queryEndRegex.ReplaceAllString(request, "")
}
