// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"encoding/json"
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Query struct {
	DatasourceUid     *string            `json:"datasourceUid,omitempty"`
	Model             variants.Dataquery `json:"model,omitempty"`
	QueryType         *string            `json:"queryType,omitempty"`
	RefId             *string            `json:"refId,omitempty"`
	RelativeTimeRange *RelativeTimeRange `json:"relativeTimeRange,omitempty"`
}

func (resource *Query) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}
	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}

	if fields["datasourceUid"] != nil {
		if err := json.Unmarshal(fields["datasourceUid"], &resource.DatasourceUid); err != nil {
			return err
		}
	}

	if fields["queryType"] != nil {
		if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
			return err
		}
	}

	if fields["refId"] != nil {
		if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
			return err
		}
	}

	if fields["relativeTimeRange"] != nil {
		if err := json.Unmarshal(fields["relativeTimeRange"], &resource.RelativeTimeRange); err != nil {
			return err
		}
	}

	dataqueryTypeHint := ""

	if fields["model"] != nil {
		model, err := cog.UnmarshalDataquery(fields["model"], dataqueryTypeHint)
		if err != nil {
			return err
		}
		resource.Model = model
	}

	return nil
}

func (resource Query) Equals(other Query) bool {
	if resource.DatasourceUid == nil && other.DatasourceUid != nil || resource.DatasourceUid != nil && other.DatasourceUid == nil {
		return false
	}

	if resource.DatasourceUid != nil {
		if *resource.DatasourceUid != *other.DatasourceUid {
			return false
		}
	}
	if resource.Model == nil && other.Model != nil || resource.Model != nil && other.Model == nil {
		return false
	}

	if resource.Model != nil {
		if !resource.Model.Equals(other.Model) {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.RefId == nil && other.RefId != nil || resource.RefId != nil && other.RefId == nil {
		return false
	}

	if resource.RefId != nil {
		if *resource.RefId != *other.RefId {
			return false
		}
	}
	if resource.RelativeTimeRange == nil && other.RelativeTimeRange != nil || resource.RelativeTimeRange != nil && other.RelativeTimeRange == nil {
		return false
	}

	if resource.RelativeTimeRange != nil {
		if !resource.RelativeTimeRange.Equals(*other.RelativeTimeRange) {
			return false
		}
	}

	return true
}

type RuleGroup struct {
	FolderUid *string `json:"folderUid,omitempty"`
	// The interval, in seconds, at which all rules in the group are evaluated.
	// If a group contains many rules, the rules are evaluated sequentially.
	Interval *Duration `json:"interval,omitempty"`
	Rules    []Rule    `json:"rules,omitempty"`
	Title    *string   `json:"title,omitempty"`
}

func (resource RuleGroup) Equals(other RuleGroup) bool {
	if resource.FolderUid == nil && other.FolderUid != nil || resource.FolderUid != nil && other.FolderUid == nil {
		return false
	}

	if resource.FolderUid != nil {
		if *resource.FolderUid != *other.FolderUid {
			return false
		}
	}
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}

	if len(resource.Rules) != len(other.Rules) {
		return false
	}

	for i1 := range resource.Rules {
		if !resource.Rules[i1].Equals(other.Rules[i1]) {
			return false
		}
	}
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}

	return true
}

type NotificationSettings struct {
	GroupBy           []string `json:"group_by,omitempty"`
	GroupInterval     *string  `json:"group_interval,omitempty"`
	GroupWait         *string  `json:"group_wait,omitempty"`
	MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
	Receiver          string   `json:"receiver"`
	RepeatInterval    *string  `json:"repeat_interval,omitempty"`
}

func (resource NotificationSettings) Equals(other NotificationSettings) bool {

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if resource.GroupBy[i1] != other.GroupBy[i1] {
			return false
		}
	}
	if resource.GroupInterval == nil && other.GroupInterval != nil || resource.GroupInterval != nil && other.GroupInterval == nil {
		return false
	}

	if resource.GroupInterval != nil {
		if *resource.GroupInterval != *other.GroupInterval {
			return false
		}
	}
	if resource.GroupWait == nil && other.GroupWait != nil || resource.GroupWait != nil && other.GroupWait == nil {
		return false
	}

	if resource.GroupWait != nil {
		if *resource.GroupWait != *other.GroupWait {
			return false
		}
	}

	if len(resource.MuteTimeIntervals) != len(other.MuteTimeIntervals) {
		return false
	}

	for i1 := range resource.MuteTimeIntervals {
		if resource.MuteTimeIntervals[i1] != other.MuteTimeIntervals[i1] {
			return false
		}
	}
	if resource.Receiver != other.Receiver {
		return false
	}
	if resource.RepeatInterval == nil && other.RepeatInterval != nil || resource.RepeatInterval != nil && other.RepeatInterval == nil {
		return false
	}

	if resource.RepeatInterval != nil {
		if *resource.RepeatInterval != *other.RepeatInterval {
			return false
		}
	}

	return true
}

