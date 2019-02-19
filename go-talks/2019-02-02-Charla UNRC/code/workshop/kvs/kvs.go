package kvs

import "github.com/mercadolibre/go-meli-toolkit/gokvsclient"

var kvsClient gokvsclient.Client

func init() {
	config := gokvsclient.MakeKvsConfig()
	kvsClient = gokvsclient.MakeKvsClient("GO_SANLUIS_KVS", config)
}
