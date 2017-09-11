package parser

import (
	"github.com/demizer/go-rst/pkg/log"
	"github.com/demizer/go-rst/pkg/testutil"
	tok "github.com/demizer/go-rst/pkg/token"
	klog "github.com/go-kit/kit/log"
)

var (
	initialCapacity = 100
)

type tokenBuffer struct {
	index int
	token *tok.Item
	buf   []*tok.Item
	lex   *tok.Lexer
	log.Logger
}

func newTokenBuffer(l *tok.Lexer, logr klog.Logger, logCallDepth int) tokenBuffer {
	return tokenBuffer{
		buf:    make([]*tok.Item, initialCapacity),
		lex:    l,
		index:  -1, // Index is unset until next() is called
		Logger: log.NewLogger("buffer", true, logCallDepth, testutil.LogExcludes, logr),
	}
}

// append a new token to the buffer but do not set the index or current token. This function should not be used directly,
// instead use next(). Returns the index where the token was set, or -1 if nothing was set.
func (t *tokenBuffer) append(item *tok.Item) int {
	if item == nil {
		return -1
	}
	for i := 0; i < len(t.buf)-1; i++ {
		if t.buf[i] == nil {
			t.buf[i] = item
			return i
		}
	}
	// The buffer has reached capacity
	t.buf = append(t.buf, item)
	return len(t.buf) - 1
}

// backup shifts the token buf right one position.
func (t *tokenBuffer) backup() (tok *tok.Item) {
	if t.index > 0 {
		t.index--
	}
	t.token = t.buf[t.index]
	tok = t.token
	t.Msgr("buffer index item", "index", t.index, "token", t.token)
	return
}

// peekBack uses the token buf to "look back" a number of positions (pos). Looking back more positions than the
// Parser.token buf allows (3) will generate a panic.
func (t *tokenBuffer) peekBack(pos int) (tok *tok.Item) {
	if t.index-pos >= 0 {
		tok = t.buf[t.index-pos]
	}
	return
}

func (t *tokenBuffer) peekBackTo(item tok.Type) (tok *tok.Item) {
	for i := t.index - 1; i >= 0; i-- {
		if t.buf[i] != nil && t.buf[i].Type == item {
			if i >= 0 {
				tok = t.buf[i]
			}
			break
		}
	}
	return
}

// peek looks ahead in the token stream a number of positions (pos) and gets the next token from the lexer. A pointer to the
// token is kept in the Parser.token buf. If a token pointer already exists in the buf, that token is used instead
// and no buf are received the the lexer stream (channel).
func (t *tokenBuffer) peek(pos int) (pi *tok.Item) {
	for i := t.index + 1; i <= t.index+pos; i++ {
		if t.buf[i] != nil {
			pi = t.buf[i]
			continue
		} else {
			ind := t.append(t.lex.NextItem())
			if ind >= 0 {
				pi = t.buf[ind]
			}
		}
	}
	t.Msgr("peek token", "index", t.index, "token", pi)

	// XXX: remove this before merging to master
	// t.Dump(t.buf)
	// t.Msgr("haz index", "index", pos)

	return
}

// peekSkip looks ahead one position skipiing a specified itemElement. If that element is found, a pointer is returned,
// otherwise nil is returned.
func (t *tokenBuffer) peekSkip(iSkip tok.Type) *tok.Item {
	var nItem *tok.Item
	count := 1
	for {
		nItem = t.peek(count)
		if nItem.Type != iSkip {
			break
		}
		count++
	}
	return nItem
}

func (t *tokenBuffer) next(pos int) *tok.Item {
	if pos == 0 {
		return t.token
	}

	if t.index+1 < len(t.buf) && t.buf[t.index+1] != nil {
		t.index++
		t.token = t.buf[t.index]
	} else {
		if ind := t.append(t.lex.NextItem()); ind != -1 {
			t.Msgr("got token from lexer", "token", t.buf[ind])
			t.index = ind
			t.token = t.buf[t.index]
		}
	}

	pos--
	if pos > 0 {
		t.next(pos)
	}

	// XXX: Remove this before merging to master
	// t.Dump(t.buf)

	return t.token
}

func (t *tokenBuffer) peekLine(line int) (toks []*tok.Item) {
	for x := 0; x < len(t.buf)-1; x++ {
		if t.buf[x] != nil && t.buf[x].Line == line {
			toks = append(toks, t.buf[x])
		}
	}
	return
}

// clearTokens sets tokens from begin to end to nil.
func (t *tokenBuffer) clearTokens(begin, end int) {
	for i := begin; i <= end; i++ {
		t.buf[i] = nil
	}
}

func (t *tokenBuffer) indexFromToken(tok *tok.Item) int {
	for k, v := range t.buf {
		if tok == v {
			t.Msgr("indexFromToken found match", "index", k, "token", v)
			return k
		}
	}
	return -1
}
