import {
  DashboardBuilder,
  PanelBuilder,
  QueryGroupBuilder,
  TargetBuilder,
  manifest,
  tabs,
  tab,
  grid,
  gridItem,
  rows,
  row,
  autoGrid,
} from '@grafana/grafana-foundation-sdk/dashboardv2beta1';
import { QueryBuilder as TestDataQueryBuilder } from '@grafana/grafana-foundation-sdk/testdata';
import { VisualizationBuilder as TimeSeriesVisualizationBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

const TESTDATA_DATASOURCE_NAME = 'grafana-testdata-datasource';

function randomWalkPanel(panelId: number, seriesCount: number): PanelBuilder {
  return new PanelBuilder()
    .title(`New panel ${panelId}`)
    .data(
      new QueryGroupBuilder().target(
        new TargetBuilder()
          .refId('A')
          .query(
            new TestDataQueryBuilder()
              .datasource({ name: TESTDATA_DATASOURCE_NAME })
              .scenarioId('random_walk')
              .seriesCount(seriesCount)
          )
      )
    )
    .visualization(new TimeSeriesVisualizationBuilder());
}

const dashboard = new DashboardBuilder('[Example] Dashboard with tabs and rows')
  .description('Dashboard with tabs and rows generated with grafana-foundation-sdk')
	// Available panels
  .element('panel-1', randomWalkPanel(1, 4))
  .element('panel-2', randomWalkPanel(2, 5))
  .element('panel-3', randomWalkPanel(3, 1))
  .element('panel-4', randomWalkPanel(4, 1))
	// Layout building
  .layout(tabs()
    .tab(
      tab('Tab without rows').layout(
        grid().item(
          gridItem('panel-1')
            .width(12)
            .height(8)
        )
      )
    )
    .tab(
      tab('Tab With Rows').layout(rows()
        .row(
          row('Row without tabs')
            .collapse(false)
            .layout(autoGrid()
              .withItem('panel-2')
            )
        )
        .row(
          row('Row with tabs')
            .collapse(true)
            .layout(
              tabs()
                .tab(
                  tab("First tab")
                    .layout(autoGrid()
                      .withItem('panel-3')
                    )
                )
                .tab(
                  tab("Second tab")
                    .layout(autoGrid()
                      .withItem('panel-4')
                    )
                )
            )
        )
        .row(
          row('Empty row')
            .layout(autoGrid())
        )
        .row(
          row('Hide header row')
            .hideHeader(true)
            .layout(autoGrid())
        )
      )
    )
  );

const dashboardManifest = manifest('example-dashboard-with-tabs-and-rows', dashboard);

console.log(JSON.stringify(dashboardManifest.build(), null, 2));
