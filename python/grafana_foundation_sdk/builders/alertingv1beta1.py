# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import alertingv1beta1
import re
from ..builders import resource as resource_builder
from ..models import resource


class IntervalTrigger(cogbuilder.Builder[alertingv1beta1.IntervalTrigger]):
    _internal: alertingv1beta1.IntervalTrigger

    def __init__(self) -> None:
        self._internal = alertingv1beta1.IntervalTrigger()

    def build(self) -> alertingv1beta1.IntervalTrigger:
        """
        Builds the object.
        """
        return self._internal    
    
    def interval(self, interval: alertingv1beta1.PromDuration) -> typing.Self:    
        self._internal.interval = interval
    
        return self
    


class SimplifiedRouting(cogbuilder.Builder[alertingv1beta1.SimplifiedRouting]):
    _internal: alertingv1beta1.SimplifiedRouting

    def __init__(self) -> None:
        self._internal = alertingv1beta1.SimplifiedRouting()

    def build(self) -> alertingv1beta1.SimplifiedRouting:
        """
        Builds the object.
        """
        return self._internal    
    
    def receiver(self, receiver: str) -> typing.Self:    
        self._internal.receiver = receiver
    
        return self
    
    def group_by(self, group_by: list[str]) -> typing.Self:    
        self._internal.group_by = group_by
    
        return self
    
    def group_wait(self, group_wait: alertingv1beta1.PromDuration) -> typing.Self:    
        self._internal.group_wait = group_wait
    
        return self
    
    def group_interval(self, group_interval: alertingv1beta1.PromDuration) -> typing.Self:    
        self._internal.group_interval = group_interval
    
        return self
    
    def repeat_interval(self, repeat_interval: alertingv1beta1.PromDuration) -> typing.Self:    
        self._internal.repeat_interval = repeat_interval
    
        return self
    
    def mute_time_intervals(self, mute_time_intervals: list[alertingv1beta1.TimeIntervalRef]) -> typing.Self:    
        self._internal.mute_time_intervals = mute_time_intervals
    
        return self
    
    def active_time_intervals(self, active_time_intervals: list[alertingv1beta1.TimeIntervalRef]) -> typing.Self:    
        self._internal.active_time_intervals = active_time_intervals
    
        return self
    


class NamedRoutingTree(cogbuilder.Builder[alertingv1beta1.NamedRoutingTree]):
    _internal: alertingv1beta1.NamedRoutingTree

    def __init__(self) -> None:
        self._internal = alertingv1beta1.NamedRoutingTree()

    def build(self) -> alertingv1beta1.NamedRoutingTree:
        """
        Builds the object.
        """
        return self._internal    
    
    def routing_tree(self, routing_tree: str) -> typing.Self:    
        self._internal.routing_tree = routing_tree
    
        return self
    


class Expression(cogbuilder.Builder[alertingv1beta1.Expression]):
    _internal: alertingv1beta1.Expression

    def __init__(self) -> None:
        self._internal = alertingv1beta1.Expression()

    def build(self) -> alertingv1beta1.Expression:
        """
        Builds the object.
        """
        return self._internal    
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        The type of query if this is a query expression
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def relative_time_range(self, relative_time_range: cogbuilder.Builder[alertingv1beta1.RelativeTimeRange]) -> typing.Self:    
        relative_time_range_resource = relative_time_range.build()
        self._internal.relative_time_range = relative_time_range_resource
    
        return self
    
    def datasource_uid(self, datasource_uid: alertingv1beta1.DatasourceUID) -> typing.Self:    
        """
        The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
        """
            
        self._internal.datasource_uid = datasource_uid
    
        return self
    
    def model(self, model: object) -> typing.Self:    
        self._internal.model = model
    
        return self
    
    def source(self, source: bool) -> typing.Self:    
        """
        Used to mark the expression to be used as the final source for the rule evaluation
        Only one expression in a rule can be marked as the source
        For AlertRules, this is the expression that will be evaluated against the alerting condition
        For RecordingRules, this is the expression that will be recorded
        """
            
        self._internal.source = source
    
        return self
    


class RelativeTimeRange(cogbuilder.Builder[alertingv1beta1.RelativeTimeRange]):
    _internal: alertingv1beta1.RelativeTimeRange

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RelativeTimeRange()

    def build(self) -> alertingv1beta1.RelativeTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: alertingv1beta1.PromDurationWMillis) -> typing.Self:    
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: alertingv1beta1.PromDurationWMillis) -> typing.Self:    
        self._internal.to = to
    
        return self
    


