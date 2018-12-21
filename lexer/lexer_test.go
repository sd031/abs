package lexer

import (
	"testing"

	"abs/token"
)

func TestNextToken(t *testing.T) {
	input := `five = 5;
ten = 10;

add = f(x, y) {
  x + y;
};

result = add(five, ten);
&&||!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

while (1 > 0) {
	echo("hello")
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
$(echo "()");
{"foo": "bar"}

$(curl icanhazip.com -X POST)
$(ls *.go);
a = [1]
a.first()
# Comment
// Comment
hello
$(command; command)
$(command2; command2);
one | two | tree
"hel\"lo"
"hel\lo"
"hel\\\\lo"
"\"hello\""
"\"he\"\"llo\""
"hello\\"
"hello\\\\"
"\\\\hello"
**
1..10
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "f"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.GT, ">"},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "echo"},
		{token.LPAREN, "("},
		{token.STRING, "hello"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.COMMAND, `echo "()"`},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.COMMAND, "curl icanhazip.com -X POST"},
		{token.COMMAND, "ls *.go"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.RBRACKET, "]"},
		{token.IDENT, "a"},
		{token.DOT, "."},
		{token.IDENT, "first"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMENT, "# Comment"},
		{token.COMMENT, "// Comment"},
		{token.IDENT, "hello"},
		{token.COMMAND, "command; command"},
		{token.COMMAND, "command2; command2"},
		{token.IDENT, "one"},
		{token.PIPE, "|"},
		{token.IDENT, "two"},
		{token.PIPE, "|"},
		{token.IDENT, "tree"},
		{token.STRING, "hel\"lo"},
		{token.STRING, "hel\\lo"},
		{token.STRING, "hel\\\\lo"},
		{token.STRING, "\"hello\""},
		{token.STRING, "\"he\"\"llo\""},
		{token.STRING, "hello\\"},
		{token.STRING, "hello\\\\"},
		{token.STRING, "\\\\hello"},
		{token.EXPONENT, "**"},
		{token.INT, "1"},
		{token.RANGE, ".."},
		{token.INT, "10"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
