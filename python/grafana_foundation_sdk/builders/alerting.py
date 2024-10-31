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
    

class NotificationSettings(cogbuilder.Builder[alerting.NotificationSettings]):    
    _internal: alerting.NotificationSettings

    def __init__(self):
        self._internal = alerting.NotificationSettings()

    def build(self) -> alerting.NotificationSettings:
        return self._internal    
    
    def group_by(self, group_by: list[str]) -> typing.Self:    
        """
        Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
        cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
        use the special value '...' as the sole label name.
        This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
        you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
        Must include 'alertname' and 'grafana_folder' if not using '...'.
        """
            
        self._internal.group_by = group_by
    
        return self
    
    def group_interval(self, group_interval: str) -> typing.Self:    
        """
        Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
        which an initial notification has already been sent. (Usually ~5m or more.)
        """
            
        self._internal.group_interval = group_interval
    
        return self
    
    def group_wait(self, group_wait: str) -> typing.Self:    
        """
        Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
        inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
        """
            
        self._internal.group_wait = group_wait
    
        return self
    
    def mute_time_intervals(self, mute_time_intervals: list[str]) -> typing.Self:    
        """
        Override the times when notifications should be muted. These must match the name of a mute time interval defined
        in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
        otherwise acts normally.
        """
            
        self._internal.mute_time_intervals = mute_time_intervals
    
        return self
    
    def receiver(self, receiver: str) -> typing.Self:    
        """
        Name of the receiver to send notifications to.
        """
            
        self._internal.receiver = receiver
    
        return self
    
    def repeat_interval(self, repeat_interval: str) -> typing.Self:    
        """
        Override how long to wait before sending a notification again if it has already been sent successfully for an
        alert. (Usually ~3h or more).
        Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
        Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
        occurs first. `repeat_interval` should not be less than `group_interval`.
        """
            
        self._internal.repeat_interval = repeat_interval
    
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
            
        if not len(uid) >= 1:
            raise ValueError("len(uid) must be >= 1")
        if not len(uid) <= 40:
            raise ValueError("len(uid) must be <= 40")
        self._internal.uid = uid
    
        return self
    

class Matcher(cogbuilder.Builder[alerting.Matcher]):    
    _internal: alerting.Matcher

    def __init__(self):
        self._internal = alerting.Matcher()

    def build(self) -> alerting.Matcher:
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
    
    def notification_settings(self, notification_settings: cogbuilder.Builder[alerting.NotificationSettings]) -> typing.Self:        
        notification_settings_resource = notification_settings.build()
        self._internal.notification_settings = notification_settings_resource
    
        return self
    
    def org_id(self, org_id: int) -> typing.Self:        
        self._internal.org_id = org_id
    
        return self
    
    def provenance(self, provenance: alerting.Provenance) -> typing.Self:        
        self._internal.provenance = provenance
    
        return self
    
    def record(self, record: cogbuilder.Builder[alerting.RecordRule]) -> typing.Self:        
        record_resource = record.build()
        self._internal.record = record_resource
    
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
        if not len(uid) >= 1:
            raise ValueError("len(uid) must be >= 1")
        if not len(uid) <= 40:
            raise ValueError("len(uid) must be <= 40")
        self._internal.uid = uid
    
        return self
    
    def updated(self, updated: str) -> typing.Self:        
        self._internal.updated = updated
    
        return self
    

class RecordRule(cogbuilder.Builder[alerting.RecordRule]):    
    _internal: alerting.RecordRule

    def __init__(self):
        self._internal = alerting.RecordRule()

    def build(self) -> alerting.RecordRule:
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        Which expression node should be used as the input for the recorded metric.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def metric(self, metric: str) -> typing.Self:    
        """
        Name of the recorded metric.
        """
            
        self._internal.metric = metric
    
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
    _internal: alerting.TimeInterval

    def __init__(self):
        self._internal = alerting.TimeInterval()

    def build(self) -> alerting.TimeInterval:
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def time_intervals(self, time_intervals: list[cogbuilder.Builder[alerting.TimeIntervalItem]]) -> typing.Self:        
        time_intervals_resources = [r1.build() for r1 in time_intervals]
        self._internal.time_intervals = time_intervals_resources
    
        return self
    

class TimeIntervalItem(cogbuilder.Builder[alerting.TimeIntervalItem]):    
    _internal: alerting.TimeIntervalItem

    def __init__(self):
        self._internal = alerting.TimeIntervalItem()

    def build(self) -> alerting.TimeIntervalItem:
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
    
    def times(self, times: list[cogbuilder.Builder[alerting.TimeIntervalTimeRange]]) -> typing.Self:        
        times_resources = [r1.build() for r1 in times]
        self._internal.times = times_resources
    
        return self
    
    def weekdays(self, weekdays: list[str]) -> typing.Self:        
        self._internal.weekdays = weekdays
    
        return self
    
    def years(self, years: list[str]) -> typing.Self:        
        self._internal.years = years
    
        return self
    

class TimeIntervalTimeRange(cogbuilder.Builder[alerting.TimeIntervalTimeRange]):    
    _internal: alerting.TimeIntervalTimeRange

    def __init__(self):
        self._internal = alerting.TimeIntervalTimeRange()

    def build(self) -> alerting.TimeIntervalTimeRange:
        return self._internal    
    
    def end_time(self, end_time: str) -> typing.Self:        
        self._internal.end_time = end_time
    
        return self
    
    def start_time(self, start_time: str) -> typing.Self:        
        self._internal.start_time = start_time
    
        return self
    