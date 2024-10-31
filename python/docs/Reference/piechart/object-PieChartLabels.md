---
title: <span class="badge object-type-enum"></span> PieChartLabels
---
# <span class="badge object-type-enum"></span> PieChartLabels

Select labels to display on the pie chart.

 - Name - The series or field name.

 - Percent - The percentage of the whole.

 - Value - The raw numerical value.

## Definition

```python
class PieChartLabels(enum.StrEnum):
    """
    Select labels to display on the pie chart.
     - Name - The series or field name.
     - Percent - The percentage of the whole.
     - Value - The raw numerical value.
    """

    NAME = "name"
    VALUE = "value"
    PERCENT = "percent"
```
