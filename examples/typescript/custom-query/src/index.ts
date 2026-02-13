import { DashboardBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as TimeSeriesBuilder } from "@grafana/grafana-foundation-sdk/timeseries";
import { CustomQueryBuilder } from "./customQuery";

const builder = new DashboardBuilder('[Example] Custom query')
  .uid('example-custom-query')
  .withPanel(
    new TimeSeriesBuilder()
      .title('Sample panel')
      .withTarget(
        new CustomQueryBuilder("query here")
      )
  );

const dashboard = builder.build();

console.log(JSON.stringify({
    apiVersion: 'dashboard.grafana.app/v1beta1',
    kind: 'Dashboard',
    metadata: {
        name: dashboard.uid!,
    },
    spec: dashboard,
}, null, 2));
