# Building a dashboard

```typescript
import { DashboardBuilder, RowBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { DataqueryBuilder } from '@grafana/grafana-foundation-sdk/prometheus';
import { PanelBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

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
      .unit('bps')
      .min(0)
      .withTarget(
        new DataqueryBuilder()
          .expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
          .legendFormat({{ `"{{ device }}"` }})
      )
  )
;

console.log(JSON.stringify(builder.build(), null, 2));
```
