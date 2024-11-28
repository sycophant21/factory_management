package repository

import (
	"errors"
	logg "factory_management_go/app/log"
	"factory_management_go/app/util"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type RepoEngine[T any] struct {
	eng *xorm.Engine
}

var eng *xorm.Engine

func InitialiseEngine[T any]() (*RepoEngine[T], error) {
	var err error
	if eng == nil {
		engineName, err := fetchProperty[string](util.DatasourceEngineName)
		if err != nil {
			return nil, err
		}
		username, err := fetchProperty[string](util.DatasourceUsername)
		if err != nil {
			return nil, err
		}
		password, err := fetchProperty[string](util.DatasourcePassword)
		if err != nil {
			return nil, err
		}
		url, err := fetchProperty[string](util.DatasourceUrl)
		if err != nil {
			return nil, err
		}
		databaseName, err := fetchProperty[string](util.DatasourceDatabaseName)
		if err != nil {
			return nil, err
		}
		eng, err = xorm.NewEngine(engineName, username+":"+password+"@tcp("+url+")/"+databaseName)
		if err != nil {
			return nil, err
		}
	}
	eng.SetTableMapper(names.SnakeMapper{})
	engine := RepoEngine[T]{eng: eng}
	logg.Logger.Info("Initialised new repo engine of type " + reflect.TypeOf(new(T)).Name())
	return &engine, err
}

func fetchProperty[T any](propertyName string) (T, error) {
	propertyInterface, er := util.GetProperty(propertyName)
	var zero T
	if er != nil {
		return zero, er
	}
	property, ok := propertyInterface.(T)
	if !ok {
		return zero, errors.New("unable to fetch engine name")
	}
	return property, nil
}

func (re *RepoEngine[T]) Create() {}
func (re *RepoEngine[T]) ReadOne(model *T, conditions ...*T) error {
	modelMap, err := readOne(re.eng, conditions...)
	if err != nil {
		return err
	}
	for _, value := range modelMap {
		*model = *value
		break
	}
	return nil
}
func (re *RepoEngine[T]) ReadAll(conditions ...*T) ([]*T, error) {
	modelArr, err := readAll(re.eng, conditions...)
	if err != nil {
		return nil, err
	}
	return modelArr, nil
}

func readOne[T any](eng *xorm.Engine, conditions ...*T) (map[string]*T, error) {
	modelMap := make(map[string]*T)
	return modelMap, eng.Find(&modelMap, conditions)
}
func readAll[T any](eng *xorm.Engine, conditions ...*T) ([]*T, error) {
	modelArr := make([]*T, 0)
	condiBeans := make([]interface{}, len(conditions))
	for i, cond := range conditions {
		condiBeans[i] = cond
	}

	return modelArr, eng.Find(&modelArr, condiBeans...)
}
func (re *RepoEngine[T]) Update() {}
func (re *RepoEngine[T]) Delete() {}
