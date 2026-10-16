---
title: <span class="badge object-type-class"></span> NotebookSpec
---
# <span class="badge object-type-class"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Definition

```java
public class NotebookSpec {
  public String title;
  public String description;
  public List<String> tags;
  public TimeSettingsSpec timeSettings;
  public Map<String, NotebookElement> elements;
  public NotebookLayoutKind layout;
}
```
## See also

 * <span class="badge builder"></span> [NotebookSpecBuilder](./builder-NotebookSpecBuilder.md)
