# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import alertingv0alpha1
import re
from ..builders import resource as resource_builder
from ..models import resource


class IntervalTrigger(cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]):
    _internal: alertingv0alpha1.IntervalTrigger

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.IntervalTrigger()

    def build(self) -> alertingv0alpha1.IntervalTrigger:
        """
        Builds the object.
        """
        return self._internal    
    
    def interval(self, interval: alertingv0alpha1.PromDuration) -> typing.Self:    
        self._internal.interval = interval
    
        return self
    


class SimplifiedRouting(cogbuilder.Builder[alertingv0alpha1.SimplifiedRouting]):
    _internal: alertingv0alpha1.SimplifiedRouting

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.SimplifiedRouting()

    def build(self) -> alertingv0alpha1.SimplifiedRouting:
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
    
    def group_wait(self, group_wait: alertingv0alpha1.PromDuration) -> typing.Self:    
        self._internal.group_wait = group_wait
    
        return self
    
    def group_interval(self, group_interval: alertingv0alpha1.PromDuration) -> typing.Self:    
        self._internal.group_interval = group_interval
    
        return self
    
    def repeat_interval(self, repeat_interval: alertingv0alpha1.PromDuration) -> typing.Self:    
        self._internal.repeat_interval = repeat_interval
    
        return self
    
    def mute_time_intervals(self, mute_time_intervals: list[alertingv0alpha1.TimeIntervalRef]) -> typing.Self:    
        self._internal.mute_time_intervals = mute_time_intervals
    
        return self
    
    def active_time_intervals(self, active_time_intervals: list[alertingv0alpha1.TimeIntervalRef]) -> typing.Self:    
        self._internal.active_time_intervals = active_time_intervals
    
        return self
    


class NamedRoutingTree(cogbuilder.Builder[alertingv0alpha1.NamedRoutingTree]):
    _internal: alertingv0alpha1.NamedRoutingTree

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.NamedRoutingTree()

    def build(self) -> alertingv0alpha1.NamedRoutingTree:
        """
        Builds the object.
        """
        return self._internal    
    
    def routing_tree(self, routing_tree: str) -> typing.Self:    
        self._internal.routing_tree = routing_tree
    
        return self
    


class Expression(cogbuilder.Builder[alertingv0alpha1.Expression]):
    _internal: alertingv0alpha1.Expression

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.Expression()

    def build(self) -> alertingv0alpha1.Expression:
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
    
    def relative_time_range(self, relative_time_range: cogbuilder.Builder[alertingv0alpha1.RelativeTimeRange]) -> typing.Self:    
        relative_time_range_resource = relative_time_range.build()
        self._internal.relative_time_range = relative_time_range_resource
    
        return self
    
    def datasource_uid(self, datasource_uid: alertingv0alpha1.DatasourceUID) -> typing.Self:    
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
    


class RelativeTimeRange(cogbuilder.Builder[alertingv0alpha1.RelativeTimeRange]):
    _internal: alertingv0alpha1.RelativeTimeRange

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RelativeTimeRange()

    def build(self) -> alertingv0alpha1.RelativeTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: alertingv0alpha1.PromDurationWMillis) -> typing.Self:    
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: alertingv0alpha1.PromDurationWMillis) -> typing.Self:    
        self._internal.to = to
    
        return self
    


class PanelRef(cogbuilder.Builder[alertingv0alpha1.PanelRef]):
    _internal: alertingv0alpha1.PanelRef

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.PanelRef()

    def build(self) -> alertingv0alpha1.PanelRef:
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
    


class RuleRef(cogbuilder.Builder[alertingv0alpha1.RuleRef]):
    _internal: alertingv0alpha1.RuleRef

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RuleRef()

    def build(self) -> alertingv0alpha1.RuleRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: alertingv0alpha1.RuleUID) -> typing.Self:    
        """
        name is the metadata.name of an AlertRule or RecordingRule resource.
        """
            
        self._internal.name = name
    
        return self
    


