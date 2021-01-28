package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/maapteh/graphql-golang/generated"
	"github.com/maapteh/graphql-golang/model"
)

func (r *mutationResolver) AddCrocodile(ctx context.Context, input model.CrocodileInput) (*model.Crocodile, error) {
	defer errorHandler(ctx, "Unable to add crocodile, %v")
	return r.zoo().AddCrocodile(&input), nil
}

func (r *queryResolver) Crocodiles(ctx context.Context) ([]*model.Crocodile, error) {
	defer errorHandler(ctx, "Unable to list crocodiles, %v")
	return r.zoo().ListCrocodiles(), nil
}

func (r *queryResolver) Crocodile(ctx context.Context, id int) (*model.Crocodile, error) {
	defer errorHandler(ctx, "Unable to get crocodile, %v")
	return r.zoo().GetCrocodile(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func errorHandler(ctx context.Context, msg string) {
	if r := recover(); r != nil {
		graphql.AddErrorf(ctx, msg, r)
	}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
