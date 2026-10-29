// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum NoDataState {
	NoData = "NoData",
	Ok = "Ok",
	Alerting = "Alerting",
	KeepLast = "KeepLast",
}

export const defaultNoDataState = (): NoDataState => (NoDataState.NoData);

export enum ExecErrState {
	Error = "Error",
	Ok = "Ok",
	Alerting = "Alerting",
	KeepLast = "KeepLast",
}

export const defaultExecErrState = (): ExecErrState => (ExecErrState.Error);

export type TimeIntervalRef = string;

export const defaultTimeIntervalRef = (): TimeIntervalRef => ("");

export interface AlertRuleSpec {
	title: string;
	paused?: boolean;
	trigger: IntervalTrigger;
	labels?: Record<string, TemplateString>;
	annotations?: Record<string, TemplateString>;
	for?: string;
	keepFiringFor?: string;
	missingSeriesEvalsToResolve?: number;
	noDataState: NoDataState;
	execErrState: ExecErrState;
	notificationSettings?: NotificationSettings;
	expressions: ExpressionMap;
	panelRef?: PanelRef;
}

export const defaultAlertRuleSpec = (): AlertRuleSpec => ({
	title: "",
	trigger: defaultIntervalTrigger(),
	noDataState: NoDataState.NoData,
	execErrState: ExecErrState.Error,
	expressions: defaultExpressionMap(),
});

export interface IntervalTrigger {
	interval: PromDuration;
}

export const defaultIntervalTrigger = (): IntervalTrigger => ({
	interval: defaultPromDuration(),
});

export type PromDuration = string;

export const defaultPromDuration = (): PromDuration => ("");

export type TemplateString = string;

export const defaultTemplateString = (): TemplateString => ("");

export type NotificationSettings = SimplifiedRouting | NamedRoutingTree;

export const defaultNotificationSettings = (): NotificationSettings => (defaultSimplifiedRouting());

export interface SimplifiedRouting {
	type: NotificationSettingsType.SimplifiedRouting;
	receiver: string;
	groupBy?: string[];
	groupWait?: PromDuration;
	groupInterval?: PromDuration;
	repeatInterval?: PromDuration;
	muteTimeIntervals?: TimeIntervalRef[];
	activeTimeIntervals?: TimeIntervalRef[];
}

export const defaultSimplifiedRouting = (): SimplifiedRouting => ({
	type: NotificationSettingsType.SimplifiedRouting,
	receiver: "",
});

export enum NotificationSettingsType {
	SimplifiedRouting = "SimplifiedRouting",
	NamedRoutingTree = "NamedRoutingTree",
}

export const defaultNotificationSettingsType = (): NotificationSettingsType => (NotificationSettingsType.SimplifiedRouting);

export interface NamedRoutingTree {
	type: NotificationSettingsType.NamedRoutingTree;
	routingTree: string;
}

export const defaultNamedRoutingTree = (): NamedRoutingTree => ({
	type: NotificationSettingsType.NamedRoutingTree,
	routingTree: "",
});

export type ExpressionMap = Record<string, Expression>;

export const defaultExpressionMap = (): ExpressionMap => ({});

export interface Expression {
	// The type of query if this is a query expression
	queryType?: string;
	relativeTimeRange?: RelativeTimeRange;
	// The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
	datasourceUID?: DatasourceUID;
	model: any;
	// Used to mark the expression to be used as the final source for the rule evaluation
	// Only one expression in a rule can be marked as the source
	// For AlertRules, this is the expression that will be evaluated against the alerting condition
	// For RecordingRules, this is the expression that will be recorded
	source?: boolean;
}

export const defaultExpression = (): Expression => ({
	model: {},
});

export interface RelativeTimeRange {
	from: PromDurationWMillis;
	to: PromDurationWMillis;
}

export const defaultRelativeTimeRange = (): RelativeTimeRange => ({
	from: defaultPromDurationWMillis(),
	to: defaultPromDurationWMillis(),
});

export type PromDurationWMillis = string;

export const defaultPromDurationWMillis = (): PromDurationWMillis => ("");

export type DatasourceUID = string;

export const defaultDatasourceUID = (): DatasourceUID => ("");

export interface PanelRef {
	dashboardUID: string;
	panelID: number;
}

export const defaultPanelRef = (): PanelRef => ({
	dashboardUID: "",
	panelID: 0,
});

