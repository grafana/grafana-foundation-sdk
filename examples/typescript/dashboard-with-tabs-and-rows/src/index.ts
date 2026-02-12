import {
  AutoGridLayoutBuilder,
  AutoGridLayoutItemBuilder,
  DashboardBuilder,
  GridLayoutBuilder,
  GridLayoutItemBuilder,
  PanelBuilder,
  QueryGroupBuilder,
  RowsLayoutBuilder,
  RowsLayoutRowBuilder,
  TabsLayoutBuilder,
  TabsLayoutTabBuilder,
  TargetBuilder,
} from '@grafana/grafana-foundation-sdk/dashboardv2beta1';
import {
  DashboardKind,
  DashboardV2Beta1,
  ManifestBuilder,
  MetadataBuilder,
} from '@grafana/grafana-foundation-sdk/resource';
import { QueryBuilder as TestDataQueryBuilder } from '@grafana/grafana-foundation-sdk/testdata';
import { VisualizationBuilder as TimeSeriesVisualizationBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

const TESTDATA_DATASOURCE_NAME = 'grafana-testdata-datasource';

function randomWalkPanel(panelId: number, seriesCount: number): PanelBuilder {
  return new PanelBuilder()
    .id(panelId)
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
  .element('panel-1', randomWalkPanel(1, 4))
  .element('panel-2', randomWalkPanel(2, 5))
  .element('panel-3', randomWalkPanel(3, 1))
  .element('panel-4', randomWalkPanel(4, 1))
  .tabsLayout(
    new TabsLayoutBuilder()
      .tab(
        new TabsLayoutTabBuilder('Tab without rows').gridLayout(
          new GridLayoutBuilder().item(
            new GridLayoutItemBuilder('panel-1')
              .width(12)
              .height(8)
          )
        )
      )
      .tab(
        new TabsLayoutTabBuilder('Tab With Rows').rowsLayout(
          new RowsLayoutBuilder()
            .row(
              new RowsLayoutRowBuilder('Row without tabs')
                .collapse(true)
                .autoGridLayout(
                  new AutoGridLayoutBuilder().item(
                    new AutoGridLayoutItemBuilder('panel-2')
                  )
                )
            )
            .row(
              new RowsLayoutRowBuilder('Row with tabs')
                .collapse(true)
                .autoGridLayout(
                  new AutoGridLayoutBuilder()
                    .item(new AutoGridLayoutItemBuilder('panel-3'))
                    .item(new AutoGridLayoutItemBuilder('panel-4'))
                )
            )
            .row(
              new RowsLayoutRowBuilder('Empty row').autoGridLayout(
                new AutoGridLayoutBuilder()
              )
            )
            .row(
              new RowsLayoutRowBuilder('Hide header row')
                .hideHeader(true)
                .autoGridLayout(new AutoGridLayoutBuilder())
            )
        )
      )
  );

const dashboardManifest = new ManifestBuilder()
  .apiVersion(DashboardV2Beta1)
  .kind(DashboardKind)
  .metadata(
    new MetadataBuilder()
      .name('example-dashboard-with-tabs-and-rows')
      .annotations({
        'grafana.app/folder': '',
      })
  )
  .spec(dashboard.build());

console.log(JSON.stringify(dashboardManifest.build(), null, 2));
