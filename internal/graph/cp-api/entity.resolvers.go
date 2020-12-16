package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"bitbucket.org/antuitinc/esp-cp-api/internal/datamodels"
	"bitbucket.org/antuitinc/esp-cp-api/internal/graph/cp-api/generated"
)

func (r *entityResolver) FindCPPinByID(ctx context.Context, id string) (*datamodels.CPPin, error) {
	return r.Query().GetCPPin(ctx, id)
}

func (r *entityResolver) FindCPTemplateByID(ctx context.Context, id string) (*datamodels.CPTemplate, error) {
	return r.Query().GetCPTemplate(ctx, id)
}

func (r *entityResolver) FindCPUserByID(ctx context.Context, id string) (*datamodels.CPUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindCPWorkbookByID(ctx context.Context, id string) (*datamodels.CPWorkbook, error) {
	return r.Query().GetCPWorkbook(ctx, id)
}

func (r *entityResolver) FindCPWorkbookCommentByWorkbookID(ctx context.Context, workbookID string) (*datamodels.CPWorkbookComment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
