package repository

import (
	"errors"
	logg "factory_management_go/app/log"
	"factory_management_go/app/util/program"
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
		engineName, err := fetchProperty[string](program.DatasourceEngineName)
		if err != nil {
			return nil, err
		}
		username, err := fetchProperty[string](program.DatasourceUsername)
		if err != nil {
			return nil, err
		}
		password, err := fetchProperty[string](program.DatasourcePassword)
		if err != nil {
			return nil, err
		}
		url, err := fetchProperty[string](program.DatasourceUrl)
		if err != nil {
			return nil, err
		}
		databaseName, err := fetchProperty[string](program.DatasourceDatabaseName)
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
	//eng.ShowSQL(true)
	logg.Logger.Info("Initialised new repo engine of type " + reflect.TypeOf(*new(T)).Name())
	return &engine, err
}

func fetchProperty[T any](propertyName string) (T, error) {
	propertyInterface, er := program.GetProperty(propertyName)
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

type JoinClause struct {
	JoinType  JoinType
	TableName string
	Condition interface{}
	Args      []interface{}
}

type JoinType string

type WhereClause struct {
	ParamName  string
	ParamValue interface{}
}

const (
	INNER JoinType = "INNER"
	OUTER JoinType = "OUTER"
	FULL  JoinType = "FULL"
	LEFT  JoinType = "LEFT"
	RIGHT JoinType = "RIGHT"
)

func (re *RepoEngine[T]) Create() {}
func (re *RepoEngine[T]) ReadOne(model *T, conditions ...*T) error {
	modelMap, err := readOne(re.eng.Asc("id"), conditions...)
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
	modelArr, err := readAll(re.eng.Asc("id"), conditions...)
	if err != nil {
		return nil, err
	}
	return modelArr, nil
}

//func (re *RepoEngine[T]) ReadFromNestedField()

func (re *RepoEngine[T]) ReadFromNestedJoinField(joinClause JoinClause, whereClauses []WhereClause, model *T, conditions ...*T) error {
	session := re.eng.Table(&model).Join(string(joinClause.JoinType), joinClause.TableName, joinClause.Condition, joinClause.Args...)
	var whereInit bool = false
	for _, whereClause := range whereClauses {
		if !whereInit {
			session = session.Where(joinClause.TableName+"."+whereClause.ParamName+"= ?", whereClause.ParamValue)
			whereInit = true
		} else {
			session = session.And(joinClause.TableName+"."+whereClause.ParamName+"= ?", whereClause.ParamValue)
		}
	}
	modelMap, err := readOne(session, conditions...)
	if err != nil {
		return err
	}
	for _, value := range modelMap {
		*model = *value
		break
	}
	return nil
}

func (re *RepoEngine[T]) ReadAllFromNestedJoinField(joinClause JoinClause, whereClauses []WhereClause, model *T, conditions ...*T) ([]*T, error) {
	session := re.eng.Table(model).Join(string(joinClause.JoinType), joinClause.TableName, joinClause.Condition, joinClause.Args...)
	var whereInit bool = false
	for _, whereClause := range whereClauses {
		if !whereInit {
			session = session.Where(joinClause.TableName+"."+whereClause.ParamName+"= ?", whereClause.ParamValue)
			whereInit = true
		} else {
			session = session.And(joinClause.TableName+"."+whereClause.ParamName+"= ?", whereClause.ParamValue)
		}
	}
	modelArr, err := readAll(session, conditions...)
	if err != nil {
		return nil, err
	}
	return modelArr, nil
}

func readOne[T any](session *xorm.Session, conditions ...*T) (map[string]*T, error) {
	modelMap := make(map[string]*T)
	return modelMap, session.Find(&modelMap, conditions)
}
func readAll[T any](session *xorm.Session, conditions ...*T) ([]*T, error) {
	modelArr := make([]*T, 0)
	condiBeans := make([]interface{}, len(conditions))
	for i, cond := range conditions {
		condiBeans[i] = cond
	}

	return modelArr, session.Find(&modelArr, condiBeans...)
}
func (re *RepoEngine[T]) Update() {}
func (re *RepoEngine[T]) Delete() {}
