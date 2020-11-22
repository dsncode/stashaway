package model

import "github.com/google/uuid"

type Deposit struct {
	Amount int16
	Month  int
	Year   int
}

type Portfolio struct {
	ID            string
	Name          string
	Total         int16
	PortfolioType PortfolioType
}

// PortfolioType indicate if is a standard of default porfolio
type PortfolioType int

// DepositPlanType single time or montly
type DepositPlanType int

const (
	// Default all remaining money goes here
	Default PortfolioType = 0
	// Standard portolio type
	Standard PortfolioType = 1

	// SingleTime once it fills up. stop depositing
	SingleTime DepositPlanType = 1
	// Montly it captures income per month
	Montly DepositPlanType = 2
)

// DespositPlan for a customer
type DespositPlan struct {
	Name                string
	DepositPlanType     DepositPlanType
	CustomerID          string
	PortfolioPlan       []*PortfolioPlan
	FirstUpdateComplete bool
}

// PortfolioPlan indicates a porfolio and its max funding limits
type PortfolioPlan struct {
	Portfolio          *Portfolio
	MaxAmountToDeposit int16
}

// CreatePortfolio builds a portfolio for a customer
func CreatePortfolio(name string, portfolioType PortfolioType) (portfolio *Portfolio) {

	portfolio = &Portfolio{
		ID:            uuid.New().String(),
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
