---
title: <span class="badge object-type-enum"></span> TableCellDisplayMode
---
# <span class="badge object-type-enum"></span> TableCellDisplayMode

Internally, this is the "type" of cell that's being displayed

in the table such as colored text, JSON, gauge, etc.

The color-background-solid, gradient-gauge, and lcd-gauge

modes are deprecated in favor of new cell subOptions

## Definition

```python
class TableCellDisplayMode(enum.StrEnum):
    """
    Internally, this is the "type" of cell that's being displayed
    in the table such as colored text, JSON, gauge, etc.
    The color-background-solid, gradient-gauge, and lcd-gauge
    modes are deprecated in favor of new cell subOptions
    """

    AUTO = "auto"
    COLOR_TEXT = "color-text"
    COLOR_BACKGROUND = "color-background"
    COLOR_BACKGROUND_SOLID = "color-background-solid"
    GRADIENT_GAUGE = "gradient-gauge"
    LCD_GAUGE = "lcd-gauge"
    JSON_VIEW = "json-view"
    BASIC_GAUGE = "basic"
    IMAGE = "image"
    GAUGE = "gauge"
    SPARKLINE = "sparkline"
    DATA_LINKS = "data-links"
    CUSTOM = "custom"
```
