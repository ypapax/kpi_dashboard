package kpi_dashboard

import (
	"time"

	"math/rand"
)

func timeMonthlyHandler(af apiCmd) (*Response, error) {
	var results ChartResult
	roles := []string{"QA Engineers", "Backend Engineers", "Frontend Engineers", "Machine learning Engineers"}
	for i := 0; i <= 100; i++ {
		var values []map[string]interface{}
		for _, role := range roles {
			value := map[string]interface{}{}
			value["role"] = role
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		results.Result = append(results.Result, ChartResultItem{Value: values})
	}

	for i := range results.Result {
		m := (i % 12) + 1
		d := time.Date(2017, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		results.Result[i].Timeframe.Start = d.Format(time.RFC3339)
		results.Result[i].Timeframe.Start = d.Format(time.RFC3339)
	}
	return &Response{Result: results}, nil
}
