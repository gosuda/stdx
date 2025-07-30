package cond

// Clause represents a single condition-result pair in a Cond expression
type Clause[T any] struct {
	condition func() bool
	result    func() T
}

// NewClause creates a new clause with a condition and result
func NewClause[T any](condition func() bool, result func() T) Clause[T] {
	return Clause[T]{
		condition: condition,
		result:    result,
	}
}

// When creates a clause with a condition and a lazy-evaluated result
func When[T any](condition func() bool, result func() T) Clause[T] {
	return NewClause(condition, result)
}

// WhenValue creates a clause with a condition and an immediate value
func WhenValue[T any](condition func() bool, value T) Clause[T] {
	return NewClause(condition, func() T { return value })
}

// WhenTrue creates a clause that's always true (used as else clause)
func WhenTrue[T any](result func() T) Clause[T] {
	return NewClause(func() bool { return true }, result)
}

// WhenTrueValue creates an else clause with an immediate value
func WhenTrueValue[T any](value T) Clause[T] {
	return NewClause(func() bool { return true }, func() T { return value })
}

// Cond represents a conditional expression similar to Lisp's cond
type Cond[T any] struct {
	clauses []Clause[T]
}

// New creates a new Cond expression
func New[T any]() *Cond[T] {
	return &Cond[T]{clauses: make([]Clause[T], 0)}
}

// Add adds a clause to the Cond expression
func (c *Cond[T]) Add(clause Clause[T]) *Cond[T] {
	c.clauses = append(c.clauses, clause)
	return c
}

// When adds a conditional clause
func (c *Cond[T]) When(condition func() bool, result func() T) *Cond[T] {
	return c.Add(When(condition, result))
}

// WhenValue adds a conditional clause with an immediate value
func (c *Cond[T]) WhenValue(condition func() bool, value T) *Cond[T] {
	return c.Add(WhenValue(condition, value))
}

// Else adds a default clause (always true)
func (c *Cond[T]) Else(result func() T) *Cond[T] {
	return c.Add(WhenTrue(result))
}

// ElseValue adds a default clause with an immediate value
func (c *Cond[T]) ElseValue(value T) *Cond[T] {
	return c.Add(WhenTrueValue(value))
}

// Unless adds a clause that executes when the condition is false
func (c *Cond[T]) Unless(condition func() bool, result func() T) *Cond[T] {
	return c.Add(Unless(condition, result))
}

// UnlessValue adds an unless clause with an immediate value
func (c *Cond[T]) UnlessValue(condition func() bool, value T) *Cond[T] {
	return c.Add(UnlessValue(condition, value))
}

// Eval evaluates the Cond expression and returns the first matching result
func (c *Cond[T]) Eval() (T, bool) {
	for _, clause := range c.clauses {
		if clause.condition() {
			return clause.result(), true
		}
	}
	var zero T
	return zero, false
}

// MustEval evaluates the Cond expression and panics if no condition matches
func (c *Cond[T]) MustEval() T {
	result, ok := c.Eval()
	if !ok {
		panic("cond: no condition matched")
	}
	return result
}

// EvalOr evaluates the Cond expression and returns a default value if no condition matches
func (c *Cond[T]) EvalOr(defaultValue T) T {
	result, ok := c.Eval()
	if !ok {
		return defaultValue
	}
	return result
}

// CondFunc is a standalone function to create and evaluate a Cond expression in one go
func CondFunc[T any](clauses ...Clause[T]) (T, bool) {
	cond := New[T]()
	for _, clause := range clauses {
		cond.Add(clause)
	}
	return cond.Eval()
}

// MustCondFunc is like CondFunc but panics if no condition matches
func MustCondFunc[T any](clauses ...Clause[T]) T {
	result, ok := CondFunc(clauses...)
	if !ok {
		panic("cond: no condition matched")
	}
	return result
}

// Switch creates a Cond expression for switch-like behavior
func Switch[T, U any](value T, equal func(T, T) bool) *SwitchCond[T, U] {
	return &SwitchCond[T, U]{
		value: value,
		equal: equal,
		cond:  New[U](),
	}
}

// SwitchCond represents a switch-like conditional expression
type SwitchCond[T, U any] struct {
	value T
	equal func(T, T) bool
	cond  *Cond[U]
}

// Case adds a case to the switch
func (s *SwitchCond[T, U]) Case(caseValue T, result func() U) *SwitchCond[T, U] {
	s.cond.When(func() bool { return s.equal(s.value, caseValue) }, result)
	return s
}

// CaseValue adds a case with an immediate value
func (s *SwitchCond[T, U]) CaseValue(caseValue T, result U) *SwitchCond[T, U] {
	return s.Case(caseValue, func() U { return result })
}

// Default adds a default case
func (s *SwitchCond[T, U]) Default(result func() U) *SwitchCond[T, U] {
	s.cond.Else(result)
	return s
}

// DefaultValue adds a default case with an immediate value
func (s *SwitchCond[T, U]) DefaultValue(result U) *SwitchCond[T, U] {
	return s.Default(func() U { return result })
}

// Eval evaluates the switch expression
func (s *SwitchCond[T, U]) Eval() (U, bool) {
	return s.cond.Eval()
}

// MustEval evaluates the switch expression and panics if no case matches
func (s *SwitchCond[T, U]) MustEval() U {
	return s.cond.MustEval()
}

// EvalOr evaluates the switch expression with a default value
func (s *SwitchCond[T, U]) EvalOr(defaultValue U) U {
	return s.cond.EvalOr(defaultValue)
}

// Guard creates a guard clause (condition with no result, for side effects)
func Guard(condition func() bool, action func()) Clause[struct{}] {
	return NewClause(condition, func() struct{} {
		action()
		return struct{}{}
	})
}

// Unless creates a clause that executes when the condition is false
func Unless[T any](condition func() bool, result func() T) Clause[T] {
	return NewClause(func() bool { return !condition() }, result)
}

// UnlessValue creates an unless clause with an immediate value
func UnlessValue[T any](condition func() bool, value T) Clause[T] {
	return Unless(condition, func() T { return value })
}
