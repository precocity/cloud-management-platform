package resources

// ResourceType defines the resources that we can manage
type ResourceType int

const (
	//Instance is a Raw VM (EC2, Compute Engine, etc)
	Instance ResourceType = 0
)
