package main

import (
	"encoding/json"

	linuxNodeOverview "github.com/grafana/grafana-foundation-sdk/examples/go/linux-node-overview/builder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		grafanaProvider, err := grafana.NewProvider(ctx, "grafana", &grafana.ProviderArgs{
			Auth: pulumi.String("admin:admin"),
			Url:  pulumi.String("http://localhost:3000"),
		})
		if err != nil {
			return err
		}

		sampleDashboard, err := linuxNodeOverview.Build()
		if err != nil {
			return err
		}

		dashboardJson, err := json.MarshalIndent(sampleDashboard, "", "  ")
		if err != nil {
			return err
		}

		pulumiDashboard, err := grafana.NewDashboard(ctx, *sampleDashboard.Uid, &grafana.DashboardArgs{
			ConfigJson: pulumi.String(dashboardJson),
		}, pulumi.Provider(grafanaProvider))
		if err != nil {
			return err
		}

		ctx.Export("dashboardUID", pulumiDashboard.Uid)

		return nil
	})
}
