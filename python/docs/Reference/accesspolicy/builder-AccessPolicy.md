---
title: <span class="badge builder"></span> AccessPolicy
---
# <span class="badge builder"></span> AccessPolicy

## Constructor

```python
AccessPolicy()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> accesspolicy.AccessPolicy
```

### <span class="badge object-method"></span> role

The role that must apply this policy

```python
def role(role: cogbuilder.Builder[accesspolicy.RoleRef]) -> typing.Self
```

### <span class="badge object-method"></span> rules

The set of rules to apply.  Note that * is required to modify

access policy rules, and that "none" will reject all actions

```python
def rules(rule: cogbuilder.Builder[accesspolicy.AccessRule]) -> typing.Self
```

### <span class="badge object-method"></span> scope

The scope where these policies should apply

```python
def scope(scope: cogbuilder.Builder[accesspolicy.ResourceRef]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AccessPolicy](./object-AccessPolicy.md)
