tests := []struct {
	...
	apiMock  *rest.Mock
}{
	{
		...
		&rest.Mock{
			URL:          "http://api.internal.ml.com/shipments/1?caller.id=2"
			HTTPMethod:   http.MethodGet, RespHTTPCode: http.StatusOK,
			RespBody:     `{ "some" : "response" }`,
		},
	}
}
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		if tt.apiMock != nil {
			rest.AddMockups(tt.apiMock)
			defer rest.FlushMockups()
		}
		if got := GetShipment(tt.args.shippingID, tt.args.senderID); (got != nil) != tt.wantResp {
			t.Errorf("GetShipment() = %v, want %v", got, tt.wantResp)
		}
	})
}