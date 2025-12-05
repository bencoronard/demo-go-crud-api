package resource

type ResourceService interface {
	ListResources() ([]Resource, error)
}
