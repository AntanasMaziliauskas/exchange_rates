package rates_test

import (
	"testing"

	"github.com/AntanasMaziliauskas/exchange_rates/rates"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestPercentage(t *testing.T) {
	testTable := []struct {
		firstVar  float64
		secondVar float64
		expect    float64
	}{
		{
			firstVar:  5.4465840,
			secondVar: 5.5465484,
			expect:    1.8023,
		},
		{
			firstVar:  100,
			secondVar: 150,
			expect:    33.3333,
		},
		{
			firstVar:  3,
			secondVar: 2,
			expect:    -50,
		},
	}
	for _, v := range testTable {
		result := rates.Percentage(v.firstVar, v.secondVar)
		assert.Equal(t, result, v.expect)
	}
}
