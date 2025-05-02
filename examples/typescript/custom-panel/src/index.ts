import { DashboardBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { CustomPanelBuilder } from "./customPanel";

const builder = new DashboardBuilder('[Example] Custom panel')
  .uid('example-custom-panel')
  .withPanel(
    new CustomPanelBuilder()
      .title('Sample custom panel')
      .makeBeautiful()
  );

console.log(JSON.stringify(builder.build(), null, 2));
