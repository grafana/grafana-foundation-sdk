# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import variants as cogvariants
from ..cog import runtime as cogruntime
import enum


class Query:
    # Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    datasource_uid: typing.Optional[str]
    # JSON is the raw JSON query and includes the above properties as well as custom properties.
    model: typing.Optional[cogvariants.Dataquery]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # RelativeTimeRange is the per query start and end time
    # for requests.
    relative_time_range: typing.Optional['RelativeTimeRange']

    def __init__(self, datasource_uid: typing.Optional[str] = None, model: typing.Optional[cogvariants.Dataquery] = None, query_type: typing.Optional[str] = None, ref_id: typing.Optional[str] = None, relative_time_range: typing.Optional['RelativeTimeRange'] = None):
        self.datasource_uid = datasource_uid
        self.model = model
        self.query_type = query_type
        self.ref_id = ref_id
        self.relative_time_range = relative_time_range

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.datasource_uid is not None:
            payload["datasourceUid"] = self.datasource_uid
        if self.model is not None:
            payload["model"] = self.model
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.relative_time_range is not None:
            payload["relativeTimeRange"] = self.relative_time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "datasourceUid" in data:
            args["datasource_uid"] = data["datasourceUid"]
        if "model" in data:
            args["model"] = cogruntime.dataquery_from_json(data["model"], "")
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "relativeTimeRange" in data:
            args["relative_time_range"] = RelativeTimeRange.from_json(data["relativeTimeRange"])        

        return cls(**args)


class RuleGroup:
    folder_uid: typing.Optional[str]
    # The interval, in seconds, at which all rules in the group are evaluated.
    # If a group contains many rules, the rules are evaluated sequentially.
    interval: typing.Optional['Duration']
    rules: typing.Optional[list['Rule']]
    title: typing.Optional[str]

    def __init__(self, folder_uid: typing.Optional[str] = None, interval: typing.Optional['Duration'] = None, rules: typing.Optional[list['Rule']] = None, title: typing.Optional[str] = None):
        self.folder_uid = folder_uid
        self.interval = interval
        self.rules = rules
        self.title = title

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.folder_uid is not None:
            payload["folderUid"] = self.folder_uid
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.rules is not None:
            payload["rules"] = self.rules
        if self.title is not None:
            payload["title"] = self.title
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "folderUid" in data:
            args["folder_uid"] = data["folderUid"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "rules" in data:
            args["rules"] = data["rules"]
        if "title" in data:
            args["title"] = data["title"]        

        return cls(**args)


# Duration in seconds.
Duration: typing.TypeAlias = int


class ContactPoint:
    """
    EmbeddedContactPoint is the contact point type that is used
    by grafanas embedded alertmanager implementation.
    """

    disable_resolve_message: typing.Optional[bool]
    # Name is used as grouping key in the UI. Contact points with the
    # same name will be grouped in the UI.
    name: typing.Optional[str]
    provenance: typing.Optional[str]
    settings: 'Json'
    type_val: typing.Literal["alertmanager", " dingding", " discord", " email", " googlechat", " kafka", " line", " opsgenie", " pagerduty", " pushover", " sensugo", " slack", " teams", " telegram", " threema", " victorops", " webhook", " wecom"]
    # UID is the unique identifier of the contact point. The UID can be
    # set by the user.
    uid: typing.Optional[str]

    def __init__(self, disable_resolve_message: typing.Optional[bool] = None, name: typing.Optional[str] = None, provenance: typing.Optional[str] = None, settings: typing.Optional['Json'] = None, type_val: typing.Optional[typing.Literal["alertmanager", " dingding", " discord", " email", " googlechat", " kafka", " line", " opsgenie", " pagerduty", " pushover", " sensugo", " slack", " teams", " telegram", " threema", " victorops", " webhook", " wecom"]] = None, uid: typing.Optional[str] = None):
        self.disable_resolve_message = disable_resolve_message
        self.name = name
        self.provenance = provenance
        self.settings = settings if settings is not None else Json()
        self.type_val = type_val if type_val is not None else "alertmanager"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "settings": self.settings,
            "type": self.type_val,
        }
        if self.disable_resolve_message is not None:
            payload["disableResolveMessage"] = self.disable_resolve_message
        if self.name is not None:
            payload["name"] = self.name
        if self.provenance is not None:
            payload["provenance"] = self.provenance
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "disableResolveMessage" in data:
            args["disable_resolve_message"] = data["disableResolveMessage"]
        if "name" in data:
            args["name"] = data["name"]
        if "provenance" in data:
            args["provenance"] = data["provenance"]
        if "settings" in data:
            args["settings"] = data["settings"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


Json: typing.TypeAlias = object


MatchRegexps: typing.TypeAlias = dict[str, str]


class MatchType(enum.StrEnum):
    EQUAL = "="
    NOT_EQUAL = "!="
    EQUAL_REGEX = "=~"
    NOT_EQUAL_REGEX = "!~"


class Matcher:
    name: typing.Optional[str]
    type: typing.Optional['MatchType']
    value: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None, type: typing.Optional['MatchType'] = None, value: typing.Optional[str] = None):
        self.name = name
        self.type = type
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["Name"] = self.name
        if self.type is not None:
            payload["Type"] = self.type
        if self.value is not None:
            payload["Value"] = self.value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "Name" in data:
            args["name"] = data["Name"]
        if "Type" in data:
            args["type"] = data["Type"]
        if "Value" in data:
            args["value"] = data["Value"]        

        return cls(**args)


