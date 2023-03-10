package router

import (
	"net/http"

	"github.com/PereRohit/util/constant"
	"github.com/PereRohit/util/middleware"
	"github.com/gorilla/mux"

	"github.com/vatsal278/AccountManagmentSvc/internal/config"
	"github.com/vatsal278/AccountManagmentSvc/internal/handler"
	middleware2 "github.com/vatsal278/AccountManagmentSvc/internal/middleware"
	"github.com/vatsal278/AccountManagmentSvc/internal/repo/datasource"
)

func Register(svcCfg *config.SvcConfig) *mux.Router {
	m := mux.NewRouter()

	// group all routes for specific version. e.g.: /v1
	if svcCfg.ServiceRouteVersion != "" {
		m = m.PathPrefix("/" + svcCfg.ServiceRouteVersion).Subrouter()
	}

	m.StrictSlash(true)
	m.Use(middleware.RequestHijacker)
	m.Use(middleware.RecoverPanic)

	commons := handler.NewCommonSvc()
	m.HandleFunc(constant.HealthRoute, commons.HealthCheck).Methods(http.MethodGet)
	m.NotFoundHandler = http.HandlerFunc(commons.RouteNotFound)
	m.MethodNotAllowedHandler = http.HandlerFunc(commons.MethodNotAllowed)

	// attach routes for services below
	m = attachAccountManagmentSvcRoutes(m, svcCfg)

	return m
}

func attachAccountManagmentSvcRoutes(m *mux.Router, svcCfg *config.SvcConfig) *mux.Router {
	dataSource := datasource.NewSql(svcCfg.DbSvc, svcCfg.Cfg.DataBase.TableName)
	svc := handler.NewAccountManagmentSvc(dataSource, svcCfg.JwtSvc.JwtSvc, svcCfg.MsgBrokerSvc, svcCfg.Cfg.Cookie)
	middleware := middleware2.NewAccMgmtMiddleware(svcCfg)

	route1 := m.PathPrefix("").Subrouter()
	route1.HandleFunc("", svc.CreateAccount).Methods(http.MethodPost)
	route1.Use(middleware.ScreenRequest)

	route2 := m.PathPrefix("").Subrouter()
	route2.HandleFunc("/update/service", svc.UpdateService).Methods(http.MethodPut)
	route2.Use(middleware.ExtractUser)

	route4 := m.PathPrefix("").Subrouter()
	route4.HandleFunc("", svc.AccountSummary).Methods(http.MethodGet)
	route4.Use(middleware.ExtractUser)
	route4.Use(middleware.Cacher(true))

	route3 := m.PathPrefix("").Subrouter()
	route3.HandleFunc("/update/transaction", svc.UpdateTransaction).Methods(http.MethodPut)

	return m
}
