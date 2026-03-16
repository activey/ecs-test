package loader

type ProgressMonitor struct {
	stepsTotal     int
	stepsLeft      int
	updateFunction func(int)
}

func Progress[V any](m *ProgressMonitor, f func() V) V {
	result := f()
	m.stepsLeft--

	percentage := int((1.0 - float64(m.stepsLeft)/float64(m.stepsTotal)) * 100)
	m.updateFunction(percentage)

	return result
}

func NewProgressMonitor(steps int, updateFunction func(int)) *ProgressMonitor {
	return &ProgressMonitor{
		stepsTotal:     steps,
		stepsLeft:      steps,
		updateFunction: updateFunction,
	}
}