# Matchers is a slice of Matchers that is sortable, implements Stringer, and
# provides a Matches method to match a LabelSet against all Matchers in the
# slice. Note that some users of Matchers might require it to be sorted.
Matchers: typing.TypeAlias = list['Matcher']


class MuteTiming:
    name: typing.Optional[str]
    time_intervals: typing.Optional[list['TimeInterval']]

    def __init__(self, name: typing.Optional[str] = None, time_intervals: typing.Optional[list['TimeInterval']] = None):
        self.name = name
        self.time_intervals = time_intervals

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.time_intervals is not None:
            payload["time_intervals"] = self.time_intervals
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "time_intervals" in data:
            args["time_intervals"] = data["time_intervals"]        

        return cls(**args)


class NotificationTemplate:
    name: typing.Optional[str]
    provenance: typing.Optional['Provenance']
    template: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None, provenance: typing.Optional['Provenance'] = None, template: typing.Optional[str] = None):
        self.name = name
        self.provenance = provenance
        self.template = template

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.provenance is not None:
            payload["provenance"] = self.provenance
        if self.template is not None:
            payload["template"] = self.template
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "provenance" in data:
            args["provenance"] = data["provenance"]
        if "template" in data:
            args["template"] = data["template"]        

        return cls(**args)


ObjectMatcher: typing.TypeAlias = list[str]


ObjectMatchers: typing.TypeAlias = list['ObjectMatcher']


Provenance: typing.TypeAlias = str


