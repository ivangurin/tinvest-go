//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package cbrf_client

type ClientMock interface {
	IClient
}
