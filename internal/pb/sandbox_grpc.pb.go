// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: sandbox.proto

package contractv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SandboxService_OpenSandboxAccount_FullMethodName           = "/tinkoff.public.invest.api.contract.v1.SandboxService/OpenSandboxAccount"
	SandboxService_GetSandboxAccounts_FullMethodName           = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxAccounts"
	SandboxService_CloseSandboxAccount_FullMethodName          = "/tinkoff.public.invest.api.contract.v1.SandboxService/CloseSandboxAccount"
	SandboxService_PostSandboxOrder_FullMethodName             = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrder"
	SandboxService_ReplaceSandboxOrder_FullMethodName          = "/tinkoff.public.invest.api.contract.v1.SandboxService/ReplaceSandboxOrder"
	SandboxService_GetSandboxOrders_FullMethodName             = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrders"
	SandboxService_CancelSandboxOrder_FullMethodName           = "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxOrder"
	SandboxService_GetSandboxOrderState_FullMethodName         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderState"
	SandboxService_GetSandboxPositions_FullMethodName          = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPositions"
	SandboxService_GetSandboxOperations_FullMethodName         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperations"
	SandboxService_GetSandboxOperationsByCursor_FullMethodName = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperationsByCursor"
	SandboxService_GetSandboxPortfolio_FullMethodName          = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPortfolio"
	SandboxService_SandboxPayIn_FullMethodName                 = "/tinkoff.public.invest.api.contract.v1.SandboxService/SandboxPayIn"
	SandboxService_GetSandboxWithdrawLimits_FullMethodName     = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxWithdrawLimits"
	SandboxService_GetSandboxMaxLots_FullMethodName            = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxMaxLots"
)

// SandboxServiceClient is the client API for SandboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandboxServiceClient interface {
	// Зарегистрировать счёт.
	OpenSandboxAccount(ctx context.Context, in *OpenSandboxAccountRequest, opts ...grpc.CallOption) (*OpenSandboxAccountResponse, error)
	// Получить счета.
	GetSandboxAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
	// Закрыть счёт.
	CloseSandboxAccount(ctx context.Context, in *CloseSandboxAccountRequest, opts ...grpc.CallOption) (*CloseSandboxAccountResponse, error)
	// Выставить торговое поручение.
	PostSandboxOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error)
	// Изменить выставленную заявку.
	ReplaceSandboxOrder(ctx context.Context, in *ReplaceOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error)
	// Получить список активных заявок по счёту.
	GetSandboxOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	// Отменить торговое поручение.
	CancelSandboxOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	// Поулчить статус заявки в песочнице. Заявки хранятся в таблице 7 дней.
	GetSandboxOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error)
	// Получить позиции по виртуальному счёту.
	GetSandboxPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error)
	// Получить операции по номеру счёта.
	GetSandboxOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
	// Получить операции по номеру счёта с пагинацией.
	GetSandboxOperationsByCursor(ctx context.Context, in *GetOperationsByCursorRequest, opts ...grpc.CallOption) (*GetOperationsByCursorResponse, error)
	// Получить портфель.
	GetSandboxPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error)
	// Пополнить счёт.
	SandboxPayIn(ctx context.Context, in *SandboxPayInRequest, opts ...grpc.CallOption) (*SandboxPayInResponse, error)
	// Получить доступный остаток для вывода средств.
	GetSandboxWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error)
	// Расчёт количества доступных для покупки/продажи лотов в песочнице.
	GetSandboxMaxLots(ctx context.Context, in *GetMaxLotsRequest, opts ...grpc.CallOption) (*GetMaxLotsResponse, error)
}

type sandboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSandboxServiceClient(cc grpc.ClientConnInterface) SandboxServiceClient {
	return &sandboxServiceClient{cc}
}