class Rule:
    annotations: typing.Optional[dict[str, str]]
    condition: str
    data: list['Query']
    exec_err_state: typing.Literal["OK", "Alerting", "Error"]
    folder_uid: str
    # The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    # Before this time has elapsed, the rule is only considered to be Pending.
    for_val: str
    id_val: typing.Optional[int]
    is_paused: typing.Optional[bool]
    labels: typing.Optional[dict[str, str]]
    no_data_state: typing.Literal["Alerting", "NoData", "OK"]
    org_id: int
    provenance: typing.Optional['Provenance']
    rule_group: str
    title: str
    uid: typing.Optional[str]
    updated: typing.Optional[str]

    def __init__(self, annotations: typing.Optional[dict[str, str]] = None, condition: str = "", data: typing.Optional[list['Query']] = None, exec_err_state: typing.Optional[typing.Literal["OK", "Alerting", "Error"]] = None, folder_uid: str = "", for_val: str = "", id_val: typing.Optional[int] = None, is_paused: typing.Optional[bool] = None, labels: typing.Optional[dict[str, str]] = None, no_data_state: typing.Optional[typing.Literal["Alerting", "NoData", "OK"]] = None, org_id: int = 0, provenance: typing.Optional['Provenance'] = None, rule_group: str = "", title: str = "", uid: typing.Optional[str] = None, updated: typing.Optional[str] = None):
        self.annotations = annotations
        self.condition = condition
        self.data = data if data is not None else []
        self.exec_err_state = exec_err_state if exec_err_state is not None else "OK"
        self.folder_uid = folder_uid
        self.for_val = for_val
        self.id_val = id_val
        self.is_paused = is_paused
        self.labels = labels
        self.no_data_state = no_data_state if no_data_state is not None else "Alerting"
        self.org_id = org_id
        self.provenance = provenance
        self.rule_group = rule_group
        self.title = title
        self.uid = uid
        self.updated = updated

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "condition": self.condition,
            "data": self.data,
            "execErrState": self.exec_err_state,
            "folderUID": self.folder_uid,
            "for": self.for_val,
            "noDataState": self.no_data_state,
            "orgID": self.org_id,
            "ruleGroup": self.rule_group,
            "title": self.title,
        }
        if self.annotations is not None:
            payload["annotations"] = self.annotations
        if self.id_val is not None:
            payload["id"] = self.id_val
        if self.is_paused is not None:
            payload["isPaused"] = self.is_paused
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.provenance is not None:
            payload["provenance"] = self.provenance
        if self.uid is not None:
            payload["uid"] = self.uid
        if self.updated is not None:
            payload["updated"] = self.updated
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "annotations" in data:
            args["annotations"] = data["annotations"]
        if "condition" in data:
            args["condition"] = data["condition"]
        if "data" in data:
            args["data"] = data["data"]
        if "execErrState" in data:
            args["exec_err_state"] = data["execErrState"]
        if "folderUID" in data:
            args["folder_uid"] = data["folderUID"]
        if "for" in data:
            args["for_val"] = data["for"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "isPaused" in data:
            args["is_paused"] = data["isPaused"]
        if "labels" in data:
            args["labels"] = data["labels"]
        if "noDataState" in data:
            args["no_data_state"] = data["noDataState"]
        if "orgID" in data:
            args["org_id"] = data["orgID"]
        if "provenance" in data:
            args["provenance"] = data["provenance"]
        if "ruleGroup" in data:
            args["rule_group"] = data["ruleGroup"]
        if "title" in data:
            args["title"] = data["title"]
        if "uid" in data:
            args["uid"] = data["uid"]
        if "updated" in data:
            args["updated"] = data["updated"]        

        return cls(**args)


class RelativeTimeRange:
    """
    RelativeTimeRange is the per query start and end time
    for requests.
    """

    # A Duration represents the elapsed time between two instants
    # as an int64 nanosecond count. The representation limits the
    # largest representable duration to approximately 290 years.
    from_val: typing.Optional['Duration']
    # A Duration represents the elapsed time between two instants
    # as an int64 nanosecond count. The representation limits the
    # largest representable duration to approximately 290 years.
    to: typing.Optional['Duration']

    def __init__(self, from_val: typing.Optional['Duration'] = None, to: typing.Optional['Duration'] = None):
        self.from_val = from_val
        self.to = to

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.from_val is not None:
            payload["from"] = self.from_val
        if self.to is not None:
            payload["to"] = self.to
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]        

        return cls(**args)


