package spec


import (
"fmt"
)

type counter struct {
  value int
}

func (c *counter) Get () int {
  c.value++
  return c.value
}

type Expression interface {
  Value(*counter) string
  Data() []interface{}
}

type emptyExpression struct {
}

func (e emptyExpression) Value(*counter) string {
  return ""
}

func (e emptyExpression) Data() []interface{} {
  return nil
}

func Empty() Expression {
  return emptyExpression{}
}

type evaluationExpression struct {
  label    string
  value    string
  operator string
}

func (e *evaluationExpression) Value(_ *counter) string {
  switch e.operator {
  case "like":
    return fmt.Sprintf("%s ILIKE ?", e.label)
  default:
    return fmt.Sprintf("%s %s ?", e.label, e.operator)
  }
}

func (e *evaluationExpression) Data() []interface{} {
  switch e.operator {
  case "like":
    return []interface{}{ "%" + e.value + "%"}
  default:
    return []interface{}{e.value}
  }
}

func NewEval(label string, value string, operator string) Expression {
  return &evaluationExpression{
    label:    label,
    value:    value,
    operator: operator,
  }
}

type andExpression struct {
  left  Expression
  right Expression
}

func (e *andExpression) Data() []interface{} {
  data := make([]interface{},0)
  left := e.left.Data()
  data = append(data, left...)

  right := e.right.Data()
  for _, v := range right {
    data = append(data, v)
  }
  return data
}

func (e *andExpression) Value(c *counter) string {
  return fmt.Sprintf("(%s AND %s)", e.left.Value(c), e.right.Value(c))
}

type orExpression struct {
  left  Expression
  right Expression
}

func (e *orExpression) Data() []interface{} {
  data := make([]interface{},0)
  left := e.left.Data()
  for _, v := range left {
    data = append(data, v)
  }

  right := e.right.Data()
  for _, v := range right {
    data = append(data, v)
  }
  return data
}

func (e *orExpression) Value(c *counter) string {
  return fmt.Sprintf(" ( %s OR %s )", e.left.Value(c), e.right.Value(c))
}

type notExpression struct {
  expression Expression
}

func (e *notExpression) Data() []interface{} {
  return e.Data()
}

func (e *notExpression) Value(*counter) string {
  return fmt.Sprintf(" NOT ( %s )", e.expression)
}

func And(l Expression, r Expression) Expression {
  return &andExpression{
    left:  l,
    right: r,
  }
}

func Or(l Expression, r Expression) Expression {
  return &orExpression{
    left:  l,
    right: r,
  }
}

func Not(negate Expression) Expression {
  return &notExpression{expression: negate}
}

type Spec interface {
  Query(string) string
  Data() []interface{}
}

func Specify(e Expression) Spec {
  return &query{expression: e}
}

type query struct {
  expression Expression
}

func (q *query) Data() []interface{} {
  return q.expression.Data()
}

func (q *query) Query(qr string) string {
  count := &counter{value: 0}
  val := q.expression.Value(count)
  if val != "" {
    return fmt.Sprintf("%s WHERE %s", qr, val)
  }
  return qr
}