class PanelRef(cogbuilder.Builder[alertingv1beta1.PanelRef]):
    _internal: alertingv1beta1.PanelRef

    def __init__(self) -> None:
        self._internal = alertingv1beta1.PanelRef()

    def build(self) -> alertingv1beta1.PanelRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def dashboard_uid(self, dashboard_uid: str) -> typing.Self:    
        if not len(dashboard_uid) >= 1:
            raise ValueError("len(dashboard_uid) must be >= 1")
        if not re.match("^[a-zA-Z0-9_-]+$", dashboard_uid):
            raise ValueError("dashboard_uid must match regex ^[a-zA-Z0-9_-]+$")
        self._internal.dashboard_uid = dashboard_uid
    
        return self
    
    def panel_id(self, panel_id: int) -> typing.Self:    
        if not panel_id > 0:
            raise ValueError("panel_id must be > 0")
        self._internal.panel_id = panel_id
    
        return self
    


class AlertRuleInstanceTotals(cogbuilder.Builder[alertingv1beta1.AlertRuleInstanceTotals]):    
    """
    Count of alert instances per state. error also counts instances whose evaluation
    errored but were mapped to another state via execErrState, so it can overlap.
    """
    
    _internal: alertingv1beta1.AlertRuleInstanceTotals

    def __init__(self) -> None:
        self._internal = alertingv1beta1.AlertRuleInstanceTotals()

    def build(self) -> alertingv1beta1.AlertRuleInstanceTotals:
        """
        Builds the object.
        """
        return self._internal    
    
    def healthy(self, healthy: int) -> typing.Self:    
        self._internal.healthy = healthy
    
        return self
    
    def firing(self, firing: int) -> typing.Self:    
        self._internal.firing = firing
    
        return self
    
    def pending(self, pending: int) -> typing.Self:    
        self._internal.pending = pending
    
        return self
    
    def recovering(self, recovering: int) -> typing.Self:    
        self._internal.recovering = recovering
    
        return self
    
    def nodata(self, nodata: int) -> typing.Self:    
        self._internal.nodata = nodata
    
        return self
    
    def error(self, error: int) -> typing.Self:    
        self._internal.error = error
    
        return self
    


class RuleRef(cogbuilder.Builder[alertingv1beta1.RuleRef]):
    _internal: alertingv1beta1.RuleRef

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RuleRef()

    def build(self) -> alertingv1beta1.RuleRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: alertingv1beta1.RuleUID) -> typing.Self:    
        """
        name is the metadata.name of an AlertRule or RecordingRule resource.
        """
            
        self._internal.name = name
    
        return self
    


