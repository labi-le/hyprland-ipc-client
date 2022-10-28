package client

type ArgvQueue []string

func (a *ArgvQueue) String() string {
	var s string
	for _, arg := range *a {
		s += arg + " "
	}

	// remove last space
	if len(s) > 0 {
		s = s[:len(s)-1]
	}

	return s
}

func (a *ArgvQueue) Add(arg string) {
	*a = append(*a, arg)
}

func (a *ArgvQueue) Len() int {
	return len(*a)
}
