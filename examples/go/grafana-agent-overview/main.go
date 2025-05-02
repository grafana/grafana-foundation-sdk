package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func main() {
	builder := dashboard.NewDashboardBuilder("[Example] Grafana Agent Overview").
		Uid("example-integration-agent").
		Tags([]string{"generated", "grafana-foundation-sdk", "grafana-agent-integration"}).
		Editable().
		Tooltip(dashboard.DashboardCursorSyncOff).
		Refresh("30s").
		Time("now-30m", "now").
		Timezone(common.TimeZoneBrowser).
		Timepicker(dashboard.NewTimePickerBuilder().
			RefreshIntervals([]string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}),
		).
		// Links to other agent-related dashboards
		Link(dashboard.NewDashboardLinkBuilder("Grafana Agent Dashboards").
			Type("dashboards").
			Icon("external link").
			Tags([]string{"grafana-agent-integration"}).
			IncludeVars(true).
			KeepTime(true),
		).
		// "Data source" variable
		WithVariable(dashboard.NewDatasourceVariableBuilder("prometheus_datasource").
			Label("Data source").
			Type("prometheus").
			Regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+").
			Multi(false),
		).
		// "Job" variable
		WithVariable(dashboard.NewQueryVariableBuilder("job").
			Label("Job").
			Query(dashboard.StringOrMap{String: cog.ToPtr("label_values(agent_build_info, job)")}).
			Datasource(datasourceRef("$prometheus_datasource")).
			Current(dashboard.VariableOption{
				Selected: cog.ToPtr(true),
				Text:     dashboard.StringOrArrayOfString{ArrayOfString: []string{"All"}},
				Value:    dashboard.StringOrArrayOfString{ArrayOfString: []string{"$__all"}},
			}).
			Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
			Sort(dashboard.VariableSortAlphabeticalAsc).
			Multi(true).
			IncludeAll(true),
		).
		// "Instance" variable
		WithVariable(dashboard.NewQueryVariableBuilder("instance").
			Label("Instance").
			Query(dashboard.StringOrMap{String: cog.ToPtr("label_values(agent_build_info{job=~\"$job\"}, instance)")}).
			Datasource(datasourceRef("$prometheus_datasource")).
			Current(dashboard.VariableOption{
				Selected: cog.ToPtr(true),
				Text:     dashboard.StringOrArrayOfString{ArrayOfString: []string{"All"}},
				Value:    dashboard.StringOrArrayOfString{ArrayOfString: []string{"$__all"}},
			}).
			Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
			Sort(dashboard.VariableSortAlphabeticalAsc).
			Multi(true).
			IncludeAll(true),
		).
		// Panels
		WithRow(dashboard.NewRowBuilder("Overview")).
		WithPanel(runningInstancesTable()).
		WithRow(dashboard.NewRowBuilder("Prometheus Discovery")).
		WithPanel(targetSyncTimeseries()).
		WithPanel(targetsTimeseries()).
		WithRow(dashboard.NewRowBuilder("Prometheus Retrieval")).
		WithPanel(avgScrapeIntervalDurationTimeseries()).
		WithPanel(scrapeFailuresTimeseries()).
		WithPanel(appendedSamplesTimeseries())

	sampleDashboard, err := builder.Build()
	if err != nil {
		panic(err)
	}
	dashboardJson, err := json.MarshalIndent(sampleDashboard, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dashboardJson))
}
