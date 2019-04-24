package fields

import (
	"regexp"
)

const (
	schemaRegexString = `__schema`

	paramsRegexString = `\(.*?\)`
	commasRegexString = `,`
	aliasMergeRegexString  = `.*?:`
	aliasReplaceRegexString  = `:\s{0,}\w{1,}\b`
	spacesRegexString = `\s{1,}`

	fragmentsStartRegexString  = `(?m)fragment(.|\n)*?{`
	partialFragmentRegexString = `\.\.\.`

	fragmentNameStartRegexString = `(?m)fragment(\s|\n){1,}`
	fragmentNameEndRegexString   = `(?m)\s{1,}on(.|\n)*`
	fragmentBodyStartRegexString = `(?m)fragment(.|\n)*?{`
	fragmentBodyEndRegexString   = `(?m)}(\s|\n)*\z`

	queryStartRegexString = `(?m)\A(.|\n)*?{`
	queryEndRegexString   = `(?m)}(\s|\n)*\z`

	levelUpRegexString   = `{`
	levelDownRegexString = `}`
)

var (
	schemaRegex = regexp.MustCompile(schemaRegexString)

	paramsRegex = regexp.MustCompile(paramsRegexString)
	commasRegex = regexp.MustCompile(commasRegexString)
	aliasMergeRegex  = regexp.MustCompile(aliasMergeRegexString)
	aliasReplaceRegex  = regexp.MustCompile(aliasReplaceRegexString)
	spacesRegex = regexp.MustCompile(spacesRegexString)

	fragmentsStartRegex  = regexp.MustCompile(fragmentsStartRegexString)
	partialFragmentRegex = regexp.MustCompile(partialFragmentRegexString)

	fragmentNameStartRegex = regexp.MustCompile(fragmentNameStartRegexString)
	fragmentNameEndRegex   = regexp.MustCompile(fragmentNameEndRegexString)
	fragmentBodyStartRegex = regexp.MustCompile(fragmentBodyStartRegexString)
	fragmentBodyEndRegex   = regexp.MustCompile(fragmentBodyEndRegexString)

	queryStartRegex = regexp.MustCompile(queryStartRegexString)
	queryEndRegex   = regexp.MustCompile(queryEndRegexString)

	levelUpRegex   = regexp.MustCompile(levelUpRegexString)
	levelDownRegex = regexp.MustCompile(levelDownRegexString)
)
