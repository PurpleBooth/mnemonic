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
	"github.com/lloyd/wnram"
	"math/rand"
)

// WnramWordGenerator is a word generator that pulls random words from a WordNet dictionary
type WnramWordGenerator struct {
	wordList     map[string][]wnram.Lookup
	partOfSpeech wnram.PartOfSpeech
}

// NewWnramWordGenerator returns a word generator that pulls random words from a WordNet dictionary
//
// See the library http://github.com/lloyd/wnram
//
// Get dictionary files from http://wordnet.princeton.edu/
//
// Could be used like
//   wn, _ := wnram.New(dictDir)
//   mnemonic.NewWnramWordGenerator(wn, wnram.Adjective)
func NewWnramWordGenerator(wn *wnram.Handle, partOfSpeech wnram.PartOfSpeech) *WnramWordGenerator {
	wordList := make(map[string][]wnram.Lookup)

	wn.Iterate(wnram.PartOfSpeechList{partOfSpeech}, func(word wnram.Lookup) error {
		firstLetter := getCharAt(word.Word(), 0)

		wordList[firstLetter] = append(wordList[firstLetter], word)
		return nil
	})

	return &WnramWordGenerator{wordList: wordList, partOfSpeech: partOfSpeech}
}

// GetFuncName the function name
func (w *WnramWordGenerator) GetFuncName() string {
	return w.partOfSpeech.String()
}

// Generate returns a random word beginning with a given letter
func (w *WnramWordGenerator) Generate(letter string) string {

	for {
		randomWord := w.wordList[letter][rand.Intn(len(w.wordList[letter]))].Word()

		if getCharAt(randomWord, 0) == letter {
			return randomWord
		}
	}
}