export interface RecordingRuleSpec {
	title: string;
	paused?: boolean;
	trigger: IntervalTrigger;
	labels?: Record<string, TemplateString>;
	metric: MetricName;
	expressions: ExpressionMap;
	targetDatasourceUID: DatasourceUID;
}

export const defaultRecordingRuleSpec = (): RecordingRuleSpec => ({
	title: "",
	trigger: defaultIntervalTrigger(),
	metric: defaultMetricName(),
	expressions: defaultExpressionMap(),
	targetDatasourceUID: defaultDatasourceUID(),
});

export type MetricName = string;

export const defaultMetricName = (): MetricName => ("");

export interface RuleSequenceSpec {
	trigger: IntervalTrigger;
	recordingRules: RuleRef[];
	alertingRules?: RuleRef[];
}

export const defaultRuleSequenceSpec = (): RuleSequenceSpec => ({
	trigger: defaultIntervalTrigger(),
	recordingRules: [],
});

export interface RuleRef {
	// name is the metadata.name of an AlertRule or RecordingRule resource.
	name: RuleUID;
}

export const defaultRuleRef = (): RuleRef => ({
	name: defaultRuleUID(),
});

export type RuleUID = string;

export const defaultRuleUID = (): RuleUID => ("");

export interface RuleSpec {
	title: string;
	paused?: boolean;
	trigger: IntervalTrigger;
	labels?: Record<string, TemplateString>;
	expressions: ExpressionMap;
}

export const defaultRuleSpec = (): RuleSpec => ({
	title: "",
	trigger: defaultIntervalTrigger(),
	expressions: defaultExpressionMap(),
});

/**
 * #SearchResultResource is the full identity of a hit.
 */
export interface SearchResultResource {
	group: string;
	resource: string;
	kind: string;
	name: string;
}

export const defaultSearchResultResource = (): SearchResultResource => ({
	group: "",
	resource: "",
	kind: "",
	name: "",
});

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 */
export interface RuleSearchHitFields {
	title?: string;
	folder?: string;
	type?: string;
	interval?: string;
	paused?: boolean;
	labels?: Record<string, string>;
	datasourceUIDs?: string[];
	// Alert-rule fields.
	annotations?: Record<string, string>;
	for?: string;
	keepFiringFor?: string;
	dashboardUID?: string;
	panelID?: number;
	receiver?: string;
	notificationType?: string;
	routingTree?: string;
	// Recording-rule fields.
	metric?: string;
	targetDatasourceUID?: string;
}

export const defaultRuleSearchHitFields = (): RuleSearchHitFields => ({
});

/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 */
export interface SearchResultHit {
	resource: SearchResultResource;
	score?: number;
	fields: RuleSearchHitFields;
}

export const defaultSearchResultHit = (): SearchResultHit => ({
	resource: defaultSearchResultResource(),
	fields: defaultRuleSearchHitFields(),
});

/**
 * #TotalHitsRelation says how totalHits relates to the real number of matching
 * rules the caller may see: "eq" when it is exact, "lte" when it is an upper
 * bound because authorisation was applied after the search ranked its results.
 */
export enum TotalHitsRelation {
	Eq = "eq",
	Lte = "lte",
}

export const defaultTotalHitsRelation = (): TotalHitsRelation => (TotalHitsRelation.Eq);

/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 */
export interface SearchResultsMetadata {
	continue?: string;
	totalHits?: number;
	// Always read totalHits together with totalHitsRelation.
	totalHitsRelation?: TotalHitsRelation;
}

export const defaultSearchResultsMetadata = (): SearchResultsMetadata => ({
});

/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 */
export interface FacetValue {
	value: string;
	count: number;
}

export const defaultFacetValue = (): FacetValue => ({
	value: "",
	count: 0,
});

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 */
export interface SearchResults {
	metadata: SearchResultsMetadata;
	items: SearchResultHit[];
	facets?: Record<string, FacetValue[]>;
}

export const defaultSearchResults = (): SearchResults => ({
	metadata: defaultSearchResultsMetadata(),
	items: [],
});

export const ApiVersion = "rules.alerting.grafana.app/v0alpha1";

export const AlertRuleKind = "AlertRule";

export const RecordingRuleKind = "RecordingRule";

export const RuleSequenceKind = "RuleSequence";

