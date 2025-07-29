package transformer

type Context byte

const (
	BaseContext     Context = 0
	GlobalContext   Context = 1
	FunctionContext Context = iota << 1
	LoopContext
)

func (ctx Context) Matches(other Context) bool {
	return ctx&other != 0
}

func (ctx Context) Extending() Context {
	return ctx.Excluding(GlobalContext)
}

func (ctx Context) Choose(cond, a, b Context) Context {
	if ctx.Matches(cond) {
		return a
	} else {
		return b
	}
}

func (ctx Context) Including(other Context) Context {
	return ctx | other
}

func (ctx Context) Excluding(other Context) Context {
	return ctx &^ other
}

func (c Context) String() string {
	parts := c.Representation()
	if len(parts) == 0 {
		return "BaseContext"
	}
	return joinWithAmpersand(parts)
}

func (c Context) Representation() []string {
	contextNames := map[Context]string{
		GlobalContext:   "GlobalContext",
		FunctionContext: "FunctionContext",
		LoopContext:     "LoopContext",
	}

	var parts []string
	for bit, name := range contextNames {
		if c.Matches(bit) {
			parts = append(parts, name)
		}
	}

	return parts
}

func joinWithAmpersand(parts []string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += "&"
		}
		result += p
	}
	return result
}
