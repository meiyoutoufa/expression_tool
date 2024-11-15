package expression_tool

import (
	"strconv"
	"strings"
)

// 运算符优先级
var precedence = map[string]int{
	"||": 1,
	"&&": 2,
	"==": 3, "!=": 3, "<": 3, ">": 3, "<=": 3, ">=": 3,
	"+": 4, "-": 4,
	"*": 5, "/": 5,
	"!": 6,
	"(": 0, ")": 0,
}

// ExpressionCalculator 表达式计算器
type ExpressionCalculator struct {
	operatorStack *Stack
	operandStack  *Stack
}

// NewExpressionCalculator 创建新的表达式计算器
func NewExpressionCalculator() *ExpressionCalculator {
	return &ExpressionCalculator{
		operatorStack: NewStack(),
		operandStack:  NewStack(),
	}
}

// 函数1：计算表达式
func (ec *ExpressionCalculator) EvaluateExpression(expr string) bool {
	postfix := ec.InfixToPostfix(expr)
	return ec.EvaluatePostfix(postfix)
}

// 函数2：中缀转后缀
func (ec *ExpressionCalculator) InfixToPostfix(expr string) []string {
	tokens := strings.Fields(expr)
	output := make([]string, 0)

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)
		case token == "(":
			ec.operatorStack.Push(token)
		case token == ")":
			for !ec.operatorStack.IsEmpty() && ec.operatorStack.Peek().(string) != "(" {
				output = append(output, ec.operatorStack.Pop().(string))
			}
			if !ec.operatorStack.IsEmpty() {
				ec.operatorStack.Pop() // 弹出 "("
			}
		default: // 运算符
			for !ec.operatorStack.IsEmpty() {
				top := ec.operatorStack.Peek().(string)
				if top != "(" && precedence[top] >= precedence[token] {
					output = append(output, ec.operatorStack.Pop().(string))
				} else {
					break
				}
			}
			ec.operatorStack.Push(token)
		}
	}

	// 处理栈中剩余的运算符
	for !ec.operatorStack.IsEmpty() {
		output = append(output, ec.operatorStack.Pop().(string))
	}

	return output
}

// 函数3：计算后缀表达式
func (ec *ExpressionCalculator) EvaluatePostfix(tokens []string) bool {
	ec.operandStack = NewStack() // 清空操作数栈

	for _, token := range tokens {
		switch {
		case isNumber(token):
			num, _ := strconv.ParseFloat(token, 64)
			ec.operandStack.Push(num)
		case token == "!":
			operand := ec.operandStack.Pop().(float64)
			ec.operandStack.Push(boolToFloat(operand == 0))
		default:
			// 二元运算符
			b := ec.operandStack.Pop().(float64)
			a := ec.operandStack.Pop().(float64)
			result := ec.calculateOperation(a, b, token)
			ec.operandStack.Push(result)
		}
	}

	return ec.operandStack.Pop().(float64) != 0
}

// calculateOperation 执行具体的运算操作
func (ec *ExpressionCalculator) calculateOperation(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "&&":
		return boolToFloat(a != 0 && b != 0)
	case "||":
		return boolToFloat(a != 0 || b != 0)
	case ">":
		return boolToFloat(a > b)
	case "<":
		return boolToFloat(a < b)
	case "==":
		return boolToFloat(a == b)
	case "!=":
		return boolToFloat(a != b)
	case ">=":
		return boolToFloat(a >= b)
	case "<=":
		return boolToFloat(a <= b)
	default:
		return 0
	}
}

// 辅助函数：判断是否为数字
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// 辅助函数：布尔值转浮点数
func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
