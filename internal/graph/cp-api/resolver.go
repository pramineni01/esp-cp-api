package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr/v2"
	"github.com/jmoiron/sqlx"

	db "bitbucket.org/antuitinc/esp-cp-api/internal/db"
)

const (
	ADD_PIN_STMT                       = "add-pin.tmpl"
	ADD_WORKBOOK_COMMENT_STMT          = "add-workbook-comment.tmpl"
	GET_PIN_STMT                       = "get-pin.tmpl"
	GET_TEMPLATE_STMT                  = "get-template.tmpl"
	GET_WORKBOOK_STMT                  = "get-workbook.tmpl"
	GET_WORKBOOK_COMMENTS_STMT         = "get-workbook-comments.tmpl"
	GET_WORKBOOK_COMMENTS_FOR_WORKBOOK = "get-workbook-comments-for-workbook.tmpl"
	UPDATE_WORKBOOK_STMT               = "update-workbook.tmpl"
)

type Resolver struct {
	// box holding all the sql templates
	tmplBox  *packr.Box
	DBClient db.Interface
}

// LoadTemplates: Initialize all templates on appl start
func (r *Resolver) LoadTemplates() error {
	// load db query stmt templates
	r.tmplBox = packr.New("dbtemplates", "../../db/dbtemplates")
	if len(r.tmplBox.List()) == 0 {
		return errors.New("Error while loading templates")
	}

	return nil
}

// func (r *Resolver) GetStmtTemplate(templateName string) string, error {
// 	if tmpl, err := r.tmplBox.FindString(tmplFile); err != nil {
// 		return "", errors.New("Template lookup failed")
// 	} else {
// 		return tmpl, nil
// 	}
// }

// GetDBStmt: Given template name and data, returns db statement
// 	Input: stmt template and data
//  Output: DB statement ready for execution for execution
func (r *Resolver) GetDBStmt(templateName string, data interface{}) (string, error) {
	rawStr, err := r.tmplBox.FindString(templateName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error looking up the template: %s, error: %v\n", templateName, err))
	}

	tmpl, err := template.New(templateName).Funcs(sprig.TxtFuncMap()).Parse(rawStr)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing the template: %s, error: %v, content: %s\n", templateName, err, rawStr))
	}

	var stmt strings.Builder
	if err = tmpl.Execute(&stmt, data); err != nil {
		return "", errors.New(fmt.Sprintf("Error while preparing statement with provided data. Stmt: %s, data: %v \n", rawStr, data))
	}

	return stmt.String(), nil
}

func (r *Resolver) getUserID(ctx context.Context) (string, bool) {
	return "user_1", true
	//	val, ok := ctx.Value("userId").(string)
	//	if ok && len(strings.TrimSpace(val)) == 0 {
	//		return "", false
	//	}
	//	return val, ok
}

func (r *Resolver) querySingleRow(ctx context.Context, tmplName string, data interface{}) (*sqlx.Row, error) {
	stmt, err := r.GetDBStmt(tmplName, data)
	if err != nil {
		return nil, err
	}
	fmt.Println("Stmt: ", stmt)
	// TBD: execute stmt
	// return r.DBClient.QueryRowxContext(ctx, stmt), nil
	if row := r.DBClient.QueryRowxContext(ctx, stmt); row.Err() != nil {
		fmt.Printf("querySingleRow: row.Err(): %v", row.Err())
		return nil, row.Err()
	} else {
		return row, err
	}
}

func (r *Resolver) queryRows(ctx context.Context, tmplName string, data interface{}) (*sqlx.Rows, error) {
	stmt, err := r.GetDBStmt(tmplName, data)
	if err != nil {
		return nil, err
	}
	fmt.Println("queryRows() Stmt: ", stmt)
	if rows, err := r.DBClient.QueryxContext(ctx, stmt); err != nil {
		fmt.Printf("queryRows: err: %v", err)
		return nil, err
	} else {
		if rows.Err() != nil {
			fmt.Printf("queryRows: rows.Err(): %v", rows.Err())
			return nil, rows.Err()
		}
		return rows, nil
	}
}

func (r *Resolver) execStmt(ctx context.Context, tmplName string, data interface{}) (sql.Result, error) {
	stmt, err := r.GetDBStmt(tmplName, data)
	if err != nil {
		return nil, err
	}
	fmt.Println("Stmt: ", stmt)
	// TBD: execute stmt
	res, err := r.DBClient.ExecContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	count, _ := res.RowsAffected()
	if count == int64(0) {
		return nil, errors.New(fmt.Sprintf("Zero rows impacted."))
	}

	return res, nil
}
