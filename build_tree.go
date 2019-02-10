package fields

import (
	"strings"
)

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

	current_level := 0

	var current_tree []string

	var last_field string

	for _, line := range strings.Split(strings.TrimSuffix(request, "\n"), "\n") {
		field := levelUpRegex.ReplaceAllString(line, "")
		field = levelDownRegex.ReplaceAllString(field, "")
		field = strings.TrimSpace(field)

		if field != "" {
			last_field = field

			path := strings.Join(current_tree, ".")

			tree[path] = append(tree[path], field)
		}

		if levelUpRegex.MatchString(line) {
			current_tree = append(current_tree, last_field)
			current_level += 1
		} else if levelDownRegex.MatchString(line) {
			current_tree = current_tree[:len(current_tree)-1]
			current_level -= 1
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
	fragment_name := fragmentNameStartRegex.ReplaceAllString(fragment, "")

	return "..." + fragmentNameEndRegex.ReplaceAllString(fragment_name, "")
}

func extractFragmentBody(fragment string) string {
	fragment_body := fragmentBodyStartRegex.ReplaceAllString(fragment, "")

	return fragmentBodyEndRegex.ReplaceAllString(fragment_body, "")
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
