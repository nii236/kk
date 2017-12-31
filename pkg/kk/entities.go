package k

// EntitiesReducer contains the data for Entities
type EntitiesReducer struct {
	Pods        *PodEntities
	Debug       *DebugEntities
	Errors      *ErrorEntities
	Namespaces  *NamespaceEntities
	Resources   *ResourceEntities
	Deployments *DeploymentEntities
}

// ResourceEntities contains the data for Resources
type ResourceEntities struct {
	Resources []string
}

// ErrorEntities contains the data for Errors
type ErrorEntities struct {
	Lines        []string
	Acknowledged bool
}

// DebugEntities contains the data for Debugs
type DebugEntities struct {
	Lines []interface{}
}
