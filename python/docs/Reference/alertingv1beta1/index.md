# alertingv1beta1

## Objects

 * <span class="badge object-type-enum"></span> [AlertRuleHealth](./object-AlertRuleHealth.md)
 * <span class="badge object-type-class"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
 * <span class="badge object-type-scalar"></span> [AlertRuleKind](./object-AlertRuleKind.md)
 * <span class="badge object-type-class"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
 * <span class="badge object-type-enum"></span> [AlertRuleState](./object-AlertRuleState.md)
 * <span class="badge object-type-enum"></span> [AlertRuleStateReason](./object-AlertRuleStateReason.md)
 * <span class="badge object-type-class"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
 * <span class="badge object-type-scalar"></span> [ApiVersion](./object-ApiVersion.md)
 * <span class="badge object-type-scalar"></span> [DatasourceUID](./object-DatasourceUID.md)
 * <span class="badge object-type-enum"></span> [ExecErrState](./object-ExecErrState.md)
 * <span class="badge object-type-class"></span> [Expression](./object-Expression.md)
 * <span class="badge object-type-map"></span> [ExpressionMap](./object-ExpressionMap.md)
 * <span class="badge object-type-class"></span> [IntervalTrigger](./object-IntervalTrigger.md)
 * <span class="badge object-type-scalar"></span> [MetricName](./object-MetricName.md)
 * <span class="badge object-type-class"></span> [NamedRoutingTree](./object-NamedRoutingTree.md)
 * <span class="badge object-type-enum"></span> [NoDataState](./object-NoDataState.md)
 * <span class="badge object-type-disjunction"></span> [NotificationSettings](./object-NotificationSettings.md)
 * <span class="badge object-type-enum"></span> [NotificationSettingsType](./object-NotificationSettingsType.md)
 * <span class="badge object-type-class"></span> [PanelRef](./object-PanelRef.md)
 * <span class="badge object-type-scalar"></span> [PromDuration](./object-PromDuration.md)
 * <span class="badge object-type-scalar"></span> [PromDurationWMillis](./object-PromDurationWMillis.md)
 * <span class="badge object-type-enum"></span> [RecordingRuleHealth](./object-RecordingRuleHealth.md)
 * <span class="badge object-type-scalar"></span> [RecordingRuleKind](./object-RecordingRuleKind.md)
 * <span class="badge object-type-class"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
 * <span class="badge object-type-class"></span> [RecordingRuleStatus](./object-RecordingRuleStatus.md)
 * <span class="badge object-type-class"></span> [RelativeTimeRange](./object-RelativeTimeRange.md)
 * <span class="badge object-type-class"></span> [RuleRef](./object-RuleRef.md)
 * <span class="badge object-type-scalar"></span> [RuleSequenceKind](./object-RuleSequenceKind.md)
 * <span class="badge object-type-class"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
 * <span class="badge object-type-class"></span> [RuleSpec](./object-RuleSpec.md)
 * <span class="badge object-type-scalar"></span> [RuleUID](./object-RuleUID.md)
 * <span class="badge object-type-class"></span> [SimplifiedRouting](./object-SimplifiedRouting.md)
 * <span class="badge object-type-scalar"></span> [TemplateString](./object-TemplateString.md)
 * <span class="badge object-type-scalar"></span> [TimeIntervalRef](./object-TimeIntervalRef.md)
## Builders

 * <span class="badge builder"></span> [AlertRule](./builder-AlertRule.md)
 * <span class="badge builder"></span> [AlertRuleInstanceTotals](./builder-AlertRuleInstanceTotals.md)
 * <span class="badge builder"></span> [AlertRuleStatus](./builder-AlertRuleStatus.md)
 * <span class="badge builder"></span> [Expression](./builder-Expression.md)
 * <span class="badge builder"></span> [IntervalTrigger](./builder-IntervalTrigger.md)
 * <span class="badge builder"></span> [NamedRoutingTree](./builder-NamedRoutingTree.md)
 * <span class="badge builder"></span> [PanelRef](./builder-PanelRef.md)
 * <span class="badge builder"></span> [RecordingRule](./builder-RecordingRule.md)
 * <span class="badge builder"></span> [RecordingRuleStatus](./builder-RecordingRuleStatus.md)
 * <span class="badge builder"></span> [RelativeTimeRange](./builder-RelativeTimeRange.md)
 * <span class="badge builder"></span> [RuleRef](./builder-RuleRef.md)
 * <span class="badge builder"></span> [RuleSequence](./builder-RuleSequence.md)
 * <span class="badge builder"></span> [RuleSpec](./builder-RuleSpec.md)
 * <span class="badge builder"></span> [SimplifiedRouting](./builder-SimplifiedRouting.md)
## Functions

### <span class="badge function"></span> alert_rule_manifest

Creates a resource manifest from an AlertRule.

```python
def alert_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.AlertRuleSpec]) -> AlertRule
```

### <span class="badge function"></span> recording_rule_manifest

Creates a resource manifest from a RecordingRule.

```python
def recording_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.RecordingRuleSpec]) -> RecordingRule
```

### <span class="badge function"></span> rule_sequence_manifest

Creates a resource manifest from a RuleSequence.

```python
def rule_sequence_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.RuleSequenceSpec]) -> RuleSequence
```

