---
title: <span class="badge object-type-class"></span> TableFieldOptions
---
# <span class="badge object-type-class"></span> TableFieldOptions

Field options for each field within a table (e.g 10, "The String", 64.20, etc.)

Generally defines alignment, filtering capabilties, display options, etc.

## Definition

```python
class TableFieldOptions:
    """
    Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
    Generally defines alignment, filtering capabilties, display options, etc.
    """

    width: typing.Optional[float]
    min_width: typing.Optional[float]
    align: common.FieldTextAlignment
    # This field is deprecated in favor of using cellOptions
    display_mode: typing.Optional[common.TableCellDisplayMode]
    cell_options: typing.Optional[common.TableCellOptions]
    # ?? default is missing or false ??
    hidden: typing.Optional[bool]
    inspect: bool
    filterable: typing.Optional[bool]
    # Hides any header for a column, usefull for columns that show some static content or buttons.
    hide_header: typing.Optional[bool]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [TableFieldOptions](./builder-TableFieldOptions.md)
