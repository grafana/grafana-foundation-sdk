from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models import (
    alerting,
    expr,
    prometheus,
    dashboard
)
from grafana_foundation_sdk.builders import (
    alerting as alert_builder,
    expr as expr_builder,
    prometheus as prom_builder
)


if __name__ == "__main__":
    datasource_uid = "DS_PROMETHEUS_UID"
    rule_group_uid = "10m"
    folder_uid = "folder-foo-uid"

    query_a = alert_builder.Query("A") \
        .datasource_uid(datasource_uid) \
        .relative_time_range(
            alerting.RelativeTimeRange(600, 0)
        ) \
        .model(prom_builder.Dataquery()
               .expr("vector(1)")
               .instant()
               .ref_id("A"))

    threshold_condition = expr_builder.ExprTypeThresholdConditions() \
        .evaluator(expr_builder.ExprTypeThresholdConditionsEvaluator()
                   .type("gt")
                   .params([0]))

    query_b = alert_builder.Query("B") \
        .datasource_uid("__expr__") \
        .model(expr_builder.TypeThreshold()
               .conditions([threshold_condition])
               .datasource(dashboard.DataSourceRef("__expr__", "__expr__"))
               .expression("A")
               .ref_id("B"))

    rule = alert_builder.Rule("[Example] Rule") \
        .queries([query_a, query_b]) \
        .condition("B") \
        .for_val("0m") \
        .rule_group(rule_group_uid) \
        .folder_uid(folder_uid) \
        .labels({"severity": "critical"}) \
        .annotations({"summary": "Demo rule"})

    rule_group = alert_builder.RuleGroup(rule_group_uid) \
        .interval(60) \
        .folder_uid(folder_uid) \
        .rules([rule])
    print(JSONEncoder(sort_keys=True, indent=2).encode(rule_group.build()))