class RuleSpec(cogbuilder.Builder[alertingv0alpha1.RuleSpec]):
    _internal: alertingv0alpha1.RuleSpec

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RuleSpec()

    def build(self) -> alertingv0alpha1.RuleSpec:
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
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def expressions(self, expressions: alertingv0alpha1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    


class SearchResultResource(cogbuilder.Builder[alertingv0alpha1.SearchResultResource]):    
    """
    #SearchResultResource is the full identity of a hit.
    """
    
    _internal: alertingv0alpha1.SearchResultResource

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.SearchResultResource()

    def build(self) -> alertingv0alpha1.SearchResultResource:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def resource(self, resource: str) -> typing.Self:    
        self._internal.resource = resource
    
        return self
    
    def kind(self, kind: str) -> typing.Self:    
        self._internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class RuleSearchHitFields(cogbuilder.Builder[alertingv0alpha1.RuleSearchHitFields]):    
    """
    #RuleSearchHitFields is the per-kind field payload returned on each hit.
    It carries the union of alert- and recording-rule search fields; only the
    fields relevant to a hit's kind are populated. This maps to the kind's
    declared searchFields.
    """
    
    _internal: alertingv0alpha1.RuleSearchHitFields

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RuleSearchHitFields()

    def build(self) -> alertingv0alpha1.RuleSearchHitFields:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def folder(self, folder: str) -> typing.Self:    
        self._internal.folder = folder
    
        return self
    
    def type(self, type_val: str) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        self._internal.interval = interval
    
        return self
    
    def paused(self, paused: bool) -> typing.Self:    
        self._internal.paused = paused
    
        return self
    
    def labels(self, labels: dict[str, str]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def datasource_ui_ds(self, datasource_ui_ds: list[str]) -> typing.Self:    
        self._internal.datasource_ui_ds = datasource_ui_ds
    
        return self
    
    def annotations(self, annotations: dict[str, str]) -> typing.Self:    
        """
        Alert-rule fields.
        """
            
        self._internal.annotations = annotations
    
        return self
    
    def for_val(self, for_val: str) -> typing.Self:    
        self._internal.for_val = for_val
    
        return self
    
    def keep_firing_for(self, keep_firing_for: str) -> typing.Self:    
        self._internal.keep_firing_for = keep_firing_for
    
        return self
    
    def dashboard_uid(self, dashboard_uid: str) -> typing.Self:    
        self._internal.dashboard_uid = dashboard_uid
    
        return self
    
    def panel_id(self, panel_id: int) -> typing.Self:    
        self._internal.panel_id = panel_id
    
        return self
    
    def receiver(self, receiver: str) -> typing.Self:    
        self._internal.receiver = receiver
    
        return self
    
    def notification_type(self, notification_type: str) -> typing.Self:    
        self._internal.notification_type = notification_type
    
        return self
    
    def routing_tree(self, routing_tree: str) -> typing.Self:    
        self._internal.routing_tree = routing_tree
    
        return self
    
    def metric(self, metric: str) -> typing.Self:    
        """
        Recording-rule fields.
        """
            
        self._internal.metric = metric
    
        return self
    
    def target_datasource_uid(self, target_datasource_uid: str) -> typing.Self:    
        self._internal.target_datasource_uid = target_datasource_uid
    
        return self
    


class SearchResultHit(cogbuilder.Builder[alertingv0alpha1.SearchResultHit]):    
    """
    #SearchResultHit is a single match: its identity, an optional relevance
    score (present only when the query included free text), and the requested
    fields.
    """
    
    _internal: alertingv0alpha1.SearchResultHit

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.SearchResultHit()

    def build(self) -> alertingv0alpha1.SearchResultHit:
        """
        Builds the object.
        """
        return self._internal    
    
    def resource(self, resource: cogbuilder.Builder[alertingv0alpha1.SearchResultResource]) -> typing.Self:    
        resource_resource = resource.build()
        self._internal.resource = resource_resource
    
        return self
    
    def score(self, score: float) -> typing.Self:    
        self._internal.score = score
    
        return self
    
    def fields(self, fields: cogbuilder.Builder[alertingv0alpha1.RuleSearchHitFields]) -> typing.Self:    
        fields_resource = fields.build()
        self._internal.fields = fields_resource
    
        return self
    


class SearchResultsMetadata(cogbuilder.Builder[alertingv0alpha1.SearchResultsMetadata]):    
    """
    #SearchResultsMetadata carries the paging token and total authorised match
    count.
    """
    
    _internal: alertingv0alpha1.SearchResultsMetadata

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.SearchResultsMetadata()

    def build(self) -> alertingv0alpha1.SearchResultsMetadata:
        """
        Builds the object.
        """
        return self._internal    
    
    def continue_val(self, continue_val: str) -> typing.Self:    
        self._internal.continue_val = continue_val
    
        return self
    
    def total_hits(self, total_hits: int) -> typing.Self:    
        self._internal.total_hits = total_hits
    
        return self
    
    def total_hits_relation(self, total_hits_relation: alertingv0alpha1.TotalHitsRelation) -> typing.Self:    
        """
        Always read totalHits together with totalHitsRelation.
        """
            
        self._internal.total_hits_relation = total_hits_relation
    
        return self
    


class FacetValue(cogbuilder.Builder[alertingv0alpha1.FacetValue]):    
    """
    #FacetValue is a single value/count pair in a facet breakdown.
    """
    
    _internal: alertingv0alpha1.FacetValue

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.FacetValue()

    def build(self) -> alertingv0alpha1.FacetValue:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def count(self, count: int) -> typing.Self:    
        self._internal.count = count
    
        return self
    


class SearchResults(cogbuilder.Builder[alertingv0alpha1.SearchResults]):    
    """
    #SearchResults is the response envelope, mirroring
    search.grafana.app SearchResults.
    """
    
    _internal: alertingv0alpha1.SearchResults

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.SearchResults()

    def build(self) -> alertingv0alpha1.SearchResults:
        """
        Builds the object.
        """
        return self._internal    
    
    def metadata(self, metadata: cogbuilder.Builder[alertingv0alpha1.SearchResultsMetadata]) -> typing.Self:    
        metadata_resource = metadata.build()
        self._internal.metadata = metadata_resource
    
        return self
    
    def items(self, items: list[cogbuilder.Builder[alertingv0alpha1.SearchResultHit]]) -> typing.Self:    
        items_resources = [r1.build() for r1 in items]
        self._internal.items = items_resources
    
        return self
    
    def facets(self, facets: dict[str, list[cogbuilder.Builder[alertingv0alpha1.FacetValue]]]) -> typing.Self:    
        facets_resources = { key1: [r2.build() for r2 in val1] for (key1, val1) in facets.items() }
        self._internal.facets = facets_resources
    
        return self
    


class AlertRule(cogbuilder.Builder[alertingv0alpha1.AlertRuleSpec]):
    _internal: alertingv0alpha1.AlertRuleSpec

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.AlertRuleSpec()

    def build(self) -> alertingv0alpha1.AlertRuleSpec:
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
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def annotations(self, annotations: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self:    
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
    
    def no_data_state(self, no_data_state: alertingv0alpha1.NoDataState) -> typing.Self:    
        self._internal.no_data_state = no_data_state
    
        return self
    
    def exec_err_state(self, exec_err_state: alertingv0alpha1.ExecErrState) -> typing.Self:    
        self._internal.exec_err_state = exec_err_state
    
        return self
    
    def notification_settings(self, notification_settings: cogbuilder.Builder[alertingv0alpha1.NotificationSettings]) -> typing.Self:    
        notification_settings_resource = notification_settings.build()
        self._internal.notification_settings = notification_settings_resource
    
        return self
    
    def expressions(self, expressions: alertingv0alpha1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    
    def panel_ref(self, panel_ref: cogbuilder.Builder[alertingv0alpha1.PanelRef]) -> typing.Self:    
        panel_ref_resource = panel_ref.build()
        self._internal.panel_ref = panel_ref_resource
    
        return self
    

"""
Creates a resource manifest from an AlertRule.
"""
def alert_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv0alpha1.AlertRuleSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v0alpha1")
    builder.kind("AlertRule")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder


class RecordingRule(cogbuilder.Builder[alertingv0alpha1.RecordingRuleSpec]):
    _internal: alertingv0alpha1.RecordingRuleSpec

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RecordingRuleSpec()

    def build(self) -> alertingv0alpha1.RecordingRuleSpec:
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
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def labels(self, labels: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def metric(self, metric: alertingv0alpha1.MetricName) -> typing.Self:    
        self._internal.metric = metric
    
        return self
    
    def expressions(self, expressions: alertingv0alpha1.ExpressionMap) -> typing.Self:    
        self._internal.expressions = expressions
    
        return self
    
    def target_datasource_uid(self, target_datasource_uid: alertingv0alpha1.DatasourceUID) -> typing.Self:    
        self._internal.target_datasource_uid = target_datasource_uid
    
        return self
    

"""
Creates a resource manifest from a RecordingRule.
"""
def recording_rule_manifest(name: str, folder: cogbuilder.Builder[alertingv0alpha1.RecordingRuleSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v0alpha1")
    builder.kind("RecordingRule")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder


class RuleSequence(cogbuilder.Builder[alertingv0alpha1.RuleSequenceSpec]):
    _internal: alertingv0alpha1.RuleSequenceSpec

    def __init__(self) -> None:
        self._internal = alertingv0alpha1.RuleSequenceSpec()

    def build(self) -> alertingv0alpha1.RuleSequenceSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def trigger(self, trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self:    
        trigger_resource = trigger.build()
        self._internal.trigger = trigger_resource
    
        return self
    
    def recording_rules(self, recording_rules: list[cogbuilder.Builder[alertingv0alpha1.RuleRef]]) -> typing.Self:    
        recording_rules_resources = [r1.build() for r1 in recording_rules]
        self._internal.recording_rules = recording_rules_resources
    
        return self
    
    def alerting_rules(self, alerting_rules: list[cogbuilder.Builder[alertingv0alpha1.RuleRef]]) -> typing.Self:    
        alerting_rules_resources = [r1.build() for r1 in alerting_rules]
        self._internal.alerting_rules = alerting_rules_resources
    
        return self
    

"""
Creates a resource manifest from a RuleSequence.
"""
def rule_sequence_manifest(name: str, folder: cogbuilder.Builder[alertingv0alpha1.RuleSequenceSpec]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("rules.alerting.grafana.app/v0alpha1")
    builder.kind("RuleSequence")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder
