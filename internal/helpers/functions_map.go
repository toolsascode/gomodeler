package helpers

import (
	"html/template"
	"strings"

	sprig "github.com/Masterminds/sprig/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var funcMapInternal = template.FuncMap{
	"arrayStr":         ToArrayString,
	"bool":             ToBool,
	"camelCase":        ToCamelCase,
	"capitalCase":      cases.Title(language.English, cases.Compact).String,
	"compare":          strings.Compare,
	"contains":         strings.Contains,
	"containsAny":      strings.ContainsAny,
	"count":            strings.Count,
	"dirName":          GetDirectory,
	"extension":        Extension,
	"fileName":         GetFileName,
	"hasPrefix":        strings.HasPrefix,
	"hasSuffix":        strings.HasSuffix,
	"include":          IncludeFile,
	"join":             strings.Join,
	"json":             ToJson,
	"kebabCaseLower":   ToKebabCaseLower,
	"kebabCaseUpper":   ToKebabCaseUpper,
	"lastIndex":        strings.LastIndex,
	"lowerCase":        strings.ToLower,
	"pascalCase":       ToPascalCase,
	"replace":          strings.Replace,
	"replaceAll":       strings.ReplaceAll,
	"snakeCaseLower":   ToSnakeCaseLower,
	"snakeCaseUpper":   ToSnakeCaseUpper,
	"split":            strings.Split,
	"splitAfter":       strings.SplitAfter,
	"splitAfterN":      strings.SplitAfterN,
	"splitN":           strings.SplitN,
	"sum":              SumFunc,
	"titleCase":        strings.ToTitle,
	"upperCase":        strings.ToUpper,
	"withoutExtension": WithoutExtension,
	"yaml":             ToYaml,
}

// Funcs internals and Sprig
// Reference: https://masterminds.github.io/sprig/
func GetFuncMap() template.FuncMap {
	return MergeMaps(funcMapInternal, sprig.GenericFuncMap())
}
