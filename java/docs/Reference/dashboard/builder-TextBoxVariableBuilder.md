---
title: <span class="badge builder"></span> TextBoxVariableBuilder
---
# <span class="badge builder"></span> TextBoxVariableBuilder

## Constructor

```java
new TextBoxVariableBuilder(String name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TextBoxVariable build()
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```java
public TextBoxVariableBuilder current(VariableOption current)
```

### <span class="badge object-method"></span> defaultValue

Query used to fetch values for a variable

```java
public TextBoxVariableBuilder defaultValue(StringOrMap query)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```java
public TextBoxVariableBuilder description(String description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```java
public TextBoxVariableBuilder hide(VariableHide hide)
```

### <span class="badge object-method"></span> label

Optional display name

```java
public TextBoxVariableBuilder label(String label)
```

### <span class="badge object-method"></span> name

Name of variable

```java
public TextBoxVariableBuilder name(String name)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

```java
public TextBoxVariableBuilder options(List<VariableOption> options)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
