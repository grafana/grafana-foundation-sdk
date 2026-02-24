# <span class="badge package-core"></span> alerting

## Objects

 * <span class="badge object-type-struct"></span> [ContactPoint](./object-ContactPoint.md)
 * <span class="badge object-type-enum"></span> [ContactPointType](./object-ContactPointType.md)
 * <span class="badge object-type-struct"></span> [DayOfMonthRange](./object-DayOfMonthRange.md)
 * <span class="badge object-type-scalar"></span> [Duration](./object-Duration.md)
 * <span class="badge object-type-scalar"></span> [Json](./object-Json.md)
 * <span class="badge object-type-scalar"></span> [Location](./object-Location.md)
 * <span class="badge object-type-map"></span> [MatchRegexps](./object-MatchRegexps.md)
 * <span class="badge object-type-enum"></span> [MatchType](./object-MatchType.md)
 * <span class="badge object-type-struct"></span> [Matcher](./object-Matcher.md)
 * <span class="badge object-type-array"></span> [Matchers](./object-Matchers.md)
 * <span class="badge object-type-struct"></span> [MonthRange](./object-MonthRange.md)
 * <span class="badge object-type-struct"></span> [MuteTiming](./object-MuteTiming.md)
 * <span class="badge object-type-struct"></span> [NotificationPolicy](./object-NotificationPolicy.md)
 * <span class="badge object-type-struct"></span> [NotificationSettings](./object-NotificationSettings.md)
 * <span class="badge object-type-struct"></span> [NotificationTemplate](./object-NotificationTemplate.md)
 * <span class="badge object-type-array"></span> [ObjectMatcher](./object-ObjectMatcher.md)
 * <span class="badge object-type-array"></span> [ObjectMatchers](./object-ObjectMatchers.md)
 * <span class="badge object-type-scalar"></span> [Provenance](./object-Provenance.md)
 * <span class="badge object-type-struct"></span> [Query](./object-Query.md)
 * <span class="badge object-type-struct"></span> [RecordRule](./object-RecordRule.md)
 * <span class="badge object-type-struct"></span> [RelativeTimeRange](./object-RelativeTimeRange.md)
 * <span class="badge object-type-struct"></span> [Rule](./object-Rule.md)
 * <span class="badge object-type-enum"></span> [RuleExecErrState](./object-RuleExecErrState.md)
 * <span class="badge object-type-struct"></span> [RuleGroup](./object-RuleGroup.md)
 * <span class="badge object-type-enum"></span> [RuleNoDataState](./object-RuleNoDataState.md)
 * <span class="badge object-type-struct"></span> [TimeInterval](./object-TimeInterval.md)
 * <span class="badge object-type-struct"></span> [TimeRange](./object-TimeRange.md)
 * <span class="badge object-type-struct"></span> [WeekdayRange](./object-WeekdayRange.md)
 * <span class="badge object-type-struct"></span> [YearRange](./object-YearRange.md)
## Builders

 * <span class="badge builder"></span> [ContactPointBuilder](./builder-ContactPointBuilder.md)
 * <span class="badge builder"></span> [DayOfMonthRangeBuilder](./builder-DayOfMonthRangeBuilder.md)
 * <span class="badge builder"></span> [MatcherBuilder](./builder-MatcherBuilder.md)
 * <span class="badge builder"></span> [MonthRangeBuilder](./builder-MonthRangeBuilder.md)
 * <span class="badge builder"></span> [MuteTimingBuilder](./builder-MuteTimingBuilder.md)
 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
 * <span class="badge builder"></span> [NotificationSettingsBuilder](./builder-NotificationSettingsBuilder.md)
 * <span class="badge builder"></span> [NotificationTemplateBuilder](./builder-NotificationTemplateBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
 * <span class="badge builder"></span> [RecordRuleBuilder](./builder-RecordRuleBuilder.md)
 * <span class="badge builder"></span> [RuleBuilder](./builder-RuleBuilder.md)
 * <span class="badge builder"></span> [RuleGroupBuilder](./builder-RuleGroupBuilder.md)
 * <span class="badge builder"></span> [TimeIntervalBuilder](./builder-TimeIntervalBuilder.md)
 * <span class="badge builder"></span> [TimeRangeBuilder](./builder-TimeRangeBuilder.md)
 * <span class="badge builder"></span> [WeekdayRangeBuilder](./builder-WeekdayRangeBuilder.md)
 * <span class="badge builder"></span> [YearRangeBuilder](./builder-YearRangeBuilder.md)
## Functions

### <span class="badge function"></span> NewQuery

NewQuery creates a new Query object.

```go
func NewQuery() *Query
```

### <span class="badge function"></span> NewRuleGroup

NewRuleGroup creates a new RuleGroup object.

```go
func NewRuleGroup() *RuleGroup
```

### <span class="badge function"></span> NewNotificationSettings

NewNotificationSettings creates a new NotificationSettings object.

```go
func NewNotificationSettings() *NotificationSettings
```

### <span class="badge function"></span> NewContactPoint

NewContactPoint creates a new ContactPoint object.

```go
func NewContactPoint() *ContactPoint
```

### <span class="badge function"></span> NewMatcher

NewMatcher creates a new Matcher object.

```go
func NewMatcher() *Matcher
```

### <span class="badge function"></span> NewMuteTiming

NewMuteTiming creates a new MuteTiming object.

```go
func NewMuteTiming() *MuteTiming
```

### <span class="badge function"></span> NewNotificationTemplate

NewNotificationTemplate creates a new NotificationTemplate object.

```go
func NewNotificationTemplate() *NotificationTemplate
```

### <span class="badge function"></span> NewRule

NewRule creates a new Rule object.

```go
func NewRule() *Rule
```

### <span class="badge function"></span> NewRecordRule

NewRecordRule creates a new RecordRule object.

```go
func NewRecordRule() *RecordRule
```

### <span class="badge function"></span> NewRelativeTimeRange

NewRelativeTimeRange creates a new RelativeTimeRange object.

```go
func NewRelativeTimeRange() *RelativeTimeRange
```

### <span class="badge function"></span> NewNotificationPolicy

NewNotificationPolicy creates a new NotificationPolicy object.

```go
func NewNotificationPolicy() *NotificationPolicy
```

### <span class="badge function"></span> NewTimeInterval

NewTimeInterval creates a new TimeInterval object.

```go
func NewTimeInterval() *TimeInterval
```

### <span class="badge function"></span> NewTimeRange

NewTimeRange creates a new TimeRange object.

```go
func NewTimeRange() *TimeRange
```

### <span class="badge function"></span> NewWeekdayRange

NewWeekdayRange creates a new WeekdayRange object.

```go
func NewWeekdayRange() *WeekdayRange
```

### <span class="badge function"></span> NewDayOfMonthRange

NewDayOfMonthRange creates a new DayOfMonthRange object.

```go
func NewDayOfMonthRange() *DayOfMonthRange
```

### <span class="badge function"></span> NewYearRange

NewYearRange creates a new YearRange object.

```go
func NewYearRange() *YearRange
```

### <span class="badge function"></span> NewMonthRange

NewMonthRange creates a new MonthRange object.

```go
func NewMonthRange() *MonthRange
```

### <span class="badge function"></span> NotificationSettingsConverter

NotificationSettingsConverter accepts a `NotificationSettings` object and generates the Go code to build this object using builders.

```go
func NotificationSettingsConverter(input NotificationSettings) string
```

### <span class="badge function"></span> ContactPointConverter

ContactPointConverter accepts a `ContactPoint` object and generates the Go code to build this object using builders.

```go
func ContactPointConverter(input ContactPoint) string
```

### <span class="badge function"></span> MatcherConverter

MatcherConverter accepts a `Matcher` object and generates the Go code to build this object using builders.

```go
func MatcherConverter(input Matcher) string
```

### <span class="badge function"></span> MuteTimingConverter

MuteTimingConverter accepts a `MuteTiming` object and generates the Go code to build this object using builders.

```go
func MuteTimingConverter(input MuteTiming) string
```

### <span class="badge function"></span> NotificationTemplateConverter

NotificationTemplateConverter accepts a `NotificationTemplate` object and generates the Go code to build this object using builders.

```go
func NotificationTemplateConverter(input NotificationTemplate) string
```

### <span class="badge function"></span> RecordRuleConverter

RecordRuleConverter accepts a `RecordRule` object and generates the Go code to build this object using builders.

```go
func RecordRuleConverter(input RecordRule) string
```

### <span class="badge function"></span> NotificationPolicyConverter

NotificationPolicyConverter accepts a `NotificationPolicy` object and generates the Go code to build this object using builders.

```go
func NotificationPolicyConverter(input NotificationPolicy) string
```

### <span class="badge function"></span> TimeIntervalConverter

TimeIntervalConverter accepts a `TimeInterval` object and generates the Go code to build this object using builders.

```go
func TimeIntervalConverter(input TimeInterval) string
```

### <span class="badge function"></span> TimeRangeConverter

TimeRangeConverter accepts a `TimeRange` object and generates the Go code to build this object using builders.

```go
func TimeRangeConverter(input TimeRange) string
```

### <span class="badge function"></span> WeekdayRangeConverter

WeekdayRangeConverter accepts a `WeekdayRange` object and generates the Go code to build this object using builders.

```go
func WeekdayRangeConverter(input WeekdayRange) string
```

### <span class="badge function"></span> DayOfMonthRangeConverter

DayOfMonthRangeConverter accepts a `DayOfMonthRange` object and generates the Go code to build this object using builders.

```go
func DayOfMonthRangeConverter(input DayOfMonthRange) string
```

### <span class="badge function"></span> YearRangeConverter

YearRangeConverter accepts a `YearRange` object and generates the Go code to build this object using builders.

```go
func YearRangeConverter(input YearRange) string
```

### <span class="badge function"></span> MonthRangeConverter

MonthRangeConverter accepts a `MonthRange` object and generates the Go code to build this object using builders.

```go
func MonthRangeConverter(input MonthRange) string
```

### <span class="badge function"></span> RuleGroupConverter

RuleGroupConverter accepts a `RuleGroup` object and generates the Go code to build this object using builders.

```go
func RuleGroupConverter(input RuleGroup) string
```

### <span class="badge function"></span> RuleConverter

RuleConverter accepts a `Rule` object and generates the Go code to build this object using builders.

```go
func RuleConverter(input Rule) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input Query) string
```

