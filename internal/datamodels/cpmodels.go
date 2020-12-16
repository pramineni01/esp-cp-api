package datamodels

import (
	"database/sql"
)

type User struct {
	UserID    string  `json:"userId"`
	FirstName *string `json:"firstName"`
	Username  *string `json:"username"`
}

func (User) IsEntity() {}

type CPWorkbookStatus string

type CPPin struct {
	ID                string       `db:"id"`
	Title             *string      `db:"title"`
	Description       *string      `db:"description"`
	CreationDate      sql.NullTime `db:"creation_date"`
	VisualizationFlag bool         `db:"visualization_flag"`
	WorkbookID        string       `db:"workbook_id"`
	Filters           string       `db:"filters"`
	Context           string       `db:"context"`
	Workbook          *CPWorkbook  `json:"workbook"`
}

func (CPPin) IsEntity() {}

type CPTemplate struct {
	ID         string  `db:"id"`
	Name       string  `db:"name"`
	Definition *string `db:"definition"`
}

func (CPTemplate) IsEntity() {}

type CPWorkbook struct {
	ID             string              `db:"id"`
	Scope          string              `db:"scope"`
	DadatasetID    string              `db:"da_dataset_id"`
	Status         CPWorkbookStatus    `db:"status"`
	LastModified   sql.NullTime        `db:"last_modified"`
	LastModifiedBy string              `db:"last_modified_by"`
	Comments       []CPWorkbookComment `db:"comments"`
	Template_id    string              `db:"template_id"`
	Template       *CPTemplate
}

func (CPWorkbook) IsEntity() {}

type CPWorkbookComment struct {
	ID         string `db:"id"`
	WorkbookID string `db:"workbook_id"`
	Comment    string `db:"comment"`
	UserID     string `db:"user_id"`
}

func (CPWorkbookComment) IsEntity() {}

type CPUser struct {
	ID         string `db:"id"`
	WorkbookID string `db:"workbookID"`
	UserID     string `db:"user_id"`
}

func (CPUser) IsEntity() {}
