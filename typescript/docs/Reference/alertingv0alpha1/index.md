# alertingv0alpha1

## Objects

 * <span class="badge object-type-scalar"></span> [AlertRuleKind](./object-AlertRuleKind.md)
 * <span class="badge object-type-interface"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
 * <span class="badge object-type-scalar"></span> [ApiVersion](./object-ApiVersion.md)
 * <span class="badge object-type-scalar"></span> [DatasourceUID](./object-DatasourceUID.md)
 * <span class="badge object-type-enum"></span> [ExecErrState](./object-ExecErrState.md)
 * <span class="badge object-type-interface"></span> [Expression](./object-Expression.md)
 * <span class="badge object-type-map"></span> [ExpressionMap](./object-ExpressionMap.md)
 * <span class="badge object-type-interface"></span> [FacetValue](./object-FacetValue.md)
 * <span class="badge object-type-interface"></span> [IntervalTrigger](./object-IntervalTrigger.md)
 * <span class="badge object-type-scalar"></span> [MetricName](./object-MetricName.md)
 * <span class="badge object-type-interface"></span> [NamedRoutingTree](./object-NamedRoutingTree.md)
 * <span class="badge object-type-enum"></span> [NoDataState](./object-NoDataState.md)
 * <span class="badge object-type-disjunction"></span> [NotificationSettings](./object-NotificationSettings.md)
 * <span class="badge object-type-enum"></span> [NotificationSettingsType](./object-NotificationSettingsType.md)
 * <span class="badge object-type-interface"></span> [PanelRef](./object-PanelRef.md)
 * <span class="badge object-type-scalar"></span> [PromDuration](./object-PromDuration.md)
 * <span class="badge object-type-scalar"></span> [PromDurationWMillis](./object-PromDurationWMillis.md)
 * <span class="badge object-type-scalar"></span> [RecordingRuleKind](./object-RecordingRuleKind.md)
 * <span class="badge object-type-interface"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
 * <span class="badge object-type-interface"></span> [RelativeTimeRange](./object-RelativeTimeRange.md)
 * <span class="badge object-type-interface"></span> [RuleRef](./object-RuleRef.md)
 * <span class="badge object-type-interface"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
 * <span class="badge object-type-scalar"></span> [RuleSequenceKind](./object-RuleSequenceKind.md)
 * <span class="badge object-type-interface"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
 * <span class="badge object-type-interface"></span> [RuleSpec](./object-RuleSpec.md)
 * <span class="badge object-type-scalar"></span> [RuleUID](./object-RuleUID.md)
 * <span class="badge object-type-interface"></span> [SearchResultHit](./object-SearchResultHit.md)
 * <span class="badge object-type-interface"></span> [SearchResultResource](./object-SearchResultResource.md)
 * <span class="badge object-type-interface"></span> [SearchResults](./object-SearchResults.md)
 * <span class="badge object-type-interface"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)
 * <span class="badge object-type-interface"></span> [SimplifiedRouting](./object-SimplifiedRouting.md)
 * <span class="badge object-type-scalar"></span> [TemplateString](./object-TemplateString.md)
 * <span class="badge object-type-scalar"></span> [TimeIntervalRef](./object-TimeIntervalRef.md)
 * <span class="badge object-type-enum"></span> [TotalHitsRelation](./object-TotalHitsRelation.md)
## Builders

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
 * <span class="badge builder"></span> [ExpressionBuilder](./builder-ExpressionBuilder.md)
 * <span class="badge builder"></span> [FacetValueBuilder](./builder-FacetValueBuilder.md)
 * <span class="badge builder"></span> [IntervalTriggerBuilder](./builder-IntervalTriggerBuilder.md)
 * <span class="badge builder"></span> [NamedRoutingTreeBuilder](./builder-NamedRoutingTreeBuilder.md)
 * <span class="badge builder"></span> [PanelRefBuilder](./builder-PanelRefBuilder.md)
 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
 * <span class="badge builder"></span> [RelativeTimeRangeBuilder](./builder-RelativeTimeRangeBuilder.md)
 * <span class="badge builder"></span> [RuleRefBuilder](./builder-RuleRefBuilder.md)
 * <span class="badge builder"></span> [RuleSearchHitFieldsBuilder](./builder-RuleSearchHitFieldsBuilder.md)
 * <span class="badge builder"></span> [RuleSequenceBuilder](./builder-RuleSequenceBuilder.md)
 * <span class="badge builder"></span> [RuleSpecBuilder](./builder-RuleSpecBuilder.md)
 * <span class="badge builder"></span> [SearchResultHitBuilder](./builder-SearchResultHitBuilder.md)
 * <span class="badge builder"></span> [SearchResultResourceBuilder](./builder-SearchResultResourceBuilder.md)
 * <span class="badge builder"></span> [SearchResultsBuilder](./builder-SearchResultsBuilder.md)
 * <span class="badge builder"></span> [SearchResultsMetadataBuilder](./builder-SearchResultsMetadataBuilder.md)
 * <span class="badge builder"></span> [SimplifiedRoutingBuilder](./builder-SimplifiedRoutingBuilder.md)
## Functions

### <span class="badge function"></span> alertRuleManifest

Creates a resource manifest from an AlertRule.

```typescript
alertRuleManifest(name: string, folder: cog.Builder<alertingv0alpha1.AlertRuleSpec>)
```

### <span class="badge function"></span> recordingRuleManifest

Creates a resource manifest from a RecordingRule.

```typescript
recordingRuleManifest(name: string, folder: cog.Builder<alertingv0alpha1.RecordingRuleSpec>)
```

### <span class="badge function"></span> ruleSequenceManifest

Creates a resource manifest from a RuleSequence.

```typescript
ruleSequenceManifest(name: string, folder: cog.Builder<alertingv0alpha1.RuleSequenceSpec>)
```

