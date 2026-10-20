package dashboard_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
)

func TestPanelDataLinks(t *testing.T) {
	for _, input := range []string{
		`{"title":"Open","url":"https://example.com/${__field.labels.id}"}`,
		`{"title":"Open","url":"https://example.com","targetBlank":true,"oneClick":true}`,
		`{"title":"Open","url":"https://example.com","targetBlank":false,"oneClick":false}`,
	} {
		t.Run(input, func(t *testing.T) {
			var link dashboard.DataLink
			if err := json.Unmarshal([]byte(input), &link); err != nil {
				t.Fatal(err)
			}
			builder := dashboard.NewDataLinkBuilder().Title(link.Title).Url(link.Url)
			if link.TargetBlank != nil {
				builder.TargetBlank(*link.TargetBlank)
			}
			if link.OneClick != nil {
				builder.OneClick(*link.OneClick)
			}
			panel, err := stat.NewPanelBuilder().DataLinks([]cog.Builder[dashboard.DataLink]{builder}).Build()
			if err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(panel.FieldConfig.Defaults.Links[0])
			if err != nil {
				t.Fatal(err)
			}
			var got, want map[string]interface{}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(input), &want); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("data link changed during panel construction: got %s, want %s", encoded, input)
			}
		})
	}
}
