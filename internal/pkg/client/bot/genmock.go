//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package bot_client

type ClientMock interface {
	IClient
}
