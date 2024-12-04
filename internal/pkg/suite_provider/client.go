package suite_provider

import (
	cbrf_client "tinvest-go/internal/pkg/client/cbrf"
	cbrf_client_mock "tinvest-go/internal/pkg/client/cbrf/mocks"
	tinvest_client "tinvest-go/internal/pkg/client/tinvest"
	tinvest_client_mock "tinvest-go/internal/pkg/client/tinvest/mocks"
)

type clients struct {
	cbrfClientMock *cbrf_client_mock.ClientMock
	cbrfClient     cbrf_client.IClient

	tinvestInstrumentsAPIMock *tinvest_client_mock.InstrumentsAPIMock
	tinvestMarketDataAPIMock  *tinvest_client_mock.MarketDataAPIMock
	tinvestOperationsAPIMock  *tinvest_client_mock.OperationsAPIMock
	tinvestOrdersAPIMock      *tinvest_client_mock.OrdersAPIMock
	tinvestSignalsAPIMock     *tinvest_client_mock.SignalsAPIMock
	tinvestSUsersAPIMock      *tinvest_client_mock.UsersAPIMock
	tinvestClient             tinvest_client.IClient
}

func (sp *suiteProvider) GetCbrfClientMock() *cbrf_client_mock.ClientMock {
	if sp.clients.cbrfClientMock == nil {
		sp.clients.cbrfClientMock = &cbrf_client_mock.ClientMock{}
	}
	return sp.clients.cbrfClientMock
}

func (sp *suiteProvider) GetCbrfClient() cbrf_client.IClient {
	if sp.clients.cbrfClient == nil {
		sp.clients.cbrfClient = sp.GetCbrfClientMock()
	}
	return sp.clients.cbrfClient
}

func (sp *suiteProvider) GetTinvestInstrumentsAPIMockMock() *tinvest_client_mock.InstrumentsAPIMock {
	if sp.clients.tinvestInstrumentsAPIMock == nil {
		sp.clients.tinvestInstrumentsAPIMock = &tinvest_client_mock.InstrumentsAPIMock{}
	}
	return sp.clients.tinvestInstrumentsAPIMock
}

func (sp *suiteProvider) GetTinvestMarketDataAPIMockMock() *tinvest_client_mock.MarketDataAPIMock {
	if sp.clients.tinvestMarketDataAPIMock == nil {
		sp.clients.tinvestMarketDataAPIMock = &tinvest_client_mock.MarketDataAPIMock{}
	}
	return sp.clients.tinvestMarketDataAPIMock
}

func (sp *suiteProvider) GetTinvestOperationsAPIMockMock() *tinvest_client_mock.OperationsAPIMock {
	if sp.clients.tinvestOperationsAPIMock == nil {
		sp.clients.tinvestOperationsAPIMock = &tinvest_client_mock.OperationsAPIMock{}
	}
	return sp.clients.tinvestOperationsAPIMock
}

func (sp *suiteProvider) GetTinvestOrdersAPIMockMock() *tinvest_client_mock.OrdersAPIMock {
	if sp.clients.tinvestOrdersAPIMock == nil {
		sp.clients.tinvestOrdersAPIMock = &tinvest_client_mock.OrdersAPIMock{}
	}
	return sp.clients.tinvestOrdersAPIMock
}

func (sp *suiteProvider) GetTinvestSignalsAPIMockMock() *tinvest_client_mock.SignalsAPIMock {
	if sp.clients.tinvestSignalsAPIMock == nil {
		sp.clients.tinvestSignalsAPIMock = &tinvest_client_mock.SignalsAPIMock{}
	}
	return sp.clients.tinvestSignalsAPIMock
}

func (sp *suiteProvider) GetTinvestUsersAPIMockMock() *tinvest_client_mock.UsersAPIMock {
	if sp.clients.tinvestSUsersAPIMock == nil {
		sp.clients.tinvestSUsersAPIMock = &tinvest_client_mock.UsersAPIMock{}
	}
	return sp.clients.tinvestSUsersAPIMock
}

func (sp *suiteProvider) GetTinvestClient() tinvest_client.IClient {
	if sp.clients.tinvestClient == nil {
		sp.clients.tinvestClient = &tinvest_client.Client{
			InstrumentsAPI: sp.GetTinvestInstrumentsAPIMockMock(),
			MarketDataAPI:  sp.GetTinvestMarketDataAPIMockMock(),
			OperationsAPI:  sp.GetTinvestOperationsAPIMockMock(),
			OrdersAPI:      sp.GetTinvestOrdersAPIMockMock(),
			SignalsAPI:     sp.GetTinvestSignalsAPIMockMock(),
			UsersAPI:       sp.GetTinvestUsersAPIMockMock(),
		}
	}
	return sp.clients.tinvestClient
}
