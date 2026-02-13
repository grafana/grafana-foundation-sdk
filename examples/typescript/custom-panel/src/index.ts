import { DashboardBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { CustomPanelBuilder } from "./customPanel";

const builder = new DashboardBuilder('[Example] Custom panel')
  .uid('example-custom-panel')
  .withPanel(
    new CustomPanelBuilder()
      .title('Sample custom panel')
      .makeBeautiful()
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
