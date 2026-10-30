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

export enum AlertRuleHealth {
	Unknown = "Unknown",
	OK = "OK",
	Paused = "Paused",
	Error = "Error",
	NoData = "NoData",
}

export const defaultAlertRuleHealth = (): AlertRuleHealth => (AlertRuleHealth.Unknown);

export enum AlertRuleState {
	Inactive = "Inactive",
	Healthy = "Healthy",
	Firing = "Firing",
	Pending = "Pending",
	Recovering = "Recovering",
}

export const defaultAlertRuleState = (): AlertRuleState => (AlertRuleState.Inactive);

export enum AlertRuleStateReason {
	Evaluated = "Evaluated",
	KeepLast = "KeepLast",
}

export const defaultAlertRuleStateReason = (): AlertRuleStateReason => (AlertRuleStateReason.Evaluated);

/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 */
export interface AlertRuleInstanceTotals {
	healthy?: number;
	firing?: number;
	pending?: number;
	recovering?: number;
	nodata?: number;
	error?: number;
}

export const defaultAlertRuleInstanceTotals = (): AlertRuleInstanceTotals => ({
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

export enum RecordingRuleHealth {
	Unknown = "Unknown",
	Recording = "Recording",
	Paused = "Paused",
	Error = "Error",
	NoData = "NoData",
}

export const defaultRecordingRuleHealth = (): RecordingRuleHealth => (RecordingRuleHealth.Unknown);

export interface RuleSequenceSpec {
	trigger: IntervalTrigger;
	recordingRules: RuleRef[];
	alertRules?: RuleRef[];
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

export interface AlertRuleStatus {
	health?: AlertRuleHealth;
	state?: AlertRuleState;
	stateReason?: AlertRuleStateReason;
	lastEvaluationTime?: string;
	// duration of the last evaluation in seconds
	evaluationDuration?: number;
	lastError?: string;
	totals?: AlertRuleInstanceTotals;
}

export const defaultAlertRuleStatus = (): AlertRuleStatus => ({
});

export interface RecordingRuleStatus {
	health?: RecordingRuleHealth;
	lastEvaluationTime?: string;
	// duration of the last evaluation in seconds
	evaluationDuration?: number;
	lastError?: string;
}

export const defaultRecordingRuleStatus = (): RecordingRuleStatus => ({
});

export const ApiVersion = "rules.alerting.grafana.app/v1beta1";

export const AlertRuleKind = "AlertRule";

export const RecordingRuleKind = "RecordingRule";

export const RuleSequenceKind = "RuleSequence";