class NotificationPolicy:
    """
    A Route is a node that contains definitions of how to handle alerts. This is modified
    from the upstream alertmanager in that it adds the ObjectMatchers property.
    """

    continue_val: typing.Optional[bool]
    group_by: typing.Optional[list[str]]
    group_interval: typing.Optional[str]
    group_wait: typing.Optional[str]
    # Deprecated. Remove before v1.0 release.
    match: typing.Optional[dict[str, str]]
    match_re: typing.Optional['MatchRegexps']
    # Matchers is a slice of Matchers that is sortable, implements Stringer, and
    # provides a Matches method to match a LabelSet against all Matchers in the
    # slice. Note that some users of Matchers might require it to be sorted.
    matchers: typing.Optional['Matchers']
    mute_time_intervals: typing.Optional[list[str]]
    object_matchers: typing.Optional['ObjectMatchers']
    provenance: typing.Optional['Provenance']
    receiver: typing.Optional[str]
    repeat_interval: typing.Optional[str]
    routes: typing.Optional[list['NotificationPolicy']]

    def __init__(self, continue_val: typing.Optional[bool] = None, group_by: typing.Optional[list[str]] = None, group_interval: typing.Optional[str] = None, group_wait: typing.Optional[str] = None, match: typing.Optional[dict[str, str]] = None, match_re: typing.Optional['MatchRegexps'] = None, matchers: typing.Optional['Matchers'] = None, mute_time_intervals: typing.Optional[list[str]] = None, object_matchers: typing.Optional['ObjectMatchers'] = None, provenance: typing.Optional['Provenance'] = None, receiver: typing.Optional[str] = None, repeat_interval: typing.Optional[str] = None, routes: typing.Optional[list['NotificationPolicy']] = None):
        self.continue_val = continue_val
        self.group_by = group_by
        self.group_interval = group_interval
        self.group_wait = group_wait
        self.match = match
        self.match_re = match_re
        self.matchers = matchers
        self.mute_time_intervals = mute_time_intervals
        self.object_matchers = object_matchers
        self.provenance = provenance
        self.receiver = receiver
        self.repeat_interval = repeat_interval
        self.routes = routes

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.continue_val is not None:
            payload["continue"] = self.continue_val
        if self.group_by is not None:
            payload["group_by"] = self.group_by
        if self.group_interval is not None:
            payload["group_interval"] = self.group_interval
        if self.group_wait is not None:
            payload["group_wait"] = self.group_wait
        if self.match is not None:
            payload["match"] = self.match
        if self.match_re is not None:
            payload["match_re"] = self.match_re
        if self.matchers is not None:
            payload["matchers"] = self.matchers
        if self.mute_time_intervals is not None:
            payload["mute_time_intervals"] = self.mute_time_intervals
        if self.object_matchers is not None:
            payload["object_matchers"] = self.object_matchers
        if self.provenance is not None:
            payload["provenance"] = self.provenance
        if self.receiver is not None:
            payload["receiver"] = self.receiver
        if self.repeat_interval is not None:
            payload["repeat_interval"] = self.repeat_interval
        if self.routes is not None:
            payload["routes"] = self.routes
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "continue" in data:
            args["continue_val"] = data["continue"]
        if "group_by" in data:
            args["group_by"] = data["group_by"]
        if "group_interval" in data:
            args["group_interval"] = data["group_interval"]
        if "group_wait" in data:
            args["group_wait"] = data["group_wait"]
        if "match" in data:
            args["match"] = data["match"]
        if "match_re" in data:
            args["match_re"] = data["match_re"]
        if "matchers" in data:
            args["matchers"] = data["matchers"]
        if "mute_time_intervals" in data:
            args["mute_time_intervals"] = data["mute_time_intervals"]
        if "object_matchers" in data:
            args["object_matchers"] = data["object_matchers"]
        if "provenance" in data:
            args["provenance"] = data["provenance"]
        if "receiver" in data:
            args["receiver"] = data["receiver"]
        if "repeat_interval" in data:
            args["repeat_interval"] = data["repeat_interval"]
        if "routes" in data:
            args["routes"] = data["routes"]        

        return cls(**args)


