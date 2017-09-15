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

package mnemonic_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/purplebooth/mnemonic/mnemonic"
)

type testTemplate struct {
	template      string
	parameters    map[string]string
	usedFunctions []string
}

// returns a template string compatible with the go template engine
func (t testTemplate) GetTemplate() string {
	return t.template
}

// returns a map with the parameters for this template in
func (t testTemplate) GetParameters() map[string]string {
	return t.parameters
}

// returns the used functions
func (t testTemplate) GetUsedFunctions() []string {
	return t.usedFunctions
}

type testWordGenerator struct {
	funcName string
	function func(letter string) string
}

func (w *testWordGenerator) Generate(letter string) string {
	return w.function(letter)
}

func (w *testWordGenerator) GetFuncName() string {
	return w.funcName
}

var _ = Describe("TemplateParser", func() {
	Context("Generating", func() {
		It("Returns without parameters or arguments", func() {
			parser := NewTemplateParser()
			actual := &bytes.Buffer{}
			writer := bufio.NewWriter(actual)

			template := testTemplate{
				template: "Testing",
			}

			parser.Parse(template, []string{}, writer)
			Expect(actual.String()).To(Equal("Testing"))
		})
		It("You can set parameters", func() {
			parser := NewTemplateParser()
			actual := &bytes.Buffer{}
			writer := bufio.NewWriter(actual)

			parameters := make(map[string]string)
			parameters["Param1"] = "a"
			parameters["Param2"] = "b"

			template := testTemplate{
				template:   "{{ .Param1 }} {{ .Param2 }}",
				parameters: parameters,
			}

			parser.Parse(template, []string{}, writer)
			Expect(actual.String()).To(Equal("a b"))
		})
		It("You can set use methods", func() {
			parser := NewTemplateParser(&testWordGenerator{
				funcName: "upper",
				function: strings.ToUpper,
			})
			actual := &bytes.Buffer{}
			writer := bufio.NewWriter(actual)

			parameters := make(map[string]string)

			template := testTemplate{
				template:      "{{ \"Testing\" | upper }}",
				parameters:    parameters,
				usedFunctions: []string{"upper"},
			}

			parser.Parse(template, []string{}, writer)
			Expect(actual.String()).To(Equal("TESTING"))
		})
		It("You can set parameters and methods", func() {
			parser := NewTemplateParser(&testWordGenerator{
				funcName: "upper",
				function: strings.ToUpper,
			})
			actual := &bytes.Buffer{}
			writer := bufio.NewWriter(actual)

			parameters := make(map[string]string)
			parameters["Param1"] = "a"
			parameters["Param2"] = "b"

			template := testTemplate{
				template:   "{{ .Param1 | upper }} {{ .Param2 | upper }}",
				parameters: parameters,
			}

			parser.Parse(template, []string{}, writer)
			Expect(actual.String()).To(Equal("A B"))
		})
	})
})

func ExampleNewTemplateParser() {
	NewTemplateParser(
		NewStaticWordGenerator("dancing", "adj"),
		NewStaticWordGenerator("eggs", "noun"),
		NewStaticWordGenerator("move", "verb"),
		NewStaticWordGenerator("outward", "adv"),
	)
}

func ExampleTemplateParserBase_Parse() {
	letters := strings.Split("demo", "")
	template := NewTemplate(letters)
	generator := NewTemplateParser(
		NewStaticWordGenerator("dancing", "adj"),
		NewStaticWordGenerator("eggs", "noun"),
		NewStaticWordGenerator("move", "verb"),
		NewStaticWordGenerator("outward", "adv"),
	)

	buffer := &bytes.Buffer{}
	writer := bufio.NewWriter(buffer)
	err := generator.Parse(template, letters, writer)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(buffer.String())
	// Output: dancing eggs move outward.
}
