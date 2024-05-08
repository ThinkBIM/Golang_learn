package maths

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// OperatorType ...
type OperatorType int32

// Operator define
const (
	OperatorAdd OperatorType = 0
	OperatorSub OperatorType = 1
)

// operator string
var operatorMap = map[OperatorType]string{
	OperatorAdd: "+",
	OperatorSub: "-",
}

func (o OperatorType) String() string {
	return operatorMap[o]
}

// Expression ...
type Expression struct {
	NumberA  int
	NumberB  int
	Operator OperatorType
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gap(max, number int, operator OperatorType) int {
	switch operator {
	case OperatorAdd:
		return maxInt(max-number, 2)
	case OperatorSub:
		return max
	default:
		return max
	}
}

// Generate ...
func Generate(max, min, count int) (list []Expression) {
	rand.Seed(time.Now().Unix())
	maps := make(map[string]Expression)
	for {
		a := rand.Intn(max)
		o := OperatorType(rand.Intn(2))
		t := gap(max, a, o)
		b := rand.Intn(t)
		if (o == OperatorSub && b > a) || (a < min || b < min) {
			continue
		}
		key := fmt.Sprintf("%d.%d.%d", a, 0, b)
		if _, ok := maps[key]; !ok {
			exp := Expression{
				NumberA:  a,
				Operator: o,
				NumberB:  b,
			}
			maps[key] = exp
			list = append(list, exp)
		}
		if len(list) >= count {
			break
		}
	}
	return
}

// Result calc result
func Result(exp Expression) int {
	switch exp.Operator {
	case OperatorAdd:
		return exp.NumberA + exp.NumberB
	case OperatorSub:
		return exp.NumberA - exp.NumberB
	default:
		return math.MinInt32
	}
}
