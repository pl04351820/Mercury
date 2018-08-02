package States

type ParallelElement struct {
	Type string
	Resource string
	End bool
}

type ParallelTransition struct{
	StateAt string
	States map[string]ParallelElement
}

type Parallel struct {
	End bool
	Branches []ParallelTransition
}

