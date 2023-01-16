package logic

import (
	"fmt"
	"github.com/PereRohit/util/log"
	respModel "github.com/PereRohit/util/model"
	"github.com/vatsal278/AccountManagmentSvc/internal/codes"
	"github.com/vatsal278/AccountManagmentSvc/internal/config"
	"github.com/vatsal278/AccountManagmentSvc/internal/model"
	jwtSvc "github.com/vatsal278/AccountManagmentSvc/internal/repo/authentication"
	"github.com/vatsal278/AccountManagmentSvc/internal/repo/datasource"
	"net/http"
)

//go:generate mockgen --build_flags=--mod=mod --destination=./../../pkg/mock/mock_logic.go --package=mock github.com/vatsal278/AccountManagmentSvc/internal/logic AccountManagmentSvcLogicIer

type AccountManagmentSvcLogicIer interface {
	HealthCheck() bool
	CreateAccount(account model.NewAccount) *respModel.Response
	AccountDetails(id string) *respModel.Response
}

type accountManagmentSvcLogic struct {
	DsSvc      datasource.DataSourceI
	jwtService jwtSvc.JWTService
	msgQueue   config.MsgQueue
	cookie     config.CookieStruct
}

func NewAccountManagmentSvcLogic(ds datasource.DataSourceI, jwtService jwtSvc.JWTService, msgQueue config.MsgQueue, cookie config.CookieStruct) AccountManagmentSvcLogicIer {
	return &accountManagmentSvcLogic{
		DsSvc:      ds,
		jwtService: jwtService,
		msgQueue:   msgQueue,
		cookie:     cookie,
	}
}

func (l accountManagmentSvcLogic) HealthCheck() bool {
	// check all internal services are working fine
	return l.DsSvc.HealthCheck()
}

func (l accountManagmentSvcLogic) CreateAccount(account model.NewAccount) *respModel.Response {
	result, err := l.DsSvc.Get(map[string]interface{}{"user_id": account.UserId})
	if err != nil {
		log.Error(err.Error())
		return &respModel.Response{
			Status:  http.StatusInternalServerError,
			Message: codes.GetErr(codes.ErrCreatingAccount),
			Data:    nil,
		}
	}
	if len(result) != 0 {
		log.Error(codes.GetErr(codes.ErrAccExists))
		return &respModel.Response{
			Status:  http.StatusBadRequest,
			Message: codes.GetErr(codes.ErrAccExists),
			Data:    nil,
		}
	}
	err = l.DsSvc.Insert(model.Account{Id: account.UserId})
	if err != nil {
		log.Error(codes.GetErr(codes.ErrCreatingAccount))
		return &respModel.Response{
			Status:  http.StatusBadRequest,
			Message: codes.GetErr(codes.ErrCreatingAccount),
			Data:    nil,
		}
	}
	go func(userId string, pubId string, channel string) {
		userID := fmt.Sprintf(`{"user_id":"%s"}`, userId)
		err := l.msgQueue.MsgBroker.PushMsg(userID, pubId, channel)
		if err != nil {
			log.Error(err)
			return
		}
		return
	}(account.UserId, l.msgQueue.PubId, l.msgQueue.Channel)
	return &respModel.Response{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
		Data:    nil,
	}
}

func (l accountManagmentSvcLogic) AccountDetails(id string) *respModel.Response {
	acc, err := l.DsSvc.Get(map[string]interface{}{"user_id": id})
	if err != nil {
		log.Error(err)
		return &respModel.Response{
			Status:  http.StatusInternalServerError,
			Message: codes.GetErr(codes.ErrFetchingUser),
			Data:    nil,
		}
	}
	if len(acc) == 0 {
		return &respModel.Response{
			Status:  http.StatusBadRequest,
			Message: codes.GetErr(codes.AccNotFound),
			Data:    nil,
		}
	}
	resp := model.AccountSummary{
		AccountNumber:    acc[0].AccountNumber,
		Income:           acc[0].Income,
		Spends:           acc[0].Spends,
		ActiveServices:   acc[0].ActiveServices,
		InactiveServices: acc[0].InactiveServices,
	}
	return &respModel.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    resp,
	}
}
