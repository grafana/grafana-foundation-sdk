package alertrule;

import java.util.List;
import java.util.Map;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.alerting.QueryBuilder;
import com.grafana.foundation.alerting.RelativeTimeRange;
import com.grafana.foundation.alerting.RuleBuilder;
import com.grafana.foundation.alerting.RuleGroupBuilder;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.expr.ExprTypeThresholdConditionsBuilder;
import com.grafana.foundation.expr.ExprTypeThresholdConditionsEvaluatorBuilder;
import com.grafana.foundation.expr.ExprTypeThresholdConditionsEvaluatorType;
import com.grafana.foundation.expr.TypeThresholdBuilder;
import com.grafana.foundation.prometheus.DataqueryBuilder;

public class Main {

    public static void main(String[] args) {
        String datasourceUid = "DS_PROMETHEUS_UID";
        String ruleGroupUid = "1m";
        String folderUid = "folder-foo-uid";

        QueryBuilder queryA = new QueryBuilder("A")
                .datasourceUid(datasourceUid)
                .relativeTimeRange(new RelativeTimeRange(600L, 0L))
                .model(new DataqueryBuilder()
                        .expr("vector(1)")
                        .instant()
                        .refId("A")
                );

        ExprTypeThresholdConditionsBuilder thresholdCondition = new ExprTypeThresholdConditionsBuilder()
                .evaluator(new ExprTypeThresholdConditionsEvaluatorBuilder()
                        .type(ExprTypeThresholdConditionsEvaluatorType.GT)
                        .params(List.of(0.0)));
        
        QueryBuilder queryB = new QueryBuilder("B")
                .datasourceUid("__expr__")
                .model(new TypeThresholdBuilder()
                        .conditions(List.of(thresholdCondition))
                        .datasource(new DataSourceRef("__expr__", "__expr__"))
                        .expression("A")
                        .refId("B"));

        RuleBuilder ruleBuilder = new RuleBuilder("[Example] Rule")
                .queries(List.of(queryA, queryB))
                .condition("B")
                .forArg("0m")
                .ruleGroup(ruleGroupUid)
                .folderUID(folderUid)
                .labels(Map.of("severity", "critical"))
                .annotations(Map.of("summary", "Demo rule"));

        RuleGroupBuilder ruleGroupBuilder = new RuleGroupBuilder(ruleGroupUid)
                .interval(60L)
                .folderUid(folderUid)
                .rules(List.of(ruleBuilder));

        try {
            System.out.println(ruleGroupBuilder.build().toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
