package utils

import (
	"regexp"
	"strings"

	"github.com/uzixCode/gocode/models"
)

func ModelsTOJson(input string) (models.Model, error) {
	var structInfo models.Model
	lines := strings.Split(input, "\n")
	structNameRegex := regexp.MustCompile(`type\s+(\w+)\s+struct`)
	fieldRegex := regexp.MustCompile(`^\s*(\w+)\s+([\*\w\[\]]+)\s+` + "`(.*?)`")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Check for the struct name
		if structNameRegex.MatchString(line) {
			matches := structNameRegex.FindStringSubmatch(line)
			if len(matches) > 1 {
				structInfo.Name = matches[1]
			}
			continue
		}

		// Extract fields
		if fieldRegex.MatchString(line) {
			matches := fieldRegex.FindStringSubmatch(line)
			if len(matches) > 3 {
				fieldName := matches[1]
				dataType := matches[2]
				tags := matches[3]

				jsonTag, gormTag := "", ""
				tagRegex := regexp.MustCompile(`(\w+):"([^"]+)"`)
				tagMatches := tagRegex.FindAllStringSubmatch(tags, -1)
				for _, tagMatch := range tagMatches {
					if tagMatch[1] == "json" {
						jsonTag = tagMatch[2]
					} else if tagMatch[1] == "gorm" {
						gormTag = tagMatch[2]
					}
				}

				structInfo.Fields = append(structInfo.Fields, models.Field{
					Name:     fieldName,
					DataType: dataType,
					JSONTag:  jsonTag,
					GormTag:  gormTag,
				})
			}
		}
	}

	return structInfo, nil
}

// func ModelsTOJson(input string) (models.Model, error) {
// 	var structInfo models.Model
// 	lines := strings.Split(input, "\n")
// 	structNameRegex := regexp.MustCompile(`type\s+(\w+)\s+struct`)
// 	fieldRegex := regexp.MustCompile(`^\s*(\w+)\s+([\w\[\]]+)\s+` +
// 		"`(.*?)`")

// 	for _, line := range lines {
// 		line = strings.TrimSpace(line)

// 		// Check for the struct name
// 		if structNameRegex.MatchString(line) {
// 			matches := structNameRegex.FindStringSubmatch(line)
// 			if len(matches) > 1 {
// 				structInfo.Name = matches[1]
// 			}
// 			continue
// 		}

// 		// Extract fields
// 		if fieldRegex.MatchString(line) {
// 			matches := fieldRegex.FindStringSubmatch(line)
// 			if len(matches) > 3 {
// 				fieldName := matches[1]
// 				dataType := matches[2]
// 				tags := matches[3]

// 				jsonTag, gormTag := "", ""
// 				tagRegex := regexp.MustCompile(`(\w+):"([^"]+)"`)
// 				tagMatches := tagRegex.FindAllStringSubmatch(tags, -1)
// 				for _, tagMatch := range tagMatches {
// 					if tagMatch[1] == "json" {
// 						jsonTag = tagMatch[2]
// 					} else if tagMatch[1] == "gorm" {
// 						gormTag = tagMatch[2]
// 					}
// 				}

// 				structInfo.Fields = append(structInfo.Fields, models.Field{
// 					Name:     fieldName,
// 					DataType: dataType,
// 					JSONTag:  jsonTag,
// 					GormTag:  gormTag,
// 				})
// 			}
// 		}
// 	}

// 	return structInfo, nil
// }
