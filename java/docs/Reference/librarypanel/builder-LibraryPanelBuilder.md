---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```java
new LibraryPanelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public LibraryPanel build()
```

### <span class="badge object-method"></span> description

Panel description

```java
public LibraryPanelBuilder description(String description)
```

### <span class="badge object-method"></span> folderUid

Folder UID

```java
public LibraryPanelBuilder folderUid(String folderUid)
```

### <span class="badge object-method"></span> meta

Object storage metadata

```java
public LibraryPanelBuilder meta(LibraryElementDTOMeta meta)
```

### <span class="badge object-method"></span> model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

```java
public LibraryPanelBuilder model(PanelModel model)
```

### <span class="badge object-method"></span> name

Panel name (also saved in the model)

```java
public LibraryPanelBuilder name(String name)
```

### <span class="badge object-method"></span> schemaVersion

Dashboard version when this was saved (zero if unknown)

```java
public LibraryPanelBuilder schemaVersion(Short schemaVersion)
```

### <span class="badge object-method"></span> type

The panel type (from inside the model)

```java
public LibraryPanelBuilder type(String type)
```

### <span class="badge object-method"></span> uid

Library element UID

```java
public LibraryPanelBuilder uid(String uid)
```

### <span class="badge object-method"></span> version

panel version, incremented each time the dashboard is updated.

```java
public LibraryPanelBuilder version(Long version)
```

## See also

 * <span class="badge object-type-class"></span> [LibraryPanel](./object-LibraryPanel.md)
