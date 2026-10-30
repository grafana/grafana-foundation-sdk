package common_test

import (
	"encoding/json"
	"testing"

	"github.com/grafana/grafana-foundation-sdk/go/bargauge"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/gauge"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
)

func TestDefaultStatOrientation(t *testing.T) {
	for name, options := range map[string]interface{}{
		"common":   common.NewSingleStatBaseOptions(),
		"stat":     stat.NewOptions(),
		"gauge":    gauge.NewOptions(),
		"bargauge": bargauge.NewOptions(),
	} {
		t.Run(name, func(t *testing.T) {
			fields := decodeObject(t, options)
			if fields["orientation"] != string(common.VizOrientationAuto) {
				t.Fatalf("default orientation = %v, want auto", fields["orientation"])
			}
		})
	}
}

func TestMultiValuePanelOrientation(t *testing.T) {
	for name, newPanel := range map[string]func(common.VizOrientation) cog.Builder[dashboard.Panel]{
		"stat": func(orientation common.VizOrientation) cog.Builder[dashboard.Panel] {
			builder := stat.NewPanelBuilder().ReduceOptions(common.NewReduceDataOptionsBuilder().Values(true))
			if orientation != "" {
				builder.Orientation(orientation)
			}
			return builder
		},
		"gauge": func(orientation common.VizOrientation) cog.Builder[dashboard.Panel] {
			builder := gauge.NewPanelBuilder().ReduceOptions(common.NewReduceDataOptionsBuilder().Values(true))
			if orientation != "" {
				builder.Orientation(orientation)
			}
			return builder
		},
		"bargauge": func(orientation common.VizOrientation) cog.Builder[dashboard.Panel] {
			builder := bargauge.NewPanelBuilder().ReduceOptions(common.NewReduceDataOptionsBuilder().Values(true))
			if orientation != "" {
				builder.Orientation(orientation)
			}
			return builder
		},
	} {
		for _, orientation := range []common.VizOrientation{"", common.VizOrientationHorizontal, common.VizOrientationVertical} {
			t.Run(name+"/"+string(orientation), func(t *testing.T) {
				panel, err := newPanel(orientation).Build()
				if err != nil {
					t.Fatal(err)
				}
				options := decodeObject(t, panel)["options"].(map[string]interface{})
				want := orientation
				if want == "" {
					want = common.VizOrientationAuto
				}
				if options["orientation"] != string(want) {
					t.Fatalf("panel orientation = %v, want %s", options["orientation"], want)
				}
				if options["reduceOptions"].(map[string]interface{})["values"] != true {
					t.Fatal("panel must retain multi-value rendering")
				}
			})
		}
	}
}

func decodeObject(t *testing.T, value interface{}) map[string]interface{} {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]interface{}
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatal(err)
	}
	return fields
}
