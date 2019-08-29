package resolvers

import (
	"context"
	"gqlgorm/generated"
	"gqlgorm/models"
	"gqlgorm/utils/database"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
var db = database.GetInstance()
type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateTodo(ctx context.Context, input models.EditTodo) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Roles(ctx context.Context) ([]*models.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) Login(ctx context.Context, username string, password string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Userfindbyid(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	db.Find(&user)
	return user,nil
}
func (r *queryResolver) Todo(ctx context.Context, input *models.FetchTodo) (*models.Todo, error) {
	panic("not implemented")
}
