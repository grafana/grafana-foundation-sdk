---
title: <span class="badge object-type-disjunction"></span> TableCellOptions
---
# <span class="badge object-type-disjunction"></span> TableCellOptions

Table cell options. Each cell has a display mode

and other potential options for that display.

## Definition

```python
# Table cell options. Each cell has a display mode
# and other potential options for that display.
TableCellOptions: typing.TypeAlias = typing.Union[common.TableAutoCellOptions, common.TableSparklineCellOptions, common.TableBarGaugeCellOptions, common.TableColoredBackgroundCellOptions, common.TableColorTextCellOptions, common.TableImageCellOptions, common.TableJsonViewCellOptions]
```
