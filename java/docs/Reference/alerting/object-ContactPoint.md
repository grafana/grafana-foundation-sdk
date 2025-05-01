---
title: <span class="badge object-type-class"></span> ContactPoint
---
# <span class="badge object-type-class"></span> ContactPoint

EmbeddedContactPoint is the contact point type that is used

by grafanas embedded alertmanager implementation.

## Definition

```java
public class ContactPoint {
  public Boolean disableResolveMessage;
  public String name;
  public String provenance;
  public Object settings;
  public ContactPointType type;
  public String uid;
}
```
## See also

 * <span class="badge builder"></span> [ContactPointBuilder](./builder-ContactPointBuilder.md)
