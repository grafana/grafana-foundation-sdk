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
    label_selector: typing.Optional[str]
    profile_type_id: typing.Optional[str]
    ref_id: typing.Optional[str]
    hide: typing.Optional[bool]
    query_type: typing.Optional[str]
    datasource: typing.Optional[object]

    def __init__(self, label_selector: typing.Optional[str] = "{}", profile_type_id: typing.Optional[str] = None, ref_id: typing.Optional[str] = None, hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None):
        self.label_selector = label_selector
        self.profile_type_id = profile_type_id
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.label_selector is not None:
            payload["labelSelector"] = self.label_selector
        if self.profile_type_id is not None:
            payload["profileTypeId"] = self.profile_type_id
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
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



