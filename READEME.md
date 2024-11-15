## expression_tool

example
```go

func TestExp(t *testing.T) {

	expr := "10 + 5 * ( 20 - 10 / 2 + 1 ) - 100 / 10 > 30 && 30 < 100 || 20 * 3 > 100 || 1 != 100"
	expression := NewExpressionCalculator()
	sss := "5 * 4 > 100 && 1 < 4 "
	//计算表达式
	evaluateExpression := expression.EvaluateExpression(sss)
	fmt.Println(evaluateExpression)
	//中缀转后缀
	postfix := expression.InfixToPostfix(expr)
	fmt.Println(postfix)
	//计算后缀表达式
	evaluatePostfix := expression.EvaluatePostfix(postfix)
	fmt.Println(evaluatePostfix)

}

```