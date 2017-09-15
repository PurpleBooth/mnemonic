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
	"bufio"
	"html/template"
)

// TemplateParser interface returned by the NewTemplateParses
type TemplateParser interface {
	Parse(userTemplate string, input []string, writer *bufio.Writer) error
}

// TemplateParserBase is a parser that can convert a template into a string
type TemplateParserBase struct {
	funcMap template.FuncMap
}

// NewTemplateParser returns a new parser that can convert a template into a string
//
// Might be used like this
//  generator := mnemonic.NewTemplateParser(
//    mnemonic.NewWnramWordGenerator(wn, wnram.Adjective),
//    mnemonic.NewWnramWordGenerator(wn, wnram.Noun),
//    mnemonic.NewWnramWordGenerator(wn, wnram.Verb),
//    mnemonic.NewWnramWordGenerator(wn, wnram.Adverb),
//  )
func NewTemplateParser(
	generator ...WordGenerator,
) *TemplateParserBase {
	funcMap := template.FuncMap{}

	for i := range generator {
		funcMap[generator[i].GetFuncName()] = generator[i].Generate
	}

	return &TemplateParserBase{
		funcMap: funcMap,
	}
}

// Parse returns the parted template
//
// Might be used like this
//   err = generator.Parse(template, letters, bufio.NewWriter(os.Stdout))
// Or you can capture the string
//   buffer := &bytes.Buffer{}
//   writer := bufio.NewWriter(buffer)
//   _ = generator.Parse(template, letters, writer)
//   fmt.Println(buffer.String())
func (g *TemplateParserBase) Parse(userTemplate Template, input []string, writer *bufio.Writer) error {
	htmlTemplate := template.New("generator").Funcs(g.funcMap)
	templateParsed, err := htmlTemplate.Parse(userTemplate.GetTemplate())

	if err != nil {
		return err
	}

	err = templateParsed.Execute(writer, userTemplate.GetParameters())
	writer.Flush()

	if err != nil {
		return err
	}

	return nil
}
