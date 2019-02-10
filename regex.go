package fields

import (
	"regexp"
)

var schemaRegex = regexp.MustCompile(`__schema`)

var paramsRegex = regexp.MustCompile(`\(.*?\)`)
var commasRegex = regexp.MustCompile(`,`)
var aliasRegex = regexp.MustCompile(`.*?:`)
var spacesRegex = regexp.MustCompile(`\s{1,}`)

var fragmentsRegex = regexp.MustCompile(`(?m)fragment(.|\n)*?}`)

var fragmentNameStartRegex = regexp.MustCompile(`(?m)fragment(\s|\n){1,}`)
var fragmentNameEndRegex = regexp.MustCompile(`(?m)\s{1,}on(.|\n)*`)
var fragmentBodyStartRegex = regexp.MustCompile(`(?m)(.|\n)*?{`)
var fragmentBodyEndRegex = regexp.MustCompile(`(?m)}(\s|\n)*\z`)

var queryStartRegex = regexp.MustCompile(`(?m)\A(.|\n)*?{`)
var queryEndRegex = regexp.MustCompile(`(?m)}(\s|\n)*\z`)

var levelUpRegex = regexp.MustCompile(`{`)
var levelDownRegex = regexp.MustCompile(`}`)
