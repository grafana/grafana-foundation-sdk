// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type NoDataState string

const (
	NoDataStateNoData   NoDataState = "NoData"
	NoDataStateOk       NoDataState = "Ok"
	NoDataStateAlerting NoDataState = "Alerting"
	NoDataStateKeepLast NoDataState = "KeepLast"
)

type ExecErrState string

const (
	ExecErrStateError    ExecErrState = "Error"
	ExecErrStateOk       ExecErrState = "Ok"
	ExecErrStateAlerting ExecErrState = "Alerting"
	ExecErrStateKeepLast ExecErrState = "KeepLast"
)

type TimeIntervalRef string

type AlertRuleSpec struct {
	Title                       string                    `json:"title"`
	Paused                      *bool                     `json:"paused,omitempty"`
	Trigger                     IntervalTrigger           `json:"trigger"`
	Labels                      map[string]TemplateString `json:"labels,omitempty"`
	Annotations                 map[string]TemplateString `json:"annotations,omitempty"`
	For                         *string                   `json:"for,omitempty"`
	KeepFiringFor               *string                   `json:"keepFiringFor,omitempty"`
	MissingSeriesEvalsToResolve *int64                    `json:"missingSeriesEvalsToResolve,omitempty"`
	NoDataState                 NoDataState               `json:"noDataState"`
	ExecErrState                ExecErrState              `json:"execErrState"`
	NotificationSettings        *NotificationSettings     `json:"notificationSettings,omitempty"`
	Expressions                 ExpressionMap             `json:"expressions"`
	PanelRef                    *PanelRef                 `json:"panelRef,omitempty"`
}

