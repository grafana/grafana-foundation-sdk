# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..models import common
from ..cog import runtime as cogruntime


class VizDisplayMode(enum.StrEnum):
    CANDLES_VOLUME = "candles+volume"
    CANDLES = "candles"
    VOLUME = "volume"


class CandleStyle(enum.StrEnum):
    CANDLES = "candles"
    OHLC_BARS = "ohlcbars"


class ColorStrategy(enum.StrEnum):
    OPEN_CLOSE = "open-close"
    CLOSE_CLOSE = "close-close"


class CandlestickFieldMap:
    # Corresponds to the starting value of the given period
    open_val: typing.Optional[str]
    # Corresponds to the highest value of the given period
    high: typing.Optional[str]
    # Corresponds to the lowest value of the given period
    low: typing.Optional[str]
    # Corresponds to the final (end) value of the given period
    close: typing.Optional[str]
    # Corresponds to the sample count in the given period. (e.g. number of trades)
    volume: typing.Optional[str]

    def __init__(self, open_val: typing.Optional[str] = None, high: typing.Optional[str] = None, low: typing.Optional[str] = None, close: typing.Optional[str] = None, volume: typing.Optional[str] = None):
        self.open_val = open_val
        self.high = high
        self.low = low
        self.close = close
        self.volume = volume

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.open_val is not None:
            payload["open"] = self.open_val
        if self.high is not None:
            payload["high"] = self.high
        if self.low is not None:
            payload["low"] = self.low
        if self.close is not None:
            payload["close"] = self.close
        if self.volume is not None:
            payload["volume"] = self.volume
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "open" in data:
            args["open_val"] = data["open"]
        if "high" in data:
            args["high"] = data["high"]
        if "low" in data:
            args["low"] = data["low"]
        if "close" in data:
            args["close"] = data["close"]
        if "volume" in data:
            args["volume"] = data["volume"]        

        return cls(**args)


class CandlestickColors:
    up: str
    down: str
    flat: str

    def __init__(self, up: str = "green", down: str = "red", flat: str = "gray"):
        self.up = up
        self.down = down
        self.flat = flat

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "up": self.up,
            "down": self.down,
            "flat": self.flat,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "up" in data:
            args["up"] = data["up"]
        if "down" in data:
            args["down"] = data["down"]
        if "flat" in data:
            args["flat"] = data["flat"]        

        return cls(**args)


class Options:
    # Sets which dimensions are used for the visualization
    mode: 'VizDisplayMode'
    # Sets the style of the candlesticks
    candle_style: 'CandleStyle'
    # Sets the color strategy for the candlesticks
    color_strategy: 'ColorStrategy'
    # Map fields to appropriate dimension
    fields: 'CandlestickFieldMap'
    # Set which colors are used when the price movement is up or down
    colors: 'CandlestickColors'
    legend: common.VizLegendOptions
    # When enabled, all fields will be sent to the graph
    include_all_fields: typing.Optional[bool]

    def __init__(self, mode: typing.Optional['VizDisplayMode'] = None, candle_style: typing.Optional['CandleStyle'] = None, color_strategy: typing.Optional['ColorStrategy'] = None, fields: typing.Optional['CandlestickFieldMap'] = None, colors: typing.Optional['CandlestickColors'] = None, legend: typing.Optional[common.VizLegendOptions] = None, include_all_fields: typing.Optional[bool] = False):
        self.mode = mode if mode is not None else VizDisplayMode.CANDLES_VOLUME
        self.candle_style = candle_style if candle_style is not None else CandleStyle.CANDLES
        self.color_strategy = color_strategy if color_strategy is not None else ColorStrategy.OPEN_CLOSE
        self.fields = fields if fields is not None else CandlestickFieldMap()
        self.colors = colors if colors is not None else CandlestickColors(down="red", flat="gray", up="green")
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.include_all_fields = include_all_fields

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "candleStyle": self.candle_style,
            "colorStrategy": self.color_strategy,
            "fields": self.fields,
            "colors": self.colors,
            "legend": self.legend,
        }
        if self.include_all_fields is not None:
            payload["includeAllFields"] = self.include_all_fields
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "candleStyle" in data:
            args["candle_style"] = data["candleStyle"]
        if "colorStrategy" in data:
            args["color_strategy"] = data["colorStrategy"]
        if "fields" in data:
            args["fields"] = CandlestickFieldMap.from_json(data["fields"])
        if "colors" in data:
            args["colors"] = CandlestickColors.from_json(data["colors"])
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "includeAllFields" in data:
            args["include_all_fields"] = data["includeAllFields"]        

        return cls(**args)


FieldConfig: typing.TypeAlias = common.GraphFieldConfig





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="candlestick",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )
