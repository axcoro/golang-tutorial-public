import "github.com/mercadolibre/go-meli-toolkit/restful/rest"

var (
	urlBase        = "http://internal.mercadolibre.com"
	defaultTimeout = time.Duration(1000)
	restClient rest.RequestBuilder
)

func init() {
	if os.Getenv("GO_ENVIRONMENT") != "production" {
		urlBase = "http://api.internal.ml.com"
		defaultTimeout = 3000
	}
	customPool := &rest.CustomPool{MaxIdleConnsPerHost: 100}
	restClient  = rest.RequestBuilder{
		Timeout:        defaultTimeout * time.Millisecond,
		BaseURL:        urlBase,
		ContentType:    rest.JSON,
		DisableCache:   false,
		DisableTimeout: false,
		CustomPool:     customPool,
	}
}