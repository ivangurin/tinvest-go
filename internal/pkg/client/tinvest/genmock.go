//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package tinvest_client

import contractv1 "tinvest-go/internal/pb"

type InstrumentsAPIMock interface {
	contractv1.InstrumentsServiceClient
}

type MarketDataAPIMock interface {
	contractv1.MarketDataServiceClient
}

type OperationsAPIMock interface {
	contractv1.OperationsServiceClient
}

type OrdersAPIMock interface {
	contractv1.OrdersServiceClient
}

type SignalsAPIMock interface {
	contractv1.SignalServiceClient
}

type UsersAPIMock interface {
	contractv1.UsersServiceClient
}
