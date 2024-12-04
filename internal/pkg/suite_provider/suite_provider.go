package suite_provider

import (
	"os"
	"syscall"
	"tinvest-go/internal/pkg/closer"
)

type suiteProvider struct {
	closer       closer.Closer
	repositories repositories
	services     services
	clients      clients
}

func NewSuiteProvider() (*suiteProvider, func()) {
	sp := &suiteProvider{}
	return sp, sp.Close
}

func (sp *suiteProvider) GetCloser() closer.Closer {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}

func (sp *suiteProvider) Close() {
	sp.GetCloser().CloseAll()
}
