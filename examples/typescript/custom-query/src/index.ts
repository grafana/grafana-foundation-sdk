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

console.log(JSON.stringify(builder.build(), null, 2));
