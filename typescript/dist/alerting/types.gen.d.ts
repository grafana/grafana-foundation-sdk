import * as cog from '../cog';
export interface Query {
    datasourceUid?: string;
    model?: cog.Dataquery;
    queryType?: string;
    refId?: string;
    relativeTimeRange?: RelativeTimeRange;
}
export declare const defaultQuery: () => Query;
export interface RuleGroup {
    folderUid?: string;
    interval?: Duration;
    rules?: Rule[];
    title?: string;
}
export declare const defaultRuleGroup: () => RuleGroup;
export interface NotificationSettings {
    group_by?: string[];
    group_interval?: string;
    group_wait?: string;
    mute_time_intervals?: string[];
    receiver: string;
    repeat_interval?: string;
}
export declare const defaultNotificationSettings: () => NotificationSettings;
export type Duration = number;
export declare const defaultDuration: () => Duration;
export interface ContactPoint {
    disableResolveMessage?: boolean;
    name?: string;
    provenance?: string;
    settings: Json;
    type: "alertmanager" | "dingding" | "discord" | "email" | "googlechat" | "kafka" | "line" | "opsgenie" | "pagerduty" | "pushover" | "sensugo" | "slack" | "teams" | "telegram" | "threema" | "victorops" | "webhook" | "wecom";
    uid?: string;
}
export declare const defaultContactPoint: () => ContactPoint;
export type Json = any;
export declare const defaultJson: () => Json;
export type MatchRegexps = Record<string, string>;
export declare const defaultMatchRegexps: () => MatchRegexps;
export declare enum MatchType {
    Equal = "=",
    NotEqual = "!=",
    EqualRegex = "=~",
    NotEqualRegex = "!~"
}
export declare const defaultMatchType: () => MatchType;
export interface Matcher {
    Name?: string;
    Type?: MatchType;
    Value?: string;
}
export declare const defaultMatcher: () => Matcher;
export type Matchers = Matcher[];
export declare const defaultMatchers: () => Matchers;
export interface MuteTiming {
    name?: string;
    time_intervals?: TimeInterval[];
}
export declare const defaultMuteTiming: () => MuteTiming;
export interface NotificationTemplate {
    name?: string;
    provenance?: Provenance;
    template?: string;
    version?: string;
}
export declare const defaultNotificationTemplate: () => NotificationTemplate;
export type ObjectMatcher = string[];
export declare const defaultObjectMatcher: () => ObjectMatcher;
export type ObjectMatchers = ObjectMatcher[];
export declare const defaultObjectMatchers: () => ObjectMatchers;
export type Provenance = string;
export declare const defaultProvenance: () => Provenance;
export interface Rule {
    annotations?: Record<string, string>;
    condition: string;
    data: Query[];
    execErrState: "OK" | "Alerting" | "Error";
    folderUID: string;
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
export declare const defaultRule: () => Rule;
export interface RecordRule {
    from: string;
    metric: string;
}
export declare const defaultRecordRule: () => RecordRule;
export interface RelativeTimeRange {
    from?: Duration;
    to?: Duration;
}
export declare const defaultRelativeTimeRange: () => RelativeTimeRange;
export interface NotificationPolicy {
    active_time_intervals?: string[];
    continue?: boolean;
    group_by?: string[];
    group_interval?: string;
    group_wait?: string;
    match?: Record<string, string>;
    match_re?: MatchRegexps;
    matchers?: Matchers;
    mute_time_intervals?: string[];
    object_matchers?: ObjectMatchers;
    provenance?: Provenance;
    receiver?: string;
    repeat_interval?: string;
    routes?: NotificationPolicy[];
}
export declare const defaultNotificationPolicy: () => NotificationPolicy;
export interface TimeInterval {
    times?: TimeRange[];
    weekdays?: WeekdayRange[];
    days_of_month?: DayOfMonthRange[];
    months?: MonthRange[];
    years?: YearRange[];
    location?: Location;
}
export declare const defaultTimeInterval: () => TimeInterval;
export interface TimeRange {
    from?: string;
    to?: string;
}
export declare const defaultTimeRange: () => TimeRange;
export interface WeekdayRange {
    begin?: number;
    end?: number;
}
export declare const defaultWeekdayRange: () => WeekdayRange;
export interface DayOfMonthRange {
    begin?: number;
    end?: number;
}
export declare const defaultDayOfMonthRange: () => DayOfMonthRange;
export interface YearRange {
    begin?: number;
    end?: number;
}
export declare const defaultYearRange: () => YearRange;
export interface MonthRange {
    begin?: number;
    end?: number;
}
export declare const defaultMonthRange: () => MonthRange;
export type Location = string;
export declare const defaultLocation: () => Location;
