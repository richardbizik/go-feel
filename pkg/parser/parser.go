package parser

import (
	"fmt"
	"io"
)

// Parser represents a parser.
// Holy fuck who wrote this grammar
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) ParseFeelExpression() (*FeelExpression, error) {
	expression, err := p.parseExpression()

	return &FeelExpression{
		Expression: &expression,
	}, err
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit
	return
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scanWithQuotes() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.ScanWithQuotes()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scanFieldPath() string {
	tok, lit := p.scan()
	if tok == WS {
		return ""
	}
	if tok == IDENT || tok == LEFT_BRACKET || tok == RIGHT_BRACKET || tok == SLASH || tok == HYPHEN || tok == DOT {
		return lit + p.scanFieldPath()
	}
	p.unscan()
	return ""
}

func (p *Parser) parseExpression() (Expression, error) {
	expression := Expression{}
	tok, _ := p.scanIgnoreWhitespace()
	p.unscan()
	if tok == LEFT_BRACKET {
		boxedExpression, err := p.parseBoxedExpression()
		if err != nil {
			return expression, err
		}
		expression.BoxedExpression = boxedExpression
		return expression, nil
	}
	textualExpression, err := p.parseTextualExpression()
	if err != nil {
		return expression, err
	}
	expression.TextualExpression = textualExpression
	return expression, nil
}

func (p *Parser) parseTextualExpression() (*TextualExpression, error) {
	textualExpression := TextualExpression{}
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case FOR:
		forExpression, err := p.parseForExpression()
		if err != nil {
			return nil, err
		}
		textualExpression.ForExpression = forExpression
	case IF:
		ifExpression, err := p.parseIfExpression()
		if err != nil {
			return nil, err
		}
		textualExpression.IfExpression = ifExpression
	case SOME:
		fallthrough
	case EVERY:
		p.unscan()
		quantifiedExpression, err := p.parseQuantifiedExpression()
		if err != nil {
			return nil, err
		}
		textualExpression.QuantifiedExpression = quantifiedExpression
	case LEFT_BRACKET:
		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		closeBracket, _ := p.scanIgnoreWhitespace()
		if closeBracket != RIGHT_BRACKET {
			return nil, fmt.Errorf("expected closing bracket after expression")
		}
		textualExpression.Expression = &expression
	case LT:
		fallthrough
	case GT:
		fallthrough
	case EQ:
		fallthrough
	case EXCLAMATION_MARK:
		p.unscan()
		simplePositiveUnaryTest, err := p.parseSimplePositiveUnaryTest()
		if err != nil {
			return nil, err
		}
		textualExpression.SimplePositiveUnaryTest = simplePositiveUnaryTest
	case AT:
		p.unscan()
		literal, err := p.parseDateTimeLiteral()
		if err != nil {
			return nil, err
		}
		textualExpression.Literal = literal
	case DOUBLEQUOTE:
		_, strValue := p.s.scanStringInQuotes()
		textualExpression.Literal = strValue
	case NULL:
		textualExpression.Literal = NullLiteral{}
	case HYPHEN:
		// arithmetic negation
	default:
		lit, err := p.parseLiteral()
		if err == nil {
			textualExpression.Literal = lit
			return &textualExpression, nil
		}
		res, err := p.parseMultiple1()
		if err != nil {
			return nil, err
		}
		switch res.(type) {
		case InstanceOf:
		case Disjunction:
		case Conjuction:
		case ArithmeticExpression:
		case PathExpression:
		case FilterExpression:
		case FunctionInvocation:
		case Name:
		}
		return nil, fmt.Errorf("unexpected input")
	}
	return &textualExpression, nil
}

// disjunction/conjuction/arithmenticExpression/pathExpression/filterExpression/functionInvocation/name/instanceOf/
func (p *Parser) parseMultiple1() (interface{}, error) {
	// we need to parse first expression an on the next token determine which kind of textual expression we are dealing with
	expLeft, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse left expression: %w", err)
	}
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case OR:
	case AND:
	// comparison
	case EQ:
	case NEQ:
	case GT:
	case GTE:
	case LT:
	case LTE:
	case BETWEEN:
	case IN:
		// arithmentic expression
	case PLUS:
	case HYPHEN:
	case ASTERISK:
		// can be multiplication or exponentiation
	case SLASH:
	}
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseSimplePositiveUnaryTest() (*SimplePositiveUnaryTest, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseBoxedExpression() (*BoxedExpression, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseDateTimeLiteral() (*DateTimeLiteral, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseQuantifiedExpression() (*QuantifiedExpression, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseIfExpression() (*IfExpression, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseForExpression() (*ForExpression, error) {
	name, err := p.parseName()
	if err != nil {
		return nil, fmt.Errorf("failed to parse name in for expression: %w", err)
	}
	in, _ := p.scanIgnoreWhitespace()
	if in != IN {
		return nil, fmt.Errorf("expected in after name in for expression (for <name> in <iteration context> return <expression>): %w", err)
	}
	iterationContexts, err := p.parseIterationContexts()
	if err != nil {
		return nil, fmt.Errorf("failed to parse iteration context in for expression: %w", err)
	}
	ret, _ := p.scanIgnoreWhitespace()
	if ret != RETURN {
		return nil, fmt.Errorf("expected return after iteration context: %w", err)
	}
	expr, err := p.parseExpression()
	if ret != RETURN {
		return nil, fmt.Errorf("failed to parse for loop expression: %w", err)
	}
	return &ForExpression{
		IteratorName:      *name,
		IterationContexts: iterationContexts,
		ReturnExpression:  expr,
	}, nil
}

func (p *Parser) parseLiteral() (*Literal, error) {
	// TODO: Parse literal 31.
	return nil, nil
}

func (p *Parser) parseName() (*Name, error) {
	name, tok := p.scanIgnoreWhitespace()
	if name != IDENT {
		return nil, fmt.Errorf("expected token of type IDENT, got: %s", name.String())
	}
	return (*Name)(&tok), nil
}

func (p *Parser) parseIterationContexts() ([]IterationContext, error) {
	// TODO: implement
	return nil, nil
}

func (p *Parser) parseIterationContext() (*IterationContext, error) {
	// TODO: implement
	return nil, nil
}
