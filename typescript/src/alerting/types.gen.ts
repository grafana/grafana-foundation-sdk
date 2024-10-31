// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';


export interface Query {
	// Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
	datasourceUid?: string;
	// JSON is the raw JSON query and includes the above properties as well as custom properties.
	model?: cog.Dataquery;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId?: string;
	// RelativeTimeRange is the per query start and end time
	// for requests.
	relativeTimeRange?: RelativeTimeRange;
}

export const defaultQuery = (): Query => ({
});

export interface RuleGroup {
	folderUid?: string;
	// The interval, in seconds, at which all rules in the group are evaluated.
	// If a group contains many rules, the rules are evaluated sequentially.
	interval?: Duration;
	rules?: Rule[];
	title?: string;
}

export const defaultRuleGroup = (): RuleGroup => ({
});

export interface NotificationSettings {
	// Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
	// cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
	// use the special value '...' as the sole label name.
	// This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
	// you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
	// Must include 'alertname' and 'grafana_folder' if not using '...'.
	group_by?: string[];
	// Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
	// which an initial notification has already been sent. (Usually ~5m or more.)
	group_interval?: string;
	// Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
	// inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
	group_wait?: string;
	// Override the times when notifications should be muted. These must match the name of a mute time interval defined
	// in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
	// otherwise acts normally.
	mute_time_intervals?: string[];
	// Name of the receiver to send notifications to.
	receiver: string;
	// Override how long to wait before sending a notification again if it has already been sent successfully for an
	// alert. (Usually ~3h or more).
	// Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
	// Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
	// occurs first. `repeat_interval` should not be less than `group_interval`.
	repeat_interval?: string;
}

export const defaultNotificationSettings = (): NotificationSettings => ({
	group_by: [
"alertname",
"grafana_folder",
],
	receiver: "",
});

// Duration in seconds.
export type Duration = number;

export const defaultDuration = (): Duration => (0);

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
export interface ContactPoint {
	disableResolveMessage?: boolean;
	// Name is used as grouping key in the UI. Contact points with the
	// same name will be grouped in the UI.
	name?: string;
	provenance?: string;
	settings: Json;
	type: "alertmanager" | " dingding" | " discord" | " email" | " googlechat" | " kafka" | " line" | " opsgenie" | " pagerduty" | " pushover" | " sensugo" | " slack" | " teams" | " telegram" | " threema" | " victorops" | " webhook" | " wecom";
	// UID is the unique identifier of the contact point. The UID can be
	// set by the user.
	uid?: string;
}

export const defaultContactPoint = (): ContactPoint => ({
	settings: defaultJson(),
	type: "alertmanager",
});

export type Json = any;

export const defaultJson = (): Json => ({});

export type MatchRegexps = Record<string, string>;

export const defaultMatchRegexps = (): MatchRegexps => ({});

export enum MatchType {
	Equal = "=",
	NotEqual = "!=",
	EqualRegex = "=~",
	NotEqualRegex = "!~",
}

export const defaultMatchType = (): MatchType => (MatchType.Equal);

export interface Matcher {
	Name?: string;
	Type?: MatchType;
	Value?: string;
}

export const defaultMatcher = (): Matcher => ({
});

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
export type Matchers = Matcher[];

export const defaultMatchers = (): Matchers => ([]);

export interface MuteTiming {
	name?: string;
	time_intervals?: TimeInterval[];
}

export const defaultMuteTiming = (): MuteTiming => ({
});

export interface NotificationTemplate {
	name?: string;
	provenance?: Provenance;
	template?: string;
}

export const defaultNotificationTemplate = (): NotificationTemplate => ({
});

export type ObjectMatcher = string[];

export const defaultObjectMatcher = (): ObjectMatcher => ([]);

export type ObjectMatchers = ObjectMatcher[];

export const defaultObjectMatchers = (): ObjectMatchers => ([]);

export type Provenance = string;

export const defaultProvenance = (): Provenance => ("");

export interface Rule {
	annotations?: Record<string, string>;
	condition: string;
	data: Query[];
	execErrState: "OK" | "Alerting" | "Error";
	folderUID: string;
	// The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
	// Before this time has elapsed, the rule is only considered to be Pending.
	for: string;
	id?: number;
	isPaused?: boolean;
	labels?: Record<string, string>;
	noDataState: "Alerting" | "NoData" | "OK";
	notification_settings?: NotificationSettings;
	orgID: number;
	provenance?: Provenance;
	record?: RecordRule;
	ruleGroup: string;
	title: string;
	uid?: string;
	updated?: string;
}

export const defaultRule = (): Rule => ({
	condition: "",
	data: [],
	execErrState: "OK",
	folderUID: "",
	for: "",
	noDataState: "Alerting",
	orgID: 0,
	ruleGroup: "",
	title: "",
});

export interface RecordRule {
	// Which expression node should be used as the input for the recorded metric.
	from: string;
	// Name of the recorded metric.
	metric: string;
}

export const defaultRecordRule = (): RecordRule => ({
	from: "",
	metric: "",
});

// RelativeTimeRange is the per query start and end time
// for requests.
export interface RelativeTimeRange {
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	from?: Duration;
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	to?: Duration;
}

export const defaultRelativeTimeRange = (): RelativeTimeRange => ({
});

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
export interface NotificationPolicy {
	continue?: boolean;
	group_by?: string[];
	group_interval?: string;
	group_wait?: string;
	// Deprecated. Remove before v1.0 release.
	match?: Record<string, string>;
	match_re?: MatchRegexps;
	// Matchers is a slice of Matchers that is sortable, implements Stringer, and
	// provides a Matches method to match a LabelSet against all Matchers in the
	// slice. Note that some users of Matchers might require it to be sorted.
	matchers?: Matchers;
	mute_time_intervals?: string[];
	object_matchers?: ObjectMatchers;
	provenance?: Provenance;
	receiver?: string;
	repeat_interval?: string;
	routes?: NotificationPolicy[];
}

export const defaultNotificationPolicy = (): NotificationPolicy => ({
});

export interface TimeInterval {
	name?: string;
	time_intervals?: TimeIntervalItem[];
}

export const defaultTimeInterval = (): TimeInterval => ({
});

export interface TimeIntervalItem {
	days_of_month?: string[];
	location?: string;
	months?: string[];
	times?: TimeIntervalTimeRange[];
	weekdays?: string[];
	years?: string[];
}

export const defaultTimeIntervalItem = (): TimeIntervalItem => ({
});

export interface TimeIntervalTimeRange {
	end_time?: string;
	start_time?: string;
}

export const defaultTimeIntervalTimeRange = (): TimeIntervalTimeRange => ({
});

