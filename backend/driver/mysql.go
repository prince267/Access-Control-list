package driver

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/Access-Control-list/backend/config"
	"github.com/Access-Control-list/backend/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME   = "mysql"
	CONN_MAX_LIFETIME   = 30 * 60 * 60 // 30 day
	COLUMN_INGNORE_FLAG = "1"
	COLUMN_PRIMARY      = "primary"
)

func NewMysqlConnection(cfg config.MysqlConnection) (*sql.DB, error) {
	db, err := sql.Open(MYSQL_DRIVER_NAME, cfg.ConnString())
	if err != nil {
		log.Fatalf("Failed to open mysql connection: %v", err)
		return nil, err
	}

	if cfg.IdleConnection > 0 {
		db.SetMaxIdleConns(cfg.IdleConnection)
	}
	if cfg.MaxConnection > 0 {
		db.SetMaxOpenConns(cfg.MaxConnection)
	}
	db.SetConnMaxLifetime(time.Second * CONN_MAX_LIFETIME)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping mysql: %v", err)
	}

	return db, err
}

// return the placeholder string with given count
func GetPlaceHolder(count int) string {
	if count > 0 {
		str := strings.Repeat("?, ", count)
		return str[:len(str)-2]
	}

	return ""
}

/**
 * Read And Write in FILEs
 */
// *****************************************************************************************
func ReadFile(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if nil == err {
		data := string(dat)
		// log.Println(err)
		// fmt.Print(data)
		return data, err
	}

	return "", err

}

func DeleteEntity(path string) error {
	err := os.RemoveAll(path)
	return err
}

func WriteFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	log.Println(err)
	if err != nil {
		return err
	}
	d1 := []byte(content)
	werr := ioutil.WriteFile(path, d1, 0644)
	defer f.Close()
	return werr

}

// ******************************************************************************************

/**
 * Insert new row
 */
func Create(conn *sql.DB, object model.UserIModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
			COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		params = append(params, value.Interface())
		count++
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	return result, nil
}

func NewUserFile(conn *sql.DB, object model.NewFileInFolderIModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
			COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		params = append(params, value.Interface())
		count++
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	return result, nil
}

func NewUserFolder(conn *sql.DB, object model.NewFolderInFolderIModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
			COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		params = append(params, value.Interface())
		count++
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	return result, nil
}

func UpdateFolderInFolder(conn *sql.DB, object model.NewFolderInFolderIModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString(strings.Join(keyColumns, " and "))
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return result, err
}

func UpdateFileInFolder(conn *sql.DB, object model.NewFileInFolderIModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString(strings.Join(keyColumns, " and "))
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return result, err
}

/**
 * Update existing row with key column
 */

func UpdateById(conn *sql.DB, object model.UserIModel) error {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString(strings.Join(keyColumns, ", "))
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func Login(conn *sql.DB, object model.UserIModel, id int64, password string) (model.UserIModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" WHERE user_id = ? and password = ?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, id, password)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d", id))
	}

	return object, nil
}

func GetUserGroup(conn *sql.DB, object model.UserGroupIModel, id int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.UserGroupTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.GroupTable())
	queryBuffer.WriteString(" WHERE user_id = ? and user_group.group_id=groups.group_id")

	query := queryBuffer.String()
	row, err := conn.Query(query, id)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.UserGroup)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetGroupUsers(conn *sql.DB, object model.GroupUsersIModel, id int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.UserGroupTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" WHERE group_id = ? and users.user_id=user_group.user_id")

	query := queryBuffer.String()
	row, err := conn.Query(query, id)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.GroupUsers)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetParentFolders(conn *sql.DB, object model.NewFolderInFolderIModel, UserID int64, FolderID int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ? and parent_folder_id = ?")

	query := queryBuffer.String()
	row, err := conn.Query(query, UserID, FolderID)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.NewFolderInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetParentFiles(conn *sql.DB, object model.NewFileInFolderIModel, UserID int64, FolderID int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ? and parent_folder_id = ?")

	query := queryBuffer.String()
	row, err := conn.Query(query, UserID, FolderID)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.NewFileInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}
func GetFileUser(conn *sql.DB, object model.NewFileInFolderIModel, FileID int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" WHERE child_file_id = ?")

	query := queryBuffer.String()
	row, err := conn.Query(query, FileID)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.NewFileInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetFolderUser(conn *sql.DB, object model.NewFolderInFolderIModel, FolderID int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" WHERE child_folder_id = ?")

	query := queryBuffer.String()
	row, err := conn.Query(query, FolderID)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.NewFileInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func CreateFile(conn *sql.DB, object model.FilesIModel, FileName string, PathName string) (sql.Result, error) {
	file, err := os.Create(PathName)
	defer file.Close()
	if nil == err {
		rValue := reflect.ValueOf(object)
		rType := reflect.TypeOf(object)

		columns := []string{}

		// count := 0
		for idx := 0; idx < rValue.Elem().NumField(); idx++ {
			field := rType.Elem().Field(idx)
			// value := rValue.Elem().Field(idx)

			if COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
				COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
				continue
			}

			column := field.Tag.Get("column")
			columns = append(columns, column)
		}

		var queryBuffer bytes.Buffer
		queryBuffer.WriteString("INSERT INTO ")
		queryBuffer.WriteString(object.FilesTable())
		queryBuffer.WriteString("(")
		queryBuffer.WriteString(strings.Join(columns, ", "))
		queryBuffer.WriteString(") VALUE(?,?)")

		query := queryBuffer.String()
		log.Println(query)
		stmt, err := conn.Prepare(query)
		if nil != err {
			log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}

		defer stmt.Close()

		result, err := stmt.Exec(FileName, PathName)
		if nil != err {
			log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}
		return result, err

	}
	return nil, err
}

func CreateFolder(conn *sql.DB, object model.FoldersIModel, FolderName string, PathName string) (sql.Result, error) {
	err := os.Mkdir(PathName, 0755)
	if nil == err {

		rValue := reflect.ValueOf(object)
		rType := reflect.TypeOf(object)

		columns := []string{}

		// count := 0
		for idx := 0; idx < rValue.Elem().NumField(); idx++ {
			field := rType.Elem().Field(idx)
			// value := rValue.Elem().Field(idx)

			if COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
				COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
				continue
			}

			column := field.Tag.Get("column")
			columns = append(columns, column)
		}

		var queryBuffer bytes.Buffer
		queryBuffer.WriteString("INSERT INTO ")
		queryBuffer.WriteString(object.FoldersTable())
		queryBuffer.WriteString("(")
		queryBuffer.WriteString(strings.Join(columns, ", "))
		queryBuffer.WriteString(") VALUE(?,?)")

		query := queryBuffer.String()
		stmt, err := conn.Prepare(query)
		if nil != err {
			log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}

		defer stmt.Close()

		result, err := stmt.Exec(FolderName, PathName)
		if nil != err {
			log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}
		return result, err

	}
	log.Println(err)
	return nil, err
}

func GetUserFolder(conn *sql.DB, object model.FolderInFolderIModel, id int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.FolderInFolderTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.PermissionTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.FoldersTable())
	queryBuffer.WriteString(", folders as folders1 WHERE user_id = ? and parent_folder_id=folders1.folder_id and child_folder_id=folders.folder_id and folder_in_folder.permission_id=permission.permission_id")

	query := queryBuffer.String()
	// log.Println(query)
	row, err := conn.Query(query, id)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.FolderInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetUserFiles(conn *sql.DB, object model.FileInFolderIModel, id int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

	}
	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.FileInFolderTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.PermissionTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.FoldersTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object.FilesTable())
	queryBuffer.WriteString(" WHERE user_id = ? and parent_folder_id=folder_id and child_file_id=file_id and file_in_folder.permission_id=permission.permission_id")

	query := queryBuffer.String()
	// log.Println(query)
	row, err := conn.Query(query, id)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	objects := make([]interface{}, 0)

	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.FileInFolder)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {

			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects, data)
	}

	return objects, nil
}

func GetById(conn *sql.DB, object model.UserIModel, object1 model.UserGroupIModel, id int64) (model.UserGroupIModel, error) {
	log.Printf("get by id called")
	// a := model.UserGroupIModel
	a1 := reflect.ValueOf(object)
	log.Println(a1)
	rValue := reflect.ValueOf(object1)
	rType := reflect.TypeOf(object1)
	log.Println(rValue)
	log.Println(rType)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object1.UserGroupTable())
	queryBuffer.WriteString(" , ")
	queryBuffer.WriteString(object1.GroupTable())
	queryBuffer.WriteString(" WHERE user_id = ? and user_group.group_id=groups.group_id")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, id)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d", id))
	}

	return object1, nil
}

