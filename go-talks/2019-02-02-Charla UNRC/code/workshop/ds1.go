import "github.com/mercadolibre/go-meli-toolkit/godsclient"

var dsClient godsclient.Client

func init() {
    conf := godsclient.NewDsClientConfig().WithServiceName("go-sanluis-ds")
    dsClient := godsclient.NewEntityClient(conf)
}
...
document := map[string]string{"key":"value"}
err := dsClient.SaveDocument("123453", document)
if err != nil {
    ...
}
...