# alertingv0alpha1

## Objects

 * <span class="badge object-type-scalar"></span> [AlertRuleKind](./object-AlertRuleKind.md)
 * <span class="badge object-type-struct"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
 * <span class="badge object-type-scalar"></span> [ApiVersion](./object-ApiVersion.md)
 * <span class="badge object-type-scalar"></span> [DatasourceUID](./object-DatasourceUID.md)
 * <span class="badge object-type-enum"></span> [ExecErrState](./object-ExecErrState.md)
 * <span class="badge object-type-struct"></span> [Expression](./object-Expression.md)
 * <span class="badge object-type-map"></span> [ExpressionMap](./object-ExpressionMap.md)
 * <span class="badge object-type-struct"></span> [FacetValue](./object-FacetValue.md)
 * <span class="badge object-type-struct"></span> [IntervalTrigger](./object-IntervalTrigger.md)
 * <span class="badge object-type-scalar"></span> [MetricName](./object-MetricName.md)
 * <span class="badge object-type-struct"></span> [NamedRoutingTree](./object-NamedRoutingTree.md)
 * <span class="badge object-type-enum"></span> [NoDataState](./object-NoDataState.md)
 * <span class="badge object-type-ref"></span> [NotificationSettings](./object-NotificationSettings.md)
 * <span class="badge object-type-enum"></span> [NotificationSettingsType](./object-NotificationSettingsType.md)
 * <span class="badge object-type-struct"></span> [PanelRef](./object-PanelRef.md)
 * <span class="badge object-type-scalar"></span> [PromDuration](./object-PromDuration.md)
 * <span class="badge object-type-scalar"></span> [PromDurationWMillis](./object-PromDurationWMillis.md)
 * <span class="badge object-type-scalar"></span> [RecordingRuleKind](./object-RecordingRuleKind.md)
 * <span class="badge object-type-struct"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
 * <span class="badge object-type-struct"></span> [RelativeTimeRange](./object-RelativeTimeRange.md)
 * <span class="badge object-type-struct"></span> [RuleRef](./object-RuleRef.md)
 * <span class="badge object-type-struct"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
 * <span class="badge object-type-scalar"></span> [RuleSequenceKind](./object-RuleSequenceKind.md)
 * <span class="badge object-type-struct"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
 * <span class="badge object-type-struct"></span> [RuleSpec](./object-RuleSpec.md)
 * <span class="badge object-type-scalar"></span> [RuleUID](./object-RuleUID.md)
 * <span class="badge object-type-struct"></span> [SearchResultHit](./object-SearchResultHit.md)
 * <span class="badge object-type-struct"></span> [SearchResultResource](./object-SearchResultResource.md)
 * <span class="badge object-type-struct"></span> [SearchResults](./object-SearchResults.md)
 * <span class="badge object-type-struct"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)
 * <span class="badge object-type-struct"></span> [SimplifiedRouting](./object-SimplifiedRouting.md)
 * <span class="badge object-type-struct"></span> [SimplifiedRoutingOrNamedRoutingTree](./object-SimplifiedRoutingOrNamedRoutingTree.md)
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

### <span class="badge function"></span> NewAlertRuleSpec

NewAlertRuleSpec creates a new AlertRuleSpec object.

```go
func NewAlertRuleSpec() *AlertRuleSpec
```

### <span class="badge function"></span> NewIntervalTrigger

NewIntervalTrigger creates a new IntervalTrigger object.

```go
func NewIntervalTrigger() *IntervalTrigger
```

### <span class="badge function"></span> NewNotificationSettings

NewNotificationSettings creates a new NotificationSettings object.

```go
func NewNotificationSettings() *NotificationSettings
```

### <span class="badge function"></span> NewSimplifiedRouting

NewSimplifiedRouting creates a new SimplifiedRouting object.

```go
func NewSimplifiedRouting() *SimplifiedRouting
```

### <span class="badge function"></span> NewNamedRoutingTree

NewNamedRoutingTree creates a new NamedRoutingTree object.

```go
func NewNamedRoutingTree() *NamedRoutingTree
```

### <span class="badge function"></span> NewExpression

NewExpression creates a new Expression object.

```go
func NewExpression() *Expression
```

### <span class="badge function"></span> NewRelativeTimeRange

NewRelativeTimeRange creates a new RelativeTimeRange object.

```go
func NewRelativeTimeRange() *RelativeTimeRange
```

### <span class="badge function"></span> NewPanelRef

NewPanelRef creates a new PanelRef object.

```go
func NewPanelRef() *PanelRef
```

### <span class="badge function"></span> NewRecordingRuleSpec

NewRecordingRuleSpec creates a new RecordingRuleSpec object.

```go
func NewRecordingRuleSpec() *RecordingRuleSpec
```

### <span class="badge function"></span> NewRuleSequenceSpec

NewRuleSequenceSpec creates a new RuleSequenceSpec object.

```go
func NewRuleSequenceSpec() *RuleSequenceSpec
```

### <span class="badge function"></span> NewRuleRef

NewRuleRef creates a new RuleRef object.

```go
func NewRuleRef() *RuleRef
```

### <span class="badge function"></span> NewRuleSpec

NewRuleSpec creates a new RuleSpec object.

```go
func NewRuleSpec() *RuleSpec
```

### <span class="badge function"></span> NewSearchResultResource

NewSearchResultResource creates a new SearchResultResource object.

```go
func NewSearchResultResource() *SearchResultResource
```

### <span class="badge function"></span> NewRuleSearchHitFields

NewRuleSearchHitFields creates a new RuleSearchHitFields object.

```go
func NewRuleSearchHitFields() *RuleSearchHitFields
```

