package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/alerting"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/expr"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
)

func main() {
	datasourceUid := "DS_PROMETHEUS_UID"
	ruleGroupTitle := "example-group"
	folderUid := "folder-foo-uid"

	queryA := alerting.NewQueryBuilder("A").
		DatasourceUid(datasourceUid).
		RelativeTimeRange(alerting.Duration(600), alerting.Duration(0)).
		Model(
			prometheus.NewDataqueryBuilder().
				Expr("vector(1)").
				Instant().
				RefId("A"),
		)

	thresholdCondition := expr.NewExprTypeThresholdConditionsBuilder().
		Evaluator(expr.NewExprTypeThresholdConditionsEvaluatorBuilder().
			Type(expr.ExprTypeThresholdConditionsEvaluatorTypeGt).
			Params([]float64{0}))

	exprStr := "__expr__"
	queryB := alerting.NewQueryBuilder("B").
		DatasourceUid("__expr__").
		Model(
			expr.NewTypeThresholdBuilder().
				Conditions([]cog.Builder[expr.ExprTypeThresholdConditions]{thresholdCondition}).
				Datasource(dashboard.DataSourceRef{Uid: &exprStr, Type: &exprStr}).
				Expression("A").
				RefId("B"),
		)

	ruleBuilder := alerting.NewRuleBuilder("[Example] Rule").
		Queries([]cog.Builder[alerting.Query]{queryA, queryB}).
		Condition("B").
		ExecErrState(alerting.RuleExecErrStateError).
		NoDataState(alerting.RuleNoDataStateNoData).
		For("0m").
		RuleGroup(ruleGroupTitle).
		FolderUID(folderUid).
		Labels(map[string]string{"severity": "critical"}).
		Annotations(map[string]string{"summary": "Demo rule"})

	ruleGroupBuilder := alerting.NewRuleGroupBuilder(ruleGroupTitle).
		Interval(60).
		FolderUid(folderUid).
		Rules([]cog.Builder[alerting.Rule]{ruleBuilder})

	ruleGroup, err := ruleGroupBuilder.Build()
	if err != nil {
		panic(err)
	}

	ruleGroupJson, err := json.MarshalIndent(ruleGroup, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(ruleGroupJson))
}
