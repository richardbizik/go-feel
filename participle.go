package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
)

func main() {
	parser, err := participle.Build[Expression](
		participle.Unquote("String"),
	)
	if err != nil {
		panic(err)
	}
	parsed, err := parser.ParseString("", "true and true")
	fmt.Println(parsed)
}

type Expression struct {
}

// type FeelExpression struct {
// 	Expression         *Expression `@@`
// 	TextualExpressions *[]TextualExpression
// 	SimpleExpression   *SimpleExpression
// 	SimpleExpressions  *[]SimpleExpression
// }
//
// type FeelUnaryTests struct {
// 	PositiveUnaryTest  *PositiveUnaryTest
// 	PositiveUnaryTests *[]PositiveUnaryTest
// 	UnaryTests         UnaryTests
// }
//
// type Expression struct {
// 	BoxedExpression   *BoxedExpression
// 	TextualExpression *TextualExpression
// }
//
// type TextualExpression struct {
// 	ForExpression           *ForExpression
// 	IfExpression            *IfExpression
// 	QuantifiedExpression    *QuantifiedExpression
// 	Disjunction             *Disjunction
// 	Comparison              *Comparison
// 	ArithmenticExpression   *ArithmeticExpression
// 	InstanceOf              *InstanceOf
// 	PathExpression          *PathExpression
// 	FilterExpression        *FilterExpression
// 	FunctionInvocation      *FunctionInvocation
// 	Literal                 Literal
// 	SimplePositiveUnaryTest *SimplePositiveUnaryTest
// 	Name                    *Name
// 	Expression              *Expression
// }
//
// type ArithmeticExpression struct {
// 	Left      Expression // when using negation only left is filled up
// 	Right     *Expression
// 	Operation ArithmeticOperation
// }
//
// type ArithmeticOperation string
//
// const (
// 	ArithmeticOperationNone           = iota
// 	ArithmeticOperationAdd            // +
// 	ArithmeticOperationSubtract       // -
// 	ArithmeticOperationMultiplication // *
// 	ArithmeticOperationDivision       // /
// 	ArithmeticOperationExponentiation // **
// 	ArithmeticOperationNegation       // -
// )
//
// type SimpleExpression struct {
// 	ArithmeticExpression *ArithmeticExpression
// 	SimpleValue          *any
// }
//
// type SimplePositiveUnaryTestOperation int
//
// const (
// 	SimplePositiveUnaryTestOperationNone = iota
// 	SimplePositiveUnaryTestOperationLT   //<
// 	SimplePositiveUnaryTestOperationLTE  //<=
// 	SimplePositiveUnaryTestOperationGT   //>
// 	SimplePositiveUnaryTestOperationGTE  //>=
// 	SimplePositiveUnaryTestOperationEq   //=
// 	SimplePositiveUnaryTestOperationNEq  //!=
// )
//
// type SimplePositiveUnaryTest struct {
// 	Endpoint *Endpoint
// 	Interval *Interval
// }
//
// type IntervalType int
//
// const (
// 	IntervalTypeNone = iota
// 	IntervalTypeOpen
// 	IntervalTypeClosed
// )
//
// type Interval struct {
// 	Start         IntervalType
// 	End           IntervalType
// 	EndpointStart Endpoint
// 	EndpointEnd   Endpoint
// }
//
// type PositiveUnaryTest struct {
// 	Expression *Expression
// }
//
// type UnaryTests struct {
// 	PositiveUnaryTests *[]PositiveUnaryTest
// 	Negated            bool
// 	Empty              bool // "-"
// }
//
// type Endpoint struct {
// 	Expression Expression
// }
//
// type QualifiedName string
// type SimpleLiteral any
// type Name string
// type Literal any // can be nil
//
// type SimpleValue struct {
// 	QualifiedName QualifiedName
// 	SimpleLiteral any
// }
//
// type FunctionInvocation struct {
// 	Expression Expression
// 	Parameters Parameters
// }
//
// type Parameters struct {
// 	NamedParameters      []NamedParameter
// 	PositionalParameters []PositionalParameter
// }
//
// type NamedParameter struct {
// 	ParameterName Name
// 	Expression    Expression
// }
//
// type PositionalParameter struct {
// 	Expression Expression
// }
//
// type PathExpression struct {
// 	Expression Expression
// 	Name       Name
// }
//
// type ForExpression struct {
// 	IteratorName Name
// 	//When multiple iteration contexts are defined in the same for loop expression, the resulting iteration is a crossproduct of the elements of the iteration contexts. The iteration order is from the inner iteration context to the outer iteration context.
// 	IterationContexts []IterationContext
// 	ReturnExpression  Expression
// }
//
// type IfExpression struct {
// 	BooleanExpression Expression
// 	ThenExpression    Expression
// 	ElseExpression    Expression
// }
//
// type QuantifiedExpressionType int
//
// const (
// 	QuantifiedExpressionTypeNone = iota
// 	QuantifiedExpressionTypeSome
// 	QuantifiedExpressionTypeEvery
// )
//
// type QuantifiedExpressionIn struct {
// 	Name       Name
// 	Expression Expression
// }
//
// type QuantifiedExpression struct {
// 	Type          QuantifiedExpressionType
// 	InExpressions []QuantifiedExpressionIn
// 	Satisfies     Expression
// }
//
// type Disjunction struct {
// 	Left  Expression
// 	Right Expression
// }
//
// type Conjuction struct {
// 	Left  Expression
// 	Right Expression
// }
//
// type ComparisonType int
//
// const (
// 	ComparisonTypeNone = iota
// 	ComparisonTypeCommonBoolean
// 	ComparisonTypeBetween
// 	ComparisonTypeIn
// )
//
// type Comparison struct {
// 	Type               ComparisonType
// 	Left               Expression
// 	Right              *Expression // empty for IN
// 	PositiveUnaryTests []PositiveUnaryTest
// }
//
// // example: [1, 2, 3, 4][item > 2] = [3, 4]
// type FilterExpression struct {
// 	Expression Expression
// 	Filter     Expression
// }
//
// type InstanceOf struct {
// 	Expression Expression
// 	Type       string // ? 52.
// }
//
// type BoxedExpression struct {
// 	List               []Expression
// 	FunctionDefinition *FunctionDefinition
// 	Context            *Context
// }
//
// type FunctionDefinition struct {
// 	FormalParameters []FormalParameter
// 	Expression       Expression
// }
//
// type FormalParameter struct {
// 	ParameterName string
// 	Type          *string
// }
//
// type Context struct {
// 	ContextEntries []ContextEntry
// }
//
// type ContextEntry struct {
// 	Key        string
// 	Expression Expression
// }
//
// type DateTimeLiteral struct {
// 	Time               *time.Time
// 	FunctionInvocation *FunctionInvocation
// }
//
// type IterationContext struct {
// 	Start Expression
// 	End   *Expression
// }
//
// type BooleanOperation int
//
// const (
// 	BooleanOperationNone BooleanOperation = iota
// 	BooleanOperationAnd
// 	BooleanOperationOr
// )
//
// //go:generate stringer -type=Token
// const (
// 	ILLEGAL Token = iota
// 	EOF
// 	WS
// 	// Literals
// 	IDENT // fields, table_name
//
// 	// Misc characters
// 	AT                // @
// 	DOUBLEQUOTE       // "
// 	ASTERISK          // *
// 	COMMA             // ,
// 	DOT               // .
// 	COLON             // :
// 	SEMICOLON         // ;
// 	PLUS              // +
// 	HYPHEN            // -
// 	LEFT_PARENTHESIS  // (
// 	RIGHT_PARENTHESIS // )
// 	LEFT_BRACKET      // [
// 	RIGHT_BRACKET     // ]
// 	SLASH             // /
// 	EXCLAMATION_MARK  // !
//
// 	// Logical operators
// 	TRUE
// 	FALSE
// 	AND
// 	OR
// 	NOT
// 	GT
// 	GTE
// 	LT
// 	LTE
// 	EQ
// 	NEQ
//
// 	// keywords
// 	NULL
// 	BETWEEN
// 	IN
// 	RETURN
// 	RANGE
// 	LIST
// 	CONTEXT
// 	FUNCTION
// 	EXTERNAL
// 	FOR
// 	SOME
// 	EVERY
// 	IF
// 	THEN
// 	ELSE
// 	SATISFIES
// )
