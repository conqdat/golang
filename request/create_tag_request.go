package request

type CreateTagRequest struct {
	Name string `validate:"required" json:"name"`
}
