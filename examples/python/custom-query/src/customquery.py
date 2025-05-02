from typing import Any, Optional, Self

from grafana_foundation_sdk.cog import variants as cogvariants
from grafana_foundation_sdk.cog import runtime as cogruntime
from grafana_foundation_sdk.cog import builder


class CustomQuery(cogvariants.Dataquery):
    # ref_id and hide are expected on all queries
    ref_id: Optional[str]
    hide: Optional[bool]

    # query is specific to the CustomQuery type
    query: str

    def __init__(
        self, query: str, ref_id: Optional[str] = None, hide: Optional[bool] = None
    ):
        self.query = query
        self.ref_id = ref_id
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "query": self.query,
        }
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, Any]) -> Self:
        args: dict[str, Any] = {}

        if "query" in data:
            args["query"] = data["query"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]

        return cls(**args)


def custom_query_variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        # datasource plugin ID
        identifier="custom-query",
        from_json_hook=CustomQuery.from_json,
    )


class CustomQueryBuilder(builder.Builder[CustomQuery]):
    __internal: CustomQuery

    def __init__(self, query: str):
        self.__internal = CustomQuery(query=query)

    def build(self) -> CustomQuery:
        return self.__internal

    def ref_id(self, ref_id: str) -> Self:
        self.__internal.ref_id = ref_id

        return self

    def hide(self, hide: bool) -> Self:
        self.__internal.hide = hide

        return self
