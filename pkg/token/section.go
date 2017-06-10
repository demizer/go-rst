package token

import (
	"unicode"
	"unicode/utf8"
)

// isSection compares a number of positions (skipping whitespace) to determine if the runes are sectionAdornments and returns
// a true if the positions match each other. Rune comparison begins at the current lexer position. isSection returns false if
// there is a blank line between the positions or if there is a rune mismatch between positions.
func isSection(l *Lexer) bool {
	// Check two positions to see if the line contains a section adornment
	checkLine := func(input string) bool {
		var first, last rune
		for j := 0; j < len(input); j++ {
			r, _ := utf8.DecodeRuneInString(input[j:])
			if unicode.IsSpace(r) {
				l.Msg("Skipping space rune")
				continue
			}
			if first == '\x00' {
				first = r
				last = r
			}
			// l.Log.Debugf("first: %q, last: %q, r: %q, j: %d", first, last, r, j)
			if !isSectionAdornment(r) || (r != first && last != first) {
				l.Msg("Section not found")
				return false
			}
			last = r
		}
		return true
	}

	if isTransition(l) {
		l.Msg("Returning (found transition)")
		return false
	}

	l.Msg("Checking current line")
	if checkLine(l.currentLine()) {
		l.Msg("Found section adornment")
		return true
	}

	l.Msg("Checking next line")

	nLine := l.peekNextLine()
	if nLine != "" {
		if checkLine(nLine) {
			l.Msg("Found section adornment (nextline)")
			return true
		}
	}
	l.Msg("Section not found")
	return false
}

// isSectionAdornment returns true if r matches a section adornment.
func isSectionAdornment(r rune) bool {
	for _, a := range sectionAdornments {
		if a == r {
			return true
		}
	}
	return false
}

// lexSection is used after isSection() has determined that the next runes of input are section.  From here, the lexTitle()
// and lexSectionAdornment() are called based on the input.
func lexSection(l *Lexer) stateFn {
	// l.Log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width: %d, " + "l.line: %d", l.mark, l.index, l.start,
	// l.width, l.lineNumber())
	if isSectionAdornment(l.mark) {
		if l.lastItem != nil && l.lastItem.Type != Title {
			return lexSectionAdornment
		}
		lexSectionAdornment(l)
	} else if unicode.IsSpace(l.mark) {
		return lexSpace
	} else if l.mark == EOL {
		l.next()
	} else if unicode.IsPrint(l.mark) {
		return lexTitle
	}
	return lexStart
}

// lexTitle consumes input until newline and emits an Title token. If spaces are detected at the start of the line, an
// Space is emitted. Spaces after the title (and before newline) are ignored. On completion control is returned to
// lexSection.
func lexTitle(l *Lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() {
			l.emit(Title)
			break
		}
	}
	return lexSection
}

// lexSectionAdornment advances the lexer until a newline is encountered and emits a SectionAdornment token. Control is
// returned to lexSection() on completion.
func lexSectionAdornment(l *Lexer) stateFn {
	for {
		if l.isEndOfLine() {
			l.emit(SectionAdornment)
			if l.mark == EOL {
				break
			}
		}
		l.next()
	}
	return lexSection
}
