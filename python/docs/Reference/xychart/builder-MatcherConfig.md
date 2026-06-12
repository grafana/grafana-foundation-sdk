---
title: <span class="badge builder"></span> MatcherConfig
---
# <span class="badge builder"></span> MatcherConfig

NOTE: (copied from dashboard_kind.cue, since not exported)

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Constructor

```python
MatcherConfig()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> xychart.MatcherConfig
```

### <span class="badge object-method"></span> id_val

The matcher id. This is used to find the matcher implementation from registry.

```python
def id_val(id_val: str) -> typing.Self
```

### <span class="badge object-method"></span> options

The matcher options. This is specific to the matcher implementation.

```python
def options(options: object) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [MatcherConfig](./object-MatcherConfig.md)
