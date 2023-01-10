package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vatsal278/AccountManagmentSvc/internal/config"
	"github.com/vatsal278/AccountManagmentSvc/internal/model"
	"log"
	"strings"
)

type sqlDs struct {
	sqlSvc *sql.DB
	table  string
}

//docker run --rm --env MYSQL_ROOT_PASSWORD=pass --env MYSQL_DATABASE=accmgmt --publish 9085:3306 --name mysqlDb -d mysql
func NewSql(dbSvc config.DbSvc, tableName string) DataSourceI {
	return &sqlDs{
		sqlSvc: dbSvc.Db,
		table:  tableName,
	}
}

func (d sqlDs) HealthCheck() bool {
	err := d.sqlSvc.Ping()
	if err != nil {
		return false
	}
	return true
}

func (d sqlDs) Get(filter map[string]interface{}) ([]model.Account, error) {
	//order the queries based on email address
	var user model.Account
	var users []model.Account
	q := fmt.Sprintf("SELECT user_id, account_number, income, spends, created_on, updated_on, active_services, inactive_services FROM %s", d.table)

	filterClause := []string{}

	for k, v := range filter {
		switch v.(type) {
		case string:
			filterClause = append(filterClause, fmt.Sprintf("%s = '%s'", k, v))
		default:
			filterClause = append(filterClause, fmt.Sprintf("%s = %+v", k, v))
		}
	}
	if len(filterClause) > 0 {
		q += fmt.Sprintf(" WHERE %s", strings.Join(filterClause, " AND "))
	}

	q += " ORDER BY account_number;"
	rows, err := d.sqlSvc.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.AccountNumber, &user.Income, &user.Spends, &user.CreatedOn, &user.UpdatedOn, &user.ActiveServices, &user.InactiveServices)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d sqlDs) Insert(user model.Account) error {
	queryString := fmt.Sprintf("INSERT INTO %s", d.table)
	_, err := d.sqlSvc.Exec(queryString+"(user_id, active_services, inactive_services) VALUES(?,?,?)", user.Id, user.ActiveServices, user.InactiveServices)
	if err != nil {
		return err
	}
	return err
}

func (d sqlDs) Update(filterSet map[string]interface{}, filterWhere map[string]interface{}) error {
	//update accmgmt.accdatabase set active_services = JSON_INSERT(`active_services`, '$."1"','{}' ) , inactive_services = JSON_REMOVE(`inactive_services`, '$."1"') where account_number= 5;
	queryString := fmt.Sprintf("UPDATE %s ", d.table)
	filterClause := []string{}

	for k, v := range filterSet {
		switch v.(type) {
		case string:
			filterClause = append(filterClause, fmt.Sprintf("%s = '%+v'", k, v))
		case model.Svc:
			a, ok := v.(model.Svc)
			if !ok {
				return errors.New("incorrect value for model.svc")
			}
			for x, y := range a {
				filterClause = append(filterClause, fmt.Sprintf("%s = JSON_INSERT('%s', '$.%s'', '%+v')", k, k, x, y))
			}
		default:
			filterClause = append(filterClause, fmt.Sprintf("%s = %+v", k, v))
		}
	}
	if len(filterClause) > 0 {
		queryString += fmt.Sprintf(" SET %s", strings.Join(filterClause, " , "))
	}
	log.Print(queryString)
	//if strings.Contains(queryString, "active_services") || strings.Contains(queryString, "inactive_services") {
	//
	//}
	filterClauseWhere := []string{}

	for k, v := range filterWhere {
		switch v.(type) {
		case string:
			filterClauseWhere = append(filterClauseWhere, fmt.Sprintf("%s = '%+v'", k, v))
		default:
			filterClauseWhere = append(filterClauseWhere, fmt.Sprintf("%s = %+v", k, v))
		}
	}
	if len(filterClauseWhere) > 0 {
		queryString += fmt.Sprintf(" WHERE %s", strings.Join(filterClauseWhere, " AND "))
	}

	queryString += " ;"
	log.Print(queryString)
	_, err := d.sqlSvc.Exec(queryString)
	if err != nil {
		return err
	}
	return nil

}
