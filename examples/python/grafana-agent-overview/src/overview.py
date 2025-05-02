from grafana_foundation_sdk.builders import common, table
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    DataTransformerConfig,
    DynamicConfigValue,
)
from grafana_foundation_sdk.models import units

from .common import table_prometheus_query


def running_instances_table() -> table.Panel:
    return (
        table.Panel()
        .title("Running Instances")
        .description("General statistics of running grafana agent instances.")
        .height(7)
        .span(24)
        .footer(common.TableFooterOptions().count_rows(False).reducer(["sum"]))
        .datasource(DataSourceRef(uid="$prometheus_datasource"))
        .with_target(
            table_prometheus_query(
                query='count by (instance, version) (agent_build_info{job=~"$job", instance=~"$instance"})',
                ref_id="A",
            )
        )
        .with_target(
            table_prometheus_query(
                query='max by (instance) (time() - process_start_time_seconds{job=~"$job", instance=~"$instance"})',
                ref_id="B",
            )
        )
        # Transformations
        .with_transformation(DataTransformerConfig(id_val="merge", options={}))
        .with_transformation(
            DataTransformerConfig(
                id_val="organize",
                options={
                    "excludeByName": {
                        "Time": True,
                        "Value #A": True,
                    },
                    "renameByName": {
                        "Value #B": "Uptime",
                    },
                },
            )
        )
        # Overrides
        .override_by_name(
            "Value #B",
            [
                DynamicConfigValue(id_val="unit", value=units.Seconds),
            ],
        )
    )
