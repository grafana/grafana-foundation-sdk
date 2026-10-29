# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing


class NoDataState(enum.StrEnum):
    NO_DATA = "NoData"
    OK = "Ok"
    ALERTING = "Alerting"
    KEEP_LAST = "KeepLast"


class ExecErrState(enum.StrEnum):
    ERROR = "Error"
    OK = "Ok"
    ALERTING = "Alerting"
    KEEP_LAST = "KeepLast"


TimeIntervalRef: typing.TypeAlias = str


class AlertRuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: 'IntervalTrigger'
    labels: typing.Optional[dict[str, 'TemplateString']]
    annotations: typing.Optional[dict[str, 'TemplateString']]
    for_val: typing.Optional[str]
    keep_firing_for: typing.Optional[str]
    missing_series_evals_to_resolve: typing.Optional[int]
    no_data_state: 'NoDataState'
    exec_err_state: 'ExecErrState'
    notification_settings: typing.Optional['NotificationSettings']
    expressions: 'ExpressionMap'
    panel_ref: typing.Optional['PanelRef']

    def __init__(self, title: str = "", paused: typing.Optional[bool] = None, trigger: typing.Optional['IntervalTrigger'] = None, labels: typing.Optional[dict[str, 'TemplateString']] = None, annotations: typing.Optional[dict[str, 'TemplateString']] = None, for_val: typing.Optional[str] = None, keep_firing_for: typing.Optional[str] = None, missing_series_evals_to_resolve: typing.Optional[int] = None, no_data_state: typing.Optional['NoDataState'] = None, exec_err_state: typing.Optional['ExecErrState'] = None, notification_settings: typing.Optional['NotificationSettings'] = None, expressions: typing.Optional['ExpressionMap'] = None, panel_ref: typing.Optional['PanelRef'] = None) -> None:
        self.title = title
        self.paused = paused
        self.trigger = trigger if trigger is not None else IntervalTrigger()
        self.labels = labels
        self.annotations = annotations
        self.for_val = for_val
        self.keep_firing_for = keep_firing_for
        self.missing_series_evals_to_resolve = missing_series_evals_to_resolve
        self.no_data_state = no_data_state if no_data_state is not None else NoDataState.NO_DATA
        self.exec_err_state = exec_err_state if exec_err_state is not None else ExecErrState.ERROR
        self.notification_settings = notification_settings
        self.expressions = expressions if expressions is not None else ExpressionMap()
        self.panel_ref = panel_ref

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "trigger": self.trigger,
            "noDataState": self.no_data_state,
            "execErrState": self.exec_err_state,
            "expressions": self.expressions,
        }
        if self.paused is not None:
            payload["paused"] = self.paused
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.annotations is not None:
            payload["annotations"] = self.annotations
        if self.for_val is not None:
            payload["for"] = self.for_val
        if self.keep_firing_for is not None:
            payload["keepFiringFor"] = self.keep_firing_for
        if self.missing_series_evals_to_resolve is not None:
            payload["missingSeriesEvalsToResolve"] = self.missing_series_evals_to_resolve
        if self.notification_settings is not None:
            payload["notificationSettings"] = self.notification_settings
        if self.panel_ref is not None:
            payload["panelRef"] = self.panel_ref
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "paused" in data:
            args["paused"] = data["paused"]
        if "trigger" in data:
            args["trigger"] = IntervalTrigger.from_json(data["trigger"])
        if "labels" in data:
            args["labels"] = {key: data["labels"][key] for key in data["labels"].keys()}
        if "annotations" in data:
            args["annotations"] = {key: data["annotations"][key] for key in data["annotations"].keys()}
        if "for" in data:
            args["for_val"] = data["for"]
        if "keepFiringFor" in data:
            args["keep_firing_for"] = data["keepFiringFor"]
        if "missingSeriesEvalsToResolve" in data:
            args["missing_series_evals_to_resolve"] = data["missingSeriesEvalsToResolve"]
        if "noDataState" in data:
            args["no_data_state"] = data["noDataState"]
        if "execErrState" in data:
            args["exec_err_state"] = data["execErrState"]
        if "notificationSettings" in data:
            decoding_map_notificationSettings_ref_union: dict[str, typing.Union[typing.Type[NamedRoutingTree], typing.Type[SimplifiedRouting]]] = {"NamedRoutingTree": NamedRoutingTree, "SimplifiedRouting": SimplifiedRouting}
            args["notification_settings"] = decoding_map_notificationSettings_ref_union[data["notificationSettings"]["type"]].from_json(data["notificationSettings"])
        if "expressions" in data:
            args["expressions"] = {key: Expression.from_json(data["expressions"][key]) for key in data["expressions"].keys()}
        if "panelRef" in data:
            args["panel_ref"] = PanelRef.from_json(data["panelRef"])        

        return cls(**args)


class IntervalTrigger:
    interval: 'PromDuration'

    def __init__(self, interval: typing.Optional['PromDuration'] = None) -> None:
        self.interval = interval if interval is not None else PromDuration()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "interval": self.interval,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "interval" in data:
            args["interval"] = data["interval"]        

        return cls(**args)


PromDuration: typing.TypeAlias = str


TemplateString: typing.TypeAlias = str


NotificationSettings: typing.TypeAlias = typing.Union['SimplifiedRouting', 'NamedRoutingTree']


class SimplifiedRouting:
    type_val: str
    receiver: str
    group_by: typing.Optional[list[str]]
    group_wait: typing.Optional['PromDuration']
    group_interval: typing.Optional['PromDuration']
    repeat_interval: typing.Optional['PromDuration']
    mute_time_intervals: typing.Optional[list['TimeIntervalRef']]
    active_time_intervals: typing.Optional[list['TimeIntervalRef']]

    def __init__(self, receiver: str = "", group_by: typing.Optional[list[str]] = None, group_wait: typing.Optional['PromDuration'] = None, group_interval: typing.Optional['PromDuration'] = None, repeat_interval: typing.Optional['PromDuration'] = None, mute_time_intervals: typing.Optional[list['TimeIntervalRef']] = None, active_time_intervals: typing.Optional[list['TimeIntervalRef']] = None) -> None:
        self.type_val = NotificationSettingsType.SIMPLIFIED_ROUTING
        self.receiver = receiver
        self.group_by = group_by
        self.group_wait = group_wait
        self.group_interval = group_interval
        self.repeat_interval = repeat_interval
        self.mute_time_intervals = mute_time_intervals
        self.active_time_intervals = active_time_intervals

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "receiver": self.receiver,
        }
        if self.group_by is not None:
            payload["groupBy"] = self.group_by
        if self.group_wait is not None:
            payload["groupWait"] = self.group_wait
        if self.group_interval is not None:
            payload["groupInterval"] = self.group_interval
        if self.repeat_interval is not None:
            payload["repeatInterval"] = self.repeat_interval
        if self.mute_time_intervals is not None:
            payload["muteTimeIntervals"] = self.mute_time_intervals
        if self.active_time_intervals is not None:
            payload["activeTimeIntervals"] = self.active_time_intervals
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "receiver" in data:
            args["receiver"] = data["receiver"]
        if "groupBy" in data:
            args["group_by"] = data["groupBy"]
        if "groupWait" in data:
            args["group_wait"] = data["groupWait"]
        if "groupInterval" in data:
            args["group_interval"] = data["groupInterval"]
        if "repeatInterval" in data:
            args["repeat_interval"] = data["repeatInterval"]
        if "muteTimeIntervals" in data:
            args["mute_time_intervals"] = [item for item in data["muteTimeIntervals"]]
        if "activeTimeIntervals" in data:
            args["active_time_intervals"] = [item for item in data["activeTimeIntervals"]]        

        return cls(**args)


class NotificationSettingsType(enum.StrEnum):
    SIMPLIFIED_ROUTING = "SimplifiedRouting"
    NAMED_ROUTING_TREE = "NamedRoutingTree"


class NamedRoutingTree:
    type_val: str
    routing_tree: str

    def __init__(self, routing_tree: str = "") -> None:
        self.type_val = NotificationSettingsType.NAMED_ROUTING_TREE
        self.routing_tree = routing_tree

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "routingTree": self.routing_tree,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "routingTree" in data:
            args["routing_tree"] = data["routingTree"]        

        return cls(**args)


ExpressionMap: typing.TypeAlias = dict[str, 'Expression']


class Expression:
    # The type of query if this is a query expression
    query_type: typing.Optional[str]
    relative_time_range: typing.Optional['RelativeTimeRange']
    # The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
    datasource_uid: typing.Optional['DatasourceUID']
    model: object
    # Used to mark the expression to be used as the final source for the rule evaluation
    # Only one expression in a rule can be marked as the source
    # For AlertRules, this is the expression that will be evaluated against the alerting condition
    # For RecordingRules, this is the expression that will be recorded
    source: typing.Optional[bool]

    def __init__(self, query_type: typing.Optional[str] = None, relative_time_range: typing.Optional['RelativeTimeRange'] = None, datasource_uid: typing.Optional['DatasourceUID'] = None, model: object = None, source: typing.Optional[bool] = None) -> None:
        self.query_type = query_type
        self.relative_time_range = relative_time_range
        self.datasource_uid = datasource_uid
        self.model = model
        self.source = source

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
        }
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.relative_time_range is not None:
            payload["relativeTimeRange"] = self.relative_time_range
        if self.datasource_uid is not None:
            payload["datasourceUID"] = self.datasource_uid
        if self.source is not None:
            payload["source"] = self.source
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "relativeTimeRange" in data:
            args["relative_time_range"] = RelativeTimeRange.from_json(data["relativeTimeRange"])
        if "datasourceUID" in data:
            args["datasource_uid"] = data["datasourceUID"]
        if "model" in data:
            args["model"] = data["model"]
        if "source" in data:
            args["source"] = data["source"]        

        return cls(**args)


class RelativeTimeRange:
    from_val: 'PromDurationWMillis'
    to: 'PromDurationWMillis'

    def __init__(self, from_val: typing.Optional['PromDurationWMillis'] = None, to: typing.Optional['PromDurationWMillis'] = None) -> None:
        self.from_val = from_val if from_val is not None else PromDurationWMillis()
        self.to = to if to is not None else PromDurationWMillis()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "from": self.from_val,
            "to": self.to,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]        

        return cls(**args)


PromDurationWMillis: typing.TypeAlias = str


DatasourceUID: typing.TypeAlias = str


class PanelRef:
    dashboard_uid: str
    panel_id: int

    def __init__(self, dashboard_uid: str = "", panel_id: int = 0) -> None:
        self.dashboard_uid = dashboard_uid
        self.panel_id = panel_id

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "dashboardUID": self.dashboard_uid,
            "panelID": self.panel_id,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "dashboardUID" in data:
            args["dashboard_uid"] = data["dashboardUID"]
        if "panelID" in data:
            args["panel_id"] = data["panelID"]        

        return cls(**args)


class RecordingRuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: 'IntervalTrigger'
    labels: typing.Optional[dict[str, 'TemplateString']]
    metric: 'MetricName'
    expressions: 'ExpressionMap'
    target_datasource_uid: 'DatasourceUID'

    def __init__(self, title: str = "", paused: typing.Optional[bool] = None, trigger: typing.Optional['IntervalTrigger'] = None, labels: typing.Optional[dict[str, 'TemplateString']] = None, metric: typing.Optional['MetricName'] = None, expressions: typing.Optional['ExpressionMap'] = None, target_datasource_uid: typing.Optional['DatasourceUID'] = None) -> None:
        self.title = title
        self.paused = paused
        self.trigger = trigger if trigger is not None else IntervalTrigger()
        self.labels = labels
        self.metric = metric if metric is not None else MetricName()
        self.expressions = expressions if expressions is not None else ExpressionMap()
        self.target_datasource_uid = target_datasource_uid if target_datasource_uid is not None else DatasourceUID()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "trigger": self.trigger,
            "metric": self.metric,
            "expressions": self.expressions,
            "targetDatasourceUID": self.target_datasource_uid,
        }
        if self.paused is not None:
            payload["paused"] = self.paused
        if self.labels is not None:
            payload["labels"] = self.labels
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "paused" in data:
            args["paused"] = data["paused"]
        if "trigger" in data:
            args["trigger"] = IntervalTrigger.from_json(data["trigger"])
        if "labels" in data:
            args["labels"] = {key: data["labels"][key] for key in data["labels"].keys()}
        if "metric" in data:
            args["metric"] = data["metric"]
        if "expressions" in data:
            args["expressions"] = {key: Expression.from_json(data["expressions"][key]) for key in data["expressions"].keys()}
        if "targetDatasourceUID" in data:
            args["target_datasource_uid"] = data["targetDatasourceUID"]        

        return cls(**args)


MetricName: typing.TypeAlias = str


class RuleSequenceSpec:
    trigger: 'IntervalTrigger'
    recording_rules: list['RuleRef']
    alerting_rules: typing.Optional[list['RuleRef']]

    def __init__(self, trigger: typing.Optional['IntervalTrigger'] = None, recording_rules: typing.Optional[list['RuleRef']] = None, alerting_rules: typing.Optional[list['RuleRef']] = None) -> None:
        self.trigger = trigger if trigger is not None else IntervalTrigger()
        self.recording_rules = recording_rules if recording_rules is not None else []
        self.alerting_rules = alerting_rules

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "trigger": self.trigger,
            "recordingRules": self.recording_rules,
        }
        if self.alerting_rules is not None:
            payload["alertingRules"] = self.alerting_rules
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "trigger" in data:
            args["trigger"] = IntervalTrigger.from_json(data["trigger"])
        if "recordingRules" in data:
            args["recording_rules"] = [RuleRef.from_json(item) for item in data["recordingRules"]]
        if "alertingRules" in data:
            args["alerting_rules"] = [RuleRef.from_json(item) for item in data["alertingRules"]]        

        return cls(**args)


class RuleRef:
    # name is the metadata.name of an AlertRule or RecordingRule resource.
    name: 'RuleUID'

    def __init__(self, name: typing.Optional['RuleUID'] = None) -> None:
        self.name = name if name is not None else RuleUID()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


RuleUID: typing.TypeAlias = str


class RuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: 'IntervalTrigger'
    labels: typing.Optional[dict[str, 'TemplateString']]
    expressions: 'ExpressionMap'

    def __init__(self, title: str = "", paused: typing.Optional[bool] = None, trigger: typing.Optional['IntervalTrigger'] = None, labels: typing.Optional[dict[str, 'TemplateString']] = None, expressions: typing.Optional['ExpressionMap'] = None) -> None:
        self.title = title
        self.paused = paused
        self.trigger = trigger if trigger is not None else IntervalTrigger()
        self.labels = labels
        self.expressions = expressions if expressions is not None else ExpressionMap()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "trigger": self.trigger,
            "expressions": self.expressions,
        }
        if self.paused is not None:
            payload["paused"] = self.paused
        if self.labels is not None:
            payload["labels"] = self.labels
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "paused" in data:
            args["paused"] = data["paused"]
        if "trigger" in data:
            args["trigger"] = IntervalTrigger.from_json(data["trigger"])
        if "labels" in data:
            args["labels"] = {key: data["labels"][key] for key in data["labels"].keys()}
        if "expressions" in data:
            args["expressions"] = {key: Expression.from_json(data["expressions"][key]) for key in data["expressions"].keys()}        

        return cls(**args)


class SearchResultResource:
    """
    #SearchResultResource is the full identity of a hit.
    """

    group: str
    resource: str
    kind: str
    name: str

    def __init__(self, group: str = "", resource: str = "", kind: str = "", name: str = "") -> None:
        self.group = group
        self.resource = resource
        self.kind = kind
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "group": self.group,
            "resource": self.resource,
            "kind": self.kind,
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "group" in data:
            args["group"] = data["group"]
        if "resource" in data:
            args["resource"] = data["resource"]
        if "kind" in data:
            args["kind"] = data["kind"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class RuleSearchHitFields:
    """
    #RuleSearchHitFields is the per-kind field payload returned on each hit.
    It carries the union of alert- and recording-rule search fields; only the
    fields relevant to a hit's kind are populated. This maps to the kind's
    declared searchFields.
    """

    title: typing.Optional[str]
    folder: typing.Optional[str]
    type_val: typing.Optional[str]
    interval: typing.Optional[str]
    paused: typing.Optional[bool]
    labels: typing.Optional[dict[str, str]]
    datasource_ui_ds: typing.Optional[list[str]]
    # Alert-rule fields.
    annotations: typing.Optional[dict[str, str]]
    for_val: typing.Optional[str]
    keep_firing_for: typing.Optional[str]
    dashboard_uid: typing.Optional[str]
    panel_id: typing.Optional[int]
    receiver: typing.Optional[str]
    notification_type: typing.Optional[str]
    routing_tree: typing.Optional[str]
    # Recording-rule fields.
    metric: typing.Optional[str]
    target_datasource_uid: typing.Optional[str]

    def __init__(self, title: typing.Optional[str] = None, folder: typing.Optional[str] = None, type_val: typing.Optional[str] = None, interval: typing.Optional[str] = None, paused: typing.Optional[bool] = None, labels: typing.Optional[dict[str, str]] = None, datasource_ui_ds: typing.Optional[list[str]] = None, annotations: typing.Optional[dict[str, str]] = None, for_val: typing.Optional[str] = None, keep_firing_for: typing.Optional[str] = None, dashboard_uid: typing.Optional[str] = None, panel_id: typing.Optional[int] = None, receiver: typing.Optional[str] = None, notification_type: typing.Optional[str] = None, routing_tree: typing.Optional[str] = None, metric: typing.Optional[str] = None, target_datasource_uid: typing.Optional[str] = None) -> None:
        self.title = title
        self.folder = folder
        self.type_val = type_val
        self.interval = interval
        self.paused = paused
        self.labels = labels
        self.datasource_ui_ds = datasource_ui_ds
        self.annotations = annotations
        self.for_val = for_val
        self.keep_firing_for = keep_firing_for
        self.dashboard_uid = dashboard_uid
        self.panel_id = panel_id
        self.receiver = receiver
        self.notification_type = notification_type
        self.routing_tree = routing_tree
        self.metric = metric
        self.target_datasource_uid = target_datasource_uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.title is not None:
            payload["title"] = self.title
        if self.folder is not None:
            payload["folder"] = self.folder
        if self.type_val is not None:
            payload["type"] = self.type_val
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.paused is not None:
            payload["paused"] = self.paused
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.datasource_ui_ds is not None:
            payload["datasourceUIDs"] = self.datasource_ui_ds
        if self.annotations is not None:
            payload["annotations"] = self.annotations
        if self.for_val is not None:
            payload["for"] = self.for_val
        if self.keep_firing_for is not None:
            payload["keepFiringFor"] = self.keep_firing_for
        if self.dashboard_uid is not None:
            payload["dashboardUID"] = self.dashboard_uid
        if self.panel_id is not None:
            payload["panelID"] = self.panel_id
        if self.receiver is not None:
            payload["receiver"] = self.receiver
        if self.notification_type is not None:
            payload["notificationType"] = self.notification_type
        if self.routing_tree is not None:
            payload["routingTree"] = self.routing_tree
        if self.metric is not None:
            payload["metric"] = self.metric
        if self.target_datasource_uid is not None:
            payload["targetDatasourceUID"] = self.target_datasource_uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "folder" in data:
            args["folder"] = data["folder"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "paused" in data:
            args["paused"] = data["paused"]
        if "labels" in data:
            args["labels"] = data["labels"]
        if "datasourceUIDs" in data:
            args["datasource_ui_ds"] = data["datasourceUIDs"]
        if "annotations" in data:
            args["annotations"] = data["annotations"]
        if "for" in data:
            args["for_val"] = data["for"]
        if "keepFiringFor" in data:
            args["keep_firing_for"] = data["keepFiringFor"]
        if "dashboardUID" in data:
            args["dashboard_uid"] = data["dashboardUID"]
        if "panelID" in data:
            args["panel_id"] = data["panelID"]
        if "receiver" in data:
            args["receiver"] = data["receiver"]
        if "notificationType" in data:
            args["notification_type"] = data["notificationType"]
        if "routingTree" in data:
            args["routing_tree"] = data["routingTree"]
        if "metric" in data:
            args["metric"] = data["metric"]
        if "targetDatasourceUID" in data:
            args["target_datasource_uid"] = data["targetDatasourceUID"]        

        return cls(**args)


class SearchResultHit:
    """
    #SearchResultHit is a single match: its identity, an optional relevance
    score (present only when the query included free text), and the requested
    fields.
    """

    resource: 'SearchResultResource'
    score: typing.Optional[float]
    fields: 'RuleSearchHitFields'

    def __init__(self, resource: typing.Optional['SearchResultResource'] = None, score: typing.Optional[float] = None, fields: typing.Optional['RuleSearchHitFields'] = None) -> None:
        self.resource = resource if resource is not None else SearchResultResource()
        self.score = score
        self.fields = fields if fields is not None else RuleSearchHitFields()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "resource": self.resource,
            "fields": self.fields,
        }
        if self.score is not None:
            payload["score"] = self.score
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "resource" in data:
            args["resource"] = SearchResultResource.from_json(data["resource"])
        if "score" in data:
            args["score"] = data["score"]
        if "fields" in data:
            args["fields"] = RuleSearchHitFields.from_json(data["fields"])        

        return cls(**args)


class TotalHitsRelation(enum.StrEnum):
    """
    #TotalHitsRelation says how totalHits relates to the real number of matching
    rules the caller may see: "eq" when it is exact, "lte" when it is an upper
    bound because authorisation was applied after the search ranked its results.
    """

    EQ = "eq"
    LTE = "lte"


class SearchResultsMetadata:
    """
    #SearchResultsMetadata carries the paging token and total authorised match
    count.
    """

    continue_val: typing.Optional[str]
    total_hits: typing.Optional[int]
    # Always read totalHits together with totalHitsRelation.
    total_hits_relation: typing.Optional['TotalHitsRelation']

    def __init__(self, continue_val: typing.Optional[str] = None, total_hits: typing.Optional[int] = None, total_hits_relation: typing.Optional['TotalHitsRelation'] = None) -> None:
        self.continue_val = continue_val
        self.total_hits = total_hits
        self.total_hits_relation = total_hits_relation

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.continue_val is not None:
            payload["continue"] = self.continue_val
        if self.total_hits is not None:
            payload["totalHits"] = self.total_hits
        if self.total_hits_relation is not None:
            payload["totalHitsRelation"] = self.total_hits_relation
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "continue" in data:
            args["continue_val"] = data["continue"]
        if "totalHits" in data:
            args["total_hits"] = data["totalHits"]
        if "totalHitsRelation" in data:
            args["total_hits_relation"] = data["totalHitsRelation"]        

        return cls(**args)


class FacetValue:
    """
    #FacetValue is a single value/count pair in a facet breakdown.
    """

    value: str
    count: int

    def __init__(self, value: str = "", count: int = 0) -> None:
        self.value = value
        self.count = count

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "value": self.value,
            "count": self.count,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]
        if "count" in data:
            args["count"] = data["count"]        

        return cls(**args)


