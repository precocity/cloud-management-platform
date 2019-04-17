package request

//Repository Handles Data Storage For Resource Requests
type Repository struct {
}

//CreateRepository creates a new instance of the Resource Request Repository
func CreateRepository() *Repository {

	r := &Repository{}

	return r
}
