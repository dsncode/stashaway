package model

import (
	"time"
)

type Deposit struct {
	Amount int16
	Month  int
	Year   int
}

type Portfolio struct {
	Name          string
	Total         int16
	PortfolioType PortfolioType
	LastDeposit   time.Time
}

// PortfolioType indicate if is a standard of default porfolio
type PortfolioType int

// DepositPlanType single time or montly
type DepositPlanType int

// MaxAmountToDepositPorfolio for a specific plan
type MaxAmountToDepositPorfolio int16

const (
	// Default all remaining money goes here
	Default PortfolioType = 0
	// Standard portolio type
	Standard PortfolioType = 1

	// SingleTime once it fills up. stop depositing
	SingleTime DepositPlanType = 1
	// Montly it captures income per month
	Montly DepositPlanType = 2

	// UnlimitedAmount indicates that, there is no max cap for this portolio
	UnlimitedAmount MaxAmountToDepositPorfolio = -1
)

// DespositPlan for a customer
type DespositPlan struct {
	Name            string
	DepositPlanType DepositPlanType
	CustomerID      string
	PortfolioPlan   []*PortfolioPlan
}

// PortfolioPlan indicates a porfolio and its max funding limits
type PortfolioPlan struct {
	Portfolio          *Portfolio
	MaxAmountToDeposit MaxAmountToDepositPorfolio
}

// CreatePortfolio builds a portfolio for a customer
func CreatePortfolio(name string, portfolioType PortfolioType) (portfolio *Portfolio) {

	portfolio = &Portfolio{
		Name:          name,
		Total:         0,
		PortfolioType: portfolioType,
	}
	return
}

func CreateDepositPlan(name string, porfolios []*PortfolioPlan, depositPlanType DepositPlanType) (depositPlan *DespositPlan) {

	depositPlan = &DespositPlan{
		Name:            name,
		PortfolioPlan:   porfolios,
		DepositPlanType: depositPlanType,
	}
	return
}
