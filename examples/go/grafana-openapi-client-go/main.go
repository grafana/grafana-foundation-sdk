package main

import (
	"fmt"
	"net/url"

	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/testdata"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	goapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func sampleDashboard() *dashboard.DashboardBuilder {
	return dashboard.NewDashboardBuilder("[Example] Grafana HTTP OpenAPI Client for Go").
		Uid("grafana-openapi-client-go").
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Sample panel").
				Datasource(dashboard.DataSourceRef{
					Uid:  cog.ToPtr("grafana"),
					Type: cog.ToPtr("grafana"),
				}).
				WithTarget(
					testdata.NewDataqueryBuilder().
						QueryType("randomWalk"),
				),
		)
}

func main() {
	// Required to correctly unmarshal panels and dataqueries
	plugins.RegisterDefaultPlugins()

	cfg := &goapi.TransportConfig{
		// Host is the doman name or IP address of the host that serves the API.
		Host: "localhost:3000",
		// BasePath is the URL prefix for all API paths, relative to the host root.
		BasePath: "/api",
		// Schemes are the transfer protocols used by the API (http or https).
		Schemes: []string{"http"},
		// BasicAuth is optional basic auth credentials.
		BasicAuth: url.UserPassword("admin", "secret"),
	}

	client := goapi.NewHTTPClientWithConfig(strfmt.Default, cfg)

	dashboardModel, err := sampleDashboard().Build()
	if err != nil {
		panic(err)
	}

	response, err := client.Dashboards.PostDashboard(&models.SaveDashboardCommand{
		FolderUID: "adjwdwpz0zny8e",
		Dashboard: dashboardModel,
		Overwrite: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %#v\n", response.Payload)
}
