package transformer

import "fmt"

type Position struct {
	Col int16 `json:"col"`
	Ln  int16 `json:"ln"`
	Idx int32 `json:"idx"`
}

type Bounds struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

func (p Position) Next() Position {
	p.Col++
	p.Idx++
	return p
}

func (p Position) Step(char byte) Position {
	p = p.Next()
	if char == '\n' {
		p.Col = 0
		p.Ln++
	}
	return p
}

func MakeBounds(start Position, end Position) Bounds {
	return Bounds{Start: start, End: end}
}

func MakeBoundsForToken(start Position, tokenLegth int) Bounds {
	next := start
	for range tokenLegth {
		next = next.Next()
	}
	return MakeBounds(start, next)
}

type SylvaError struct {
	Bounds  Bounds `json:"bounds"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *SylvaError) String() string {
	return fmt.Sprintf("%v: %v (%v)", e.Type, e.Message, e.Bounds)
}

func (e *SylvaError) Format(text, fileName, prefix string) string {
	arrows := StringsWithArrows(text, e.Bounds, prefix)
	return fmt.Sprintf(
		"%v: %v\n%v\n%vFile %v, line %v, column %v",
		e.Type,
		e.Message,
		arrows,
		prefix,
		fileName,
		e.Bounds.Start.Ln+1,
		e.Bounds.Start.Col+1,
	)
}
