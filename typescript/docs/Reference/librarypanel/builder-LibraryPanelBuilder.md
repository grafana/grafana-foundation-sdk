---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```typescript
new LibraryPanelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> description

Panel description

```typescript
description(description: string)
```

### <span class="badge object-method"></span> folderUid

Folder UID

```typescript
folderUid(folderUid: string)
```

### <span class="badge object-method"></span> meta

Object storage metadata

```typescript
meta(meta: cog.Builder<librarypanel.LibraryElementDTOMeta>)
```

### <span class="badge object-method"></span> model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

```typescript
model(model: cog.Builder<librarypanel.PanelModel>)
```

### <span class="badge object-method"></span> name

Panel name (also saved in the model)

```typescript
name(name: string)
```

### <span class="badge object-method"></span> schemaVersion

Dashboard version when this was saved (zero if unknown)

```typescript
schemaVersion(schemaVersion: number)
```

### <span class="badge object-method"></span> type

The panel type (from inside the model)

```typescript
type(type: string)
```

### <span class="badge object-method"></span> uid

Library element UID

```typescript
uid(uid: string)
```

### <span class="badge object-method"></span> version

panel version, incremented each time the dashboard is updated.

```typescript
version(version: number)
```

## See also

 * <span class="badge object-type-interface"></span> [LibraryPanel](./object-LibraryPanel.md)
