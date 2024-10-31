# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import alerting
from ..cog import variants as cogvariants


class Query(cogbuilder.Builder[alerting.Query]):    
    _internal: alerting.Query

    def __init__(self, ref_id: str):
        self._internal = alerting.Query()        
        self._internal.ref_id = ref_id

    def build(self) -> alerting.Query:
        """
        Builds the object.
        """
        return self._internal    
    
    def datasource_uid(self, datasource_uid: str) -> typing.Self:    
        """
        Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
        """
            
        self._internal.datasource_uid = datasource_uid
    
        return self
    
    def model(self, model: cogbuilder.Builder[cogvariants.Dataquery]) -> typing.Self:    
        """
        JSON is the raw JSON query and includes the above properties as well as custom properties.
        """
            
        model_resource = model.build()
        self._internal.model = model_resource
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def relative_time_range(self, relative_time_range: alerting.RelativeTimeRange) -> typing.Self:    
        """
        RelativeTimeRange is the per query start and end time
        for requests.
        """
            
        self._internal.relative_time_range = relative_time_range
    
        return self
    

class RuleGroup(cogbuilder.Builder[alerting.RuleGroup]):    
    _internal: alerting.RuleGroup

    def __init__(self, title: str):
        self._internal = alerting.RuleGroup()        
        self._internal.title = title

    def build(self) -> alerting.RuleGroup:
        """
        Builds the object.
        """
        return self._internal    
    
    def folder_uid(self, folder_uid: str) -> typing.Self:        
        self._internal.folder_uid = folder_uid
    
        return self
    
    def interval(self, interval: alerting.Duration) -> typing.Self:    
        """
        The interval, in seconds, at which all rules in the group are evaluated.
        If a group contains many rules, the rules are evaluated sequentially.
        """
            
        self._internal.interval = interval
    
        return self
    
    def rules(self, rules: list[cogbuilder.Builder[alerting.Rule]]) -> typing.Self:        
        rules_resources = [r1.build() for r1 in rules]
        self._internal.rules = rules_resources
    
        return self
    
    def with_rule(self, rule: cogbuilder.Builder[alerting.Rule]) -> typing.Self:        
        if self._internal.rules is None:
            self._internal.rules = []
        
        rule_resource = rule.build()
        self._internal.rules.append(rule_resource)
    
        return self
    
    def title(self, title: str) -> typing.Self:        
        self._internal.title = title
    
        return self
    

class ContactPoint(cogbuilder.Builder[alerting.ContactPoint]):    
    """
    EmbeddedContactPoint is the contact point type that is used
    by grafanas embedded alertmanager implementation.
    """
    
    _internal: alerting.ContactPoint

    def __init__(self):
        self._internal = alerting.ContactPoint()

    def build(self) -> alerting.ContactPoint:
        """
        Builds the object.
        """
        return self._internal    
    
    def disable_resolve_message(self, disable_resolve_message: bool) -> typing.Self:        
        self._internal.disable_resolve_message = disable_resolve_message
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        Name is used as grouping key in the UI. Contact points with the
        same name will be grouped in the UI.
        """
            
        self._internal.name = name
    
        return self
    
    def provenance(self, provenance: str) -> typing.Self:        
        self._internal.provenance = provenance
    
        return self
    
    def settings(self, settings: alerting.Json) -> typing.Self:        
        self._internal.settings = settings
    
        return self
    
    def type_val(self, type_val: typing.Literal["alertmanager", " dingding", " discord", " email", " googlechat", " kafka", " line", " opsgenie", " pagerduty", " pushover", " sensugo", " slack", " teams", " telegram", " threema", " victorops", " webhook", " wecom"]) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        UID is the unique identifier of the contact point. The UID can be
        set by the user.
        """
            
        self._internal.uid = uid
    
        return self
    

class Matcher(cogbuilder.Builder[alerting.Matcher]):    
    _internal: alerting.Matcher

    def __init__(self):
        self._internal = alerting.Matcher()

    def build(self) -> alerting.Matcher:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def type(self, type: alerting.MatchType) -> typing.Self:        
        self._internal.type = type
    
        return self
    
    def value(self, value: str) -> typing.Self:        
        self._internal.value = value
    
        return self
    

