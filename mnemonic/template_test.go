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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/purplebooth/mnemonic/mnemonic"

	"fmt"
)

var _ = Describe("TemplateParser", func() {
	Context("No characters", func() {
		It("Returns empty template", func() {
			actual := NewTemplate([]string{})

			Expect(actual.GetTemplate()).To(Equal(""))
		})
	})
	Context("Base templates", func() {
		It("Returns a noun", func() {
			actual := NewTemplate([]string{"a"})

			Expect(actual.GetTemplate()).To(Equal("{{ .Param1 | noun }}."))
		})
		It("Returns a adjective then a noun", func() {
			actual := NewTemplate([]string{"a", "b"})

			Expect(actual.GetTemplate()).To(Equal("{{ .Param1 | adj }} {{ .Param2 | noun }}."))
		})
		It("Returns a adjective, a noun and then a verb", func() {
			actual := NewTemplate([]string{"a", "b", "c"})

			Expect(actual.GetTemplate()).To(Equal("{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }}."))
		})
		It("Returns a adjective, a noun, a verb and then a adverb", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d"})

			Expect(actual.GetTemplate()).To(Equal("{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}."))
		})
	})
	Context("Loops beyond 4 characters", func() {
		It("5 characters", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e"})

			Expect(actual.GetTemplate()).To(
				Equal("{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}. {{ .Param5 | noun }}."),
			)
		})
		It("6 characters", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e", "f"})

			Expect(actual.GetTemplate()).To(
				Equal("{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}. {{ .Param5 | adj }} {{ .Param6 | noun }}."),
			)
		})
		It("7 characters", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e", "f", "g"})

			Expect(actual.GetTemplate()).To(Equal(
				"{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}. {{ .Param5 | adj }} {{ .Param6 | noun }} {{ .Param7 | verb }}.",
			))
		})
		It("8 characters", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e", "f", "g", "h"})

			Expect(actual.GetTemplate()).To(Equal(
				"{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}. {{ .Param5 | adj }} {{ .Param6 | noun }} {{ .Param7 | verb }} {{ .Param8 | adv }}.",
			))
		})
		It("9 characters", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

			Expect(actual.GetTemplate()).To(Equal(
				"{{ .Param1 | adj }} {{ .Param2 | noun }} {{ .Param3 | verb }} {{ .Param4 | adv }}. {{ .Param5 | adj }} {{ .Param6 | noun }} {{ .Param7 | verb }} {{ .Param8 | adv }}. {{ .Param9 | noun }}.",
			))
		})
	})
	Context("Get functions", func() {
		It("Returns nothing for nothing", func() {
			actual := NewTemplate([]string{})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{}))
		})
		It("Returns a noun", func() {
			actual := NewTemplate([]string{"a"})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{"noun"}))
		})
		It("Returns a noun and adjective", func() {
			actual := NewTemplate([]string{"a", "b"})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{"adj", "noun"}))
		})
		It("Returns a noun, adjective, verb", func() {
			actual := NewTemplate([]string{"a", "b", "c"})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{"adj", "noun", "verb"}))
		})
		It("Returns a noun, adjective, verb, adverb", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d"})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{"adj", "noun", "verb", "adv"}))
		})
		It("Returns a noun, adjective, verb, adverb when more then 5", func() {
			actual := NewTemplate([]string{"a", "b", "c", "d", "e"})

			Expect(actual.GetUsedFunctions()).To(Equal([]string{"adj", "noun", "verb", "adv"}))
		})
	})
	Context("Get parameters", func() {
		It("No parameters", func() {
			actual := NewTemplate([]string{})

			parameters := make(map[string]string)

			Expect(actual.GetParameters()).To(Equal(parameters))
		})
		It("Param1 => a", func() {
			actual := NewTemplate([]string{"a"})

			parameters := make(map[string]string)
			parameters["Param1"] = "a"

			Expect(actual.GetParameters()).To(Equal(parameters))
		})
		It("Param1 => a, Param2 => b", func() {
			actual := NewTemplate([]string{"a", "b"})

			parameters := make(map[string]string)
			parameters["Param1"] = "a"
			parameters["Param2"] = "b"

			Expect(actual.GetParameters()).To(Equal(parameters))
		})
		It("Param1 => a, Param2 => b, Param3 => c", func() {
			actual := NewTemplate([]string{"a", "b", "c"})

			parameters := make(map[string]string)
			parameters["Param1"] = "a"
			parameters["Param2"] = "b"
			parameters["Param3"] = "c"

			Expect(actual.GetParameters()).To(Equal(parameters))
		})
	})
})

func ExampleNewTemplate() {
	NewTemplate([]string{"a"})
}

func ExampleTemplateBase_GetTemplate() {
	actual := NewTemplate([]string{"a"})
	fmt.Println(actual.GetTemplate())
	// Output: {{ .Param1 | noun }}.
}

func ExampleTemplateBase_GetUsedFunctions() {
	actual := NewTemplate([]string{"a"})
	fmt.Println(actual.GetUsedFunctions()[0])
	// Output: noun
}

func ExampleTemplateBase_GetParameters() {
	actual := NewTemplate([]string{"a"})

	for key, value := range actual.GetParameters() {
		fmt.Println("Key:", key, "Value:", value)
	}
	// Output: Key: Param1 Value: a
}
