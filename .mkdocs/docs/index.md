---
title: Home
hide:
  - navigation
---

# Grafana Foundation SDK

A set of libraries for manipulating and generating Grafana resources
– dashboards, alerts, … – as-code, in various languages.

**_Types_**, **_builder libraries_** and **JSON to code converters** are
provided for a range of versions of Grafana.

!!! note

    The various SDKs are generated by [`cog`][cog] from
    [schemas exposed by Grafana][kind-registry].

## Examples

=== "Go"

    ```go
    package main

    import (
        "encoding/json"
        "fmt"

        "github.com/grafana/grafana-foundation-sdk/go/common"
        "github.com/grafana/grafana-foundation-sdk/go/dashboard"
        "github.com/grafana/grafana-foundation-sdk/go/prometheus"
        "github.com/grafana/grafana-foundation-sdk/go/timeseries"
        "github.com/grafana/grafana-foundation-sdk/go/units"
    )

    func main() {
        builder := dashboard.NewDashboardBuilder("Sample dashboard").
            Uid("generated-from-go").
            Tags([]string{"generated", "from", "go"}).
            Refresh("1m").
            Time("now-30m", "now").
            Timezone(common.TimeZoneBrowser).
            WithRow(dashboard.NewRowBuilder("Overview")).
            WithPanel(
                timeseries.NewPanelBuilder().
                    Title("Network Received").
                    Unit(units.BitsPerSecondSI).
                    Min(0).
                    WithTarget(
                        prometheus.NewDataqueryBuilder().
                            Expr(`rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8`).
                            LegendFormat("{{ device }}"),
                    ),
            )

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
    ```

=== "Java"

    ```java
    package example;

    import com.fasterxml.jackson.core.JsonProcessingException;
    import com.grafana.foundation.dashboard.Dashboard;
    import com.grafana.foundation.common.Constants;
    import com.grafana.foundation.dashboard.DashboardBuilder;
    import com.grafana.foundation.dashboard.DashboardDashboardTimeBuilder;
    import com.grafana.foundation.dashboard.RowBuilder;
    import com.grafana.foundation.prometheus.DataqueryBuilder;
    import com.grafana.foundation.timeseries.PanelBuilder;

    import java.util.List;

    public class Main {
        public static void main(String[] args) throws JsonProcessingException {
            Dashboard dashboard = new DashboardBuilder("Sample Dashboard").
                    uid("generated-from-java").
                    tags(List.of("generated", "from", "java")).
                    refresh("1m").
                    time(new DashboardDashboardTimeBuilder().
                            from("now-30m").
                            to("now")
                    ).
                    timezone(Constants.TimeZoneBrowser).
                    withRow(new RowBuilder("Overview")).
                    withPanel(new PanelBuilder().
                            title("Network Received").
                            unit("bps").
                            min(0.0).
                            withTarget(new DataqueryBuilder().
                                    expr("rate(node_network_receive_bytes_total{job=\"integrations/raspberrypi-node\", device!=\"lo\"}[$__rate_interval]) * 8").
                                    legendFormat("{{ device }}")
                            )
                    ).build();

            System.out.println(dashboard.toJSON());
        }
    }
    ```

=== "PHP"

    ```php
    <?php

    use Grafana\Foundation\Common;
    use Grafana\Foundation\Dashboard\DashboardBuilder;
    use Grafana\Foundation\Dashboard\RowBuilder;
    use Grafana\Foundation\Prometheus;
    use Grafana\Foundation\Timeseries;
    use Grafana\Foundation\Units\Constants as Units;

    require_once __DIR__.'/vendor/autoload.php';

    $builder = (new DashboardBuilder(title: 'Sample dashboard'))
        ->uid('generated-from-php')
        ->tags(['generated', 'from', 'php'])
        ->refresh('1m')
        ->time('now-30m', 'now')
        ->timezone(Common\Constants::TIME_ZONE_BROWSER)
        ->withRow(new RowBuilder('Overview'))
        ->withPanel(
            (new Timeseries\PanelBuilder())
                ->title('Network received')
                ->unit(Units::BITS_PER_SECOND_SI)
                ->min(0)
                ->withTarget(
                    (new Prometheus\DataqueryBuilder())
                        ->expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
                        ->legendFormat('{{ device }}')
                )
        )
    ;

    echo(json_encode($builder->build(), JSON_PRETTY_PRINT).PHP_EOL);
    ```

=== "Python"

    ```python
    from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
    from grafana_foundation_sdk.builders import prometheus, timeseries
    from grafana_foundation_sdk.cog.encoder import JSONEncoder
    from grafana_foundation_sdk.models.common import TimeZoneBrowser
    from grafana_foundation_sdk.models import units


    def build_dashboard() -> Dashboard:
        builder = (
            Dashboard("[TEST] Node Exporter / Raspberry")
            .uid("test-dashboard-raspberry")
            .tags(["generated", "raspberrypi-node-integration"])
            .refresh("1m")
            .time("now-30m", "now")
            .timezone(TimeZoneBrowser)
            .with_row(Row("Overview"))
            .with_panel(
                timeseries.Panel()
                .title("Network Received")
                .unit(units.BitsPerSecondSI)
                .min(0)
                .with_target(
                    prometheus.Dataquery()
                    .expr(
                        'rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8'
                    )
                    .legend_format("{{ device }}")
                )
            )
        )

        return builder


    if __name__ == "__main__":
        dashboard = build_dashboard().build()
        encoder = JSONEncoder(sort_keys=True, indent=2)

        print(encoder.encode(dashboard))
    ```

=== "Typescript"

    ```typescript
    import { DashboardBuilder, RowBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
    import { DataqueryBuilder } from '@grafana/grafana-foundation-sdk/prometheus';
    import { PanelBuilder } from '@grafana/grafana-foundation-sdk/timeseries';
    import * as units from '@grafana/grafana-foundation-sdk/units';
    
    const builder = new DashboardBuilder('[TEST] Node Exporter / Raspberry')
      .uid('test-dashboard-raspberry')
      .tags(['generated', 'raspberrypi-node-integration'])
    
      .refresh('1m')
      .time({from: 'now-30m', to: 'now'})
      .timezone('browser')
    
      .withRow(new RowBuilder('Overview'))
      .withPanel(
        new PanelBuilder()
          .title('Network Received')
          .unit(units.BitsPerSecondSI)
          .min(0)
          .withTarget(
            new DataqueryBuilder()
              .expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
              .legendFormat("{{ device }}")
          )
      )
    ;
    
    console.log(JSON.stringify(builder.build(), null, 2));
    ```

## Maturity

The Grafana Foundation SDK should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

Additional information can be found in [Release life cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

!!! note

    Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

[cog]: <https://github.com/grafana/cog>
[kind-registry]: <https://github.com/grafana/kind-registry>
