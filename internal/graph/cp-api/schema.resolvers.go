package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"bitbucket.org/antuitinc/esp-cp-api/internal/datamodels"
	"bitbucket.org/antuitinc/esp-cp-api/internal/graph/cp-api/generated"
	"bitbucket.org/antuitinc/esp-cp-api/internal/graph/cp-api/model"
)

func (r *cPPinResolver) CreationDate(ctx context.Context, obj *datamodels.CPPin) (*time.Time, error) {
	if obj == nil {
		return nil, errors.New("cPPinResolver::CreationDate(): Pin parameter is empty.")
	}
	return &obj.CreationDate.Time, nil
}

func (r *cPPinResolver) Workbook(ctx context.Context, obj *datamodels.CPPin) (*datamodels.CPWorkbook, error) {
	id := obj.WorkbookID
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	ID, _ := strconv.Atoi(id)
	data := struct {
		ID     int
		UserID string
	}{ID, userID}

	row, err := r.querySingleRow(ctx, GET_WORKBOOK_STMT, data)
	if err != nil {
		return nil, err
	}

	switch {
	case err == sql.ErrNoRows:
		log.Println("No workbook with id: ", id)
		return nil, err
	case err != nil:
		log.Println("query error: ", err)
		return nil, err
	}

	out := datamodels.CPWorkbook{}
	if err := row.StructScan(&out); err != nil {
		log.Println("Error while structscaning a workbook: ", err)
		return nil, err
	}
	return &out, nil
}

func (r *cPUserResolver) User(ctx context.Context, obj *datamodels.CPUser) (*datamodels.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *cPWorkbookResolver) Template(ctx context.Context, obj *datamodels.CPWorkbook) (*datamodels.CPTemplate, error) {
	if obj != nil {
		id := obj.Template_id
		ID, _ := strconv.Atoi(id)
		data := struct {
			ID int
		}{ID}

		row, err := r.querySingleRow(ctx, GET_TEMPLATE_STMT, data)
		if err != nil {
			return nil, err
		}

		switch {
		case err == sql.ErrNoRows:
			log.Println("No template with id: ", id)
			return nil, err
		case err != nil:
			log.Println("query error: ", err)
			return nil, err
		}

		out := datamodels.CPTemplate{}
		if err := row.StructScan(&out); err != nil {
			log.Println("Error while structscaning a template: ", err)
			return nil, err
		}
		return &out, nil
	}
	return nil, errors.New("cPWorkbookResolver::Template(): Workbook parameter is empty.")
}

func (r *cPWorkbookResolver) LastModified(ctx context.Context, obj *datamodels.CPWorkbook) (*time.Time, error) {
	if obj == nil {
		return nil, errors.New("cPWorkbookResolver::LastModified(): Workbook parameter is empty.")
	}
	return &obj.LastModified.Time, nil
}

func (r *cPWorkbookResolver) LastModifiedBy(ctx context.Context, obj *datamodels.CPWorkbook) (*datamodels.User, error) {
	if obj != nil {
		user := datamodels.User{UserID: obj.LastModifiedBy}
		return &user, nil
	}
	return nil, errors.New("cPWorkbookResolver::LastModifiedBy(): Workbook parameter is empty.")
}

func (r *cPWorkbookResolver) Comments(ctx context.Context, obj *datamodels.CPWorkbook) ([]datamodels.CPWorkbookComment, error) {

	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		WorkbookID string
		UserID     string
	}{obj.ID, userID}

	rows, err := r.queryRows(ctx, GET_WORKBOOK_COMMENTS_FOR_WORKBOOK, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]datamodels.CPWorkbookComment, 0)
	for rows.Next() {
		comment := datamodels.CPWorkbookComment{}
		err = rows.StructScan(&comment)
		if err != nil {
			return out, errors.New(fmt.Sprintf("Error fetching workbook comments: %v", err))
		}
		out = append(out, comment)
	}
	return out, nil

}

