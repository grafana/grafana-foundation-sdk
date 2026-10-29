// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

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

type RuleSequenceSpec struct {
	Trigger        IntervalTrigger `json:"trigger"`
	RecordingRules []RuleRef       `json:"recordingRules"`
	AlertingRules  []RuleRef       `json:"alertingRules,omitempty"`
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
	// Field "alertingRules"
	if fields["alertingRules"] != nil {
		if string(fields["alertingRules"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["alertingRules"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 RuleRef

				result1 = RuleRef{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("alertingRules["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.AlertingRules = append(resource.AlertingRules, result1)
			}

		}
		delete(fields, "alertingRules")

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

	if len(resource.AlertingRules) != len(other.AlertingRules) {
		return false
	}

	for i1 := range resource.AlertingRules {
		if !resource.AlertingRules[i1].Equals(other.AlertingRules[i1]) {
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

	for i1 := range resource.AlertingRules {
		if err := resource.AlertingRules[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("alertingRules["+strconv.Itoa(i1)+"]", err)...)
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

// #SearchResultResource is the full identity of a hit.
type SearchResultResource struct {
	Group    string `json:"group"`
	Resource string `json:"resource"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

// NewSearchResultResource creates a new SearchResultResource object.
func NewSearchResultResource() *SearchResultResource {
	return &SearchResultResource{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultResource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SearchResultResource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
	}
	// Field "resource"
	if fields["resource"] != nil {
		if string(fields["resource"]) != "null" {
			if err := json.Unmarshal(fields["resource"], &resource.Resource); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resource", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is null"))...)

		}
		delete(fields, "resource")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is missing from input"))...)
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
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
		errs = append(errs, cog.MakeBuildErrors("SearchResultResource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `SearchResultResource` objects.
func (resource SearchResultResource) Equals(other SearchResultResource) bool {
	if resource.Group != other.Group {
		return false
	}
	if resource.Resource != other.Resource {
		return false
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SearchResultResource` fields for violations and returns them.
func (resource SearchResultResource) Validate() error {
	return nil
}

// #RuleSearchHitFields is the per-kind field payload returned on each hit.
// It carries the union of alert- and recording-rule search fields; only the
// fields relevant to a hit's kind are populated. This maps to the kind's
// declared searchFields.
type RuleSearchHitFields struct {
	Title          *string           `json:"title,omitempty"`
	Folder         *string           `json:"folder,omitempty"`
	Type           *string           `json:"type,omitempty"`
	Interval       *string           `json:"interval,omitempty"`
	Paused         *bool             `json:"paused,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	DatasourceUIDs []string          `json:"datasourceUIDs,omitempty"`
	// Alert-rule fields.
	Annotations      map[string]string `json:"annotations,omitempty"`
	For              *string           `json:"for,omitempty"`
	KeepFiringFor    *string           `json:"keepFiringFor,omitempty"`
	DashboardUID     *string           `json:"dashboardUID,omitempty"`
	PanelID          *int64            `json:"panelID,omitempty"`
	Receiver         *string           `json:"receiver,omitempty"`
	NotificationType *string           `json:"notificationType,omitempty"`
	RoutingTree      *string           `json:"routingTree,omitempty"`
	// Recording-rule fields.
	Metric              *string `json:"metric,omitempty"`
	TargetDatasourceUID *string `json:"targetDatasourceUID,omitempty"`
}

// NewRuleSearchHitFields creates a new RuleSearchHitFields object.
func NewRuleSearchHitFields() *RuleSearchHitFields {
	return &RuleSearchHitFields{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleSearchHitFields` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RuleSearchHitFields) UnmarshalJSONStrict(raw []byte) error {
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

		}
		delete(fields, "title")

	}
	// Field "folder"
	if fields["folder"] != nil {
		if string(fields["folder"]) != "null" {
			if err := json.Unmarshal(fields["folder"], &resource.Folder); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folder", err)...)
			}

		}
		delete(fields, "folder")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

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
	// Field "paused"
	if fields["paused"] != nil {
		if string(fields["paused"]) != "null" {
			if err := json.Unmarshal(fields["paused"], &resource.Paused); err != nil {
				errs = append(errs, cog.MakeBuildErrors("paused", err)...)
			}

		}
		delete(fields, "paused")

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
	// Field "datasourceUIDs"
	if fields["datasourceUIDs"] != nil {
		if string(fields["datasourceUIDs"]) != "null" {

			if err := json.Unmarshal(fields["datasourceUIDs"], &resource.DatasourceUIDs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasourceUIDs", err)...)
			}

		}
		delete(fields, "datasourceUIDs")

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
	// Field "dashboardUID"
	if fields["dashboardUID"] != nil {
		if string(fields["dashboardUID"]) != "null" {
			if err := json.Unmarshal(fields["dashboardUID"], &resource.DashboardUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dashboardUID", err)...)
			}

		}
		delete(fields, "dashboardUID")

	}
	// Field "panelID"
	if fields["panelID"] != nil {
		if string(fields["panelID"]) != "null" {
			if err := json.Unmarshal(fields["panelID"], &resource.PanelID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("panelID", err)...)
			}

		}
		delete(fields, "panelID")

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
	// Field "notificationType"
	if fields["notificationType"] != nil {
		if string(fields["notificationType"]) != "null" {
			if err := json.Unmarshal(fields["notificationType"], &resource.NotificationType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("notificationType", err)...)
			}

		}
		delete(fields, "notificationType")

	}
	// Field "routingTree"
	if fields["routingTree"] != nil {
		if string(fields["routingTree"]) != "null" {
			if err := json.Unmarshal(fields["routingTree"], &resource.RoutingTree); err != nil {
				errs = append(errs, cog.MakeBuildErrors("routingTree", err)...)
			}

		}
		delete(fields, "routingTree")

	}
	// Field "metric"
	if fields["metric"] != nil {
		if string(fields["metric"]) != "null" {
			if err := json.Unmarshal(fields["metric"], &resource.Metric); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metric", err)...)
			}

		}
		delete(fields, "metric")

	}
	// Field "targetDatasourceUID"
	if fields["targetDatasourceUID"] != nil {
		if string(fields["targetDatasourceUID"]) != "null" {
			if err := json.Unmarshal(fields["targetDatasourceUID"], &resource.TargetDatasourceUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("targetDatasourceUID", err)...)
			}

		}
		delete(fields, "targetDatasourceUID")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RuleSearchHitFields", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `RuleSearchHitFields` objects.
func (resource RuleSearchHitFields) Equals(other RuleSearchHitFields) bool {
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}
	if resource.Folder == nil && other.Folder != nil || resource.Folder != nil && other.Folder == nil {
		return false
	}

	if resource.Folder != nil {
		if *resource.Folder != *other.Folder {
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
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.Paused == nil && other.Paused != nil || resource.Paused != nil && other.Paused == nil {
		return false
	}

	if resource.Paused != nil {
		if *resource.Paused != *other.Paused {
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

	if len(resource.DatasourceUIDs) != len(other.DatasourceUIDs) {
		return false
	}

	for i1 := range resource.DatasourceUIDs {
		if resource.DatasourceUIDs[i1] != other.DatasourceUIDs[i1] {
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
	if resource.DashboardUID == nil && other.DashboardUID != nil || resource.DashboardUID != nil && other.DashboardUID == nil {
		return false
	}

	if resource.DashboardUID != nil {
		if *resource.DashboardUID != *other.DashboardUID {
			return false
		}
	}
	if resource.PanelID == nil && other.PanelID != nil || resource.PanelID != nil && other.PanelID == nil {
		return false
	}

	if resource.PanelID != nil {
		if *resource.PanelID != *other.PanelID {
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
	if resource.NotificationType == nil && other.NotificationType != nil || resource.NotificationType != nil && other.NotificationType == nil {
		return false
	}

	if resource.NotificationType != nil {
		if *resource.NotificationType != *other.NotificationType {
			return false
		}
	}
	if resource.RoutingTree == nil && other.RoutingTree != nil || resource.RoutingTree != nil && other.RoutingTree == nil {
		return false
	}

	if resource.RoutingTree != nil {
		if *resource.RoutingTree != *other.RoutingTree {
			return false
		}
	}
	if resource.Metric == nil && other.Metric != nil || resource.Metric != nil && other.Metric == nil {
		return false
	}

	if resource.Metric != nil {
		if *resource.Metric != *other.Metric {
			return false
		}
	}
	if resource.TargetDatasourceUID == nil && other.TargetDatasourceUID != nil || resource.TargetDatasourceUID != nil && other.TargetDatasourceUID == nil {
		return false
	}

	if resource.TargetDatasourceUID != nil {
		if *resource.TargetDatasourceUID != *other.TargetDatasourceUID {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RuleSearchHitFields` fields for violations and returns them.
func (resource RuleSearchHitFields) Validate() error {
	return nil
}

// #SearchResultHit is a single match: its identity, an optional relevance
// score (present only when the query included free text), and the requested
// fields.
type SearchResultHit struct {
	Resource SearchResultResource `json:"resource"`
	Score    *float64             `json:"score,omitempty"`
	Fields   RuleSearchHitFields  `json:"fields"`
}

// NewSearchResultHit creates a new SearchResultHit object.
func NewSearchResultHit() *SearchResultHit {
	return &SearchResultHit{
		Resource: *NewSearchResultResource(),
		Fields:   *NewRuleSearchHitFields(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultHit` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SearchResultHit) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "resource"
	if fields["resource"] != nil {
		if string(fields["resource"]) != "null" {

			resource.Resource = SearchResultResource{}
			if err := resource.Resource.UnmarshalJSONStrict(fields["resource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resource", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is null"))...)

		}
		delete(fields, "resource")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is missing from input"))...)
	}
	// Field "score"
	if fields["score"] != nil {
		if string(fields["score"]) != "null" {
			if err := json.Unmarshal(fields["score"], &resource.Score); err != nil {
				errs = append(errs, cog.MakeBuildErrors("score", err)...)
			}

		}
		delete(fields, "score")

	}
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {

			resource.Fields = RuleSearchHitFields{}
			if err := resource.Fields.UnmarshalJSONStrict(fields["fields"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fields", errors.New("required field is null"))...)

		}
		delete(fields, "fields")
	} else {
		errs = append(errs, cog.MakeBuildErrors("fields", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SearchResultHit", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `SearchResultHit` objects.
func (resource SearchResultHit) Equals(other SearchResultHit) bool {
	if !resource.Resource.Equals(other.Resource) {
		return false
	}
	if resource.Score == nil && other.Score != nil || resource.Score != nil && other.Score == nil {
		return false
	}

	if resource.Score != nil {
		if *resource.Score != *other.Score {
			return false
		}
	}
	if !resource.Fields.Equals(other.Fields) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SearchResultHit` fields for violations and returns them.
func (resource SearchResultHit) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Resource.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("resource", err)...)
	}
	if err := resource.Fields.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("fields", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// #TotalHitsRelation says how totalHits relates to the real number of matching
// rules the caller may see: "eq" when it is exact, "lte" when it is an upper
// bound because authorisation was applied after the search ranked its results.
type TotalHitsRelation string

const (
	TotalHitsRelationEq  TotalHitsRelation = "eq"
	TotalHitsRelationLte TotalHitsRelation = "lte"
)

// #SearchResultsMetadata carries the paging token and total authorised match
// count.
type SearchResultsMetadata struct {
	Continue  *string `json:"continue,omitempty"`
	TotalHits *int64  `json:"totalHits,omitempty"`
	// Always read totalHits together with totalHitsRelation.
	TotalHitsRelation *TotalHitsRelation `json:"totalHitsRelation,omitempty"`
}

// NewSearchResultsMetadata creates a new SearchResultsMetadata object.
func NewSearchResultsMetadata() *SearchResultsMetadata {
	return &SearchResultsMetadata{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultsMetadata` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SearchResultsMetadata) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "totalHits"
	if fields["totalHits"] != nil {
		if string(fields["totalHits"]) != "null" {
			if err := json.Unmarshal(fields["totalHits"], &resource.TotalHits); err != nil {
				errs = append(errs, cog.MakeBuildErrors("totalHits", err)...)
			}

		}
		delete(fields, "totalHits")

	}
	// Field "totalHitsRelation"
	if fields["totalHitsRelation"] != nil {
		if string(fields["totalHitsRelation"]) != "null" {
			if err := json.Unmarshal(fields["totalHitsRelation"], &resource.TotalHitsRelation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("totalHitsRelation", err)...)
			}

		}
		delete(fields, "totalHitsRelation")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SearchResultsMetadata", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `SearchResultsMetadata` objects.
func (resource SearchResultsMetadata) Equals(other SearchResultsMetadata) bool {
	if resource.Continue == nil && other.Continue != nil || resource.Continue != nil && other.Continue == nil {
		return false
	}

	if resource.Continue != nil {
		if *resource.Continue != *other.Continue {
			return false
		}
	}
	if resource.TotalHits == nil && other.TotalHits != nil || resource.TotalHits != nil && other.TotalHits == nil {
		return false
	}

	if resource.TotalHits != nil {
		if *resource.TotalHits != *other.TotalHits {
			return false
		}
	}
	if resource.TotalHitsRelation == nil && other.TotalHitsRelation != nil || resource.TotalHitsRelation != nil && other.TotalHitsRelation == nil {
		return false
	}

	if resource.TotalHitsRelation != nil {
		if *resource.TotalHitsRelation != *other.TotalHitsRelation {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SearchResultsMetadata` fields for violations and returns them.
func (resource SearchResultsMetadata) Validate() error {
	return nil
}

// #FacetValue is a single value/count pair in a facet breakdown.
type FacetValue struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

// NewFacetValue creates a new FacetValue object.
func NewFacetValue() *FacetValue {
	return &FacetValue{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FacetValue` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FacetValue) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "count"
	if fields["count"] != nil {
		if string(fields["count"]) != "null" {
			if err := json.Unmarshal(fields["count"], &resource.Count); err != nil {
				errs = append(errs, cog.MakeBuildErrors("count", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("count", errors.New("required field is null"))...)

		}
		delete(fields, "count")
	} else {
		errs = append(errs, cog.MakeBuildErrors("count", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FacetValue", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `FacetValue` objects.
func (resource FacetValue) Equals(other FacetValue) bool {
	if resource.Value != other.Value {
		return false
	}
	if resource.Count != other.Count {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FacetValue` fields for violations and returns them.
func (resource FacetValue) Validate() error {
	return nil
}

// #SearchResults is the response envelope, mirroring
// search.grafana.app SearchResults.
type SearchResults struct {
	Metadata SearchResultsMetadata   `json:"metadata"`
	Items    []SearchResultHit       `json:"items"`
	Facets   map[string][]FacetValue `json:"facets,omitempty"`
}

// NewSearchResults creates a new SearchResults object.
func NewSearchResults() *SearchResults {
	return &SearchResults{
		Metadata: *NewSearchResultsMetadata(),
		Items:    []SearchResultHit{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResults` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SearchResults) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "metadata"
	if fields["metadata"] != nil {
		if string(fields["metadata"]) != "null" {

			resource.Metadata = SearchResultsMetadata{}
			if err := resource.Metadata.UnmarshalJSONStrict(fields["metadata"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metadata", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metadata", errors.New("required field is null"))...)

		}
		delete(fields, "metadata")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metadata", errors.New("required field is missing from input"))...)
	}
	// Field "items"
	if fields["items"] != nil {
		if string(fields["items"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["items"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 SearchResultHit

				result1 = SearchResultHit{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is null"))...)

		}
		delete(fields, "items")
	} else {
		errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is missing from input"))...)
	}
	// Field "facets"
	if fields["facets"] != nil {
		if string(fields["facets"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["facets"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string][]FacetValue, len(partialMap))
			for key1 := range partialMap {
				var result1 []FacetValue

				partialArray := []json.RawMessage{}
				if err := json.Unmarshal(partialMap[key1], &partialArray); err != nil {
					return err
				}

				for i2 := range partialArray {
					var result2 FacetValue

					result2 = FacetValue{}
					if err := result2.UnmarshalJSONStrict(partialArray[i2]); err != nil {
						errs = append(errs, cog.MakeBuildErrors("facets["+key1+"]["+strconv.Itoa(i2)+"]", err)...)
					}
					result1 = append(result1, result2)
				}
				parsedMap1[key1] = result1
			}
			resource.Facets = parsedMap1

		}
		delete(fields, "facets")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SearchResults", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `SearchResults` objects.
func (resource SearchResults) Equals(other SearchResults) bool {
	if !resource.Metadata.Equals(other.Metadata) {
		return false
	}

	if len(resource.Items) != len(other.Items) {
		return false
	}

	for i1 := range resource.Items {
		if !resource.Items[i1].Equals(other.Items[i1]) {
			return false
		}
	}

	if len(resource.Facets) != len(other.Facets) {
		return false
	}

	for key1 := range resource.Facets {

		if len(resource.Facets[key1]) != len(other.Facets[key1]) {
			return false
		}

		for i2 := range resource.Facets[key1] {
			if !resource.Facets[key1][i2].Equals(other.Facets[key1][i2]) {
				return false
			}
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SearchResults` fields for violations and returns them.
func (resource SearchResults) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Metadata.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("metadata", err)...)
	}

	for i1 := range resource.Items {
		if err := resource.Items[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for key1 := range resource.Facets {

		for i2 := range resource.Facets[key1] {
			if err := resource.Facets[key1][i2].Validate(); err != nil {
				errs = append(errs, cog.MakeBuildErrors("facets["+key1+"]["+strconv.Itoa(i2)+"]", err)...)
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

const ApiVersion = "rules.alerting.grafana.app/v0alpha1"

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
