package syntax

type condition struct {
	column       string
	operator     string
	value        string
	compoundHead Condition
	next         Condition
	prev         Condition
}

type Condition interface {
	Value() string
	Next() Condition
	Prev() Condition
	SetNext(item Condition)
	SetPrev(item Condition)
	Column() string
	Operator() string
}

func (i *condition) Value() string {
	return i.value
}

func (i *condition) Next() Condition {
	return i.next
}

func (i *condition) Prev() Condition {
	return i.prev
}

func (i *condition) SetNext(condition Condition) {
	i.next = condition
}

func (i *condition) SetPrev(condition Condition) {
	i.prev = condition
}

func (i *condition) Column() string {
	return i.column
}

func (i *condition) Operator() string {
	return i.operator
}

func NewCondition(column, operator, value string) Condition {
	return &condition{
		column:       column,
		operator:     operator,
		value:        value,
		compoundHead: nil,
		next:         nil,
		prev:         nil,
	}
}
