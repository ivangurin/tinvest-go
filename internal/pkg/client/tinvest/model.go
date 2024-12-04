package tinvest_client

type PortfolioPosition struct {
	ID           string
	Type         string
	CurrentPrice float64
	CurrencyCode string
}

type PortfolioPositions []*PortfolioPosition