class MuteTiming(cogbuilder.Builder[alerting.MuteTiming]):    
    _internal: alerting.MuteTiming

    def __init__(self):
        self._internal = alerting.MuteTiming()

    def build(self) -> alerting.MuteTiming:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def time_intervals(self, time_intervals: list[cogbuilder.Builder[alerting.TimeInterval]]) -> typing.Self:        
        time_intervals_resources = [r1.build() for r1 in time_intervals]
        self._internal.time_intervals = time_intervals_resources
    
        return self
    

class NotificationTemplate(cogbuilder.Builder[alerting.NotificationTemplate]):    
    _internal: alerting.NotificationTemplate

    def __init__(self):
        self._internal = alerting.NotificationTemplate()

    def build(self) -> alerting.NotificationTemplate:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def provenance(self, provenance: alerting.Provenance) -> typing.Self:        
        self._internal.provenance = provenance
    
        return self
    
    def template(self, template: str) -> typing.Self:        
        self._internal.template = template
    
        return self
    

class Rule(cogbuilder.Builder[alerting.Rule]):    
    _internal: alerting.Rule

    def __init__(self, title: str):
        self._internal = alerting.Rule()        
        if not len(title) >= 1:
            raise ValueError("len(title) must be >= 1")
        if not len(title) <= 190:
            raise ValueError("len(title) must be <= 190")
        self._internal.title = title

    def build(self) -> alerting.Rule:
        """
        Builds the object.
        """
        return self._internal    
    
    def annotations(self, annotations: dict[str, str]) -> typing.Self:        
        self._internal.annotations = annotations
    
        return self
    
    def condition(self, condition: str) -> typing.Self:        
        self._internal.condition = condition
    
        return self
    
    def queries(self, data: list[cogbuilder.Builder[alerting.Query]]) -> typing.Self:        
        data_resources = [r1.build() for r1 in data]
        self._internal.data = data_resources
    
        return self
    
    def with_query(self, data: cogbuilder.Builder[alerting.Query]) -> typing.Self:        
        if self._internal.data is None:
            self._internal.data = []
        
        data_resource = data.build()
        self._internal.data.append(data_resource)
    
        return self
    
    def exec_err_state(self, exec_err_state: typing.Literal["OK", "Alerting", "Error"]) -> typing.Self:        
        self._internal.exec_err_state = exec_err_state
    
        return self
    
    def folder_uid(self, folder_uid: str) -> typing.Self:        
        self._internal.folder_uid = folder_uid
    
        return self
    
    def for_val(self, for_val: str) -> typing.Self:    
        """
        The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
        Before this time has elapsed, the rule is only considered to be Pending.
        """
            
        self._internal.for_val = for_val
    
        return self
    
    def id_val(self, id_val: int) -> typing.Self:        
        self._internal.id_val = id_val
    
        return self
    
    def is_paused(self, is_paused: bool) -> typing.Self:        
        self._internal.is_paused = is_paused
    
        return self
    
    def labels(self, labels: dict[str, str]) -> typing.Self:        
        self._internal.labels = labels
    
        return self
    
    def no_data_state(self, no_data_state: typing.Literal["Alerting", "NoData", "OK"]) -> typing.Self:        
        self._internal.no_data_state = no_data_state
    
        return self
    
    def org_id(self, org_id: int) -> typing.Self:        
        self._internal.org_id = org_id
    
        return self
    
    def provenance(self, provenance: alerting.Provenance) -> typing.Self:        
        self._internal.provenance = provenance
    
        return self
    
    def rule_group(self, rule_group: str) -> typing.Self:        
        if not len(rule_group) >= 1:
            raise ValueError("len(rule_group) must be >= 1")
        if not len(rule_group) <= 190:
            raise ValueError("len(rule_group) must be <= 190")
        self._internal.rule_group = rule_group
    
        return self
    
    def title(self, title: str) -> typing.Self:        
        if not len(title) >= 1:
            raise ValueError("len(title) must be >= 1")
        if not len(title) <= 190:
            raise ValueError("len(title) must be <= 190")
        self._internal.title = title
    
        return self
    
    def uid(self, uid: str) -> typing.Self:        
        self._internal.uid = uid
    
        return self
    
    def updated(self, updated: str) -> typing.Self:        
        self._internal.updated = updated
    
        return self
    

