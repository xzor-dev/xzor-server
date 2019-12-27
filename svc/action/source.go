package action

// SourceName is used to identify how actions are routed.
type SourceName string

// Source processes actions.
type Source interface {
	ProcessAction(*Action) error
}
