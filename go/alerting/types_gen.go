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

type RuleGroup struct {
	FolderUid *string `json:"folderUid,omitempty"`
	// The interval, in seconds, at which all rules in the group are evaluated.
	// If a group contains many rules, the rules are evaluated sequentially.
	Interval *Duration `json:"interval,omitempty"`
	Rules    []Rule    `json:"rules,omitempty"`
	Title    *string   `json:"title,omitempty"`
}

type NotificationSettings struct {
	GroupBy           []string `json:"group_by,omitempty"`
	GroupInterval     *string  `json:"group_interval,omitempty"`
	GroupWait         *string  `json:"group_wait,omitempty"`
	MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
	Receiver          string   `json:"receiver"`
	RepeatInterval    *string  `json:"repeat_interval,omitempty"`
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

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
type Matchers []Matcher

type MuteTiming struct {
	Name          *string        `json:"name,omitempty"`
	TimeIntervals []TimeInterval `json:"time_intervals,omitempty"`
}

type NotificationTemplate struct {
	Name       *string     `json:"name,omitempty"`
	Provenance *Provenance `json:"provenance,omitempty"`
	Template   *string     `json:"template,omitempty"`
	Version    *string     `json:"version,omitempty"`
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
	Record               *RecordRule           `json:"record,omitempty"`
	RuleGroup            string                `json:"ruleGroup"`
	Title                string                `json:"title"`
	Uid                  *string               `json:"uid,omitempty"`
	Updated              *time.Time            `json:"updated,omitempty"`
}

type RecordRule struct {
	From   string `json:"from"`
	Metric string `json:"metric"`
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

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
type NotificationPolicy struct {
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	ActiveTimeIntervals []string `json:"active_time_intervals,omitempty"`
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

type TimeInterval struct {
	Name          *string            `json:"name,omitempty"`
	TimeIntervals []TimeIntervalItem `json:"time_intervals,omitempty"`
}

type TimeIntervalItem struct {
	DaysOfMonth []string                `json:"days_of_month,omitempty"`
	Location    *string                 `json:"location,omitempty"`
	Months      []string                `json:"months,omitempty"`
	Times       []TimeIntervalTimeRange `json:"times,omitempty"`
	Weekdays    []string                `json:"weekdays,omitempty"`
	Years       []string                `json:"years,omitempty"`
}

type TimeIntervalTimeRange struct {
	EndTime   *string `json:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
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
