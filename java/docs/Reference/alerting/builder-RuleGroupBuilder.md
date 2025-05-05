---
title: <span class="badge builder"></span> RuleGroupBuilder
---
# <span class="badge builder"></span> RuleGroupBuilder

## Constructor

```java
new RuleGroupBuilder(String title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RuleGroup build()
```

### <span class="badge object-method"></span> folderUid

```java
public RuleGroupBuilder folderUid(String folderUid)
```

### <span class="badge object-method"></span> interval

The interval, in seconds, at which all rules in the group are evaluated.

If a group contains many rules, the rules are evaluated sequentially.

```java
public RuleGroupBuilder interval(Long interval)
```

### <span class="badge object-method"></span> rules

```java
public RuleGroupBuilder rules(List<com.grafana.foundation.cog.Builder<Rule>> rules)
```

### <span class="badge object-method"></span> title

```java
public RuleGroupBuilder title(String title)
```

### <span class="badge object-method"></span> withRule

```java
public RuleGroupBuilder withRule(com.grafana.foundation.cog.Builder<Rule> rule)
```

## See also

 * <span class="badge object-type-class"></span> [RuleGroup](./object-RuleGroup.md)
