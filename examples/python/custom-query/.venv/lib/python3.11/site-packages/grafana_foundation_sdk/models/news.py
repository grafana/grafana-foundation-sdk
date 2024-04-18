# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class Options:
    # empty/missing will default to grafana blog
    feed_url: typing.Optional[str]
    show_image: typing.Optional[bool]

    def __init__(self, feed_url: typing.Optional[str] = None, show_image: typing.Optional[bool] = True):
        self.feed_url = feed_url
        self.show_image = show_image

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.feed_url is not None:
            payload["feedUrl"] = self.feed_url
        if self.show_image is not None:
            payload["showImage"] = self.show_image
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "feedUrl" in data:
            args["feed_url"] = data["feedUrl"]
        if "showImage" in data:
            args["show_image"] = data["showImage"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="news",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
