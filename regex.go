package fields

import (
	"regexp"
)

const (
	schemaRegexString = `__schema`

	trueIncludesRegexString  = `\s{1,}@.*true.*?{`
	falseIncludesRegexString = `\s{1,}@.*false.*?{`
	includesRegexString      = `\s{1,}@.*?{`

	paramsRegexString       = `(?sm)\(.*?\)`
	commasRegexString       = `,`
	aliasMergeRegexString   = `.*?:`
	aliasReplaceRegexString = `:\s{0,}\w{1,}\b`
	spacesRegexString       = `\s{1,}`

	fragmentsStartRegexString  = `(?sm)fragment(.|\n)*?{`
	partialFragmentRegexString = `\.\.\.`

	fragmentNameStartRegexString = `(?sm)fragment(\s|\n){1,}`
	fragmentNameEndRegexString   = `(?sm)\s{1,}on(.|\n)*`
	fragmentBodyStartRegexString = `(?sm)fragment(.|\n)*?{`
	fragmentBodyEndRegexString   = `(?sm)}(\s|\n)*\z`

	queryStartRegexString = `(?sm)\A(.|\n)*?{`
	queryEndRegexString   = `(?sm)}(\s|\n)*\z`

	levelUpRegexString   = `{`
	levelDownRegexString = `}`
)

var (
	schemaRegex        = regexp.MustCompile(schemaRegexString)
	trueIncludesRegex  = regexp.MustCompile(trueIncludesRegexString)
	falseIncludesRegex = regexp.MustCompile(falseIncludesRegexString)
	includesRegex      = regexp.MustCompile(includesRegexString)

	paramsRegex       = regexp.MustCompile(paramsRegexString)
	commasRegex       = regexp.MustCompile(commasRegexString)
	aliasMergeRegex   = regexp.MustCompile(aliasMergeRegexString)
	aliasReplaceRegex = regexp.MustCompile(aliasReplaceRegexString)
	spacesRegex       = regexp.MustCompile(spacesRegexString)

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
