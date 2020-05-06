package OptionalPattern

type Human struct {
	Name string
	Type string
	Age  int
}

func NewHuman(opts ...HumanOption) *Human {
	const (
		defaultName = "Ben"
		defaultType = "Man"
		defaultAge  = 30
	)

	h :=  &Human{
		Name: defaultName,
		Type: defaultType,
		Age:  defaultAge,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

type HumanOption func(*Human)

func WithName(n string) HumanOption {
	return func(h *Human) {
		h.Name = n
	}
}

func WithType(t string) HumanOption {
	return func(h *Human) {
		h.Type = t
	}
}

func WithAge(a int) HumanOption {
	return func(h *Human) {
		h.Age = a
	}
}