func GetAll(conn *sql.DB, object model.UserIModel, limit, offset int64) ([]interface{}, error) {
	// err := os.RemoveAll("ACL/temp")
	// log.Println(err)
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
	}

	var queryBuffer bytes.Buffer
	var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" where user_id!=0")
	if 0 != limit && 0 != offset {
		queryBuffer.WriteString(" LIMIT ? OFFSET ?")
		params = append(params, limit)
		params = append(params, offset)
	}

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.User)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, data)

	}

	return objects, nil

}

func GetAllFiles(conn *sql.DB, object model.AllFilesIModel) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
	}

	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT distinct ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.FileInFolderTable())
	queryBuffer.WriteString(",")
	queryBuffer.WriteString(object.FoldersTable())
	queryBuffer.WriteString(",")
	queryBuffer.WriteString(object.FilesTable())
	queryBuffer.WriteString(" where parent_folder_id=folder_id and child_file_id=file_id")
	// if 0 != limit && 0 != offset {
	// 	queryBuffer.WriteString(" LIMIT ? OFFSET ?")
	// 	params = append(params, limit)
	// 	params = append(params, offset)
	// }

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.AllFiles)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, data)

	}

	return objects, nil

}

func GetAllFolders(conn *sql.DB, object model.AllFoldersIModel) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
	}

	var queryBuffer bytes.Buffer
	// var params []interface{}

	queryBuffer.WriteString("SELECT distinct ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.FolderInFolderTable())
	queryBuffer.WriteString(",")
	queryBuffer.WriteString(object.FoldersTable())
	queryBuffer.WriteString(",")
	queryBuffer.WriteString("folders as folders1 where parent_folder_id=folders1.folder_id and child_folder_id=folders.folder_id")
	// if 0 != limit && 0 != offset {
	// 	queryBuffer.WriteString(" LIMIT ? OFFSET ?")
	// 	params = append(params, limit)
	// 	params = append(params, offset)
	// }

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		data := new(model.AllFolders)

		rValue1 := reflect.ValueOf(data)
		pointers := make([]interface{}, 0)
		for idx := 0; idx < rValue1.Elem().NumField(); idx++ {
			pointers = append(pointers, rValue1.Elem().Field(idx).Addr().Interface())
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, data)

	}

	return objects, nil

}

func DeleteById(conn *sql.DB, object model.UserIModel, id int64) error {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" WHERE user_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		// return nil, err
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	log.Println(result)
	return err
}

func DeleteFileInFolderById(conn *sql.DB, object model.NewFileInFolderIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" WHERE child_file_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		// return nil, err
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func DeleteFileInFolderByUserId(conn *sql.DB, object model.NewFileInFolderIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		// return nil, err
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func DeleteFolderInFolderById(conn *sql.DB, object model.NewFolderInFolderIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" WHERE child_folder_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err

	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func DeleteFolderInFolderByUserId(conn *sql.DB, object model.NewFolderInFolderIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err

	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func DeleteFolderById(conn *sql.DB, object model.FoldersIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.FoldersTable())
	queryBuffer.WriteString(" WHERE folder_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func DeleteFileById(conn *sql.DB, object model.FilesIModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.FilesTable())
	queryBuffer.WriteString(" WHERE file_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return result, err
}

func SoftDeleteById(conn *sql.DB, object model.UserIModel, id int64) error {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.UserTable())
	queryBuffer.WriteString(" SET deleted = 1  WHERE user_id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func CheckIsFileUser(conn *sql.DB, object model.NewFileInFolderIModel, userId int64, fileId int64) (model.NewFileInFolderIModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFileInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ? and child_file_id=?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, userId, fileId)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d %d", userId, fileId))
	}

	return object, nil
}

func CheckIsFolderUser(conn *sql.DB, object model.NewFolderInFolderIModel, userId int64, folderId int64) (model.NewFolderInFolderIModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.NewFolderInFolderTable())
	queryBuffer.WriteString(" WHERE user_id = ? and child_folder_id = ?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, userId, folderId)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d %d", userId, folderId))
	}

	return object, nil
}
