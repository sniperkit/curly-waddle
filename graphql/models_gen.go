// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package graphql

type CreatePostInput struct {
	UserId string `json:"userId"`
	Body   string `json:"body"`
}
type CreateUserInput struct {
	Name string `json:"name"`
}
type Pagination struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}
