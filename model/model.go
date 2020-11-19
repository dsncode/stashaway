package model

type Transacction struct {
	transacctionID string
	amount         int16
}

type Portfolio struct {
	portfolioID string
	name        string
	total       int16
	maxAmount   int16
}

type DespositPlan struct {
	depositPlanID   string
	depositPlanType string // montly / one-time
	customerID      string
	portfolio       []*Portfolio
}
