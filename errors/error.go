package errors

type UnknownProviderError struct {
	Provider string
}

func (e *UnknownProviderError) Error() string {
	return "unknown provider %s" + e.Provider
}

type MetricsNotInitializedError struct {
}

func (e *MetricsNotInitializedError) Error() string {
	return "metrics not initialized"
}
