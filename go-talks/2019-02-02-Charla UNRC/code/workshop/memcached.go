import "github.com/mercadolibre/go-meli-toolkit/gomemcached"
...
var memcached gomemcached.Client
...
func init() {

	if os.Getenv("GO_ENVIRONMENT") != "production" {
		gomemcached.RegisterCluster("go-sanluis-cache")
	} else {
		gomemcached.RegisterCluster("go-sanluis-cache", os.Getenv("CACHE_KEYS_CONFIG_ENDPOINT"))
	}

	//Create a new client:
	client, err := gomemcached.NewClient("go-sanluis-cache")
	if err != nil {
		logger.Panicf(
			"No se puede conectar a los servers %s - ENV:%s - SCOPE:%s", err, os.Getenv("CACHE_KEYS_CONFIG_ENDPOINT"), os.Getenv("GO_ENVIROMENT"), os.Getenv("SCOPE"),
		)
	}

	//Set default expiration time:
	client.SetExpiration(600) // 10 minutos
	memcached = client        // expongo el cliente recien creado
}