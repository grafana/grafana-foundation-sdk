// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';


// Duration in seconds.
export type Duration = number;

export const defaultDuration = (): Duration => (0);

export type Json = any;

export const defaultJson = (): Json => ({});

export type MatchRegexps = Record<string, string>;

export const defaultMatchRegexps = (): MatchRegexps => ({});

export type MatchType = number;

export const defaultMatchType = (): MatchType => (0);

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

// RelativeTimeRange is the per query start and end time
// for requests.
export interface RelativeTimeRange {
	// RelativeTimeRange is the per query start and end time
	// for requests.
	from?: Duration;
	// RelativeTimeRange is the per query start and end time
	// for requests.
	to?: Duration;
}

export const defaultRelativeTimeRange = (): RelativeTimeRange => ({
});

export interface TimeInterval {
	name?: string;
	time_intervals?: TimeInterval[];
}

export const defaultTimeInterval = (): TimeInterval => ({
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

export interface NotificationSettings {
	group_by?: string[];
	group_interval?: string;
	group_wait?: string;
	mute_time_intervals?: string[];
	receiver: string;
	repeat_interval?: string;
}

export const defaultNotificationSettings = (): NotificationSettings => ({
	group_by: [
"alertname",
"grafana_folder",
],
	receiver: "",
});

export interface Query {
	datasourceUid?: string;
	model?: cog.Dataquery;
	queryType?: string;
	refId?: string;
	relativeTimeRange?: RelativeTimeRange;
}

export const defaultQuery = (): Query => ({
});

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
export interface ContactPoint {
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	disableResolveMessage?: boolean;
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	name?: string;
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	provenance?: string;
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	settings: Json;
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	type: "alertmanager" | " dingding" | " discord" | " email" | " googlechat" | " kafka" | " line" | " opsgenie" | " pagerduty" | " pushover" | " sensugo" | " slack" | " teams" | " telegram" | " threema" | " victorops" | " webhook" | " wecom";
	// EmbeddedContactPoint is the contact point type that is used
	// by grafanas embedded alertmanager implementation.
	uid?: string;
}

export const defaultContactPoint = (): ContactPoint => ({
	settings: defaultJson(),
	type: "alertmanager",
});

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
export interface NotificationPolicy {
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	continue?: boolean;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	group_by?: string[];
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	group_interval?: string;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	group_wait?: string;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	match?: Record<string, string>;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	match_re?: MatchRegexps;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	matchers?: Matchers;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	mute_time_intervals?: string[];
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	object_matchers?: ObjectMatchers;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	provenance?: Provenance;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	receiver?: string;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	repeat_interval?: string;
	// A Route is a node that contains definitions of how to handle alerts. This is modified
	// from the upstream alertmanager in that it adds the ObjectMatchers property.
	routes?: NotificationPolicy[];
}

export const defaultNotificationPolicy = (): NotificationPolicy => ({
});

export interface MuteTiming {
	name?: string;
	time_intervals?: TimeInterval[];
}

export const defaultMuteTiming = (): MuteTiming => ({
});

