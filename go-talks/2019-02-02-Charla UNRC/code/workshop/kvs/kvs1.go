package kvs

import "github.com/mercadolibre/go-meli-toolkit/gokvsclient"

func Get(key string) (gokvsclient.Item, error) {
	return kvsClient.Get(key)
}

func Update(kvsItem gokvsclient.Item) error {
	return kvsClient.Update(kvsItem)
}

func Save(kvsItem gokvsclient.Item) error {
	return kvsClient.Save(kvsItem)
}

func Delete(key string) error {
	return kvsClient.Delete(key)
}
