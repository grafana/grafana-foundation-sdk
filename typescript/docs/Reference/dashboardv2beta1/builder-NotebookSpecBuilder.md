---
title: <span class="badge builder"></span> NotebookSpecBuilder
---
# <span class="badge builder"></span> NotebookSpecBuilder

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Constructor

```typescript
new NotebookSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> description

```typescript
description(description: string)
```

### <span class="badge object-method"></span> elements

```typescript
elements(elements: Record<string, cog.Builder<dashboardv2beta1.NotebookElement>>)
```

### <span class="badge object-method"></span> layout

```typescript
layout(layout: cog.Builder<dashboardv2beta1.NotebookLayoutKind>)
```

### <span class="badge object-method"></span> tags

```typescript
tags(tags: string[])
```

### <span class="badge object-method"></span> timeSettings

```typescript
timeSettings(timeSettings: cog.Builder<dashboardv2beta1.TimeSettingsSpec>)
```

### <span class="badge object-method"></span> title

```typescript
title(title: string)
```

## See also

 * <span class="badge object-type-interface"></span> [NotebookSpec](./object-NotebookSpec.md)
