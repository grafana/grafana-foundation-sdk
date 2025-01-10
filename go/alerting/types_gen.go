// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Query struct {
	// Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
	DatasourceUid *string `json:"datasourceUid,omitempty"`
	// JSON is the raw JSON query and includes the above properties as well as custom properties.
	Model variants.Dataquery `json:"model,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId *string `json:"refId,omitempty"`
	// RelativeTimeRange is the per query start and end time
	// for requests.
	RelativeTimeRange *RelativeTimeRange `json:"relativeTimeRange,omitempty"`
}

// NewQuery creates a new Query object.
func NewQuery() *Query {
	return &Query{}
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode Query from JSON.
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
			return fmt.Errorf("error decoding field 'datasourceUid': %w", err)
		}
	}

	if fields["queryType"] != nil {
		if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
			return fmt.Errorf("error decoding field 'queryType': %w", err)
		}
	}

	if fields["refId"] != nil {
		if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
			return fmt.Errorf("error decoding field 'refId': %w", err)
		}
	}

	if fields["relativeTimeRange"] != nil {
		if err := json.Unmarshal(fields["relativeTimeRange"], &resource.RelativeTimeRange); err != nil {
			return fmt.Errorf("error decoding field 'relativeTimeRange': %w", err)
		}
	}

	if fields["model"] != nil {
		dataqueryTypeHint := ""

		model, err := cog.UnmarshalDataquery(fields["model"], dataqueryTypeHint)
		if err != nil {
			return err
		}
		resource.Model = model
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Query` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Query) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "datasourceUid"
	if fields["datasourceUid"] != nil {
		if string(fields["datasourceUid"]) != "null" {
			if err := json.Unmarshal(fields["datasourceUid"], &resource.DatasourceUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasourceUid", err)...)
			}

		}
		delete(fields, "datasourceUid")

	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {

			dataquery, err := cog.StrictUnmarshalDataquery(fields["model"], "")
			if err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			} else {
				resource.Model = dataquery
			}

		}
		delete(fields, "model")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}

		}
		delete(fields, "refId")

	}
	// Field "relativeTimeRange"
	if fields["relativeTimeRange"] != nil {
		if string(fields["relativeTimeRange"]) != "null" {

			resource.RelativeTimeRange = &RelativeTimeRange{}
			if err := resource.RelativeTimeRange.UnmarshalJSONStrict(fields["relativeTimeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("relativeTimeRange", err)...)
			}

		}
		delete(fields, "relativeTimeRange")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Query", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Query` objects.
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

// Validate checks all the validation constraints that may be defined on `Query` fields for violations and returns them.
func (resource Query) Validate() error {
	var errs cog.BuildErrors
	if resource.Model != nil {
		if err := resource.Model.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("model", err)...)
		}
	}
	if resource.RelativeTimeRange != nil {
		if err := resource.RelativeTimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("relativeTimeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RuleGroup struct {
	FolderUid *string `json:"folderUid,omitempty"`
	// The interval, in seconds, at which all rules in the group are evaluated.
	// If a group contains many rules, the rules are evaluated sequentially.
	Interval *Duration `json:"interval,omitempty"`
	Rules    []Rule    `json:"rules,omitempty"`
	Title    *string   `json:"title,omitempty"`
}

// NewRuleGroup creates a new RuleGroup object.
func NewRuleGroup() *RuleGroup {
	return &RuleGroup{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleGroup` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RuleGroup) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "folderUid"
	if fields["folderUid"] != nil {
		if string(fields["folderUid"]) != "null" {
			if err := json.Unmarshal(fields["folderUid"], &resource.FolderUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderUid", err)...)
			}

		}
		delete(fields, "folderUid")

	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "rules"
	if fields["rules"] != nil {
		if string(fields["rules"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["rules"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Rule

				result1 = Rule{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("rules["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Rules = append(resource.Rules, result1)
			}

		}
		delete(fields, "rules")

	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}

		}
		delete(fields, "title")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RuleGroup", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RuleGroup` objects.
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

// Validate checks all the validation constraints that may be defined on `RuleGroup` fields for violations and returns them.
func (resource RuleGroup) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Rules {
		if err := resource.Rules[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("rules["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type NotificationSettings struct {
	// Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
	// cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
	// use the special value '...' as the sole label name.
	// This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
	// you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
	// Must include 'alertname' and 'grafana_folder' if not using '...'.
	GroupBy []string `json:"group_by,omitempty"`
	// Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
	// which an initial notification has already been sent. (Usually ~5m or more.)
	GroupInterval *string `json:"group_interval,omitempty"`
	// Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
	// inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
	GroupWait *string `json:"group_wait,omitempty"`
	// Override the times when notifications should be muted. These must match the name of a mute time interval defined
	// in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
	// otherwise acts normally.
	MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
	// Name of the receiver to send notifications to.
	Receiver string `json:"receiver"`
	// Override how long to wait before sending a notification again if it has already been sent successfully for an
	// alert. (Usually ~3h or more).
	// Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
	// Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
	// occurs first. `repeat_interval` should not be less than `group_interval`.
	RepeatInterval *string `json:"repeat_interval,omitempty"`
}

// NewNotificationSettings creates a new NotificationSettings object.
func NewNotificationSettings() *NotificationSettings {
	return &NotificationSettings{
		GroupBy: []string{"alertname", "grafana_folder"},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NotificationSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "group_by"
	if fields["group_by"] != nil {
		if string(fields["group_by"]) != "null" {

			if err := json.Unmarshal(fields["group_by"], &resource.GroupBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_by", err)...)
			}

		}
		delete(fields, "group_by")

	}
	// Field "group_interval"
	if fields["group_interval"] != nil {
		if string(fields["group_interval"]) != "null" {
			if err := json.Unmarshal(fields["group_interval"], &resource.GroupInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_interval", err)...)
			}

		}
		delete(fields, "group_interval")

	}
	// Field "group_wait"
	if fields["group_wait"] != nil {
		if string(fields["group_wait"]) != "null" {
			if err := json.Unmarshal(fields["group_wait"], &resource.GroupWait); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_wait", err)...)
			}

		}
		delete(fields, "group_wait")

	}
	// Field "mute_time_intervals"
	if fields["mute_time_intervals"] != nil {
		if string(fields["mute_time_intervals"]) != "null" {

			if err := json.Unmarshal(fields["mute_time_intervals"], &resource.MuteTimeIntervals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mute_time_intervals", err)...)
			}

		}
		delete(fields, "mute_time_intervals")

	}
	// Field "receiver"
	if fields["receiver"] != nil {
		if string(fields["receiver"]) != "null" {
			if err := json.Unmarshal(fields["receiver"], &resource.Receiver); err != nil {
				errs = append(errs, cog.MakeBuildErrors("receiver", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("receiver", errors.New("required field is null"))...)

		}
		delete(fields, "receiver")
	} else {
		errs = append(errs, cog.MakeBuildErrors("receiver", errors.New("required field is missing from input"))...)
	}
	// Field "repeat_interval"
	if fields["repeat_interval"] != nil {
		if string(fields["repeat_interval"]) != "null" {
			if err := json.Unmarshal(fields["repeat_interval"], &resource.RepeatInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat_interval", err)...)
			}

		}
		delete(fields, "repeat_interval")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NotificationSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NotificationSettings` objects.
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

// Validate checks all the validation constraints that may be defined on `NotificationSettings` fields for violations and returns them.
func (resource NotificationSettings) Validate() error {
	return nil
}

// Duration in seconds.
type Duration int64

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
type ContactPoint struct {
	DisableResolveMessage *bool `json:"disableResolveMessage,omitempty"`
	// Name is used as grouping key in the UI. Contact points with the
	// same name will be grouped in the UI.
	Name       *string          `json:"name,omitempty"`
	Provenance *string          `json:"provenance,omitempty"`
	Settings   Json             `json:"settings"`
	Type       ContactPointType `json:"type"`
	// UID is the unique identifier of the contact point. The UID can be
	// set by the user.
	Uid *string `json:"uid,omitempty"`
}

// NewContactPoint creates a new ContactPoint object.
func NewContactPoint() *ContactPoint {
	return &ContactPoint{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ContactPoint` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ContactPoint) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "disableResolveMessage"
	if fields["disableResolveMessage"] != nil {
		if string(fields["disableResolveMessage"]) != "null" {
			if err := json.Unmarshal(fields["disableResolveMessage"], &resource.DisableResolveMessage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("disableResolveMessage", err)...)
			}

		}
		delete(fields, "disableResolveMessage")

	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "provenance"
	if fields["provenance"] != nil {
		if string(fields["provenance"]) != "null" {
			if err := json.Unmarshal(fields["provenance"], &resource.Provenance); err != nil {
				errs = append(errs, cog.MakeBuildErrors("provenance", err)...)
			}

		}
		delete(fields, "provenance")

	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {
			if err := json.Unmarshal(fields["settings"], &resource.Settings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is null"))...)

		}
		delete(fields, "settings")
	} else {
		errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}

		}
		delete(fields, "uid")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ContactPoint", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ContactPoint` objects.
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

// Validate checks all the validation constraints that may be defined on `ContactPoint` fields for violations and returns them.
func (resource ContactPoint) Validate() error {
	var errs cog.BuildErrors
	if resource.Uid != nil {
		if !(len([]rune(*resource.Uid)) >= 0x1) {
			errs = append(errs, cog.MakeBuildErrors(
				"uid",
				errors.New("must be >= 1"),
			)...)
		}
		if !(len([]rune(*resource.Uid)) <= 0x28) {
			errs = append(errs, cog.MakeBuildErrors(
				"uid",
				errors.New("must be <= 40"),
			)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// NewMatcher creates a new Matcher object.
func NewMatcher() *Matcher {
	return &Matcher{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Matcher` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Matcher) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "Name"
	if fields["Name"] != nil {
		if string(fields["Name"]) != "null" {
			if err := json.Unmarshal(fields["Name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("Name", err)...)
			}

		}
		delete(fields, "Name")

	}
	// Field "Type"
	if fields["Type"] != nil {
		if string(fields["Type"]) != "null" {
			if err := json.Unmarshal(fields["Type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("Type", err)...)
			}

		}
		delete(fields, "Type")

	}
	// Field "Value"
	if fields["Value"] != nil {
		if string(fields["Value"]) != "null" {
			if err := json.Unmarshal(fields["Value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("Value", err)...)
			}

		}
		delete(fields, "Value")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Matcher", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Matcher` objects.
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

// Validate checks all the validation constraints that may be defined on `Matcher` fields for violations and returns them.
func (resource Matcher) Validate() error {
	return nil
}

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
type Matchers []Matcher

type MuteTiming struct {
	Name          *string        `json:"name,omitempty"`
	TimeIntervals []TimeInterval `json:"time_intervals,omitempty"`
}

// NewMuteTiming creates a new MuteTiming object.
func NewMuteTiming() *MuteTiming {
	return &MuteTiming{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MuteTiming` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MuteTiming) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "time_intervals"
	if fields["time_intervals"] != nil {
		if string(fields["time_intervals"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["time_intervals"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TimeInterval

				result1 = TimeInterval{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("time_intervals["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.TimeIntervals = append(resource.TimeIntervals, result1)
			}

		}
		delete(fields, "time_intervals")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MuteTiming", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MuteTiming` objects.
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

// Validate checks all the validation constraints that may be defined on `MuteTiming` fields for violations and returns them.
func (resource MuteTiming) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.TimeIntervals {
		if err := resource.TimeIntervals[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("time_intervals["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type NotificationTemplate struct {
	Name       *string     `json:"name,omitempty"`
	Provenance *Provenance `json:"provenance,omitempty"`
	Template   *string     `json:"template,omitempty"`
}

// NewNotificationTemplate creates a new NotificationTemplate object.
func NewNotificationTemplate() *NotificationTemplate {
	return &NotificationTemplate{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationTemplate` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NotificationTemplate) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "provenance"
	if fields["provenance"] != nil {
		if string(fields["provenance"]) != "null" {
			if err := json.Unmarshal(fields["provenance"], &resource.Provenance); err != nil {
				errs = append(errs, cog.MakeBuildErrors("provenance", err)...)
			}

		}
		delete(fields, "provenance")

	}
	// Field "template"
	if fields["template"] != nil {
		if string(fields["template"]) != "null" {
			if err := json.Unmarshal(fields["template"], &resource.Template); err != nil {
				errs = append(errs, cog.MakeBuildErrors("template", err)...)
			}

		}
		delete(fields, "template")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NotificationTemplate", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NotificationTemplate` objects.
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

// Validate checks all the validation constraints that may be defined on `NotificationTemplate` fields for violations and returns them.
func (resource NotificationTemplate) Validate() error {
	return nil
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

// NewRule creates a new Rule object.
func NewRule() *Rule {
	return &Rule{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Rule` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Rule) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "annotations"
	if fields["annotations"] != nil {
		if string(fields["annotations"]) != "null" {

			if err := json.Unmarshal(fields["annotations"], &resource.Annotations); err != nil {
				errs = append(errs, cog.MakeBuildErrors("annotations", err)...)
			}

		}
		delete(fields, "annotations")

	}
	// Field "condition"
	if fields["condition"] != nil {
		if string(fields["condition"]) != "null" {
			if err := json.Unmarshal(fields["condition"], &resource.Condition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("condition", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("condition", errors.New("required field is null"))...)

		}
		delete(fields, "condition")
	} else {
		errs = append(errs, cog.MakeBuildErrors("condition", errors.New("required field is missing from input"))...)
	}
	// Field "data"
	if fields["data"] != nil {
		if string(fields["data"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["data"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Query

				result1 = Query{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("data["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Data = append(resource.Data, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("data", errors.New("required field is null"))...)

		}
		delete(fields, "data")
	} else {
		errs = append(errs, cog.MakeBuildErrors("data", errors.New("required field is missing from input"))...)
	}
	// Field "execErrState"
	if fields["execErrState"] != nil {
		if string(fields["execErrState"]) != "null" {
			if err := json.Unmarshal(fields["execErrState"], &resource.ExecErrState); err != nil {
				errs = append(errs, cog.MakeBuildErrors("execErrState", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("execErrState", errors.New("required field is null"))...)

		}
		delete(fields, "execErrState")
	} else {
		errs = append(errs, cog.MakeBuildErrors("execErrState", errors.New("required field is missing from input"))...)
	}
	// Field "folderUID"
	if fields["folderUID"] != nil {
		if string(fields["folderUID"]) != "null" {
			if err := json.Unmarshal(fields["folderUID"], &resource.FolderUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderUID", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("folderUID", errors.New("required field is null"))...)

		}
		delete(fields, "folderUID")
	} else {
		errs = append(errs, cog.MakeBuildErrors("folderUID", errors.New("required field is missing from input"))...)
	}
	// Field "for"
	if fields["for"] != nil {
		if string(fields["for"]) != "null" {
			if err := json.Unmarshal(fields["for"], &resource.For); err != nil {
				errs = append(errs, cog.MakeBuildErrors("for", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("for", errors.New("required field is null"))...)

		}
		delete(fields, "for")
	} else {
		errs = append(errs, cog.MakeBuildErrors("for", errors.New("required field is missing from input"))...)
	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}

		}
		delete(fields, "id")

	}
	// Field "isPaused"
	if fields["isPaused"] != nil {
		if string(fields["isPaused"]) != "null" {
			if err := json.Unmarshal(fields["isPaused"], &resource.IsPaused); err != nil {
				errs = append(errs, cog.MakeBuildErrors("isPaused", err)...)
			}

		}
		delete(fields, "isPaused")

	}
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {

			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}

		}
		delete(fields, "labels")

	}
	// Field "noDataState"
	if fields["noDataState"] != nil {
		if string(fields["noDataState"]) != "null" {
			if err := json.Unmarshal(fields["noDataState"], &resource.NoDataState); err != nil {
				errs = append(errs, cog.MakeBuildErrors("noDataState", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("noDataState", errors.New("required field is null"))...)

		}
		delete(fields, "noDataState")
	} else {
		errs = append(errs, cog.MakeBuildErrors("noDataState", errors.New("required field is missing from input"))...)
	}
	// Field "notification_settings"
	if fields["notification_settings"] != nil {
		if string(fields["notification_settings"]) != "null" {

			resource.NotificationSettings = &NotificationSettings{}
			if err := resource.NotificationSettings.UnmarshalJSONStrict(fields["notification_settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("notification_settings", err)...)
			}

		}
		delete(fields, "notification_settings")

	}
	// Field "orgID"
	if fields["orgID"] != nil {
		if string(fields["orgID"]) != "null" {
			if err := json.Unmarshal(fields["orgID"], &resource.OrgID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orgID", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("orgID", errors.New("required field is null"))...)

		}
		delete(fields, "orgID")
	} else {
		errs = append(errs, cog.MakeBuildErrors("orgID", errors.New("required field is missing from input"))...)
	}
	// Field "provenance"
	if fields["provenance"] != nil {
		if string(fields["provenance"]) != "null" {
			if err := json.Unmarshal(fields["provenance"], &resource.Provenance); err != nil {
				errs = append(errs, cog.MakeBuildErrors("provenance", err)...)
			}

		}
		delete(fields, "provenance")

	}
	// Field "ruleGroup"
	if fields["ruleGroup"] != nil {
		if string(fields["ruleGroup"]) != "null" {
			if err := json.Unmarshal(fields["ruleGroup"], &resource.RuleGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("ruleGroup", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("ruleGroup", errors.New("required field is null"))...)

		}
		delete(fields, "ruleGroup")
	} else {
		errs = append(errs, cog.MakeBuildErrors("ruleGroup", errors.New("required field is missing from input"))...)
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}

		}
		delete(fields, "uid")

	}
	// Field "updated"
	if fields["updated"] != nil {
		if string(fields["updated"]) != "null" {
			if err := json.Unmarshal(fields["updated"], &resource.Updated); err != nil {
				errs = append(errs, cog.MakeBuildErrors("updated", err)...)
			}

		}
		delete(fields, "updated")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Rule", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Rule` objects.
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

// Validate checks all the validation constraints that may be defined on `Rule` fields for violations and returns them.
func (resource Rule) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Data {
		if err := resource.Data[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("data["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.NotificationSettings != nil {
		if err := resource.NotificationSettings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("notification_settings", err)...)
		}
	}
	if !(len([]rune(resource.RuleGroup)) >= 0x1) {
		errs = append(errs, cog.MakeBuildErrors(
			"ruleGroup",
			errors.New("must be >= 1"),
		)...)
	}
	if !(len([]rune(resource.RuleGroup)) <= 0xbe) {
		errs = append(errs, cog.MakeBuildErrors(
			"ruleGroup",
			errors.New("must be <= 190"),
		)...)
	}
	if !(len([]rune(resource.Title)) >= 0x1) {
		errs = append(errs, cog.MakeBuildErrors(
			"title",
			errors.New("must be >= 1"),
		)...)
	}
	if !(len([]rune(resource.Title)) <= 0xbe) {
		errs = append(errs, cog.MakeBuildErrors(
			"title",
			errors.New("must be <= 190"),
		)...)
	}
	if resource.Uid != nil {
		if !(len([]rune(*resource.Uid)) >= 0x1) {
			errs = append(errs, cog.MakeBuildErrors(
				"uid",
				errors.New("must be >= 1"),
			)...)
		}
		if !(len([]rune(*resource.Uid)) <= 0x28) {
			errs = append(errs, cog.MakeBuildErrors(
				"uid",
				errors.New("must be <= 40"),
			)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// RelativeTimeRange is the per query start and end time
// for requests.
type RelativeTimeRange struct {
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	From *Duration `json:"from,omitempty"`
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	To *Duration `json:"to,omitempty"`
}

// NewRelativeTimeRange creates a new RelativeTimeRange object.
func NewRelativeTimeRange() *RelativeTimeRange {
	return &RelativeTimeRange{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RelativeTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RelativeTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}

		}
		delete(fields, "from")

	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}

		}
		delete(fields, "to")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RelativeTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RelativeTimeRange` objects.
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

// Validate checks all the validation constraints that may be defined on `RelativeTimeRange` fields for violations and returns them.
func (resource RelativeTimeRange) Validate() error {
	return nil
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
type NotificationPolicy struct {
	Continue      *bool    `json:"continue,omitempty"`
	GroupBy       []string `json:"group_by,omitempty"`
	GroupInterval *string  `json:"group_interval,omitempty"`
	GroupWait     *string  `json:"group_wait,omitempty"`
	// Deprecated. Remove before v1.0 release.
	Match   map[string]string `json:"match,omitempty"`
	MatchRe *MatchRegexps     `json:"match_re,omitempty"`
	// Matchers is a slice of Matchers that is sortable, implements Stringer, and
	// provides a Matches method to match a LabelSet against all Matchers in the
	// slice. Note that some users of Matchers might require it to be sorted.
	Matchers          *Matchers            `json:"matchers,omitempty"`
	MuteTimeIntervals []string             `json:"mute_time_intervals,omitempty"`
	ObjectMatchers    *ObjectMatchers      `json:"object_matchers,omitempty"`
	Provenance        *Provenance          `json:"provenance,omitempty"`
	Receiver          *string              `json:"receiver,omitempty"`
	RepeatInterval    *string              `json:"repeat_interval,omitempty"`
	Routes            []NotificationPolicy `json:"routes,omitempty"`
}

// NewNotificationPolicy creates a new NotificationPolicy object.
func NewNotificationPolicy() *NotificationPolicy {
	return &NotificationPolicy{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationPolicy` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NotificationPolicy) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "continue"
	if fields["continue"] != nil {
		if string(fields["continue"]) != "null" {
			if err := json.Unmarshal(fields["continue"], &resource.Continue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("continue", err)...)
			}

		}
		delete(fields, "continue")

	}
	// Field "group_by"
	if fields["group_by"] != nil {
		if string(fields["group_by"]) != "null" {

			if err := json.Unmarshal(fields["group_by"], &resource.GroupBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_by", err)...)
			}

		}
		delete(fields, "group_by")

	}
	// Field "group_interval"
	if fields["group_interval"] != nil {
		if string(fields["group_interval"]) != "null" {
			if err := json.Unmarshal(fields["group_interval"], &resource.GroupInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_interval", err)...)
			}

		}
		delete(fields, "group_interval")

	}
	// Field "group_wait"
	if fields["group_wait"] != nil {
		if string(fields["group_wait"]) != "null" {
			if err := json.Unmarshal(fields["group_wait"], &resource.GroupWait); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group_wait", err)...)
			}

		}
		delete(fields, "group_wait")

	}
	// Field "match"
	if fields["match"] != nil {
		if string(fields["match"]) != "null" {

			if err := json.Unmarshal(fields["match"], &resource.Match); err != nil {
				errs = append(errs, cog.MakeBuildErrors("match", err)...)
			}

		}
		delete(fields, "match")

	}
	// Field "match_re"
	if fields["match_re"] != nil {
		if string(fields["match_re"]) != "null" {

			if err := json.Unmarshal(fields["match_re"], &resource.MatchRe); err != nil {
				errs = append(errs, cog.MakeBuildErrors("match_re", err)...)
			}

		}
		delete(fields, "match_re")

	}
	// Field "matchers"
	if fields["matchers"] != nil {
		if string(fields["matchers"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["matchers"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Matcher

				result1 = Matcher{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("matchers["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Matchers = cog.ToPtr(append(*resource.Matchers, result1))
			}

		}
		delete(fields, "matchers")

	}
	// Field "mute_time_intervals"
	if fields["mute_time_intervals"] != nil {
		if string(fields["mute_time_intervals"]) != "null" {

			if err := json.Unmarshal(fields["mute_time_intervals"], &resource.MuteTimeIntervals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mute_time_intervals", err)...)
			}

		}
		delete(fields, "mute_time_intervals")

	}
	// Field "object_matchers"
	if fields["object_matchers"] != nil {
		if string(fields["object_matchers"]) != "null" {

			if err := json.Unmarshal(fields["object_matchers"], &resource.ObjectMatchers); err != nil {
				errs = append(errs, cog.MakeBuildErrors("object_matchers", err)...)
			}

		}
		delete(fields, "object_matchers")

	}
	// Field "provenance"
	if fields["provenance"] != nil {
		if string(fields["provenance"]) != "null" {
			if err := json.Unmarshal(fields["provenance"], &resource.Provenance); err != nil {
				errs = append(errs, cog.MakeBuildErrors("provenance", err)...)
			}

		}
		delete(fields, "provenance")

	}
	// Field "receiver"
	if fields["receiver"] != nil {
		if string(fields["receiver"]) != "null" {
			if err := json.Unmarshal(fields["receiver"], &resource.Receiver); err != nil {
				errs = append(errs, cog.MakeBuildErrors("receiver", err)...)
			}

		}
		delete(fields, "receiver")

	}
	// Field "repeat_interval"
	if fields["repeat_interval"] != nil {
		if string(fields["repeat_interval"]) != "null" {
			if err := json.Unmarshal(fields["repeat_interval"], &resource.RepeatInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat_interval", err)...)
			}

		}
		delete(fields, "repeat_interval")

	}
	// Field "routes"
	if fields["routes"] != nil {
		if string(fields["routes"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["routes"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 NotificationPolicy

				result1 = NotificationPolicy{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("routes["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Routes = append(resource.Routes, result1)
			}

		}
		delete(fields, "routes")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NotificationPolicy", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NotificationPolicy` objects.
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

// Validate checks all the validation constraints that may be defined on `NotificationPolicy` fields for violations and returns them.
func (resource NotificationPolicy) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Routes {
		if err := resource.Routes[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("routes["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TimeInterval struct {
	Name          *string            `json:"name,omitempty"`
	TimeIntervals []TimeIntervalItem `json:"time_intervals,omitempty"`
}

// NewTimeInterval creates a new TimeInterval object.
func NewTimeInterval() *TimeInterval {
	return &TimeInterval{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeInterval` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeInterval) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "time_intervals"
	if fields["time_intervals"] != nil {
		if string(fields["time_intervals"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["time_intervals"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TimeIntervalItem

				result1 = TimeIntervalItem{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("time_intervals["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.TimeIntervals = append(resource.TimeIntervals, result1)
			}

		}
		delete(fields, "time_intervals")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeInterval", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeInterval` objects.
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

// Validate checks all the validation constraints that may be defined on `TimeInterval` fields for violations and returns them.
func (resource TimeInterval) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.TimeIntervals {
		if err := resource.TimeIntervals[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("time_intervals["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TimeIntervalItem struct {
	DaysOfMonth []string                `json:"days_of_month,omitempty"`
	Location    *string                 `json:"location,omitempty"`
	Months      []string                `json:"months,omitempty"`
	Times       []TimeIntervalTimeRange `json:"times,omitempty"`
	Weekdays    []string                `json:"weekdays,omitempty"`
	Years       []string                `json:"years,omitempty"`
}

// NewTimeIntervalItem creates a new TimeIntervalItem object.
func NewTimeIntervalItem() *TimeIntervalItem {
	return &TimeIntervalItem{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeIntervalItem` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeIntervalItem) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "days_of_month"
	if fields["days_of_month"] != nil {
		if string(fields["days_of_month"]) != "null" {

			if err := json.Unmarshal(fields["days_of_month"], &resource.DaysOfMonth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("days_of_month", err)...)
			}

		}
		delete(fields, "days_of_month")

	}
	// Field "location"
	if fields["location"] != nil {
		if string(fields["location"]) != "null" {
			if err := json.Unmarshal(fields["location"], &resource.Location); err != nil {
				errs = append(errs, cog.MakeBuildErrors("location", err)...)
			}

		}
		delete(fields, "location")

	}
	// Field "months"
	if fields["months"] != nil {
		if string(fields["months"]) != "null" {

			if err := json.Unmarshal(fields["months"], &resource.Months); err != nil {
				errs = append(errs, cog.MakeBuildErrors("months", err)...)
			}

		}
		delete(fields, "months")

	}
	// Field "times"
	if fields["times"] != nil {
		if string(fields["times"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["times"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TimeIntervalTimeRange

				result1 = TimeIntervalTimeRange{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("times["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Times = append(resource.Times, result1)
			}

		}
		delete(fields, "times")

	}
	// Field "weekdays"
	if fields["weekdays"] != nil {
		if string(fields["weekdays"]) != "null" {

			if err := json.Unmarshal(fields["weekdays"], &resource.Weekdays); err != nil {
				errs = append(errs, cog.MakeBuildErrors("weekdays", err)...)
			}

		}
		delete(fields, "weekdays")

	}
	// Field "years"
	if fields["years"] != nil {
		if string(fields["years"]) != "null" {

			if err := json.Unmarshal(fields["years"], &resource.Years); err != nil {
				errs = append(errs, cog.MakeBuildErrors("years", err)...)
			}

		}
		delete(fields, "years")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeIntervalItem", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeIntervalItem` objects.
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

// Validate checks all the validation constraints that may be defined on `TimeIntervalItem` fields for violations and returns them.
func (resource TimeIntervalItem) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Times {
		if err := resource.Times[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("times["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TimeIntervalTimeRange struct {
	EndTime   *string `json:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

// NewTimeIntervalTimeRange creates a new TimeIntervalTimeRange object.
func NewTimeIntervalTimeRange() *TimeIntervalTimeRange {
	return &TimeIntervalTimeRange{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeIntervalTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeIntervalTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "end_time"
	if fields["end_time"] != nil {
		if string(fields["end_time"]) != "null" {
			if err := json.Unmarshal(fields["end_time"], &resource.EndTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("end_time", err)...)
			}

		}
		delete(fields, "end_time")

	}
	// Field "start_time"
	if fields["start_time"] != nil {
		if string(fields["start_time"]) != "null" {
			if err := json.Unmarshal(fields["start_time"], &resource.StartTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("start_time", err)...)
			}

		}
		delete(fields, "start_time")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeIntervalTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeIntervalTimeRange` objects.
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

// Validate checks all the validation constraints that may be defined on `TimeIntervalTimeRange` fields for violations and returns them.
func (resource TimeIntervalTimeRange) Validate() error {
	return nil
}

type ContactPointType string

const (
	ContactPointTypeAlertmanager ContactPointType = "alertmanager"
	ContactPointTypeDingding     ContactPointType = "dingding"
	ContactPointTypeDiscord      ContactPointType = "discord"
	ContactPointTypeEmail        ContactPointType = "email"
	ContactPointTypeGooglechat   ContactPointType = "googlechat"
	ContactPointTypeKafka        ContactPointType = "kafka"
	ContactPointTypeLine         ContactPointType = "line"
	ContactPointTypeOpsgenie     ContactPointType = "opsgenie"
	ContactPointTypePagerduty    ContactPointType = "pagerduty"
	ContactPointTypePushover     ContactPointType = "pushover"
	ContactPointTypeSensugo      ContactPointType = "sensugo"
	ContactPointTypeSlack        ContactPointType = "slack"
	ContactPointTypeTeams        ContactPointType = "teams"
	ContactPointTypeTelegram     ContactPointType = "telegram"
	ContactPointTypeThreema      ContactPointType = "threema"
	ContactPointTypeVictorops    ContactPointType = "victorops"
	ContactPointTypeWebhook      ContactPointType = "webhook"
	ContactPointTypeWecom        ContactPointType = "wecom"
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
