package fields

import (
	"regexp"
)

const (
	schemaRegexString = `__schema`

	paramsRegexString = `\(.*?\)`
	commasRegexString = `,`
	aliasRegexString  = `.*?:`
	spacesRegexString = `\s{1,}`

	fragmentsRegexString       = `(?m)fragment(.|\n)*?}`
	partialFragmentRegexString = `\.\.\.`

	fragmentNameStartRegexString = `(?m)fragment(\s|\n){1,}`
	fragmentNameEndRegexString   = `(?m)\s{1,}on(.|\n)*`
	fragmentBodyStartRegexString = `(?m)(.|\n)*?{`
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
	aliasRegex  = regexp.MustCompile(aliasRegexString)
	spacesRegex = regexp.MustCompile(spacesRegexString)

	fragmentsRegex       = regexp.MustCompile(fragmentsRegexString)
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
