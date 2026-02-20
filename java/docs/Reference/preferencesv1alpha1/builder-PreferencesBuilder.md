---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

## Constructor

```java
new PreferencesBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Preferences build()
```

### <span class="badge object-method"></span> cookiePreferences

Cookie preferences

```java
public PreferencesBuilder cookiePreferences(com.grafana.foundation.cog.Builder<CookiePreferences> cookiePreferences)
```

### <span class="badge object-method"></span> homeDashboardUID

UID for the home dashboard

```java
public PreferencesBuilder homeDashboardUID(String homeDashboardUID)
```

### <span class="badge object-method"></span> language

Selected language (beta)

```java
public PreferencesBuilder language(String language)
```

### <span class="badge object-method"></span> navbar

Navigation preferences

```java
public PreferencesBuilder navbar(com.grafana.foundation.cog.Builder<NavbarPreference> navbar)
```

### <span class="badge object-method"></span> queryHistory

Explore query history preferences

```java
public PreferencesBuilder queryHistory(com.grafana.foundation.cog.Builder<QueryHistoryPreference> queryHistory)
```

### <span class="badge object-method"></span> regionalFormat

Selected locale (beta)

```java
public PreferencesBuilder regionalFormat(String regionalFormat)
```

### <span class="badge object-method"></span> theme

light, dark, empty is default

```java
public PreferencesBuilder theme(String theme)
```

### <span class="badge object-method"></span> timezone

The timezone selection

TODO: this should use the timezone defined in common

```java
public PreferencesBuilder timezone(String timezone)
```

### <span class="badge object-method"></span> weekStart

day of the week (sunday, monday, etc)

```java
public PreferencesBuilder weekStart(String weekStart)
```

## See also

 * <span class="badge object-type-class"></span> [Preferences](./object-Preferences.md)
