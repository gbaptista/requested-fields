package fields

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func BuildHash(query string, variables map[string]interface{}, normalizeVariables bool) (string, map[string]string) {
	query = applyFragments(query, extractAndGroupFragments(query))
	query = removeFragments(query, extractAndGroupFragments(query))
	query = removeCommas(query)
	query = normalizeSpaces(query)

	query = regexp.MustCompile(`\n`).ReplaceAllString(query, " ")

	variablesMap := make(map[string]string)

	if normalizeVariables {
		variablesLetters := map[int]string{
			0: "$a", 1: "$b", 2: "$c", 3: "$d", 4: "$e",
			5: "$f", 6: "$g", 7: "$g", 8: "$g", 9: "$g"}

		index := 0

		for _, variable := range regexp.MustCompile(`\$.+?\b`).FindAllString(query, -1) {
			if _, exists := variablesMap[variable]; !exists {
				variablesMap[variable] = variablesLetters[index]
				index++
			}
		}

		for originalVariable, newVariable := range variablesMap {
			query = strings.ReplaceAll(
				query, originalVariable+")", newVariable+")")

			query = strings.ReplaceAll(
				query, originalVariable+":", newVariable+":")

			query = strings.ReplaceAll(
				query, originalVariable+" ", newVariable+" ")
		}
	}

	query = regexp.MustCompile(`query.*?\(`).ReplaceAllString(query, "query(")
	query = regexp.MustCompile(`\s{1,}`).ReplaceAllString(query, "")
	query = regexp.MustCompile(`__.*?\b`).ReplaceAllString(query, "")

	fmt.Printf("\n%s\n", query)
	

	letters := strings.Split(query, "")

	sort.Strings(letters)

	query = strings.Join(letters, "")

	fmt.Printf("\n%s\n", query)

	return fmt.Sprintf("%x", md5.Sum([]byte(query))), variablesMap
}
