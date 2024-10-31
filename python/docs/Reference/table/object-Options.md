---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Represents the index of the selected frame
    frame_index: float
    # Controls whether the panel should show the header
    show_header: bool
    # Controls whether the header should show icons for the column types
    show_type_icons: typing.Optional[bool]
    # Used to control row sorting
    sort_by: typing.Optional[list[common.TableSortByFieldState]]
    # Controls footer options
    footer: typing.Optional[common.TableFooterOptions]
    # Controls the height of the rows
    cell_height: typing.Optional[common.TableCellHeight]
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

