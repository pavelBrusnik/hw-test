package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := do(in, done)
	for _, stage := range stages {
		if stage != nil {
			out = stage(do(out, done))
		} else {
			panic("stage is nil!")
		}
	}

	return out
}

func do(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()

	return out
}
