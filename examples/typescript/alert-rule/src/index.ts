import {
  QueryBuilder,
  RuleBuilder,
  RuleGroupBuilder,
} from '@grafana/grafana-foundation-sdk/alerting';

import { DataqueryBuilder} from '@grafana/grafana-foundation-sdk/prometheus';
import { TypeThresholdBuilder } from '@grafana/grafana-foundation-sdk/expr';

const datasourceUid = 'DS_PROMETHEUS_UID';
const ruleGroupUid = '1m';
const folderUid = 'folder-foo-uid';

const queryA = new QueryBuilder('A')
  .datasourceUid(datasourceUid)
  .relativeTimeRange({ from: 600, to: 0 })
  .model(
    new DataqueryBuilder()
      .expr('vector(1)')
      .instant()
      .refId('A')
  );

const queryB = new QueryBuilder('B')
  .datasourceUid('__expr__')
  .model(
    new TypeThresholdBuilder().conditions([{
      evaluator: { type: 'gt', params: [0] }, 
    },
  ])
      .datasource({ uid: '__expr__', type: '__expr__' })
      .expression('A')
      .refId('B')
  );

const ruleBuilder = new RuleBuilder('[Example] Rule')
  .queries([queryA, queryB])
  .condition('B')
  .forDuration('0m')
  .ruleGroup(ruleGroupUid)
  .folderUID(folderUid)
  .labels({ severity: 'critical' })
  .annotations({ summary: 'Demo rule' });

const ruleGroupBuilder = new RuleGroupBuilder(ruleGroupUid)
  .interval(60)
  .folderUid(folderUid)
  .rules([ruleBuilder]);
console.log(JSON.stringify(ruleGroupBuilder.build(), null, 2));

