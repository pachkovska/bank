package bank

import (
	"fmt"
	"time"

	"github.com/CodingSquire/bank/pkg/api"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	svc    Bank
	logger *api.Logger
}

func (s *loggingMiddleware) AddToBell(request api.AddToBellRequest) (response api.AddToBellResponse, err error) {
	defer func(begin time.Time) {
		fmt.Println("___________________________")
		fmt.Println("method", "AddToBell")
		fmt.Println("timestamp", time.Now())
		fmt.Println("response", response)
		fmt.Println("elapsed", time.Since(begin))
		fmt.Println("err", err)
		fmt.Println("End.")
		fmt.Println("___________________________")
	}(time.Now())
	fmt.Println("___________________________")
	fmt.Println("method", "AddToBell")
	fmt.Println("timestamp", time.Now())
	fmt.Println("request", request)
	fmt.Println("Start.")
	fmt.Println("___________________________")
	response, err = s.svc.AddToBell(request)
	return response, err
}

func (s *loggingMiddleware) DeductFromBell(request api.DeductFromBellRequest) (response api.DeductFromBellResponse, err error) {
	defer func(begin time.Time) {
		fmt.Println("___________________________")
		fmt.Println("method", "DeductFromBell")
		fmt.Println("timestamp", time.Now())
		fmt.Println("response", response)
		fmt.Println("elapsed", time.Since(begin))
		fmt.Println("err", err)
		fmt.Println("End.")
		fmt.Println("___________________________")
	}(time.Now())
	fmt.Println("___________________________")
	fmt.Println("method", "DeductFromBell")
	fmt.Println("timestamp", time.Now())
	fmt.Println("request", request)
	fmt.Println("Start.")
	fmt.Println("___________________________")
	response, err = s.svc.DeductFromBell(request)
	return response, err
}

func (s *loggingMiddleware) CreateAcc(request api.CreateAccRequest) (response api.CreateAccResponse, err error) {
	defer func(begin time.Time) {
		fmt.Println("___________________________")
		fmt.Println("method", "CreateAcc")
		fmt.Println("timestamp", time.Now())
		fmt.Println("response", response)
		fmt.Println("elapsed", time.Since(begin))
		fmt.Println("err", err)
		fmt.Println("End.")
		fmt.Println("___________________________")
	}(time.Now())
	fmt.Println("___________________________")
	fmt.Println("method", "CreateAcc")
	fmt.Println("timestamp", time.Now())
	fmt.Println("request", request)
	fmt.Println("Start.")
	fmt.Println("___________________________")
	response, err = s.svc.CreateAcc(request)
	return response, err
}

func (s *loggingMiddleware) GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error) {
	defer func(begin time.Time) {
		fmt.Println("___________________________")
		fmt.Println("method", "GetBalance")
		fmt.Println("timestamp", time.Now())
		fmt.Println("response", response)
		fmt.Println("elapsed", time.Since(begin))
		fmt.Println("err", err)
		fmt.Println("End.")
		fmt.Println("___________________________")
	}(time.Now())
	fmt.Println("___________________________")
	fmt.Println("method", "GetBalance")
	fmt.Println("timestamp", time.Now())
	fmt.Println("request", request)
	fmt.Println("Start.")
	fmt.Println("___________________________")
	response, err = s.svc.GetBalance(request)
	return response, err
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(svc Bank, logger *api.Logger) Bank {
	return &loggingMiddleware{
		svc:    svc,
		logger: logger,
	}
}
