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

// StaticWordGenerator is a word generator that returns a static word
type StaticWordGenerator struct {
	word     string
	funcName string
}

// NewStaticWordGenerator returns a word generator that returns a static word
//
// Could be used like
//   mnemonic.NewWnramWordGenerator("soft", wnram.Adjective.String())
func NewStaticWordGenerator(word string, funcName string) *StaticWordGenerator {

	return &StaticWordGenerator{word: word, funcName: funcName}
}

// GetFuncName gets the function name
func (w *StaticWordGenerator) GetFuncName() string {
	return w.funcName
}

// Generate returns a single word
func (w *StaticWordGenerator) Generate(letter string) string {
	return w.word
}