// NewAlertRuleSpec creates a new AlertRuleSpec object.
func NewAlertRuleSpec() *AlertRuleSpec {
	return &AlertRuleSpec{
		Trigger:      *NewIntervalTrigger(),
		NoDataState:  NoDataStateNoData,
		ExecErrState: ExecErrStateError,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AlertRuleSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "paused"
	if fields["paused"] != nil {
		if string(fields["paused"]) != "null" {
			if err := json.Unmarshal(fields["paused"], &resource.Paused); err != nil {
				errs = append(errs, cog.MakeBuildErrors("paused", err)...)
			}

		}
		delete(fields, "paused")

	}
	// Field "trigger"
	if fields["trigger"] != nil {
		if string(fields["trigger"]) != "null" {

			resource.Trigger = IntervalTrigger{}
			if err := resource.Trigger.UnmarshalJSONStrict(fields["trigger"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is null"))...)

		}
		delete(fields, "trigger")
	} else {
		errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is missing from input"))...)
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
	// Field "annotations"
	if fields["annotations"] != nil {
		if string(fields["annotations"]) != "null" {

			if err := json.Unmarshal(fields["annotations"], &resource.Annotations); err != nil {
				errs = append(errs, cog.MakeBuildErrors("annotations", err)...)
			}

		}
		delete(fields, "annotations")

	}
	// Field "for"
	if fields["for"] != nil {
		if string(fields["for"]) != "null" {
			if err := json.Unmarshal(fields["for"], &resource.For); err != nil {
				errs = append(errs, cog.MakeBuildErrors("for", err)...)
			}

		}
		delete(fields, "for")

	}
	// Field "keepFiringFor"
	if fields["keepFiringFor"] != nil {
		if string(fields["keepFiringFor"]) != "null" {
			if err := json.Unmarshal(fields["keepFiringFor"], &resource.KeepFiringFor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("keepFiringFor", err)...)
			}

		}
		delete(fields, "keepFiringFor")

	}
	// Field "missingSeriesEvalsToResolve"
	if fields["missingSeriesEvalsToResolve"] != nil {
		if string(fields["missingSeriesEvalsToResolve"]) != "null" {
			if err := json.Unmarshal(fields["missingSeriesEvalsToResolve"], &resource.MissingSeriesEvalsToResolve); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missingSeriesEvalsToResolve", err)...)
			}

		}
		delete(fields, "missingSeriesEvalsToResolve")

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
	// Field "notificationSettings"
	if fields["notificationSettings"] != nil {
		if string(fields["notificationSettings"]) != "null" {

			resource.NotificationSettings = &NotificationSettings{}
			if err := resource.NotificationSettings.UnmarshalJSONStrict(fields["notificationSettings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("notificationSettings", err)...)
			}

		}
		delete(fields, "notificationSettings")

	}
	// Field "expressions"
	if fields["expressions"] != nil {
		if string(fields["expressions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["expressions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]Expression, len(partialMap))
			for key1 := range partialMap {
				var result1 Expression

				result1 = Expression{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("expressions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Expressions = ExpressionMap(parsedMap1)
		} else {
			errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is null"))...)

		}
		delete(fields, "expressions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is missing from input"))...)
	}
	// Field "panelRef"
	if fields["panelRef"] != nil {
		if string(fields["panelRef"]) != "null" {

			resource.PanelRef = &PanelRef{}
			if err := resource.PanelRef.UnmarshalJSONStrict(fields["panelRef"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("panelRef", err)...)
			}

		}
		delete(fields, "panelRef")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AlertRuleSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `AlertRuleSpec` objects.
func (resource AlertRuleSpec) Equals(other AlertRuleSpec) bool {
	if resource.Title != other.Title {
		return false
	}
	if resource.Paused == nil && other.Paused != nil || resource.Paused != nil && other.Paused == nil {
		return false
	}

	if resource.Paused != nil {
		if *resource.Paused != *other.Paused {
			return false
		}
	}
	if !resource.Trigger.Equals(other.Trigger) {
		return false
	}

	if len(resource.Labels) != len(other.Labels) {
		return false
	}

	for key1 := range resource.Labels {
		if resource.Labels[key1] != other.Labels[key1] {
			return false
		}
	}

	if len(resource.Annotations) != len(other.Annotations) {
		return false
	}

	for key1 := range resource.Annotations {
		if resource.Annotations[key1] != other.Annotations[key1] {
			return false
		}
	}
	if resource.For == nil && other.For != nil || resource.For != nil && other.For == nil {
		return false
	}

	if resource.For != nil {
		if *resource.For != *other.For {
			return false
		}
	}
	if resource.KeepFiringFor == nil && other.KeepFiringFor != nil || resource.KeepFiringFor != nil && other.KeepFiringFor == nil {
		return false
	}

	if resource.KeepFiringFor != nil {
		if *resource.KeepFiringFor != *other.KeepFiringFor {
			return false
		}
	}
	if resource.MissingSeriesEvalsToResolve == nil && other.MissingSeriesEvalsToResolve != nil || resource.MissingSeriesEvalsToResolve != nil && other.MissingSeriesEvalsToResolve == nil {
		return false
	}

	if resource.MissingSeriesEvalsToResolve != nil {
		if *resource.MissingSeriesEvalsToResolve != *other.MissingSeriesEvalsToResolve {
			return false
		}
	}
	if resource.NoDataState != other.NoDataState {
		return false
	}
	if resource.ExecErrState != other.ExecErrState {
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

	if len(resource.Expressions) != len(other.Expressions) {
		return false
	}

	for key1 := range resource.Expressions {
		if !resource.Expressions[key1].Equals(other.Expressions[key1]) {
			return false
		}
	}
	if resource.PanelRef == nil && other.PanelRef != nil || resource.PanelRef != nil && other.PanelRef == nil {
		return false
	}

	if resource.PanelRef != nil {
		if !resource.PanelRef.Equals(*other.PanelRef) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AlertRuleSpec` fields for violations and returns them.
func (resource AlertRuleSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Trigger.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
	}
	if resource.MissingSeriesEvalsToResolve != nil {
		if !(*resource.MissingSeriesEvalsToResolve >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"missingSeriesEvalsToResolve",
				errors.New("must be >= 0"),
			)...)
		}
	}
	if resource.NotificationSettings != nil {
		if err := resource.NotificationSettings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("notificationSettings", err)...)
		}
	}
	if resource.PanelRef != nil {
		if err := resource.PanelRef.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("panelRef", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type IntervalTrigger struct {
	Interval PromDuration `json:"interval"`
}

// NewIntervalTrigger creates a new IntervalTrigger object.
func NewIntervalTrigger() *IntervalTrigger {
	return &IntervalTrigger{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalTrigger` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *IntervalTrigger) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("interval", errors.New("required field is null"))...)

		}
		delete(fields, "interval")
	} else {
		errs = append(errs, cog.MakeBuildErrors("interval", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("IntervalTrigger", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `IntervalTrigger` objects.
func (resource IntervalTrigger) Equals(other IntervalTrigger) bool {
	if resource.Interval != other.Interval {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `IntervalTrigger` fields for violations and returns them.
func (resource IntervalTrigger) Validate() error {
	return nil
}

type PromDuration string

type TemplateString string

type NotificationSettings = SimplifiedRoutingOrNamedRoutingTree

// NewNotificationSettings creates a new NotificationSettings object.
func NewNotificationSettings() *NotificationSettings {
	return NewSimplifiedRoutingOrNamedRoutingTree()
}

type SimplifiedRouting struct {
	Type                NotificationSettingsType `json:"type"`
	Receiver            string                   `json:"receiver"`
	GroupBy             []string                 `json:"groupBy,omitempty"`
	GroupWait           *PromDuration            `json:"groupWait,omitempty"`
	GroupInterval       *PromDuration            `json:"groupInterval,omitempty"`
	RepeatInterval      *PromDuration            `json:"repeatInterval,omitempty"`
	MuteTimeIntervals   []TimeIntervalRef        `json:"muteTimeIntervals,omitempty"`
	ActiveTimeIntervals []TimeIntervalRef        `json:"activeTimeIntervals,omitempty"`
}

// NewSimplifiedRouting creates a new SimplifiedRouting object.
func NewSimplifiedRouting() *SimplifiedRouting {
	return &SimplifiedRouting{
		Type: NotificationSettingsTypeSimplifiedRouting,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimplifiedRouting` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SimplifiedRouting) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "groupBy"
	if fields["groupBy"] != nil {
		if string(fields["groupBy"]) != "null" {

			if err := json.Unmarshal(fields["groupBy"], &resource.GroupBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupBy", err)...)
			}

		}
		delete(fields, "groupBy")

	}
	// Field "groupWait"
	if fields["groupWait"] != nil {
		if string(fields["groupWait"]) != "null" {
			if err := json.Unmarshal(fields["groupWait"], &resource.GroupWait); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupWait", err)...)
			}

		}
		delete(fields, "groupWait")

	}
	// Field "groupInterval"
	if fields["groupInterval"] != nil {
		if string(fields["groupInterval"]) != "null" {
			if err := json.Unmarshal(fields["groupInterval"], &resource.GroupInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupInterval", err)...)
			}

		}
		delete(fields, "groupInterval")

	}
	// Field "repeatInterval"
	if fields["repeatInterval"] != nil {
		if string(fields["repeatInterval"]) != "null" {
			if err := json.Unmarshal(fields["repeatInterval"], &resource.RepeatInterval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeatInterval", err)...)
			}

		}
		delete(fields, "repeatInterval")

	}
	// Field "muteTimeIntervals"
	if fields["muteTimeIntervals"] != nil {
		if string(fields["muteTimeIntervals"]) != "null" {

			if err := json.Unmarshal(fields["muteTimeIntervals"], &resource.MuteTimeIntervals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("muteTimeIntervals", err)...)
			}

		}
		delete(fields, "muteTimeIntervals")

	}
	// Field "activeTimeIntervals"
	if fields["activeTimeIntervals"] != nil {
		if string(fields["activeTimeIntervals"]) != "null" {

			if err := json.Unmarshal(fields["activeTimeIntervals"], &resource.ActiveTimeIntervals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("activeTimeIntervals", err)...)
			}

		}
		delete(fields, "activeTimeIntervals")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SimplifiedRouting", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `SimplifiedRouting` objects.
func (resource SimplifiedRouting) Equals(other SimplifiedRouting) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Receiver != other.Receiver {
		return false
	}

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if resource.GroupBy[i1] != other.GroupBy[i1] {
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
	if resource.GroupInterval == nil && other.GroupInterval != nil || resource.GroupInterval != nil && other.GroupInterval == nil {
		return false
	}

	if resource.GroupInterval != nil {
		if *resource.GroupInterval != *other.GroupInterval {
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

	if len(resource.MuteTimeIntervals) != len(other.MuteTimeIntervals) {
		return false
	}

	for i1 := range resource.MuteTimeIntervals {
		if resource.MuteTimeIntervals[i1] != other.MuteTimeIntervals[i1] {
			return false
		}
	}

	if len(resource.ActiveTimeIntervals) != len(other.ActiveTimeIntervals) {
		return false
	}

	for i1 := range resource.ActiveTimeIntervals {
		if resource.ActiveTimeIntervals[i1] != other.ActiveTimeIntervals[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SimplifiedRouting` fields for violations and returns them.
func (resource SimplifiedRouting) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "SimplifiedRouting" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be SimplifiedRouting"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type NotificationSettingsType string

const (
	NotificationSettingsTypeSimplifiedRouting NotificationSettingsType = "SimplifiedRouting"
	NotificationSettingsTypeNamedRoutingTree  NotificationSettingsType = "NamedRoutingTree"
)

type NamedRoutingTree struct {
	Type        NotificationSettingsType `json:"type"`
	RoutingTree string                   `json:"routingTree"`
}

// NewNamedRoutingTree creates a new NamedRoutingTree object.
func NewNamedRoutingTree() *NamedRoutingTree {
	return &NamedRoutingTree{
		Type: NotificationSettingsTypeNamedRoutingTree,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NamedRoutingTree` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NamedRoutingTree) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "routingTree"
	if fields["routingTree"] != nil {
		if string(fields["routingTree"]) != "null" {
			if err := json.Unmarshal(fields["routingTree"], &resource.RoutingTree); err != nil {
				errs = append(errs, cog.MakeBuildErrors("routingTree", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("routingTree", errors.New("required field is null"))...)

		}
		delete(fields, "routingTree")
	} else {
		errs = append(errs, cog.MakeBuildErrors("routingTree", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NamedRoutingTree", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `NamedRoutingTree` objects.
func (resource NamedRoutingTree) Equals(other NamedRoutingTree) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.RoutingTree != other.RoutingTree {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `NamedRoutingTree` fields for violations and returns them.
func (resource NamedRoutingTree) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "NamedRoutingTree" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be NamedRoutingTree"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ExpressionMap map[string]Expression

type Expression struct {
	// The type of query if this is a query expression
	QueryType         *string            `json:"queryType,omitempty"`
	RelativeTimeRange *RelativeTimeRange `json:"relativeTimeRange,omitempty"`
	// The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
	DatasourceUID *DatasourceUID `json:"datasourceUID,omitempty"`
	Model         any            `json:"model"`
	// Used to mark the expression to be used as the final source for the rule evaluation
	// Only one expression in a rule can be marked as the source
	// For AlertRules, this is the expression that will be evaluated against the alerting condition
	// For RecordingRules, this is the expression that will be recorded
	Source *bool `json:"source,omitempty"`
}

// NewExpression creates a new Expression object.
func NewExpression() *Expression {
	return &Expression{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Expression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Expression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "datasourceUID"
	if fields["datasourceUID"] != nil {
		if string(fields["datasourceUID"]) != "null" {
			if err := json.Unmarshal(fields["datasourceUID"], &resource.DatasourceUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasourceUID", err)...)
			}

		}
		delete(fields, "datasourceUID")

	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "source"
	if fields["source"] != nil {
		if string(fields["source"]) != "null" {
			if err := json.Unmarshal(fields["source"], &resource.Source); err != nil {
				errs = append(errs, cog.MakeBuildErrors("source", err)...)
			}

		}
		delete(fields, "source")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Expression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `Expression` objects.
func (resource Expression) Equals(other Expression) bool {
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
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
	if resource.DatasourceUID == nil && other.DatasourceUID != nil || resource.DatasourceUID != nil && other.DatasourceUID == nil {
		return false
	}

	if resource.DatasourceUID != nil {
		if *resource.DatasourceUID != *other.DatasourceUID {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Model, other.Model) {
		return false
	}
	if resource.Source == nil && other.Source != nil || resource.Source != nil && other.Source == nil {
		return false
	}

	if resource.Source != nil {
		if *resource.Source != *other.Source {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Expression` fields for violations and returns them.
func (resource Expression) Validate() error {
	var errs cog.BuildErrors
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

type RelativeTimeRange struct {
	From PromDurationWMillis `json:"from"`
	To   PromDurationWMillis `json:"to"`
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
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RelativeTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RelativeTimeRange` objects.
func (resource RelativeTimeRange) Equals(other RelativeTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RelativeTimeRange` fields for violations and returns them.
func (resource RelativeTimeRange) Validate() error {
	return nil
}

type PromDurationWMillis string

type DatasourceUID string

type PanelRef struct {
	DashboardUID string `json:"dashboardUID"`
	PanelID      int64  `json:"panelID"`
}

// NewPanelRef creates a new PanelRef object.
func NewPanelRef() *PanelRef {
	return &PanelRef{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelRef) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "dashboardUID"
	if fields["dashboardUID"] != nil {
		if string(fields["dashboardUID"]) != "null" {
			if err := json.Unmarshal(fields["dashboardUID"], &resource.DashboardUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dashboardUID", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dashboardUID", errors.New("required field is null"))...)

		}
		delete(fields, "dashboardUID")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dashboardUID", errors.New("required field is missing from input"))...)
	}
	// Field "panelID"
	if fields["panelID"] != nil {
		if string(fields["panelID"]) != "null" {
			if err := json.Unmarshal(fields["panelID"], &resource.PanelID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("panelID", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("panelID", errors.New("required field is null"))...)

		}
		delete(fields, "panelID")
	} else {
		errs = append(errs, cog.MakeBuildErrors("panelID", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PanelRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `PanelRef` objects.
func (resource PanelRef) Equals(other PanelRef) bool {
	if resource.DashboardUID != other.DashboardUID {
		return false
	}
	if resource.PanelID != other.PanelID {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelRef` fields for violations and returns them.
func (resource PanelRef) Validate() error {
	var errs cog.BuildErrors
	if !(len([]rune(resource.DashboardUID)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"dashboardUID",
			errors.New("must be >= 1"),
		)...)
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(resource.DashboardUID) {
		errs = append(errs, cog.MakeBuildErrors(
			"dashboardUID",
			errors.New("must match regex ^[a-zA-Z0-9_-]+$"),
		)...)
	}
	if !(resource.PanelID > 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"panelID",
			errors.New("must be > 0"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AlertRuleHealth string

const (
	AlertRuleHealthUnknown AlertRuleHealth = "Unknown"
	AlertRuleHealthOK      AlertRuleHealth = "OK"
	AlertRuleHealthPaused  AlertRuleHealth = "Paused"
	AlertRuleHealthError   AlertRuleHealth = "Error"
	AlertRuleHealthNoData  AlertRuleHealth = "NoData"
)

type AlertRuleState string

const (
	AlertRuleStateInactive   AlertRuleState = "Inactive"
	AlertRuleStateHealthy    AlertRuleState = "Healthy"
	AlertRuleStateFiring     AlertRuleState = "Firing"
	AlertRuleStatePending    AlertRuleState = "Pending"
	AlertRuleStateRecovering AlertRuleState = "Recovering"
)

type AlertRuleStateReason string

const (
	AlertRuleStateReasonEvaluated AlertRuleStateReason = "Evaluated"
	AlertRuleStateReasonKeepLast  AlertRuleStateReason = "KeepLast"
)

// Count of alert instances per state. error also counts instances whose evaluation
// errored but were mapped to another state via execErrState, so it can overlap.
type AlertRuleInstanceTotals struct {
	Healthy    *int64 `json:"healthy,omitempty"`
	Firing     *int64 `json:"firing,omitempty"`
	Pending    *int64 `json:"pending,omitempty"`
	Recovering *int64 `json:"recovering,omitempty"`
	Nodata     *int64 `json:"nodata,omitempty"`
	Error      *int64 `json:"error,omitempty"`
}

// NewAlertRuleInstanceTotals creates a new AlertRuleInstanceTotals object.
func NewAlertRuleInstanceTotals() *AlertRuleInstanceTotals {
	return &AlertRuleInstanceTotals{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleInstanceTotals` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AlertRuleInstanceTotals) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "healthy"
	if fields["healthy"] != nil {
		if string(fields["healthy"]) != "null" {
			if err := json.Unmarshal(fields["healthy"], &resource.Healthy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("healthy", err)...)
			}

		}
		delete(fields, "healthy")

	}
	// Field "firing"
	if fields["firing"] != nil {
		if string(fields["firing"]) != "null" {
			if err := json.Unmarshal(fields["firing"], &resource.Firing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("firing", err)...)
			}

		}
		delete(fields, "firing")

	}
	// Field "pending"
	if fields["pending"] != nil {
		if string(fields["pending"]) != "null" {
			if err := json.Unmarshal(fields["pending"], &resource.Pending); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pending", err)...)
			}

		}
		delete(fields, "pending")

	}
	// Field "recovering"
	if fields["recovering"] != nil {
		if string(fields["recovering"]) != "null" {
			if err := json.Unmarshal(fields["recovering"], &resource.Recovering); err != nil {
				errs = append(errs, cog.MakeBuildErrors("recovering", err)...)
			}

		}
		delete(fields, "recovering")

	}
	// Field "nodata"
	if fields["nodata"] != nil {
		if string(fields["nodata"]) != "null" {
			if err := json.Unmarshal(fields["nodata"], &resource.Nodata); err != nil {
				errs = append(errs, cog.MakeBuildErrors("nodata", err)...)
			}

		}
		delete(fields, "nodata")

	}
	// Field "error"
	if fields["error"] != nil {
		if string(fields["error"]) != "null" {
			if err := json.Unmarshal(fields["error"], &resource.Error); err != nil {
				errs = append(errs, cog.MakeBuildErrors("error", err)...)
			}

		}
		delete(fields, "error")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AlertRuleInstanceTotals", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `AlertRuleInstanceTotals` objects.
func (resource AlertRuleInstanceTotals) Equals(other AlertRuleInstanceTotals) bool {
	if resource.Healthy == nil && other.Healthy != nil || resource.Healthy != nil && other.Healthy == nil {
		return false
	}

	if resource.Healthy != nil {
		if *resource.Healthy != *other.Healthy {
			return false
		}
	}
	if resource.Firing == nil && other.Firing != nil || resource.Firing != nil && other.Firing == nil {
		return false
	}

	if resource.Firing != nil {
		if *resource.Firing != *other.Firing {
			return false
		}
	}
	if resource.Pending == nil && other.Pending != nil || resource.Pending != nil && other.Pending == nil {
		return false
	}

	if resource.Pending != nil {
		if *resource.Pending != *other.Pending {
			return false
		}
	}
	if resource.Recovering == nil && other.Recovering != nil || resource.Recovering != nil && other.Recovering == nil {
		return false
	}

	if resource.Recovering != nil {
		if *resource.Recovering != *other.Recovering {
			return false
		}
	}
	if resource.Nodata == nil && other.Nodata != nil || resource.Nodata != nil && other.Nodata == nil {
		return false
	}

	if resource.Nodata != nil {
		if *resource.Nodata != *other.Nodata {
			return false
		}
	}
	if resource.Error == nil && other.Error != nil || resource.Error != nil && other.Error == nil {
		return false
	}

	if resource.Error != nil {
		if *resource.Error != *other.Error {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AlertRuleInstanceTotals` fields for violations and returns them.
func (resource AlertRuleInstanceTotals) Validate() error {
	return nil
}

type RecordingRuleSpec struct {
	Title               string                    `json:"title"`
	Paused              *bool                     `json:"paused,omitempty"`
	Trigger             IntervalTrigger           `json:"trigger"`
	Labels              map[string]TemplateString `json:"labels,omitempty"`
	Metric              MetricName                `json:"metric"`
	Expressions         ExpressionMap             `json:"expressions"`
	TargetDatasourceUID DatasourceUID             `json:"targetDatasourceUID"`
}

// NewRecordingRuleSpec creates a new RecordingRuleSpec object.
func NewRecordingRuleSpec() *RecordingRuleSpec {
	return &RecordingRuleSpec{
		Trigger: *NewIntervalTrigger(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RecordingRuleSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RecordingRuleSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "paused"
	if fields["paused"] != nil {
		if string(fields["paused"]) != "null" {
			if err := json.Unmarshal(fields["paused"], &resource.Paused); err != nil {
				errs = append(errs, cog.MakeBuildErrors("paused", err)...)
			}

		}
		delete(fields, "paused")

	}
	// Field "trigger"
	if fields["trigger"] != nil {
		if string(fields["trigger"]) != "null" {

			resource.Trigger = IntervalTrigger{}
			if err := resource.Trigger.UnmarshalJSONStrict(fields["trigger"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is null"))...)

		}
		delete(fields, "trigger")
	} else {
		errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is missing from input"))...)
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
	// Field "metric"
	if fields["metric"] != nil {
		if string(fields["metric"]) != "null" {
			if err := json.Unmarshal(fields["metric"], &resource.Metric); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metric", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metric", errors.New("required field is null"))...)

		}
		delete(fields, "metric")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metric", errors.New("required field is missing from input"))...)
	}
	// Field "expressions"
	if fields["expressions"] != nil {
		if string(fields["expressions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["expressions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]Expression, len(partialMap))
			for key1 := range partialMap {
				var result1 Expression

				result1 = Expression{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("expressions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Expressions = ExpressionMap(parsedMap1)
		} else {
			errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is null"))...)

		}
		delete(fields, "expressions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is missing from input"))...)
	}
	// Field "targetDatasourceUID"
	if fields["targetDatasourceUID"] != nil {
		if string(fields["targetDatasourceUID"]) != "null" {
			if err := json.Unmarshal(fields["targetDatasourceUID"], &resource.TargetDatasourceUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("targetDatasourceUID", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("targetDatasourceUID", errors.New("required field is null"))...)

		}
		delete(fields, "targetDatasourceUID")
	} else {
		errs = append(errs, cog.MakeBuildErrors("targetDatasourceUID", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RecordingRuleSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RecordingRuleSpec` objects.
func (resource RecordingRuleSpec) Equals(other RecordingRuleSpec) bool {
	if resource.Title != other.Title {
		return false
	}
	if resource.Paused == nil && other.Paused != nil || resource.Paused != nil && other.Paused == nil {
		return false
	}

	if resource.Paused != nil {
		if *resource.Paused != *other.Paused {
			return false
		}
	}
	if !resource.Trigger.Equals(other.Trigger) {
		return false
	}

	if len(resource.Labels) != len(other.Labels) {
		return false
	}

	for key1 := range resource.Labels {
		if resource.Labels[key1] != other.Labels[key1] {
			return false
		}
	}
	if resource.Metric != other.Metric {
		return false
	}

	if len(resource.Expressions) != len(other.Expressions) {
		return false
	}

	for key1 := range resource.Expressions {
		if !resource.Expressions[key1].Equals(other.Expressions[key1]) {
			return false
		}
	}
	if resource.TargetDatasourceUID != other.TargetDatasourceUID {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RecordingRuleSpec` fields for violations and returns them.
func (resource RecordingRuleSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Trigger.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MetricName string

type RecordingRuleHealth string

const (
	RecordingRuleHealthUnknown   RecordingRuleHealth = "Unknown"
	RecordingRuleHealthRecording RecordingRuleHealth = "Recording"
	RecordingRuleHealthPaused    RecordingRuleHealth = "Paused"
	RecordingRuleHealthError     RecordingRuleHealth = "Error"
	RecordingRuleHealthNoData    RecordingRuleHealth = "NoData"
)

type RuleSequenceSpec struct {
	Trigger        IntervalTrigger `json:"trigger"`
	RecordingRules []RuleRef       `json:"recordingRules"`
	AlertRules     []RuleRef       `json:"alertRules,omitempty"`
}

// NewRuleSequenceSpec creates a new RuleSequenceSpec object.
func NewRuleSequenceSpec() *RuleSequenceSpec {
	return &RuleSequenceSpec{
		Trigger:        *NewIntervalTrigger(),
		RecordingRules: []RuleRef{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleSequenceSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RuleSequenceSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "trigger"
	if fields["trigger"] != nil {
		if string(fields["trigger"]) != "null" {

			resource.Trigger = IntervalTrigger{}
			if err := resource.Trigger.UnmarshalJSONStrict(fields["trigger"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is null"))...)

		}
		delete(fields, "trigger")
	} else {
		errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is missing from input"))...)
	}
	// Field "recordingRules"
	if fields["recordingRules"] != nil {
		if string(fields["recordingRules"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["recordingRules"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 RuleRef

				result1 = RuleRef{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("recordingRules["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.RecordingRules = append(resource.RecordingRules, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("recordingRules", errors.New("required field is null"))...)

		}
		delete(fields, "recordingRules")
	} else {
		errs = append(errs, cog.MakeBuildErrors("recordingRules", errors.New("required field is missing from input"))...)
	}
	// Field "alertRules"
	if fields["alertRules"] != nil {
		if string(fields["alertRules"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["alertRules"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 RuleRef

				result1 = RuleRef{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("alertRules["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.AlertRules = append(resource.AlertRules, result1)
			}

		}
		delete(fields, "alertRules")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RuleSequenceSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RuleSequenceSpec` objects.
func (resource RuleSequenceSpec) Equals(other RuleSequenceSpec) bool {
	if !resource.Trigger.Equals(other.Trigger) {
		return false
	}

	if len(resource.RecordingRules) != len(other.RecordingRules) {
		return false
	}

	for i1 := range resource.RecordingRules {
		if !resource.RecordingRules[i1].Equals(other.RecordingRules[i1]) {
			return false
		}
	}

	if len(resource.AlertRules) != len(other.AlertRules) {
		return false
	}

	for i1 := range resource.AlertRules {
		if !resource.AlertRules[i1].Equals(other.AlertRules[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RuleSequenceSpec` fields for violations and returns them.
func (resource RuleSequenceSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Trigger.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
	}

	for i1 := range resource.RecordingRules {
		if err := resource.RecordingRules[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("recordingRules["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.AlertRules {
		if err := resource.AlertRules[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("alertRules["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RuleRef struct {
	// name is the metadata.name of an AlertRule or RecordingRule resource.
	Name RuleUID `json:"name"`
}

// NewRuleRef creates a new RuleRef object.
func NewRuleRef() *RuleRef {
	return &RuleRef{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RuleRef) UnmarshalJSONStrict(raw []byte) error {
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
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RuleRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RuleRef` objects.
func (resource RuleRef) Equals(other RuleRef) bool {
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RuleRef` fields for violations and returns them.
func (resource RuleRef) Validate() error {
	return nil
}

type RuleUID string

type RuleSpec struct {
	Title       string                    `json:"title"`
	Paused      *bool                     `json:"paused,omitempty"`
	Trigger     IntervalTrigger           `json:"trigger"`
	Labels      map[string]TemplateString `json:"labels,omitempty"`
	Expressions ExpressionMap             `json:"expressions"`
}

// NewRuleSpec creates a new RuleSpec object.
func NewRuleSpec() *RuleSpec {
	return &RuleSpec{
		Trigger: *NewIntervalTrigger(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RuleSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "paused"
	if fields["paused"] != nil {
		if string(fields["paused"]) != "null" {
			if err := json.Unmarshal(fields["paused"], &resource.Paused); err != nil {
				errs = append(errs, cog.MakeBuildErrors("paused", err)...)
			}

		}
		delete(fields, "paused")

	}
	// Field "trigger"
	if fields["trigger"] != nil {
		if string(fields["trigger"]) != "null" {

			resource.Trigger = IntervalTrigger{}
			if err := resource.Trigger.UnmarshalJSONStrict(fields["trigger"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is null"))...)

		}
		delete(fields, "trigger")
	} else {
		errs = append(errs, cog.MakeBuildErrors("trigger", errors.New("required field is missing from input"))...)
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
	// Field "expressions"
	if fields["expressions"] != nil {
		if string(fields["expressions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["expressions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]Expression, len(partialMap))
			for key1 := range partialMap {
				var result1 Expression

				result1 = Expression{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("expressions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Expressions = ExpressionMap(parsedMap1)
		} else {
			errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is null"))...)

		}
		delete(fields, "expressions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RuleSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RuleSpec` objects.
func (resource RuleSpec) Equals(other RuleSpec) bool {
	if resource.Title != other.Title {
		return false
	}
	if resource.Paused == nil && other.Paused != nil || resource.Paused != nil && other.Paused == nil {
		return false
	}

	if resource.Paused != nil {
		if *resource.Paused != *other.Paused {
			return false
		}
	}
	if !resource.Trigger.Equals(other.Trigger) {
		return false
	}

	if len(resource.Labels) != len(other.Labels) {
		return false
	}

	for key1 := range resource.Labels {
		if resource.Labels[key1] != other.Labels[key1] {
			return false
		}
	}

	if len(resource.Expressions) != len(other.Expressions) {
		return false
	}

	for key1 := range resource.Expressions {
		if !resource.Expressions[key1].Equals(other.Expressions[key1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RuleSpec` fields for violations and returns them.
func (resource RuleSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Trigger.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("trigger", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AlertRuleStatus struct {
	Health             *AlertRuleHealth      `json:"health,omitempty"`
	State              *AlertRuleState       `json:"state,omitempty"`
	StateReason        *AlertRuleStateReason `json:"stateReason,omitempty"`
	LastEvaluationTime *time.Time            `json:"lastEvaluationTime,omitempty"`
	// duration of the last evaluation in seconds
	EvaluationDuration *float64                 `json:"evaluationDuration,omitempty"`
	LastError          *string                  `json:"lastError,omitempty"`
	Totals             *AlertRuleInstanceTotals `json:"totals,omitempty"`
}

// NewAlertRuleStatus creates a new AlertRuleStatus object.
func NewAlertRuleStatus() *AlertRuleStatus {
	return &AlertRuleStatus{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleStatus` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AlertRuleStatus) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "health"
	if fields["health"] != nil {
		if string(fields["health"]) != "null" {
			if err := json.Unmarshal(fields["health"], &resource.Health); err != nil {
				errs = append(errs, cog.MakeBuildErrors("health", err)...)
			}

		}
		delete(fields, "health")

	}
	// Field "state"
	if fields["state"] != nil {
		if string(fields["state"]) != "null" {
			if err := json.Unmarshal(fields["state"], &resource.State); err != nil {
				errs = append(errs, cog.MakeBuildErrors("state", err)...)
			}

		}
		delete(fields, "state")

	}
	// Field "stateReason"
	if fields["stateReason"] != nil {
		if string(fields["stateReason"]) != "null" {
			if err := json.Unmarshal(fields["stateReason"], &resource.StateReason); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stateReason", err)...)
			}

		}
		delete(fields, "stateReason")

	}
	// Field "lastEvaluationTime"
	if fields["lastEvaluationTime"] != nil {
		if string(fields["lastEvaluationTime"]) != "null" {
			if err := json.Unmarshal(fields["lastEvaluationTime"], &resource.LastEvaluationTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lastEvaluationTime", err)...)
			}

		}
		delete(fields, "lastEvaluationTime")

	}
	// Field "evaluationDuration"
	if fields["evaluationDuration"] != nil {
		if string(fields["evaluationDuration"]) != "null" {
			if err := json.Unmarshal(fields["evaluationDuration"], &resource.EvaluationDuration); err != nil {
				errs = append(errs, cog.MakeBuildErrors("evaluationDuration", err)...)
			}

		}
		delete(fields, "evaluationDuration")

	}
	// Field "lastError"
	if fields["lastError"] != nil {
		if string(fields["lastError"]) != "null" {
			if err := json.Unmarshal(fields["lastError"], &resource.LastError); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lastError", err)...)
			}

		}
		delete(fields, "lastError")

	}
	// Field "totals"
	if fields["totals"] != nil {
		if string(fields["totals"]) != "null" {

			resource.Totals = &AlertRuleInstanceTotals{}
			if err := resource.Totals.UnmarshalJSONStrict(fields["totals"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("totals", err)...)
			}

		}
		delete(fields, "totals")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AlertRuleStatus", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `AlertRuleStatus` objects.
func (resource AlertRuleStatus) Equals(other AlertRuleStatus) bool {
	if resource.Health == nil && other.Health != nil || resource.Health != nil && other.Health == nil {
		return false
	}

	if resource.Health != nil {
		if *resource.Health != *other.Health {
			return false
		}
	}
	if resource.State == nil && other.State != nil || resource.State != nil && other.State == nil {
		return false
	}

	if resource.State != nil {
		if *resource.State != *other.State {
			return false
		}
	}
	if resource.StateReason == nil && other.StateReason != nil || resource.StateReason != nil && other.StateReason == nil {
		return false
	}

	if resource.StateReason != nil {
		if *resource.StateReason != *other.StateReason {
			return false
		}
	}
	if resource.LastEvaluationTime == nil && other.LastEvaluationTime != nil || resource.LastEvaluationTime != nil && other.LastEvaluationTime == nil {
		return false
	}

	if resource.LastEvaluationTime != nil {
		if *resource.LastEvaluationTime != *other.LastEvaluationTime {
			return false
		}
	}
	if resource.EvaluationDuration == nil && other.EvaluationDuration != nil || resource.EvaluationDuration != nil && other.EvaluationDuration == nil {
		return false
	}

	if resource.EvaluationDuration != nil {
		if *resource.EvaluationDuration != *other.EvaluationDuration {
			return false
		}
	}
	if resource.LastError == nil && other.LastError != nil || resource.LastError != nil && other.LastError == nil {
		return false
	}

	if resource.LastError != nil {
		if *resource.LastError != *other.LastError {
			return false
		}
	}
	if resource.Totals == nil && other.Totals != nil || resource.Totals != nil && other.Totals == nil {
		return false
	}

	if resource.Totals != nil {
		if !resource.Totals.Equals(*other.Totals) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AlertRuleStatus` fields for violations and returns them.
func (resource AlertRuleStatus) Validate() error {
	var errs cog.BuildErrors
	if resource.Totals != nil {
		if err := resource.Totals.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("totals", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RecordingRuleStatus struct {
	Health             *RecordingRuleHealth `json:"health,omitempty"`
	LastEvaluationTime *time.Time           `json:"lastEvaluationTime,omitempty"`
	// duration of the last evaluation in seconds
	EvaluationDuration *float64 `json:"evaluationDuration,omitempty"`
	LastError          *string  `json:"lastError,omitempty"`
}

// NewRecordingRuleStatus creates a new RecordingRuleStatus object.
func NewRecordingRuleStatus() *RecordingRuleStatus {
	return &RecordingRuleStatus{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RecordingRuleStatus` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RecordingRuleStatus) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "health"
	if fields["health"] != nil {
		if string(fields["health"]) != "null" {
			if err := json.Unmarshal(fields["health"], &resource.Health); err != nil {
				errs = append(errs, cog.MakeBuildErrors("health", err)...)
			}

		}
		delete(fields, "health")

	}
	// Field "lastEvaluationTime"
	if fields["lastEvaluationTime"] != nil {
		if string(fields["lastEvaluationTime"]) != "null" {
			if err := json.Unmarshal(fields["lastEvaluationTime"], &resource.LastEvaluationTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lastEvaluationTime", err)...)
			}

		}
		delete(fields, "lastEvaluationTime")

	}
	// Field "evaluationDuration"
	if fields["evaluationDuration"] != nil {
		if string(fields["evaluationDuration"]) != "null" {
			if err := json.Unmarshal(fields["evaluationDuration"], &resource.EvaluationDuration); err != nil {
				errs = append(errs, cog.MakeBuildErrors("evaluationDuration", err)...)
			}

		}
		delete(fields, "evaluationDuration")

	}
	// Field "lastError"
	if fields["lastError"] != nil {
		if string(fields["lastError"]) != "null" {
			if err := json.Unmarshal(fields["lastError"], &resource.LastError); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lastError", err)...)
			}

		}
		delete(fields, "lastError")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RecordingRuleStatus", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RecordingRuleStatus` objects.
func (resource RecordingRuleStatus) Equals(other RecordingRuleStatus) bool {
	if resource.Health == nil && other.Health != nil || resource.Health != nil && other.Health == nil {
		return false
	}

	if resource.Health != nil {
		if *resource.Health != *other.Health {
			return false
		}
	}
	if resource.LastEvaluationTime == nil && other.LastEvaluationTime != nil || resource.LastEvaluationTime != nil && other.LastEvaluationTime == nil {
		return false
	}

	if resource.LastEvaluationTime != nil {
		if *resource.LastEvaluationTime != *other.LastEvaluationTime {
			return false
		}
	}
	if resource.EvaluationDuration == nil && other.EvaluationDuration != nil || resource.EvaluationDuration != nil && other.EvaluationDuration == nil {
		return false
	}

	if resource.EvaluationDuration != nil {
		if *resource.EvaluationDuration != *other.EvaluationDuration {
			return false
		}
	}
	if resource.LastError == nil && other.LastError != nil || resource.LastError != nil && other.LastError == nil {
		return false
	}

	if resource.LastError != nil {
		if *resource.LastError != *other.LastError {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RecordingRuleStatus` fields for violations and returns them.
func (resource RecordingRuleStatus) Validate() error {
	return nil
}

const ApiVersion = "rules.alerting.grafana.app/v1beta1"

const AlertRuleKind = "AlertRule"

const RecordingRuleKind = "RecordingRule"

const RuleSequenceKind = "RuleSequence"

type SimplifiedRoutingOrNamedRoutingTree struct {
	SimplifiedRouting *SimplifiedRouting `json:"SimplifiedRouting,omitempty"`
	NamedRoutingTree  *NamedRoutingTree  `json:"NamedRoutingTree,omitempty"`
}

// NewSimplifiedRoutingOrNamedRoutingTree creates a new SimplifiedRoutingOrNamedRoutingTree object.
func NewSimplifiedRoutingOrNamedRoutingTree() *SimplifiedRoutingOrNamedRoutingTree {
	return &SimplifiedRoutingOrNamedRoutingTree{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `SimplifiedRoutingOrNamedRoutingTree` as JSON.
func (resource SimplifiedRoutingOrNamedRoutingTree) MarshalJSON() ([]byte, error) {
	if resource.SimplifiedRouting != nil {
		return json.Marshal(resource.SimplifiedRouting)
	}
	if resource.NamedRoutingTree != nil {
		return json.Marshal(resource.NamedRoutingTree)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `SimplifiedRoutingOrNamedRoutingTree` from JSON.
func (resource *SimplifiedRoutingOrNamedRoutingTree) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return nil
	}

	switch discriminator {
	case "NamedRoutingTree":
		var namedRoutingTree NamedRoutingTree
		if err := json.Unmarshal(raw, &namedRoutingTree); err != nil {
			return err
		}

		resource.NamedRoutingTree = &namedRoutingTree
		return nil
	case "SimplifiedRouting":
		var simplifiedRouting SimplifiedRouting
		if err := json.Unmarshal(raw, &simplifiedRouting); err != nil {
			return err
		}

		resource.SimplifiedRouting = &simplifiedRouting
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimplifiedRoutingOrNamedRoutingTree` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SimplifiedRoutingOrNamedRoutingTree) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "NamedRoutingTree":
		namedRoutingTree := &NamedRoutingTree{}
		if err := namedRoutingTree.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.NamedRoutingTree = namedRoutingTree
		return nil
	case "SimplifiedRouting":
		simplifiedRouting := &SimplifiedRouting{}
		if err := simplifiedRouting.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SimplifiedRouting = simplifiedRouting
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `SimplifiedRoutingOrNamedRoutingTree` objects.
func (resource SimplifiedRoutingOrNamedRoutingTree) Equals(other SimplifiedRoutingOrNamedRoutingTree) bool {
	if resource.SimplifiedRouting == nil && other.SimplifiedRouting != nil || resource.SimplifiedRouting != nil && other.SimplifiedRouting == nil {
		return false
	}

	if resource.SimplifiedRouting != nil {
		if !resource.SimplifiedRouting.Equals(*other.SimplifiedRouting) {
			return false
		}
	}
	if resource.NamedRoutingTree == nil && other.NamedRoutingTree != nil || resource.NamedRoutingTree != nil && other.NamedRoutingTree == nil {
		return false
	}

	if resource.NamedRoutingTree != nil {
		if !resource.NamedRoutingTree.Equals(*other.NamedRoutingTree) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SimplifiedRoutingOrNamedRoutingTree` fields for violations and returns them.
func (resource SimplifiedRoutingOrNamedRoutingTree) Validate() error {
	var errs cog.BuildErrors
	if resource.SimplifiedRouting != nil {
		if err := resource.SimplifiedRouting.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SimplifiedRouting", err)...)
		}
	}
	if resource.NamedRoutingTree != nil {
		if err := resource.NamedRoutingTree.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("NamedRoutingTree", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
