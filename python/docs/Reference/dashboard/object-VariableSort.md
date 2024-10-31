---
title: <span class="badge object-type-enum"></span> VariableSort
---
# <span class="badge object-type-enum"></span> VariableSort

Sort variable options

Accepted values are:

`0`: No sorting

`1`: Alphabetical ASC

`2`: Alphabetical DESC

`3`: Numerical ASC

`4`: Numerical DESC

`5`: Alphabetical Case Insensitive ASC

`6`: Alphabetical Case Insensitive DESC

`7`: Natural ASC

`8`: Natural DESC

## Definition

```python
class VariableSort(enum.IntEnum):
    """
    Sort variable options
    Accepted values are:
    `0`: No sorting
    `1`: Alphabetical ASC
    `2`: Alphabetical DESC
    `3`: Numerical ASC
    `4`: Numerical DESC
    `5`: Alphabetical Case Insensitive ASC
    `6`: Alphabetical Case Insensitive DESC
    `7`: Natural ASC
    `8`: Natural DESC
    """

    DISABLED = 0
    ALPHABETICAL_ASC = 1
    ALPHABETICAL_DESC = 2
    NUMERICAL_ASC = 3
    NUMERICAL_DESC = 4
    ALPHABETICAL_CASE_INSENSITIVE_ASC = 5
    ALPHABETICAL_CASE_INSENSITIVE_DESC = 6
    NATURAL_ASC = 7
    NATURAL_DESC = 8
```
