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
	orgID: number;
	provenance?: Provenance;
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

// TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
// within the interval.
export interface TimeInterval {
	days_of_month?: string[];
	location?: string;
	months?: string[];
	times?: TimeRange[];
	weekdays?: string[];
	years?: string[];
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

// Redefining this to avoid an import cycle
export interface TimeRange {
	from?: string;
	to?: string;
}

export const defaultTimeRange = (): TimeRange => ({
});

