package transformer_test

import (
	"sylva_parser/transformer"
	"testing"
)

func TestContexts(t *testing.T) {
	ctxA := transformer.BaseContext
	ctxB := transformer.GlobalContext
	ctxC := transformer.FunctionContext.Including(transformer.LoopContext)

	strA := ctxA.String()
	strB := ctxB.String()
	strC := ctxC.String()

	if strA != "BaseContext" {
		t.Errorf("expected BaseContext, got %v", strA)
	}
	if strB != "GlobalContext" {
		t.Errorf("expected GlobalContext, got %v", strB)
	}
	if strC != "FunctionContext&LoopContext" {
		t.Errorf("expected FunctionContext&LoopContext, got %v", strC)
	}
}
