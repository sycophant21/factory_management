package repository

import (
	"errors"
	"factory_management_go/app/util"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Eng *xorm.Engine

func fetchProperty[T any](propertyName string) (T, error) {
	propertyInterface, er := util.GetProperty(propertyName)
	var zero T
	if er != nil {
		return zero, er
	}
	property, ok := propertyInterface.(T)
	if !ok {
		return zero, errors.New("Unable to fetch engine name")
	}
	return property, nil
}
func InitialiseEngine() error {
	var err error
	if Eng == nil {
		engineName, err := fetchProperty[string](util.DatasourceEngineName)
		if err != nil {
			return err
		}
		username, err := fetchProperty[string](util.DatasourceUsername)
		if err != nil {
			return err
		}
		password, err := fetchProperty[string](util.DatasourcePassword)
		if err != nil {
			return err
		}
		url, err := fetchProperty[string](util.DatasourceUrl)
		if err != nil {
			return err
		}
		databaseName, err := fetchProperty[string](util.DatasourceDatabaseName)
		if err != nil {
			return err
		}
		Eng, err = xorm.NewEngine(engineName, username+":"+password+"@tcp("+url+")/"+databaseName)
		if err != nil {
			return err
		}
		Eng.ShowSQL(true)
	}
	return err
}
