---
title: Home
hide:
  - navigation
---

# Grafana Foundation SDK

The Grafana Foundation SDK is a set of **_types_**, and **_builder libraries_**
that let you define Grafana dashboards and other resources using strongly typed
code. By writing your resources as code, you can:

* **Leverage strong typing**: Catch errors at compile time, ensuring more reliable configurations.
* **Enhance version control**: Track changes seamlessly using standard version control systems like Git.
* **Automate deployments**: Integrate dashboard provisioning into your CI/CD pipelines for consistent and repeatable setups.

The SDK supports multiple programming languages, including **Go**, **TypeScript**, **Python**, **PHP**, and **Java**, so you can choose the one that best fits your development environment.

!!! note

    This SDK is best suited for Grafana >= 12, but will work with Grafana >= 10.

## Overview

Here's a quick overview of how the SDK works:

* **Builder pattern**: The SDK implements the builder pattern to let you define dashboards fluently. You start with a `DashboardBuilder`, then add panels, queries, and other components step by step.
* **Strong typing**: Everything in the SDK is strongly typed. This gives you autocompletion in your IDE, catches mistakes early, and helps ensure you're always using valid configuration values.
* **Structured options**: When a configuration get complex (like data reduction or display settings), the SDK uses typed option builders to keep things readable and predictable.
* **Resources**: Multiple resources are supported by the SDK. Dashboards, alerts, …

For example, here is how dashboards can be built:

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

!!! tip

    More examples can be found in the [`./examples/`](https://github.com/grafana/grafana-foundation-sdk/tree/main/examples)
    folder of the Foundation SDK repository.

## Publishing resources

After you've defined your resource as code, call the `build()` function (its
actual name might be slightly different depending on language choice) and
output the result as a JSON.

With the JSON payload, you can:

* Call [Grafana's API](https://grafana.com/docs/grafana/latest/developer-resources/api-reference/http-api/) to programmatically manage the resource.
* Use [Grafana CLI](https://grafana.com/docs/grafana/latest/as-code/observability-as-code/grafana-cli/) publish the resource from CLI.

## Next steps

With the basics of using the Grafana Foundation SDK in mind, here are some possible next steps:

* **Explore more features**: Check out the full API reference to learn more about what the SDK can do.
* **Version control your resources**: Store your resources code in a Git repository to track changes over time.
* **Automate provisioning with CI/CD**: [Integrate the SDK into your CI/CD pipeline](https://grafana.com/docs/grafana/latest/as-code/observability-as-code/foundation-sdk/dashboard-automation/) to deploy dashboards and other resources automatically.
* **Explore a more [real-world example of using the SDK](https://www.youtube.com/watch?v=ZjWdGVsrCiQ)**

## Maturity

The Grafana Foundation SDK should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

Additional information can be found in [Release life cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

!!! note

    Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)