class NotificationPolicy(cogbuilder.Builder[alerting.NotificationPolicy]):    
    """
    A Route is a node that contains definitions of how to handle alerts. This is modified
    from the upstream alertmanager in that it adds the ObjectMatchers property.
    """
    
    _internal: alerting.NotificationPolicy

    def __init__(self):
        self._internal = alerting.NotificationPolicy()

    def build(self) -> alerting.NotificationPolicy:
        """
        Builds the object.
        """
        return self._internal    
    
    def continue_val(self, continue_val: bool) -> typing.Self:        
        self._internal.continue_val = continue_val
    
        return self
    
    def group_by(self, group_by: list[str]) -> typing.Self:        
        self._internal.group_by = group_by
    
        return self
    
    def group_interval(self, group_interval: str) -> typing.Self:        
        self._internal.group_interval = group_interval
    
        return self
    
    def group_wait(self, group_wait: str) -> typing.Self:        
        self._internal.group_wait = group_wait
    
        return self
    
    def match(self, match: dict[str, str]) -> typing.Self:    
        """
        Deprecated. Remove before v1.0 release.
        """
            
        self._internal.match = match
    
        return self
    
    def match_re(self, match_re: alerting.MatchRegexps) -> typing.Self:        
        self._internal.match_re = match_re
    
        return self
    
    def matchers(self, matchers: alerting.Matchers) -> typing.Self:    
        """
        Matchers is a slice of Matchers that is sortable, implements Stringer, and
        provides a Matches method to match a LabelSet against all Matchers in the
        slice. Note that some users of Matchers might require it to be sorted.
        """
            
        self._internal.matchers = matchers
    
        return self
    
    def mute_time_intervals(self, mute_time_intervals: list[str]) -> typing.Self:        
        self._internal.mute_time_intervals = mute_time_intervals
    
        return self
    
    def object_matchers(self, object_matchers: alerting.ObjectMatchers) -> typing.Self:    
        """
        Matchers is a slice of Matchers that is sortable, implements Stringer, and
        provides a Matches method to match a LabelSet against all Matchers in the
        slice. Note that some users of Matchers might require it to be sorted.
        """
            
        self._internal.object_matchers = object_matchers
    
        return self
    
    def provenance(self, provenance: alerting.Provenance) -> typing.Self:        
        self._internal.provenance = provenance
    
        return self
    
    def receiver(self, receiver: str) -> typing.Self:        
        self._internal.receiver = receiver
    
        return self
    
    def repeat_interval(self, repeat_interval: str) -> typing.Self:        
        self._internal.repeat_interval = repeat_interval
    
        return self
    
    def routes(self, routes: list[cogbuilder.Builder[alerting.NotificationPolicy]]) -> typing.Self:        
        routes_resources = [r1.build() for r1 in routes]
        self._internal.routes = routes_resources
    
        return self
    

class TimeInterval(cogbuilder.Builder[alerting.TimeInterval]):    
    """
    TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    within the interval.
    """
    
    _internal: alerting.TimeInterval

    def __init__(self):
        self._internal = alerting.TimeInterval()

    def build(self) -> alerting.TimeInterval:
        """
        Builds the object.
        """
        return self._internal    
    
    def days_of_month(self, days_of_month: list[str]) -> typing.Self:        
        self._internal.days_of_month = days_of_month
    
        return self
    
    def location(self, location: str) -> typing.Self:        
        self._internal.location = location
    
        return self
    
    def months(self, months: list[str]) -> typing.Self:        
        self._internal.months = months
    
        return self
    
    def times(self, times: list[cogbuilder.Builder[alerting.TimeRange]]) -> typing.Self:        
        times_resources = [r1.build() for r1 in times]
        self._internal.times = times_resources
    
        return self
    
    def weekdays(self, weekdays: list[str]) -> typing.Self:        
        self._internal.weekdays = weekdays
    
        return self
    
    def years(self, years: list[str]) -> typing.Self:        
        self._internal.years = years
    
        return self
    

class TimeRange(cogbuilder.Builder[alerting.TimeRange]):    
    """
    Redefining this to avoid an import cycle
    """
    
    _internal: alerting.TimeRange

    def __init__(self):
        self._internal = alerting.TimeRange()

    def build(self) -> alerting.TimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:        
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:        
        self._internal.to = to
    
        return self
    