// Duration in seconds.
type Duration int64

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
type ContactPoint struct {
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	DisableResolveMessage *bool `json:"disableResolveMessage,omitempty"`
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	Name *string `json:"name,omitempty"`
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	Provenance *string `json:"provenance,omitempty"`
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	Settings Json `json:"settings"`
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	Type ContactPointType `json:"type"`
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	Uid *string `json:"uid,omitempty"`
}

func (resource ContactPoint) Equals(other ContactPoint) bool {
	if resource.DisableResolveMessage == nil && other.DisableResolveMessage != nil || resource.DisableResolveMessage != nil && other.DisableResolveMessage == nil {
		return false
	}

	if resource.DisableResolveMessage != nil {
		if *resource.DisableResolveMessage != *other.DisableResolveMessage {
			return false
		}
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.Provenance == nil && other.Provenance != nil || resource.Provenance != nil && other.Provenance == nil {
		return false
	}

	if resource.Provenance != nil {
		if *resource.Provenance != *other.Provenance {
			return false
		}
	}
	if resource.Settings != other.Settings {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Uid == nil && other.Uid != nil || resource.Uid != nil && other.Uid == nil {
		return false
	}

	if resource.Uid != nil {
		if *resource.Uid != *other.Uid {
			return false
		}
	}

	return true
}

type Json any

type MatchRegexps map[string]string

type MatchType string

const (
	MatchTypeEqual         MatchType = "="
	MatchTypeNotEqual      MatchType = "!="
	MatchTypeEqualRegex    MatchType = "=~"
	MatchTypeNotEqualRegex MatchType = "!~"
)

type Matcher struct {
	Name  *string    `json:"Name,omitempty"`
	Type  *MatchType `json:"Type,omitempty"`
	Value *string    `json:"Value,omitempty"`
}

func (resource Matcher) Equals(other Matcher) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if *resource.Value != *other.Value {
			return false
		}
	}

	return true
}

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
type Matchers []Matcher

type MuteTiming struct {
	Name          *string        `json:"name,omitempty"`
	TimeIntervals []TimeInterval `json:"time_intervals,omitempty"`
}

func (resource MuteTiming) Equals(other MuteTiming) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	if len(resource.TimeIntervals) != len(other.TimeIntervals) {
		return false
	}

	for i1 := range resource.TimeIntervals {
		if !resource.TimeIntervals[i1].Equals(other.TimeIntervals[i1]) {
			return false
		}
	}

	return true
}

type NotificationTemplate struct {
	Name       *string     `json:"name,omitempty"`
	Provenance *Provenance `json:"provenance,omitempty"`
	Template   *string     `json:"template,omitempty"`
}

func (resource NotificationTemplate) Equals(other NotificationTemplate) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.Provenance == nil && other.Provenance != nil || resource.Provenance != nil && other.Provenance == nil {
		return false
	}

	if resource.Provenance != nil {
		if *resource.Provenance != *other.Provenance {
			return false
		}
	}
	if resource.Template == nil && other.Template != nil || resource.Template != nil && other.Template == nil {
		return false
	}

	if resource.Template != nil {
		if *resource.Template != *other.Template {
			return false
		}
	}

	return true
}

type ObjectMatcher []string

type ObjectMatchers []ObjectMatcher

type Provenance string

type Rule struct {
	Annotations  map[string]string `json:"annotations,omitempty"`
	Condition    string            `json:"condition"`
	Data         []Query           `json:"data"`
	ExecErrState RuleExecErrState  `json:"execErrState"`
	FolderUID    string            `json:"folderUID"`
	// The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
	// Before this time has elapsed, the rule is only considered to be Pending.
	For                  string                `json:"for"`
	Id                   *int64                `json:"id,omitempty"`
	IsPaused             *bool                 `json:"isPaused,omitempty"`
	Labels               map[string]string     `json:"labels,omitempty"`
	NoDataState          RuleNoDataState       `json:"noDataState"`
	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
	OrgID                int64                 `json:"orgID"`
	Provenance           *Provenance           `json:"provenance,omitempty"`
	RuleGroup            string                `json:"ruleGroup"`
	Title                string                `json:"title"`
	Uid                  *string               `json:"uid,omitempty"`
	Updated              *time.Time            `json:"updated,omitempty"`
}

func (resource Rule) Equals(other Rule) bool {

	if len(resource.Annotations) != len(other.Annotations) {
		return false
	}

	for key1 := range resource.Annotations {
		if resource.Annotations[key1] != other.Annotations[key1] {
			return false
		}
	}
	if resource.Condition != other.Condition {
		return false
	}

	if len(resource.Data) != len(other.Data) {
		return false
	}

	for i1 := range resource.Data {
		if !resource.Data[i1].Equals(other.Data[i1]) {
			return false
		}
	}
	if resource.ExecErrState != other.ExecErrState {
		return false
	}
	if resource.FolderUID != other.FolderUID {
		return false
	}
	if resource.For != other.For {
		return false
	}
	if resource.Id == nil && other.Id != nil || resource.Id != nil && other.Id == nil {
		return false
	}

	if resource.Id != nil {
		if *resource.Id != *other.Id {
			return false
		}
	}
	if resource.IsPaused == nil && other.IsPaused != nil || resource.IsPaused != nil && other.IsPaused == nil {
		return false
	}

	if resource.IsPaused != nil {
		if *resource.IsPaused != *other.IsPaused {
			return false
		}
	}

	if len(resource.Labels) != len(other.Labels) {
		return false
	}

	for key1 := range resource.Labels {
		if resource.Labels[key1] != other.Labels[key1] {
			return false
		}
	}
	if resource.NoDataState != other.NoDataState {
		return false
	}
	if resource.NotificationSettings == nil && other.NotificationSettings != nil || resource.NotificationSettings != nil && other.NotificationSettings == nil {
		return false
	}

	if resource.NotificationSettings != nil {
		if !resource.NotificationSettings.Equals(*other.NotificationSettings) {
			return false
		}
	}
	if resource.OrgID != other.OrgID {
		return false
	}
	if resource.Provenance == nil && other.Provenance != nil || resource.Provenance != nil && other.Provenance == nil {
		return false
	}

	if resource.Provenance != nil {
		if *resource.Provenance != *other.Provenance {
			return false
		}
	}
	if resource.RuleGroup != other.RuleGroup {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if resource.Uid == nil && other.Uid != nil || resource.Uid != nil && other.Uid == nil {
		return false
	}

	if resource.Uid != nil {
		if *resource.Uid != *other.Uid {
			return false
		}
	}
	if resource.Updated == nil && other.Updated != nil || resource.Updated != nil && other.Updated == nil {
		return false
	}

	if resource.Updated != nil {
		if *resource.Updated != *other.Updated {
			return false
		}
	}

	return true
}

// RelativeTimeRange is the per query start and end time
// for requests.
type RelativeTimeRange struct {
	// RelativeTimeRange is the per query start and end time
	// for requests.
	From *Duration `json:"from,omitempty"`
	// RelativeTimeRange is the per query start and end time
	// for requests.
	To *Duration `json:"to,omitempty"`
}

func (resource RelativeTimeRange) Equals(other RelativeTimeRange) bool {
	if resource.From == nil && other.From != nil || resource.From != nil && other.From == nil {
		return false
	}

	if resource.From != nil {
		if *resource.From != *other.From {
			return false
		}
	}
	if resource.To == nil && other.To != nil || resource.To != nil && other.To == nil {
		return false
	}

	if resource.To != nil {
		if *resource.To != *other.To {
			return false
		}
	}

	return true
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
type NotificationPolicy struct {
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Continue *bool `json:"continue,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	GroupBy []string `json:"group_by,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	GroupInterval *string `json:"group_interval,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	GroupWait *string `json:"group_wait,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Match map[string]string `json:"match,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	MatchRe *MatchRegexps `json:"match_re,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Matchers *Matchers `json:"matchers,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	ObjectMatchers *ObjectMatchers `json:"object_matchers,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Provenance *Provenance `json:"provenance,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Receiver *string `json:"receiver,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	RepeatInterval *string `json:"repeat_interval,omitempty"`
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	Routes []NotificationPolicy `json:"routes,omitempty"`
}

func (resource NotificationPolicy) Equals(other NotificationPolicy) bool {
	if resource.Continue == nil && other.Continue != nil || resource.Continue != nil && other.Continue == nil {
		return false
	}

	if resource.Continue != nil {
		if *resource.Continue != *other.Continue {
			return false
		}
	}

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if resource.GroupBy[i1] != other.GroupBy[i1] {
			return false
		}
	}
	if resource.GroupInterval == nil && other.GroupInterval != nil || resource.GroupInterval != nil && other.GroupInterval == nil {
		return false
	}

	if resource.GroupInterval != nil {
		if *resource.GroupInterval != *other.GroupInterval {
			return false
		}
	}
	if resource.GroupWait == nil && other.GroupWait != nil || resource.GroupWait != nil && other.GroupWait == nil {
		return false
	}

	if resource.GroupWait != nil {
		if *resource.GroupWait != *other.GroupWait {
			return false
		}
	}

	if len(resource.Match) != len(other.Match) {
		return false
	}

	for key1 := range resource.Match {
		if resource.Match[key1] != other.Match[key1] {
			return false
		}
	}
	if resource.MatchRe == nil && other.MatchRe != nil || resource.MatchRe != nil && other.MatchRe == nil {
		return false
	}

	if resource.MatchRe != nil {

		if len(*resource.MatchRe) != len(*other.MatchRe) {
			return false
		}

		for key1 := range *resource.MatchRe {
			if (*resource.MatchRe)[key1] != (*other.MatchRe)[key1] {
				return false
			}
		}
	}
	if resource.Matchers == nil && other.Matchers != nil || resource.Matchers != nil && other.Matchers == nil {
		return false
	}

	if resource.Matchers != nil {

		if len(*resource.Matchers) != len(*other.Matchers) {
			return false
		}

		for i1 := range *resource.Matchers {
			if !(*resource.Matchers)[i1].Equals((*other.Matchers)[i1]) {
				return false
			}
		}
	}

	if len(resource.MuteTimeIntervals) != len(other.MuteTimeIntervals) {
		return false
	}

	for i1 := range resource.MuteTimeIntervals {
		if resource.MuteTimeIntervals[i1] != other.MuteTimeIntervals[i1] {
			return false
		}
	}
	if resource.ObjectMatchers == nil && other.ObjectMatchers != nil || resource.ObjectMatchers != nil && other.ObjectMatchers == nil {
		return false
	}

	if resource.ObjectMatchers != nil {

		if len(*resource.ObjectMatchers) != len(*other.ObjectMatchers) {
			return false
		}

		for i1 := range *resource.ObjectMatchers {

			if len((*resource.ObjectMatchers)[i1]) != len((*other.ObjectMatchers)[i1]) {
				return false
			}

			for i2 := range (*resource.ObjectMatchers)[i1] {
				if (*resource.ObjectMatchers)[i1][i2] != (*other.ObjectMatchers)[i1][i2] {
					return false
				}
			}
		}
	}
	if resource.Provenance == nil && other.Provenance != nil || resource.Provenance != nil && other.Provenance == nil {
		return false
	}

	if resource.Provenance != nil {
		if *resource.Provenance != *other.Provenance {
			return false
		}
	}
	if resource.Receiver == nil && other.Receiver != nil || resource.Receiver != nil && other.Receiver == nil {
		return false
	}

	if resource.Receiver != nil {
		if *resource.Receiver != *other.Receiver {
			return false
		}
	}
	if resource.RepeatInterval == nil && other.RepeatInterval != nil || resource.RepeatInterval != nil && other.RepeatInterval == nil {
		return false
	}

	if resource.RepeatInterval != nil {
		if *resource.RepeatInterval != *other.RepeatInterval {
			return false
		}
	}

	if len(resource.Routes) != len(other.Routes) {
		return false
	}

	for i1 := range resource.Routes {
		if !resource.Routes[i1].Equals(other.Routes[i1]) {
			return false
		}
	}

	return true
}

type TimeInterval struct {
	Name          *string            `json:"name,omitempty"`
	TimeIntervals []TimeIntervalItem `json:"time_intervals,omitempty"`
}

func (resource TimeInterval) Equals(other TimeInterval) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	if len(resource.TimeIntervals) != len(other.TimeIntervals) {
		return false
	}

	for i1 := range resource.TimeIntervals {
		if !resource.TimeIntervals[i1].Equals(other.TimeIntervals[i1]) {
			return false
		}
	}

	return true
}

type TimeIntervalItem struct {
	DaysOfMonth []string                `json:"days_of_month,omitempty"`
	Location    *string                 `json:"location,omitempty"`
	Months      []string                `json:"months,omitempty"`
	Times       []TimeIntervalTimeRange `json:"times,omitempty"`
	Weekdays    []string                `json:"weekdays,omitempty"`
	Years       []string                `json:"years,omitempty"`
}

func (resource TimeIntervalItem) Equals(other TimeIntervalItem) bool {

	if len(resource.DaysOfMonth) != len(other.DaysOfMonth) {
		return false
	}

	for i1 := range resource.DaysOfMonth {
		if resource.DaysOfMonth[i1] != other.DaysOfMonth[i1] {
			return false
		}
	}
	if resource.Location == nil && other.Location != nil || resource.Location != nil && other.Location == nil {
		return false
	}

	if resource.Location != nil {
		if *resource.Location != *other.Location {
			return false
		}
	}

	if len(resource.Months) != len(other.Months) {
		return false
	}

	for i1 := range resource.Months {
		if resource.Months[i1] != other.Months[i1] {
			return false
		}
	}

	if len(resource.Times) != len(other.Times) {
		return false
	}

	for i1 := range resource.Times {
		if !resource.Times[i1].Equals(other.Times[i1]) {
			return false
		}
	}

	if len(resource.Weekdays) != len(other.Weekdays) {
		return false
	}

	for i1 := range resource.Weekdays {
		if resource.Weekdays[i1] != other.Weekdays[i1] {
			return false
		}
	}

	if len(resource.Years) != len(other.Years) {
		return false
	}

	for i1 := range resource.Years {
		if resource.Years[i1] != other.Years[i1] {
			return false
		}
	}

	return true
}

type TimeIntervalTimeRange struct {
	EndTime   *string `json:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

func (resource TimeIntervalTimeRange) Equals(other TimeIntervalTimeRange) bool {
	if resource.EndTime == nil && other.EndTime != nil || resource.EndTime != nil && other.EndTime == nil {
		return false
	}

	if resource.EndTime != nil {
		if *resource.EndTime != *other.EndTime {
			return false
		}
	}
	if resource.StartTime == nil && other.StartTime != nil || resource.StartTime != nil && other.StartTime == nil {
		return false
	}

	if resource.StartTime != nil {
		if *resource.StartTime != *other.StartTime {
			return false
		}
	}

	return true
}

type ContactPointType string

const (
	ContactPointTypeAlertmanager ContactPointType = "alertmanager"
	ContactPointTypeDingding     ContactPointType = " dingding"
	ContactPointTypeDiscord      ContactPointType = " discord"
	ContactPointTypeEmail        ContactPointType = " email"
	ContactPointTypeGooglechat   ContactPointType = " googlechat"
	ContactPointTypeKafka        ContactPointType = " kafka"
	ContactPointTypeLine         ContactPointType = " line"
	ContactPointTypeOpsgenie     ContactPointType = " opsgenie"
	ContactPointTypePagerduty    ContactPointType = " pagerduty"
	ContactPointTypePushover     ContactPointType = " pushover"
	ContactPointTypeSensugo      ContactPointType = " sensugo"
	ContactPointTypeSlack        ContactPointType = " slack"
	ContactPointTypeTeams        ContactPointType = " teams"
	ContactPointTypeTelegram     ContactPointType = " telegram"
	ContactPointTypeThreema      ContactPointType = " threema"
	ContactPointTypeVictorops    ContactPointType = " victorops"
	ContactPointTypeWebhook      ContactPointType = " webhook"
	ContactPointTypeWecom        ContactPointType = " wecom"
)

type RuleExecErrState string

const (
	RuleExecErrStateOK       RuleExecErrState = "OK"
	RuleExecErrStateAlerting RuleExecErrState = "Alerting"
	RuleExecErrStateError    RuleExecErrState = "Error"
)

type RuleNoDataState string

const (
	RuleNoDataStateAlerting RuleNoDataState = "Alerting"
	RuleNoDataStateNoData   RuleNoDataState = "NoData"
	RuleNoDataStateOK       RuleNoDataState = "OK"
)
