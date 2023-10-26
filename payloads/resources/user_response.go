package resources

type UserResource struct {
	First_name string `json:"first_name,omitempty"`
	Last_name  string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
}
