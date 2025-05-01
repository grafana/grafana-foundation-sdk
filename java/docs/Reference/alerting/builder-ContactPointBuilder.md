---
title: <span class="badge builder"></span> ContactPointBuilder
---
# <span class="badge builder"></span> ContactPointBuilder

## Constructor

```java
new ContactPointBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ContactPoint build()
```

### <span class="badge object-method"></span> disableResolveMessage

```java
public ContactPointBuilder disableResolveMessage(Boolean disableResolveMessage)
```

### <span class="badge object-method"></span> name

Name is used as grouping key in the UI. Contact points with the

same name will be grouped in the UI.

```java
public ContactPointBuilder name(String name)
```

### <span class="badge object-method"></span> provenance

```java
public ContactPointBuilder provenance(String provenance)
```

### <span class="badge object-method"></span> settings

```java
public ContactPointBuilder settings(Object settings)
```

### <span class="badge object-method"></span> type

```java
public ContactPointBuilder type(ContactPointType type)
```

### <span class="badge object-method"></span> uid

UID is the unique identifier of the contact point. The UID can be

set by the user.

```java
public ContactPointBuilder uid(String uid)
```

## See also

 * <span class="badge object-type-class"></span> [ContactPoint](./object-ContactPoint.md)
