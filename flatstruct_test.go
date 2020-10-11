package flatstruct

import (
	"encoding/json"
	"strconv"
	"testing"
)

func BenchmarkWithReflection(b *testing.B) {
	sampleStruct := sampleStruct{}
	if err := sampleStruct.FillStruct(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := StructToFlatMap(sampleStruct); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithoutReflection(b *testing.B) {
	sampleStruct := sampleStruct{}
	if err := sampleStruct.FillStruct(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = sampleStruct.ToFlatMap()
	}
}

type sampleStruct struct {
	ID            string `json:"id"`
	AppName       string `json:"app_name"`
	Time          int64  `json:"time"`
	Type          string `json:"type"`
	ClientID      string `json:"client_id"`
	OrderID       string `json:"order_id"`
	TariffGroupId int    `json:"tariff_group_id"`
}

func (event *sampleStruct) ToFlatMap() map[string]interface{} {
	return map[string]interface{}{
		"Id":      event.ID,
		"Type":    event.Type,
		"TS":      event.Time,
		"AppName": event.AppName,
		"Keys":    []string{"client_id", "order_id", "tariff_group_id"},
		"Values":  []string{event.ClientID, event.OrderID, strconv.FormatInt(int64(event.TariffGroupId), 10)},
	}
}

func (event *sampleStruct) FillStruct() error {
	jsonDataBytes := []byte(`
{
	"id": "some-event-id",
	"app_name": "some-app-name",
	"time": 1602445028,
	"type": "test_event",
	"client_id": "some-client_id",
	"order_id": "some-order_id",
	"tariff_group_id" : 45231
}
`)

	return json.Unmarshal(jsonDataBytes, event)
}
