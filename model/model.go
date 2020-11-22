package model

import (
	"time"

	"github.com/google/uuid"
)

type Deposit struct {
	TransacctionID string
	Amount         int16
}

type Portfolio struct {
	PortfolioID string
	Name        string
	Total       int16
	MaxAmount   MaxAmountToDeposit
	LastDeposit time.Time
}

// DepositPlanType default, single time or montly
type DepositPlanType int
type MaxAmountToDeposit int16

const (
	// Default all remaining money goes here
	Default DepositPlanType = 0
	// SingleTime once it fills up. stop depositing
	SingleTime DepositPlanType = 1
	// Montly it captures income per month
	Montly DepositPlanType = 2

	// UnlimitedAmount indicates that, there is no max cap for this portolio
	UnlimitedAmount MaxAmountToDeposit = -1
)

// DespositPlan for a customer
type DespositPlan struct {
	DepositPlanID   string
	Name            string
	DepositPlanType DepositPlanType
	CustomerID      string
	Portfolio       []*Portfolio
}

// CreatePortfolio builds a portfolio for a customer
func CreatePortfolio(name string, maxAmount MaxAmountToDeposit) (portfolio *Portfolio) {

	portfolio = &Portfolio{
		PortfolioID: uuid.New().String(),
		Name:        name,
		Total:       0,
		MaxAmount:   maxAmount,
	}
	return
}

func CreateDepositPlan(name string, porfolios []*Portfolio, maxAmount DepositPlanType) (depositPlan *DespositPlan) {

	depositPlan = &DespositPlan{
		DepositPlanID: uuid.New().String(),
		Name:          name,
		Portfolio:     porfolios,
	}
	return
}
