package templates

import "fmt"

import "nor/helper"

func Interface(name, face string) string {
	tmpl := `import { Document } from "mongoose";
	
export interface %s extends Document {%s}
`

	return fmt.Sprintf(tmpl, helper.Capitalize(name), face)
}