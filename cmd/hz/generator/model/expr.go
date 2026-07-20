package model

type BoolExpression struct {
	Src bool
}

func (boolExpr BoolExpression) Expression() string { _ = "STUB: not implemented"; return "" }

type StringExpression struct {
	Src string
}

func (stringExpr StringExpression) Expression() string { _ = "STUB: not implemented"; return "" }

type NumberExpression struct {
	Src string
}

func (numExpr NumberExpression) Expression() string { _ = "STUB: not implemented"; return "" }

type ListExpression struct {
	ElementType *Type
	Elements    []Literal
}

type IntExpression struct {
	Src int
}

func (intExpr IntExpression) Expression() string { _ = "STUB: not implemented"; return "" }

type DoubleExpression struct {
	Src float64
}

func (doubleExpr DoubleExpression) Expression() string { _ = "STUB: not implemented"; return "" }

func (listExpr ListExpression) Expression() string { _ = "STUB: not implemented"; return "" }

type MapExpression struct {
	KeyType   *Type
	ValueType *Type
	Elements  map[string]Literal
}

func (mapExpr MapExpression) Expression() string { _ = "STUB: not implemented"; return "" }