func (r *cPWorkbookCommentResolver) User(ctx context.Context, obj *datamodels.CPWorkbookComment) (*datamodels.User, error) {
	if obj != nil {
		user := datamodels.User{UserID: obj.UserID}
		return &user, nil
	}
	return nil, errors.New("cPWorkbookResolver::LastModifiedBy(): WorkbookComment parameter is empty.")
}

func (r *mutationResolver) UpdateCPWorkbook(ctx context.Context, workbookID string, status datamodels.CPWorkbookStatus) (*datamodels.CPWorkbook, error) {
	// Get user id from context
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		ID     string
		UserID string
		Status datamodels.CPWorkbookStatus
	}{
		workbookID,
		userID,
		status,
	}

	_, err := r.execStmt(ctx, UPDATE_WORKBOOK_STMT, data)
	if err != nil {
		return nil, err
	}

	out, err := r.Entity().FindCPWorkbookByID(ctx, workbookID)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *mutationResolver) AddCPPin(ctx context.Context, pin *model.CPPinInput) (*datamodels.CPPin, error) {
	if pin == nil {
		return nil, errors.New("Pin input empty")
	}
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		Title             string
		Description       string
		Filters           string
		Context           string
		VisualizationFlag bool
		WorkbookID        string
		UserID            string
	}{
		pin.Title,
		*pin.Description,
		pin.Filters,
		pin.Context,
		*pin.VisualizationFlag,
		pin.WorkbookID,
		userID,
	}

	res, err := r.execStmt(ctx, ADD_PIN_STMT, data)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	out, err := r.Entity().FindCPPinByID(ctx, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *mutationResolver) AddCPWorkbookComment(ctx context.Context, workbookID string, comment *string) (*bool, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		WorkbookID string
		UserID     string
		Comment    string
	}{
		workbookID,
		userID,
		*comment,
	}

	var res bool
	_, err := r.execStmt(ctx, ADD_WORKBOOK_COMMENT_STMT, data)
	if err != nil {
		return &res, err
	}
	res = true
	return &res, nil
}

func (r *queryResolver) GetCPTemplate(ctx context.Context, id string) (*datamodels.CPTemplate, error) {
	ID, _ := strconv.Atoi(id)
	data := struct {
		ID int
	}{ID}

	row, err := r.querySingleRow(ctx, GET_TEMPLATE_STMT, data)
	if err != nil {
		return nil, err
	}

	switch {
	case err == sql.ErrNoRows:
		log.Println("No template with id: ", id)
		return nil, err
	case err != nil:
		log.Println("query error: ", err)
		return nil, err
	}

	out := datamodels.CPTemplate{}
	if err := row.StructScan(&out); err != nil {
		log.Println("Error while structscaning a template: ", err)
		return nil, err
	}
	return &out, nil
}

func (r *queryResolver) GetCPTemplates(ctx context.Context, limit *int) ([]datamodels.CPTemplate, error) {
	data := struct {
		ID    int
		Limit int
	}{0, *limit}

	rows, err := r.queryRows(ctx, GET_TEMPLATE_STMT, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]datamodels.CPTemplate, 0)
	for rows.Next() {
		tmpl := datamodels.CPTemplate{}

		err := rows.StructScan(&tmpl)
		if err != nil {
			return out, errors.New(fmt.Sprintf("Error scanning rows: %v", err))
		}
		out = append(out, tmpl)
	}

	return out, nil
}

func (r *queryResolver) GetCPWorkbook(ctx context.Context, id string) (*datamodels.CPWorkbook, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	ID, _ := strconv.Atoi(id)
	data := struct {
		ID     int
		UserID string
	}{ID, userID}

	row, err := r.querySingleRow(ctx, GET_WORKBOOK_STMT, data)
	if err != nil {
		return nil, err
	}

	switch {
	case err == sql.ErrNoRows:
		log.Println("No workbook with id: ", id)
		return nil, err
	case err != nil:
		log.Println("query error: ", err)
		return nil, err
	}

	out := datamodels.CPWorkbook{}
	if err := row.StructScan(&out); err != nil {
		log.Println("Error while structscaning a workbook: ", err)
		return nil, err
	}
	return &out, nil
}