func (c *sandboxServiceClient) OpenSandboxAccount(ctx context.Context, in *OpenSandboxAccountRequest, opts ...grpc.CallOption) (*OpenSandboxAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenSandboxAccountResponse)
	err := c.cc.Invoke(ctx, SandboxService_OpenSandboxAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) CloseSandboxAccount(ctx context.Context, in *CloseSandboxAccountRequest, opts ...grpc.CallOption) (*CloseSandboxAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseSandboxAccountResponse)
	err := c.cc.Invoke(ctx, SandboxService_CloseSandboxAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) PostSandboxOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostOrderResponse)
	err := c.cc.Invoke(ctx, SandboxService_PostSandboxOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) ReplaceSandboxOrder(ctx context.Context, in *ReplaceOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostOrderResponse)
	err := c.cc.Invoke(ctx, SandboxService_ReplaceSandboxOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) CancelSandboxOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, SandboxService_CancelSandboxOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderState)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxOrderState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PositionsResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxPositions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOperationsByCursor(ctx context.Context, in *GetOperationsByCursorRequest, opts ...grpc.CallOption) (*GetOperationsByCursorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOperationsByCursorResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxOperationsByCursor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PortfolioResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxPortfolio_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) SandboxPayIn(ctx context.Context, in *SandboxPayInRequest, opts ...grpc.CallOption) (*SandboxPayInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SandboxPayInResponse)
	err := c.cc.Invoke(ctx, SandboxService_SandboxPayIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WithdrawLimitsResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxWithdrawLimits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxMaxLots(ctx context.Context, in *GetMaxLotsRequest, opts ...grpc.CallOption) (*GetMaxLotsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMaxLotsResponse)
	err := c.cc.Invoke(ctx, SandboxService_GetSandboxMaxLots_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandboxServiceServer is the server API for SandboxService service.
// All implementations must embed UnimplementedSandboxServiceServer
// for forward compatibility.
type SandboxServiceServer interface {
	// Зарегистрировать счёт.
	OpenSandboxAccount(context.Context, *OpenSandboxAccountRequest) (*OpenSandboxAccountResponse, error)
	// Получить счета.
	GetSandboxAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
	// Закрыть счёт.
	CloseSandboxAccount(context.Context, *CloseSandboxAccountRequest) (*CloseSandboxAccountResponse, error)
	// Выставить торговое поручение.
	PostSandboxOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error)
	// Изменить выставленную заявку.
	ReplaceSandboxOrder(context.Context, *ReplaceOrderRequest) (*PostOrderResponse, error)
	// Получить список активных заявок по счёту.
	GetSandboxOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	// Отменить торговое поручение.
	CancelSandboxOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	// Поулчить статус заявки в песочнице. Заявки хранятся в таблице 7 дней.
	GetSandboxOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error)
	// Получить позиции по виртуальному счёту.
	GetSandboxPositions(context.Context, *PositionsRequest) (*PositionsResponse, error)
	// Получить операции по номеру счёта.
	GetSandboxOperations(context.Context, *OperationsRequest) (*OperationsResponse, error)
	// Получить операции по номеру счёта с пагинацией.
	GetSandboxOperationsByCursor(context.Context, *GetOperationsByCursorRequest) (*GetOperationsByCursorResponse, error)
	// Получить портфель.
	GetSandboxPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error)
	// Пополнить счёт.
	SandboxPayIn(context.Context, *SandboxPayInRequest) (*SandboxPayInResponse, error)
	// Получить доступный остаток для вывода средств.
	GetSandboxWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error)
	// Расчёт количества доступных для покупки/продажи лотов в песочнице.
	GetSandboxMaxLots(context.Context, *GetMaxLotsRequest) (*GetMaxLotsResponse, error)
	mustEmbedUnimplementedSandboxServiceServer()
}

// UnimplementedSandboxServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSandboxServiceServer struct{}