class TimeInterval:
    """
    TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    within the interval.
    """

    days_of_month: typing.Optional[list[str]]
    location: typing.Optional[str]
    months: typing.Optional[list[str]]
    times: typing.Optional[list['TimeRange']]
    weekdays: typing.Optional[list[str]]
    years: typing.Optional[list[str]]

    def __init__(self, days_of_month: typing.Optional[list[str]] = None, location: typing.Optional[str] = None, months: typing.Optional[list[str]] = None, times: typing.Optional[list['TimeRange']] = None, weekdays: typing.Optional[list[str]] = None, years: typing.Optional[list[str]] = None):
        self.days_of_month = days_of_month
        self.location = location
        self.months = months
        self.times = times
        self.weekdays = weekdays
        self.years = years

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.days_of_month is not None:
            payload["days_of_month"] = self.days_of_month
        if self.location is not None:
            payload["location"] = self.location
        if self.months is not None:
            payload["months"] = self.months
        if self.times is not None:
            payload["times"] = self.times
        if self.weekdays is not None:
            payload["weekdays"] = self.weekdays
        if self.years is not None:
            payload["years"] = self.years
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "days_of_month" in data:
            args["days_of_month"] = data["days_of_month"]
        if "location" in data:
            args["location"] = data["location"]
        if "months" in data:
            args["months"] = data["months"]
        if "times" in data:
            args["times"] = data["times"]
        if "weekdays" in data:
            args["weekdays"] = data["weekdays"]
        if "years" in data:
            args["years"] = data["years"]        

        return cls(**args)


class TimeIntervalItem:
    days_of_month: typing.Optional[list[str]]
    location: typing.Optional[str]
    months: typing.Optional[list[str]]
    times: typing.Optional[list['TimeIntervalTimeRange']]
    weekdays: typing.Optional[list[str]]
    years: typing.Optional[list[str]]

    def __init__(self, days_of_month: typing.Optional[list[str]] = None, location: typing.Optional[str] = None, months: typing.Optional[list[str]] = None, times: typing.Optional[list['TimeIntervalTimeRange']] = None, weekdays: typing.Optional[list[str]] = None, years: typing.Optional[list[str]] = None):
        self.days_of_month = days_of_month
        self.location = location
        self.months = months
        self.times = times
        self.weekdays = weekdays
        self.years = years

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.days_of_month is not None:
            payload["days_of_month"] = self.days_of_month
        if self.location is not None:
            payload["location"] = self.location
        if self.months is not None:
            payload["months"] = self.months
        if self.times is not None:
            payload["times"] = self.times
        if self.weekdays is not None:
            payload["weekdays"] = self.weekdays
        if self.years is not None:
            payload["years"] = self.years
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "days_of_month" in data:
            args["days_of_month"] = data["days_of_month"]
        if "location" in data:
            args["location"] = data["location"]
        if "months" in data:
            args["months"] = data["months"]
        if "times" in data:
            args["times"] = data["times"]
        if "weekdays" in data:
            args["weekdays"] = data["weekdays"]
        if "years" in data:
            args["years"] = data["years"]        

        return cls(**args)


class TimeIntervalTimeRange:
    end_time: typing.Optional[str]
    start_time: typing.Optional[str]

    def __init__(self, end_time: typing.Optional[str] = None, start_time: typing.Optional[str] = None):
        self.end_time = end_time
        self.start_time = start_time

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.end_time is not None:
            payload["end_time"] = self.end_time
        if self.start_time is not None:
            payload["start_time"] = self.start_time
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "end_time" in data:
            args["end_time"] = data["end_time"]
        if "start_time" in data:
            args["start_time"] = data["start_time"]        

        return cls(**args)


class TimeRange:
    """
    Redefining this to avoid an import cycle
    """

    from_val: typing.Optional[str]
    to: typing.Optional[str]

    def __init__(self, from_val: typing.Optional[str] = None, to: typing.Optional[str] = None):
        self.from_val = from_val
        self.to = to

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.from_val is not None:
            payload["from"] = self.from_val
        if self.to is not None:
            payload["to"] = self.to
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]        

        return cls(**args)



