package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value with quotes.
func (s *Scanner) ScanWithQuotes() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if ch == '\'' {
		tok, lit := s.scanStringInQuotes()
		return tok, fmt.Sprintf("'%s'", lit)
	} else if isLetter(ch) || isDigit(ch) {
		s.unread()
		return s.scanIdent()
	}

	return scanCharacters(ch)
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if ch == '\'' {
		return s.scanStringInQuotes()
	} else if isLetter(ch) || isDigit(ch) {
		s.unread()
		return s.scanIdent()
	}

	return scanCharacters(ch)
}

func scanCharacters(ch rune) (Token, string) {
	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '@':
		return AT, string(ch)
	case '+':
		return PLUS, string(ch)
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case ':':
		return COLON, string(ch)
	case '.':
		return DOT, string(ch)
	case '-':
		return HYPHEN, string(ch)
	case '/':
		return SLASH, string(ch)
	case '(':
		return LEFT_PARENTHESIS, string(ch)
	case ')':
		return RIGHT_PARENTHESIS, string(ch)
	case '[':
		return LEFT_BRACKET, string(ch)
	case ']':
		return RIGHT_BRACKET, string(ch)
	case '<':
		return LT, string(ch)
	case '>':
		return GT, string(ch)
	case '=':
		return EQ, string(ch)
	case '!':
		return EXCLAMATION_MARK, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanStringInQuotes consumes all runes until next quote is found.
func (s *Scanner) scanStringInQuotes() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '"' { // we found the end of string
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return IDENT, buf.String()
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read()) // buffer for the read string

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' { // allowed characters in identifier
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "NULL":
		return NULL, buf.String()
	case "FUNCTION":
		return FUNCTION, buf.String()
	case "EXTERNAL":
		return EXTERNAL, buf.String()
	case "FOR":
		return FOR, buf.String()
	case "SOME":
		return SOME, buf.String()
	case "EVERY":
		return EVERY, buf.String()
	case "RANGE":
		return RANGE, buf.String()
	case "LIST":
		return LIST, buf.String()
	case "CONTEXT":
		return CONTEXT, buf.String()
	case "IF":
		return IF, buf.String()
	case "THEN":
		return THEN, buf.String()
	case "ELSE":
		return ELSE, buf.String()
	case "AND":
		return AND, buf.String()
	case "OR":
		return OR, buf.String()
	case "IN":
		return IN, buf.String()
	case "SATISFIES":
		return SATISFIES, buf.String()
	case "BETWEEN":
		return BETWEEN, buf.String()
	case "TRUE":
		return TRUE, buf.String()
	case "FALSE":
		return FALSE, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}
