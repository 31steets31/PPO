package postgres

import "fmt"

const (
	checkpointTable = "checkpoint"
	companyTable    = "company"
	documentTable   = "document"
	employeeTable   = "employee"
	fieldTable      = "field"
	infoCardTable   = "info_card"
	photoTable      = "photo"
	passageTable    = "passage"
)

const (
	idField                = "id"
	phoneNumberField       = "phone_number"
	nameField              = "name"
	cityField              = "city"
	infoCardIdField        = "info_card_id"
	fullNameField          = "full_name"
	companyIdField         = "company_id"
	postField              = "post"
	passwordField          = "password"
	dateOfBirthField       = "date_of_birth"
	documentIdField        = "document_id"
	typeField              = "type"
	valueField             = "value"
	createdEmployeeIdField = "created_employee_id"
	isConfirmedField       = "is_confirmed"
	createdDateField       = "created_date"
	checkpointIdField      = "checkpoint_id"
	timeField              = "time"
	keyField               = "key"
)

func fullColName(tableName, columnName string) string {
	return fmt.Sprintf("%s.%s", tableName, columnName)
}

func on(baseTable, targetTable, baseColumn, targetColumn string) string {
	return fmt.Sprintf("%s on %s.%s=%s.%s",
		targetTable,
		baseTable,
		baseColumn,
		targetTable,
		targetColumn,
	)
}