class RuleSpec(cogbuilder.Builder[alertingv1beta1.RuleSpec]):
    _internal: alertingv1beta1.RuleSpec

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RuleSpec()

    def build(self) -> alertingv1beta1.RuleSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def paused(self, paused: bool) -> typing.Self:    
        self._internal.paused = paused
    
        return self
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv1beta1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def expressions(self, expressions: alertingv1beta1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    


class AlertRuleStatus(cogbuilder.Builder[alertingv1beta1.AlertRuleStatus]):
    _internal: alertingv1beta1.AlertRuleStatus

    def __init__(self) -> None:
        self._internal = alertingv1beta1.AlertRuleStatus()

    def build(self) -> alertingv1beta1.AlertRuleStatus:
        """
        Builds the object.
        """
        return self._internal    
    
    def health(self, health: alertingv1beta1.AlertRuleHealth) -> typing.Self:    
        self._internal.health = health
    
        return self
    
    def state(self, state: alertingv1beta1.AlertRuleState) -> typing.Self:    
        self._internal.state = state
    
        return self
    
    def state_reason(self, state_reason: alertingv1beta1.AlertRuleStateReason) -> typing.Self:    
        self._internal.state_reason = state_reason
    
        return self
    
    def last_evaluation_time(self, last_evaluation_time: str) -> typing.Self:    
        self._internal.last_evaluation_time = last_evaluation_time
    
        return self
    
    def evaluation_duration(self, evaluation_duration: float) -> typing.Self:    
        """
        duration of the last evaluation in seconds
        """
            
        self._internal.evaluation_duration = evaluation_duration
    
        return self
    
    def last_error(self, last_error: str) -> typing.Self:    
        self._internal.last_error = last_error
    
        return self
    
    def totals(self, totals: cogbuilder.Builder[alertingv1beta1.AlertRuleInstanceTotals]) -> typing.Self:    
        totals_resource = totals.build()
        self._internal.totals = totals_resource
    
        return self
    


class RecordingRuleStatus(cogbuilder.Builder[alertingv1beta1.RecordingRuleStatus]):
    _internal: alertingv1beta1.RecordingRuleStatus

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RecordingRuleStatus()

    def build(self) -> alertingv1beta1.RecordingRuleStatus:
        """
        Builds the object.
        """
        return self._internal    
    
    def health(self, health: alertingv1beta1.RecordingRuleHealth) -> typing.Self:    
        self._internal.health = health
    
        return self
    
    def last_evaluation_time(self, last_evaluation_time: str) -> typing.Self:    
        self._internal.last_evaluation_time = last_evaluation_time
    
        return self
    
    def evaluation_duration(self, evaluation_duration: float) -> typing.Self:    
        """
        duration of the last evaluation in seconds
        """
            
        self._internal.evaluation_duration = evaluation_duration
    
        return self
    
    def last_error(self, last_error: str) -> typing.Self:    
        self._internal.last_error = last_error
    
        return self
    


class AlertRule(cogbuilder.Builder[alertingv1beta1.AlertRuleSpec]):
    _internal: alertingv1beta1.AlertRuleSpec

    def __init__(self) -> None:
        self._internal = alertingv1beta1.AlertRuleSpec()

    def build(self) -> alertingv1beta1.AlertRuleSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def paused(self, paused: bool) -> typing.Self:    
        self._internal.paused = paused
    
        return self
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv1beta1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def annotations(self, annotations: dict[str, alertingv1beta1.TemplateString]) -> typing.Self:    
        self._internal.annotations = annotations
    
        return self
    
    def for_val(self, for_val: str) -> typing.Self:    
        self._internal.for_val = for_val
    
        return self
    
    def keep_firing_for(self, keep_firing_for: str) -> typing.Self:    
        self._internal.keep_firing_for = keep_firing_for
    
        return self
    
    def missing_series_evals_to_resolve(self, missing_series_evals_to_resolve: int) -> typing.Self:    
        if not missing_series_evals_to_resolve >= 0:
            raise ValueError("missing_series_evals_to_resolve must be >= 0")
        self._internal.missing_series_evals_to_resolve = missing_series_evals_to_resolve
    
        return self
    
    def no_data_state(self, no_data_state: alertingv1beta1.NoDataState) -> typing.Self:    
        self._internal.no_data_state = no_data_state
    
        return self
    
    def exec_err_state(self, exec_err_state: alertingv1beta1.ExecErrState) -> typing.Self:    
        self._internal.exec_err_state = exec_err_state
    
        return self
    
    def notification_settings(self, notification_settings: cogbuilder.Builder[alertingv1beta1.NotificationSettings]) -> typing.Self:    
        notification_settings_resource = notification_settings.build()
        self._internal.notification_settings = notification_settings_resource
    
        return self
    
    def expressions(self, expressions: alertingv1beta1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    
    def panel_ref(self, panel_ref: cogbuilder.Builder[alertingv1beta1.PanelRef]) -> typing.Self:    
        panel_ref_resource = panel_ref.build()
        self._internal.panel_ref = panel_ref_resource
    
        return self
    

"""
Creates a resource manifest from an AlertRule.
"""
def alert_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.AlertRuleSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v1beta1")
    builder.kind("AlertRule")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder


class RecordingRule(cogbuilder.Builder[alertingv1beta1.RecordingRuleSpec]):
    _internal: alertingv1beta1.RecordingRuleSpec

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RecordingRuleSpec()

    def build(self) -> alertingv1beta1.RecordingRuleSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def paused(self, paused: bool) -> typing.Self:    
        self._internal.paused = paused
    
        return self
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv1beta1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def metric(self, metric: alertingv1beta1.MetricName) -> typing.Self:    
        self._internal.metric = metric
    
        return self
    
    def expressions(self, expressions: alertingv1beta1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    
    def target_datasource_uid(self, target_datasource_uid: alertingv1beta1.DatasourceUID) -> typing.Self:    
        self._internal.target_datasource_uid = target_datasource_uid
    
        return self
    

"""
Creates a resource manifest from a RecordingRule.
"""
def recording_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.RecordingRuleSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v1beta1")
    builder.kind("RecordingRule")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder


class RuleSequence(cogbuilder.Builder[alertingv1beta1.RuleSequenceSpec]):
    _internal: alertingv1beta1.RuleSequenceSpec

    def __init__(self) -> None:
        self._internal = alertingv1beta1.RuleSequenceSpec()

    def build(self) -> alertingv1beta1.RuleSequenceSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def recording_rules(self, recording_rules: list[cogbuilder.Builder[alertingv1beta1.RuleRef]]) -> typing.Self:    
        recording_rules_resources = [r1.build() for r1 in recording_rules]
        self._internal.recording_rules = recording_rules_resources
    
        return self
    
    def alert_rules(self, alert_rules: list[cogbuilder.Builder[alertingv1beta1.RuleRef]]) -> typing.Self:    
        alert_rules_resources = [r1.build() for r1 in alert_rules]
        self._internal.alert_rules = alert_rules_resources
    
        return self
    

"""
Creates a resource manifest from a RuleSequence.
"""
def rule_sequence_manifest(name: str, folder: cogbuilder.Builder[alertingv1beta1.RuleSequenceSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v1beta1")
    builder.kind("RuleSequence")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder
