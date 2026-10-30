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


class AlertRuleHealth(enum.StrEnum):
    UNKNOWN = "Unknown"
    OK = "OK"
    PAUSED = "Paused"
    ERROR = "Error"
    NO_DATA = "NoData"


class AlertRuleState(enum.StrEnum):
    INACTIVE = "Inactive"
    HEALTHY = "Healthy"
    FIRING = "Firing"
    PENDING = "Pending"
    RECOVERING = "Recovering"


class AlertRuleStateReason(enum.StrEnum):
    EVALUATED = "Evaluated"
    KEEP_LAST = "KeepLast"


class AlertRuleInstanceTotals:
    """
    Count of alert instances per state. error also counts instances whose evaluation
    errored but were mapped to another state via execErrState, so it can overlap.
    """

    healthy: typing.Optional[int]
    firing: typing.Optional[int]
    pending: typing.Optional[int]
    recovering: typing.Optional[int]
    nodata: typing.Optional[int]
    error: typing.Optional[int]

    def __init__(self, healthy: typing.Optional[int] = None, firing: typing.Optional[int] = None, pending: typing.Optional[int] = None, recovering: typing.Optional[int] = None, nodata: typing.Optional[int] = None, error: typing.Optional[int] = None) -> None:
        self.healthy = healthy
        self.firing = firing
        self.pending = pending
        self.recovering = recovering
        self.nodata = nodata
        self.error = error

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.healthy is not None:
            payload["healthy"] = self.healthy
        if self.firing is not None:
            payload["firing"] = self.firing
        if self.pending is not None:
            payload["pending"] = self.pending
        if self.recovering is not None:
            payload["recovering"] = self.recovering
        if self.nodata is not None:
            payload["nodata"] = self.nodata
        if self.error is not None:
            payload["error"] = self.error
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "healthy" in data:
            args["healthy"] = data["healthy"]
        if "firing" in data:
            args["firing"] = data["firing"]
        if "pending" in data:
            args["pending"] = data["pending"]
        if "recovering" in data:
            args["recovering"] = data["recovering"]
        if "nodata" in data:
            args["nodata"] = data["nodata"]
        if "error" in data:
            args["error"] = data["error"]        

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


class RecordingRuleHealth(enum.StrEnum):
    UNKNOWN = "Unknown"
    RECORDING = "Recording"
    PAUSED = "Paused"
    ERROR = "Error"
    NO_DATA = "NoData"


class RuleSequenceSpec:
    trigger: 'IntervalTrigger'
    recording_rules: list['RuleRef']
    alert_rules: typing.Optional[list['RuleRef']]

    def __init__(self, trigger: typing.Optional['IntervalTrigger'] = None, recording_rules: typing.Optional[list['RuleRef']] = None, alert_rules: typing.Optional[list['RuleRef']] = None) -> None:
        self.trigger = trigger if trigger is not None else IntervalTrigger()
        self.recording_rules = recording_rules if recording_rules is not None else []
        self.alert_rules = alert_rules

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "trigger": self.trigger,
            "recordingRules": self.recording_rules,
        }
        if self.alert_rules is not None:
            payload["alertRules"] = self.alert_rules
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "trigger" in data:
            args["trigger"] = IntervalTrigger.from_json(data["trigger"])
        if "recordingRules" in data:
            args["recording_rules"] = [RuleRef.from_json(item) for item in data["recordingRules"]]
        if "alertRules" in data:
            args["alert_rules"] = [RuleRef.from_json(item) for item in data["alertRules"]]        

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


class AlertRuleStatus:
    health: typing.Optional['AlertRuleHealth']
    state: typing.Optional['AlertRuleState']
    state_reason: typing.Optional['AlertRuleStateReason']
    last_evaluation_time: typing.Optional[str]
    # duration of the last evaluation in seconds
    evaluation_duration: typing.Optional[float]
    last_error: typing.Optional[str]
    totals: typing.Optional['AlertRuleInstanceTotals']

    def __init__(self, health: typing.Optional['AlertRuleHealth'] = None, state: typing.Optional['AlertRuleState'] = None, state_reason: typing.Optional['AlertRuleStateReason'] = None, last_evaluation_time: typing.Optional[str] = None, evaluation_duration: typing.Optional[float] = None, last_error: typing.Optional[str] = None, totals: typing.Optional['AlertRuleInstanceTotals'] = None) -> None:
        self.health = health
        self.state = state
        self.state_reason = state_reason
        self.last_evaluation_time = last_evaluation_time
        self.evaluation_duration = evaluation_duration
        self.last_error = last_error
        self.totals = totals

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.health is not None:
            payload["health"] = self.health
        if self.state is not None:
            payload["state"] = self.state
        if self.state_reason is not None:
            payload["stateReason"] = self.state_reason
        if self.last_evaluation_time is not None:
            payload["lastEvaluationTime"] = self.last_evaluation_time
        if self.evaluation_duration is not None:
            payload["evaluationDuration"] = self.evaluation_duration
        if self.last_error is not None:
            payload["lastError"] = self.last_error
        if self.totals is not None:
            payload["totals"] = self.totals
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "health" in data:
            args["health"] = data["health"]
        if "state" in data:
            args["state"] = data["state"]
        if "stateReason" in data:
            args["state_reason"] = data["stateReason"]
        if "lastEvaluationTime" in data:
            args["last_evaluation_time"] = data["lastEvaluationTime"]
        if "evaluationDuration" in data:
            args["evaluation_duration"] = data["evaluationDuration"]
        if "lastError" in data:
            args["last_error"] = data["lastError"]
        if "totals" in data:
            args["totals"] = AlertRuleInstanceTotals.from_json(data["totals"])        

        return cls(**args)


class RecordingRuleStatus:
    health: typing.Optional['RecordingRuleHealth']
    last_evaluation_time: typing.Optional[str]
    # duration of the last evaluation in seconds
    evaluation_duration: typing.Optional[float]
    last_error: typing.Optional[str]

    def __init__(self, health: typing.Optional['RecordingRuleHealth'] = None, last_evaluation_time: typing.Optional[str] = None, evaluation_duration: typing.Optional[float] = None, last_error: typing.Optional[str] = None) -> None:
        self.health = health
        self.last_evaluation_time = last_evaluation_time
        self.evaluation_duration = evaluation_duration
        self.last_error = last_error

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.health is not None:
            payload["health"] = self.health
        if self.last_evaluation_time is not None:
            payload["lastEvaluationTime"] = self.last_evaluation_time
        if self.evaluation_duration is not None:
            payload["evaluationDuration"] = self.evaluation_duration
        if self.last_error is not None:
            payload["lastError"] = self.last_error
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "health" in data:
            args["health"] = data["health"]
        if "lastEvaluationTime" in data:
            args["last_evaluation_time"] = data["lastEvaluationTime"]
        if "evaluationDuration" in data:
            args["evaluation_duration"] = data["evaluationDuration"]
        if "lastError" in data:
            args["last_error"] = data["lastError"]        

        return cls(**args)


ApiVersion: typing.Literal["rules.alerting.grafana.app/v1beta1"] = "rules.alerting.grafana.app/v1beta1"


AlertRuleKind: typing.Literal["AlertRule"] = "AlertRule"


RecordingRuleKind: typing.Literal["RecordingRule"] = "RecordingRule"


RuleSequenceKind: typing.Literal["RuleSequence"] = "RuleSequence"



