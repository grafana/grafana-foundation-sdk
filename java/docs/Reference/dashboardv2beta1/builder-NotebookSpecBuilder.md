---
title: <span class="badge builder"></span> NotebookSpecBuilder
---
# <span class="badge builder"></span> NotebookSpecBuilder

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Constructor

```java
new NotebookSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public NotebookSpec build()
```

### <span class="badge object-method"></span> description

```java
public NotebookSpecBuilder description(String description)
```

### <span class="badge object-method"></span> elements

```java
public NotebookSpecBuilder elements(Map<String, NotebookElement> elements)
```

### <span class="badge object-method"></span> layout

```java
public NotebookSpecBuilder layout(com.grafana.foundation.cog.Builder<NotebookLayoutKind> layout)
```

### <span class="badge object-method"></span> tags

```java
public NotebookSpecBuilder tags(List<String> tags)
```

### <span class="badge object-method"></span> timeSettings

```java
public NotebookSpecBuilder timeSettings(com.grafana.foundation.cog.Builder<TimeSettingsSpec> timeSettings)
```

### <span class="badge object-method"></span> title

```java
public NotebookSpecBuilder title(String title)
```

## See also

 * <span class="badge object-type-class"></span> [NotebookSpec](./object-NotebookSpec.md)