class SearchResults:
    """
    #SearchResults is the response envelope, mirroring
    search.grafana.app SearchResults.
    """

    metadata: 'SearchResultsMetadata'
    items: list['SearchResultHit']
    facets: typing.Optional[dict[str, list['FacetValue']]]

    def __init__(self, metadata: typing.Optional['SearchResultsMetadata'] = None, items: typing.Optional[list['SearchResultHit']] = None, facets: typing.Optional[dict[str, list['FacetValue']]] = None) -> None:
        self.metadata = metadata if metadata is not None else SearchResultsMetadata()
        self.items = items if items is not None else []
        self.facets = facets

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "metadata": self.metadata,
            "items": self.items,
        }
        if self.facets is not None:
            payload["facets"] = self.facets
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "metadata" in data:
            args["metadata"] = SearchResultsMetadata.from_json(data["metadata"])
        if "items" in data:
            args["items"] = [SearchResultHit.from_json(item) for item in data["items"]]
        if "facets" in data:
            args["facets"] = {key: [FacetValue.from_json(item) for item in data["facets"][key]] for key in data["facets"].keys()}        

        return cls(**args)


ApiVersion: typing.Literal["rules.alerting.grafana.app/v0alpha1"] = "rules.alerting.grafana.app/v0alpha1"


AlertRuleKind: typing.Literal["AlertRule"] = "AlertRule"


RecordingRuleKind: typing.Literal["RecordingRule"] = "RecordingRule"


RuleSequenceKind: typing.Literal["RuleSequence"] = "RuleSequence"



