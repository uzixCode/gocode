package utils

import (
	"fmt"
	"strings"

	"github.com/uzixCode/gocode/models"
	"github.com/uzixCode/gocode/utils/changecase"
)

func GenerateStructsFromJSON(model models.Model) (string, error) {

	// Start creating the struct definition string
	var structDef strings.Builder
	var begin []string = []string{"package models", "import ("}
	if model.IsGorm {
		begin = append(begin, "\t\"gorm.io/gorm\"")
	}
	begin = append(begin, "\t)")
	begin = append(begin, fmt.Sprintf("type %s struct {", changecase.ToPascal(model.Name)))
	if model.IsGorm {
		begin = append(begin, "\tgorm.Model\n")

	}
	structDef.WriteString(strings.Join(begin, "\n"))

	// Iterate over fields and generate struct fields
	for _, field := range model.Fields {
		structDef.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" gorm:\"%s\"`\n",
			field.Name,
			field.DataType,
			field.JSONTag,
			field.GormTag,
		))
	}

	// Close the struct definition
	structDef.WriteString("}\n")
	return structDef.String(), nil
}
