package traceinfo

type traceInfo struct {
	stats HTTPStats
}

func (r *traceInfo) Stats() HTTPStats { _ = "STUB: not implemented"; return *new(HTTPStats) }

func (r *traceInfo) Reset() { _ = "STUB: not implemented"; return }

func NewTraceInfo() TraceInfo { _ = "STUB: not implemented"; return *new(TraceInfo) }
