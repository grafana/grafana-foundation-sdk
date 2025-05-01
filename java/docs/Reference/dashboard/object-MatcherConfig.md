---
title: <span class="badge object-type-class"></span> MatcherConfig
---
# <span class="badge object-type-class"></span> MatcherConfig

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Definition

```java
public class MatcherConfig {
  public String id;
  public Object options;
}
```
