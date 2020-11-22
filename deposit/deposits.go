package deposit

import "github.com/dsncode/stash/model"

// ComputeSavingsDistribution adds up all the deposits for their respective portfolios
func ComputeSavingsDistribution(depositPlans []*model.DespositPlan, deposits []model.Deposit) (portfolios []*model.Portfolio) {

	portfolioMap := make(map[string]*model.Portfolio)

	defaultPortfolio := model.CreatePortfolio("Default", model.Default)
	portfolioMap[defaultPortfolio.ID] = defaultPortfolio
	portfolios = append(portfolios, defaultPortfolio)

	for _, deposit := range deposits {

		totalDeposit := deposit.Amount

		for _, depositPlan := range depositPlans {

			// do not consider single time deposit plan that have received a deposit already
			if totalDeposit == 0 || (depositPlan.DepositPlanType == model.SingleTime && depositPlan.FirstUpdateComplete) {
				continue
			}

			for _, plan := range depositPlan.PortfolioPlan {

				// we should not deposit here
				if plan.MaxAmountToDeposit == 0 {
					continue
				}

				if totalDeposit <= plan.MaxAmountToDeposit {
					plan.Portfolio.Total = plan.Portfolio.Total + totalDeposit
					totalDeposit = 0
				} else {
					plan.Portfolio.Total = plan.Portfolio.Total + plan.MaxAmountToDeposit
					totalDeposit = totalDeposit - plan.MaxAmountToDeposit
				}

				// aggregate portfolio to response
				if _, exists := portfolioMap[plan.Portfolio.ID]; exists == false {
					portfolioMap[plan.Portfolio.ID] = plan.Portfolio
					portfolios = append(portfolios, plan.Portfolio)
				}

			}
			depositPlan.FirstUpdateComplete = true
		}

		// which means that, not all the money was allocated on all portfolios
		// in this case, we allocate it into a risk free 'default' account.
		if totalDeposit > 0 {
			defaultPortfolio.Total = defaultPortfolio.Total + totalDeposit
		}

	}
	return
}
