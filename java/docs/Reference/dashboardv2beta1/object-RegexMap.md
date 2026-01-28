---
title: <span class="badge object-type-class"></span> RegexMap
---
# <span class="badge object-type-class"></span> RegexMap

Maps regular expressions to replacement text and a color.

For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

## Definition

```java
public class RegexMap {
  public MappingType type;
  public Dashboardv2beta1RegexMapOptions options;
}
```
