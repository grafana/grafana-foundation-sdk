---
title: <span class="badge builder"></span> RuleGroup
---
# <span class="badge builder"></span> RuleGroup

## Constructor

```python
RuleGroup(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.RuleGroup
```

### <span class="badge object-method"></span> folder_uid

```python
def folder_uid(folder_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> interval

The interval, in seconds, at which all rules in the group are evaluated.

If a group contains many rules, the rules are evaluated sequentially.

```python
def interval(interval: alerting.Duration) -> typing.Self
```

### <span class="badge object-method"></span> rules

```python
def rules(rules: list[cogbuilder.Builder[alerting.Rule]]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> with_rule

```python
def with_rule(rule: cogbuilder.Builder[alerting.Rule]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RuleGroup](./object-RuleGroup.md)
