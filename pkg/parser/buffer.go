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

func newTokenBuffer(l *tok.Lexer, logr klog.Logger) tokenBuffer {
	return tokenBuffer{
		buf:    make([]*tok.Item, initialCapacity),
		lex:    l,
		index:  -1, // Index is unset until next() is called
		Logger: log.NewLogger("token_buffer", true, testutil.LogExcludes, logr),
	}
}

// append a new token to the buffer but do not set the index or current token. This function should not be used directly,
// instead use next(). Returns the index where the token was set.
func (t *tokenBuffer) append(item *tok.Item) int {
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
	t.Msgr("have t.index", "index", t.index)
	if t.index > 0 {
		t.index--
	}
	t.Msgr("have t.index", "index", t.index)
	t.token = t.buf[t.index]
	tok = t.token
	t.Msgr("buffer index item", "index", t.index, "token", t.token)
	return
}

// peekBack uses the token buf to "look back" a number of positions (pos). Looking back more positions than the
// Parser.token buf allows (3) will generate a panic.
func (t *tokenBuffer) peekBack(pos int) (tok *tok.Item) {
	if t.index-pos > 0 {
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
			// t.Msg("PONG22222222222222222222222")
			continue
		} else {
			ind := t.append(t.lex.NextItem())
			pi = t.buf[ind]
		}
		// t.Msg("PONG!!!!!!!!!!!!!!!!!!!!!!!")
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
	if t.token != nil && t.token.Type == tok.EOF {
		return t.token
	}
	t.index = t.append(t.lex.NextItem())
	t.token = t.buf[t.index]
	pos--
	if pos > 0 {
		t.next(pos)
	}

	// XXX: Remove this before merging to master
	// t.Dump(t.buf)

	return t.token
}
