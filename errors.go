package resticgo

type NoResticBinaryError struct {
	err string
}

func (e *NoResticBinaryError) Error() string {
	return e.err
}

type ResticFailedError struct {
	err string
}

func (e *ResticFailedError) Error() string {
	return e.err
}

type UnexpectedResticResult struct {
	err string
}

func (e *UnexpectedResticResult) Error() string {
	return e.err
}
