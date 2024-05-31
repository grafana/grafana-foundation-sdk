# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime


class ParcaQueryType(enum.StrEnum):
    METRICS = "metrics"
    PROFILE = "profile"
    BOTH = "both"


class Dataquery(cogvariants.Dataquery):
    # Specifies the query label selectors.
    label_selector: str
    # Specifies the type of profile to query.
    profile_type_id: str
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]

    def __init__(self, label_selector: str = "{}", profile_type_id: str = "", ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None):
        self.label_selector = label_selector
        self.profile_type_id = profile_type_id
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "labelSelector": self.label_selector,
            "profileTypeId": self.profile_type_id,
            "refId": self.ref_id,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "labelSelector" in data:
            args["label_selector"] = data["labelSelector"]
        if "profileTypeId" in data:
            args["profile_type_id"] = data["profileTypeId"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="parca",
        from_json_hook=Dataquery.from_json,
    )



