package main

import (
	"fmt"
	"strings"
	"unicode"
)

// precedence map
var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'^': 3,
}

// check if the operator is right-associative
func isRightAssociative(op rune) bool {
	return op == '^'
}

// check if the character is an operator
func isOperator(c rune) bool {
	_, exists := precedence[c]
	return exists
}

// check if character is left parenthesis
func isLeftParenthesis(c rune) bool {
	return c == '('
}

// check if character is right parenthesis
func isRightParenthesis(c rune) bool {
	return c == ')'
}

// check precedence of operators
func hasHigherPrecedence(op1, op2 rune) bool {
	p1 := precedence[op1]
	p2 := precedence[op2]
	if p1 == p2 {
		return !isRightAssociative(op1)
	}
	return p1 > p2
}

// function to convert infix to postfix
func infixToPostfix(expression string) string {
	var result strings.Builder
	var stack []rune

	for _, token := range expression {
		// If the token is an operand, add it to the output
		if unicode.IsLetter(token) || unicode.IsDigit(token) {
			result.WriteRune(token)
		} else if isOperator(token) {
			// If the token is an operator
			for len(stack) > 0 && isOperator(stack[len(stack)-1]) && hasHigherPrecedence(stack[len(stack)-1], token) {
				result.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Push the current operator onto the stack
			stack = append(stack, token)
		} else if isLeftParenthesis(token) {
			// If the token is a left parenthesis, push it to the stack
			stack = append(stack, token)
		} else if isRightParenthesis(token) {
			// If the token is a right parenthesis, pop the stack until the left parenthesis is found
			for len(stack) > 0 && !isLeftParenthesis(stack[len(stack)-1]) {
				result.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Pop the left parenthesis
			stack = stack[:len(stack)-1]
		}
	}

	// Pop all the operators from the stack
	for len(stack) > 0 {
		result.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result.String()
}

func main() {
	expression := "a+b*(c^d-e)^(f+g*h)-i"
	fmt.Println("Infix Expression: ", expression)
	postfix := infixToPostfix(expression)
	fmt.Println("Postfix Expression: ", postfix)
}
