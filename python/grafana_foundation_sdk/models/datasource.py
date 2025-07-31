# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..cog import variants as cogvariants
import typing
from ..models import dashboard
from ..cog import runtime as cogruntime


class Dataquery(cogvariants.Dataquery):
    # Panel ID from wich the queries will be reused.
    panel_id: int
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    with_transforms: bool
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]

    def __init__(self, panel_id: int = 0, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, with_transforms: bool = False, datasource: typing.Optional[dashboard.DataSourceRef] = None):
        self.panel_id = panel_id
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.with_transforms = with_transforms
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "panelId": self.panel_id,
            "refId": self.ref_id,
            "withTransforms": self.with_transforms,
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
        
        if "panelId" in data:
            args["panel_id"] = data["panelId"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "withTransforms" in data:
            args["with_transforms"] = data["withTransforms"]
        if "datasource" in data:
            args["datasource"] = dashboard.DataSourceRef.from_json(data["datasource"])        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="datasource",
        from_json_hook=Dataquery.from_json,
    )

