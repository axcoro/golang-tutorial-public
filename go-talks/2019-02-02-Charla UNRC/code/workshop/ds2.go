import (
    "github.com/mercadolibre/go-meli-toolkit/godsclient/querybuilders"
    "github.com/mercadolibre/go-meli-toolkit/godsclient/sortbuilders"
    "github.com/mercadolibre/go-meli-toolkit/godsclient/aggbuilders"
    "github.com/mercadolibre/go-meli-toolkit/godsclient/fieldtype"
    "github.com/mercadolibre/go-meli-toolkit/godsclient/sortorder"
)

queryBuilder := querybuilders.And(
    querybuilders.Eq("my_field1", "my_value1"),
    querybuilders.Or(
            querybuilders.Eq("my_field2", "my_value2"),
            querybuilders.Not(querybuilders.Match("my_field3", "my_value3")),
    ),
)

response, err := dsClient.SearchBuilder().
    AddSort(sortbuilders.Field("creation_date", fieldtype.Date).Order(sortorder.Asc)).
    AddSort(sortbuilders.Field("other_field", fieldtype.Number).Order(sortorder.Desc)).
    AddAggregation(aggbuilders.Terms("my_aggregation").Field("test", fieldtype.Number)).
    AddProjections("field1", "field2").
    WithQuery(queryBuilder).
    WithSize(100).
    Execute()