func (r *queryResolver) GetCPWorkbooks(ctx context.Context, limit *int) ([]datamodels.CPWorkbook, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		ID     int
		UserID string
		Limit  int
	}{0, userID, *limit}

	rows, err := r.queryRows(ctx, GET_WORKBOOK_STMT, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]datamodels.CPWorkbook, 0)
	for rows.Next() {
		wb := datamodels.CPWorkbook{}
		err = rows.StructScan(&wb)
		if err != nil {
			return out, errors.New(fmt.Sprintf("Error fetching workbooks: %v", err))
		}
		out = append(out, wb)
	}
	return out, nil
}

func (r *queryResolver) GetCPPin(ctx context.Context, id string) (*datamodels.CPPin, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	ID, _ := strconv.Atoi(id)
	data := struct {
		ID     int
		UserID string
	}{ID, userID}

	row, err := r.querySingleRow(ctx, GET_PIN_STMT, data)
	if err != nil {
		return nil, err
	}

	switch {
	case err == sql.ErrNoRows:
		log.Println("No workbook with id: ", id)
		return nil, err
	case err != nil:
		log.Println("query error: ", err)
		return nil, err
	}

	out := datamodels.CPPin{}
	if err := row.StructScan(&out); err != nil {
		log.Println("Error while structscaning a workbook: ", err)
		return nil, err
	}
	return &out, nil
}

func (r *queryResolver) GetCPPins(ctx context.Context, limit *int) ([]datamodels.CPPin, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	data := struct {
		ID     int
		UserID string
		Limit  int
	}{0, userID, *limit}

	rows, err := r.queryRows(ctx, GET_PIN_STMT, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]datamodels.CPPin, 0)
	for rows.Next() {
		pin := datamodels.CPPin{}
		err = rows.StructScan(&pin)
		if err != nil {
			return out, errors.New(fmt.Sprintf("Error fetching pins: %v", err))
		}
		out = append(out, pin)
	}
	return out, nil
}

func (r *queryResolver) GetCPWorkbookComments(ctx context.Context, workbookID string, limitInp *int) ([]datamodels.CPWorkbookComment, error) {
	userID, ok := r.getUserID(ctx)
	if !ok {
		log.Printf("ERROR: [getUserID]: No userID found in context")
		return nil, errors.New("No userID found in context")
	}

	var limit int
	if limitInp != nil && *limitInp > 0 {
		limit = *limitInp
	}
	data := struct {
		WorkbookID string
		UserID     string
		Limit      int
	}{workbookID, userID, limit}

	rows, err := r.queryRows(ctx, GET_WORKBOOK_COMMENTS_STMT, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]datamodels.CPWorkbookComment, 0)
	for rows.Next() {
		comment := datamodels.CPWorkbookComment{}
		err = rows.StructScan(&comment)
		if err != nil {
			return out, errors.New(fmt.Sprintf("Error fetching workbook comments: %v", err))
		}
		out = append(out, comment)
	}
	return out, nil
}

// CPPin returns generated.CPPinResolver implementation.
func (r *Resolver) CPPin() generated.CPPinResolver { return &cPPinResolver{r} }

// CPUser returns generated.CPUserResolver implementation.
func (r *Resolver) CPUser() generated.CPUserResolver { return &cPUserResolver{r} }

// CPWorkbook returns generated.CPWorkbookResolver implementation.
func (r *Resolver) CPWorkbook() generated.CPWorkbookResolver { return &cPWorkbookResolver{r} }

// CPWorkbookComment returns generated.CPWorkbookCommentResolver implementation.
func (r *Resolver) CPWorkbookComment() generated.CPWorkbookCommentResolver {
	return &cPWorkbookCommentResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type cPPinResolver struct{ *Resolver }
type cPUserResolver struct{ *Resolver }
type cPWorkbookResolver struct{ *Resolver }
type cPWorkbookCommentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
