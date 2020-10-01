package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New creates a new pointer to a
// Lexer with the input field already
// passed.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// read position allows us to peak ahead
	// and make sure we haven't hit the EOF yet
	if l.readPosition >= len(l.input) {
		// sets l.ch to 0 which is the ASCII code for "Null"
		l.ch = 0
	} else {
		// otherwise current character is equal to next character in input
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
