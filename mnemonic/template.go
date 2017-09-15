// Copyright (C) 2017 Billie Alice Thompson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package mnemonic

import (
	"fmt"
	"strings"

	"github.com/lloyd/wnram"
)

// parameterPrefix is the prefix to give elements in the template
const parameterPrefix = "Param"

// Template to be used to generate the mnemonic
type Template interface {
	GetTemplate() string
	GetParameters() map[string]string
	GetUsedFunctions() []string
}

// TemplateBase is a template to generate a mnemonic
type TemplateBase struct {
	template      string
	parameters    map[string]string
	usedFunctions []string
}

// NewTemplate returns a template to generate a mnemonic
//
// Might be used like this:
//   template := mnemonic.NewTemplate([]string{"e", "x", "a", "m", "p", "l", "e"})
func NewTemplate(letters []string) *TemplateBase {
	templateFragments := []string{}
	letterCount := len(letters)

	for i := 0; i < letterCount/4; i++ {
		templateFragments = append(templateFragments, generateUpTo4Template(4, i*4))
	}

	templateFragments = append(templateFragments, generateUpTo4Template(letterCount%4, 4*(letterCount/4)))

	return &TemplateBase{
		usedFunctions: newUsedFunctions(letters),
		parameters:    newParameterMap(letters),
		template:      strings.Trim(strings.Join(templateFragments, " "), " "),
	}
}

// newParameterMap turns a list of letters into parameters to use in the template
func newParameterMap(parameters []string) map[string]string {
	parameterMap := make(map[string]string)

	for i := range parameters {
		parameterMap[fmt.Sprintf("%s%d", parameterPrefix, i+1)] = parameters[i]
	}

	return parameterMap
}

// newUsedFunctions returns the parameters used in this template
func newUsedFunctions(parameters []string) []string {
	usedFunctions := []string{}
	availableFunc := availableFunctions()

	switch len(parameters) {
	case 0:
	case 1:
		usedFunctions = append(usedFunctions, availableFunc[1])
	case 2:
		usedFunctions = append(usedFunctions, availableFunc[:2]...)
	case 3:
		usedFunctions = append(usedFunctions, availableFunc[:3]...)
	default:
		usedFunctions = append(usedFunctions, availableFunc...)
	}

	return usedFunctions
}

// availableFunctions returns all the functions available to use, sorted in order of preference
func availableFunctions() []string {
	return []string{
		wnram.Adjective.String(),
		wnram.Noun.String(),
		wnram.Verb.String(),
		wnram.Adverb.String(),
	}
}

// generateUpTo4Template returns a template for up to 4 parameters, you can offset the parameter number too
func generateUpTo4Template(length int, paramOffset int) string {
	availableFunc := availableFunctions()

	switch length {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("{{ .%s%d | %s }}.", parameterPrefix, paramOffset+1, availableFunc[1])
	default:
		placeholderFragments := []string{}

		for i := 0; i < length; i++ {
			placeholderFragments = append(
				placeholderFragments,
				fmt.Sprintf(
					"{{ .%s%d | %s }}",
					parameterPrefix,
					paramOffset+i+1,
					availableFunc[i],
				),
			)
		}

		template := strings.Join(placeholderFragments, " ")

		return fmt.Sprintf("%s.", template)
	}
}

// GetTemplate returns a template string compatible with the go template engine
func (t TemplateBase) GetTemplate() string {
	return t.template
}

// GetParameters returns a map with the parameters for this template in
func (t TemplateBase) GetParameters() map[string]string {
	return t.parameters
}

// GetUsedFunctions returns the used functions
func (t TemplateBase) GetUsedFunctions() []string {
	return t.usedFunctions
}