func (UnimplementedSandboxServiceServer) OpenSandboxAccount(context.Context, *OpenSandboxAccountRequest) (*OpenSandboxAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSandboxAccount not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxAccounts not implemented")
}
func (UnimplementedSandboxServiceServer) CloseSandboxAccount(context.Context, *CloseSandboxAccountRequest) (*CloseSandboxAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSandboxAccount not implemented")
}
func (UnimplementedSandboxServiceServer) PostSandboxOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSandboxOrder not implemented")
}
func (UnimplementedSandboxServiceServer) ReplaceSandboxOrder(context.Context, *ReplaceOrderRequest) (*PostOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceSandboxOrder not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOrders not implemented")
}
func (UnimplementedSandboxServiceServer) CancelSandboxOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSandboxOrder not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOrderState not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxPositions(context.Context, *PositionsRequest) (*PositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxPositions not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOperations(context.Context, *OperationsRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOperations not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOperationsByCursor(context.Context, *GetOperationsByCursorRequest) (*GetOperationsByCursorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOperationsByCursor not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxPortfolio not implemented")
}
func (UnimplementedSandboxServiceServer) SandboxPayIn(context.Context, *SandboxPayInRequest) (*SandboxPayInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SandboxPayIn not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxWithdrawLimits not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxMaxLots(context.Context, *GetMaxLotsRequest) (*GetMaxLotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxMaxLots not implemented")
}
func (UnimplementedSandboxServiceServer) mustEmbedUnimplementedSandboxServiceServer() {}
func (UnimplementedSandboxServiceServer) testEmbeddedByValue()                        {}

// UnsafeSandboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandboxServiceServer will
// result in compilation errors.
type UnsafeSandboxServiceServer interface {
	mustEmbedUnimplementedSandboxServiceServer()
}

func RegisterSandboxServiceServer(s grpc.ServiceRegistrar, srv SandboxServiceServer) {
	// If the following call pancis, it indicates UnimplementedSandboxServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SandboxService_ServiceDesc, srv)
}

func _SandboxService_OpenSandboxAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSandboxAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).OpenSandboxAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_OpenSandboxAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).OpenSandboxAccount(ctx, req.(*OpenSandboxAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_CloseSandboxAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSandboxAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).CloseSandboxAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_CloseSandboxAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).CloseSandboxAccount(ctx, req.(*CloseSandboxAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_PostSandboxOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).PostSandboxOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_PostSandboxOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).PostSandboxOrder(ctx, req.(*PostOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_ReplaceSandboxOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).ReplaceSandboxOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_ReplaceSandboxOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).ReplaceSandboxOrder(ctx, req.(*ReplaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_CancelSandboxOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).CancelSandboxOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_CancelSandboxOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).CancelSandboxOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxOrderState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOrderState(ctx, req.(*GetOrderStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxPositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxPositions(ctx, req.(*PositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOperations(ctx, req.(*OperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOperationsByCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsByCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOperationsByCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxOperationsByCursor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOperationsByCursor(ctx, req.(*GetOperationsByCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxPortfolio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxPortfolio(ctx, req.(*PortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_SandboxPayIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxPayInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).SandboxPayIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_SandboxPayIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).SandboxPayIn(ctx, req.(*SandboxPayInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxWithdrawLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxWithdrawLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxWithdrawLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxWithdrawLimits(ctx, req.(*WithdrawLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxMaxLots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxLotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxMaxLots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxService_GetSandboxMaxLots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxMaxLots(ctx, req.(*GetMaxLotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SandboxService_ServiceDesc is the grpc.ServiceDesc for SandboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SandboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.SandboxService",
	HandlerType: (*SandboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSandboxAccount",
			Handler:    _SandboxService_OpenSandboxAccount_Handler,
		},
		{
			MethodName: "GetSandboxAccounts",
			Handler:    _SandboxService_GetSandboxAccounts_Handler,
		},
		{
			MethodName: "CloseSandboxAccount",
			Handler:    _SandboxService_CloseSandboxAccount_Handler,
		},
		{
			MethodName: "PostSandboxOrder",
			Handler:    _SandboxService_PostSandboxOrder_Handler,
		},
		{
			MethodName: "ReplaceSandboxOrder",
			Handler:    _SandboxService_ReplaceSandboxOrder_Handler,
		},
		{
			MethodName: "GetSandboxOrders",
			Handler:    _SandboxService_GetSandboxOrders_Handler,
		},
		{
			MethodName: "CancelSandboxOrder",
			Handler:    _SandboxService_CancelSandboxOrder_Handler,
		},
		{
			MethodName: "GetSandboxOrderState",
			Handler:    _SandboxService_GetSandboxOrderState_Handler,
		},
		{
			MethodName: "GetSandboxPositions",
			Handler:    _SandboxService_GetSandboxPositions_Handler,
		},
		{
			MethodName: "GetSandboxOperations",
			Handler:    _SandboxService_GetSandboxOperations_Handler,
		},
		{
			MethodName: "GetSandboxOperationsByCursor",
			Handler:    _SandboxService_GetSandboxOperationsByCursor_Handler,
		},
		{
			MethodName: "GetSandboxPortfolio",
			Handler:    _SandboxService_GetSandboxPortfolio_Handler,
		},
		{
			MethodName: "SandboxPayIn",
			Handler:    _SandboxService_SandboxPayIn_Handler,
		},
		{
			MethodName: "GetSandboxWithdrawLimits",
			Handler:    _SandboxService_GetSandboxWithdrawLimits_Handler,
		},
		{
			MethodName: "GetSandboxMaxLots",
			Handler:    _SandboxService_GetSandboxMaxLots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sandbox.proto",
}
