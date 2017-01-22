package parse

import (
	"unicode"
	"unicode/utf8"
)

// isSection compares a number of positions (skipping whitespace) to determine if the runes are sectionAdornments and returns
// a true if the positions match each other. Rune comparison begins at the current lexer position. isSection returns false if
// there is a blank line between the positions or if there is a rune mismatch between positions.
func isSection(l *lexer) bool {
	// Check two positions to see if the line contains a section adornment
	checkLine := func(input string) bool {
		var first, last rune
		for j := 0; j < len(input); j++ {
			r, _ := utf8.DecodeRuneInString(input[j:])
			if unicode.IsSpace(r) {
				logl.Msg("Skipping space rune")
				continue
			}
			if first == '\x00' {
				first = r
				last = r
			}
			// logl.Log.Debugf("first: %q, last: %q, r: %q, j: %d", first, last, r, j)
			if !isSectionAdornment(r) || (r != first && last != first) {
				logl.Msg("Section not found")
				return false
			}
			last = r
		}
		return true
	}

	if isTransition(l) {
		logl.Msg("Returning (found transition)")
		return false
	}

	logl.Msg("Checking current line")
	if checkLine(l.currentLine()) {
		logl.Msg("Found section adornment")
		return true
	}

	logl.Msg("Checking next line")

	nLine := l.peekNextLine()
	if nLine != "" {
		if checkLine(nLine) {
			logl.Msg("Found section adornment (nextline)")
			return true
		}
	}
	logl.Msg("Section not found")
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
func lexSection(l *lexer) stateFn {
	// logl.Log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width: %d, " + "l.line: %d", l.mark, l.index, l.start,
	// l.width, l.lineNumber())
	if isSectionAdornment(l.mark) {
		if l.lastItem != nil && l.lastItem.Type != itemTitle {
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

// lexTitle consumes input until newline and emits an itemTitle token. If spaces are detected at the start of the line, an
// itemSpace is emitted. Spaces after the title (and before newline) are ignored. On completion control is returned to
// lexSection.
func lexTitle(l *lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() {
			l.emit(itemTitle)
			break
		}
	}
	return lexSection
}

// lexSectionAdornment advances the lexer until a newline is encountered and emits a itemSectionAdornment token. Control is
// returned to lexSection() on completion.
func lexSectionAdornment(l *lexer) stateFn {
	for {
		if l.isEndOfLine() {
			l.emit(itemSectionAdornment)
			if l.mark == EOL {
				break
			}
		}
		l.next()
	}
	return lexSection
}