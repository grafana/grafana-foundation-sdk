from grafana_foundation_sdk.builders import dashboardv2beta1 as dashboard, timeseries, testdata
from grafana_foundation_sdk.cog.encoder import JSONEncoder

def random_walk_panel(panel_id: int, series_count: int) -> dashboard.Panel:
    return (
        dashboard.Panel()
            .title("New panel {}".format(panel_id))
            .visualization(timeseries.Visualization())
            .data(dashboard.QueryGroup().target(
                dashboard.Target().query(
                    testdata.Query()
                        .scenario_id("random_walk")
                        .series_count(series_count)
                )
            ))
    )

builder = (
    dashboard.Dashboard("[Example] Dashboard with tabs and rows")
        .description("Dashboard with tabs and rows generated with grafana-foundation-sdk")
        # Panels
        .element("panel-1", random_walk_panel(1, 4))
        .element("panel-2", random_walk_panel(2, 5))
        .element("panel-3", random_walk_panel(3, 1))
        .element("panel-4", random_walk_panel(4, 1))
        # Layout building
        .layout(dashboard.tabs()
            .tab(dashboard.tab("Tab without rows")
                .layout(dashboard.grid()
                    .item(dashboard.grid_item("panel-1").width(12).height(8))
                )
            )
            .tab(dashboard.tab("Tab with rows")
                .layout(dashboard.rows()
                    .row(dashboard.row("Row without tabs")
                        .collapse(False)
                        .layout(dashboard.auto_grid()
                            .item(dashboard.auto_grid_item("panel-2"))
                        )
                    )
                    .row(dashboard.row("Row with tabs")
                        .collapse(True)
                        .layout(dashboard.tabs()
                            .tab(dashboard.tab("First tab").layout(
                                dashboard.auto_grid()
                                    .item(dashboard.auto_grid_item("panel-3"))
                            ))
                            .tab(dashboard.tab("Second tab").layout(
                                dashboard.auto_grid()
                                    .item(dashboard.auto_grid_item("panel-4"))
                            ))
                        )
                    )
                    .row(dashboard.row("Empty row")
                        .layout(dashboard.auto_grid())
                    )
                    .row(dashboard.row("Hide header row")
                        .hide_header(True)
                        .layout(dashboard.auto_grid())
                    )
                )
            )
        )
)

if __name__ == "__main__":
    manifest = dashboard.manifest(
        name="example-dashboard-with-tabs-and-rows",
        dashboard=builder,
    )

    print(JSONEncoder(sort_keys=True, indent=2).encode(manifest.build()))