### <span class="badge function"></span> NewSearchResultHit

NewSearchResultHit creates a new SearchResultHit object.

```go
func NewSearchResultHit() *SearchResultHit
```

### <span class="badge function"></span> NewSearchResultsMetadata

NewSearchResultsMetadata creates a new SearchResultsMetadata object.

```go
func NewSearchResultsMetadata() *SearchResultsMetadata
```

### <span class="badge function"></span> NewFacetValue

NewFacetValue creates a new FacetValue object.

```go
func NewFacetValue() *FacetValue
```

### <span class="badge function"></span> NewSearchResults

NewSearchResults creates a new SearchResults object.

```go
func NewSearchResults() *SearchResults
```

### <span class="badge function"></span> NewSimplifiedRoutingOrNamedRoutingTree

NewSimplifiedRoutingOrNamedRoutingTree creates a new SimplifiedRoutingOrNamedRoutingTree object.

```go
func NewSimplifiedRoutingOrNamedRoutingTree() *SimplifiedRoutingOrNamedRoutingTree
```

### <span class="badge function"></span> AlertRuleManifest

Creates a resource manifest from an AlertRule.

```go
func AlertRuleManifest(name string, folder cog.Builder[AlertRuleSpec]) *AlertRuleBuilder
```

### <span class="badge function"></span> RecordingRuleManifest

Creates a resource manifest from a RecordingRule.

```go
func RecordingRuleManifest(name string, folder cog.Builder[RecordingRuleSpec]) *RecordingRuleBuilder
```

### <span class="badge function"></span> RuleSequenceManifest

Creates a resource manifest from a RuleSequence.

```go
func RuleSequenceManifest(name string, folder cog.Builder[RuleSequenceSpec]) *RuleSequenceBuilder
```

### <span class="badge function"></span> IntervalTriggerConverter

IntervalTriggerConverter accepts a `IntervalTrigger` object and generates the Go code to build this object using builders.

```go
func IntervalTriggerConverter(input IntervalTrigger) string
```

### <span class="badge function"></span> SimplifiedRoutingConverter

SimplifiedRoutingConverter accepts a `SimplifiedRouting` object and generates the Go code to build this object using builders.

```go
func SimplifiedRoutingConverter(input SimplifiedRouting) string
```

### <span class="badge function"></span> NamedRoutingTreeConverter

NamedRoutingTreeConverter accepts a `NamedRoutingTree` object and generates the Go code to build this object using builders.

```go
func NamedRoutingTreeConverter(input NamedRoutingTree) string
```

### <span class="badge function"></span> ExpressionConverter

ExpressionConverter accepts a `Expression` object and generates the Go code to build this object using builders.

```go
func ExpressionConverter(input Expression) string
```

### <span class="badge function"></span> RelativeTimeRangeConverter

RelativeTimeRangeConverter accepts a `RelativeTimeRange` object and generates the Go code to build this object using builders.

```go
func RelativeTimeRangeConverter(input RelativeTimeRange) string
```

### <span class="badge function"></span> PanelRefConverter

PanelRefConverter accepts a `PanelRef` object and generates the Go code to build this object using builders.

```go
func PanelRefConverter(input PanelRef) string
```

### <span class="badge function"></span> RuleRefConverter

RuleRefConverter accepts a `RuleRef` object and generates the Go code to build this object using builders.

```go
func RuleRefConverter(input RuleRef) string
```

### <span class="badge function"></span> RuleSpecConverter

RuleSpecConverter accepts a `RuleSpec` object and generates the Go code to build this object using builders.

```go
func RuleSpecConverter(input RuleSpec) string
```

### <span class="badge function"></span> SearchResultResourceConverter

SearchResultResourceConverter accepts a `SearchResultResource` object and generates the Go code to build this object using builders.

```go
func SearchResultResourceConverter(input SearchResultResource) string
```

### <span class="badge function"></span> RuleSearchHitFieldsConverter

RuleSearchHitFieldsConverter accepts a `RuleSearchHitFields` object and generates the Go code to build this object using builders.

```go
func RuleSearchHitFieldsConverter(input RuleSearchHitFields) string
```

### <span class="badge function"></span> SearchResultHitConverter

SearchResultHitConverter accepts a `SearchResultHit` object and generates the Go code to build this object using builders.

```go
func SearchResultHitConverter(input SearchResultHit) string
```

### <span class="badge function"></span> SearchResultsMetadataConverter

SearchResultsMetadataConverter accepts a `SearchResultsMetadata` object and generates the Go code to build this object using builders.

```go
func SearchResultsMetadataConverter(input SearchResultsMetadata) string
```

### <span class="badge function"></span> FacetValueConverter

FacetValueConverter accepts a `FacetValue` object and generates the Go code to build this object using builders.

```go
func FacetValueConverter(input FacetValue) string
```

### <span class="badge function"></span> SearchResultsConverter

SearchResultsConverter accepts a `SearchResults` object and generates the Go code to build this object using builders.

```go
func SearchResultsConverter(input SearchResults) string
```

### <span class="badge function"></span> AlertRuleConverter

AlertRuleConverter accepts a `AlertRule` object and generates the Go code to build this object using builders.

```go
func AlertRuleConverter(input AlertRuleSpec) string
```

### <span class="badge function"></span> RecordingRuleConverter

RecordingRuleConverter accepts a `RecordingRule` object and generates the Go code to build this object using builders.

```go
func RecordingRuleConverter(input RecordingRuleSpec) string
```

### <span class="badge function"></span> RuleSequenceConverter

RuleSequenceConverter accepts a `RuleSequence` object and generates the Go code to build this object using builders.

```go
func RuleSequenceConverter(input RuleSequenceSpec) string
```

