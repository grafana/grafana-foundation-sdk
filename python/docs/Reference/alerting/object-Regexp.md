---
title: <span class="badge object-type-scalar"></span> Regexp
---
# <span class="badge object-type-scalar"></span> Regexp

A Regexp is safe for concurrent use by multiple goroutines,

except for configuration methods, such as Longest.

## Definition

```python
# A Regexp is safe for concurrent use by multiple goroutines,
# except for configuration methods, such as Longest.
Regexp: typing.TypeAlias = object
```
