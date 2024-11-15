package expression_tool

// Stack 通用栈结构
type Stack struct {
	items []interface{}
}

// NewStack 创建新栈
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push 压栈
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop 弹栈
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 获取栈大小
func (s *Stack) Size() int {
	return len(s.items)
}
