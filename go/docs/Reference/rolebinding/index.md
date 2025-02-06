# <span class="badge package-core"></span> rolebinding

## Objects

 * <span class="badge object-type-struct"></span> [BuiltinRoleRef](./object-BuiltinRoleRef.md)
 * <span class="badge object-type-enum"></span> [BuiltinRoleRefName](./object-BuiltinRoleRefName.md)
 * <span class="badge object-type-struct"></span> [BuiltinRoleRefOrCustomRoleRef](./object-BuiltinRoleRefOrCustomRoleRef.md)
 * <span class="badge object-type-struct"></span> [CustomRoleRef](./object-CustomRoleRef.md)
 * <span class="badge object-type-struct"></span> [RoleBinding](./object-RoleBinding.md)
 * <span class="badge object-type-struct"></span> [RoleBindingSubject](./object-RoleBindingSubject.md)
 * <span class="badge object-type-enum"></span> [RoleBindingSubjectKind](./object-RoleBindingSubjectKind.md)
## Builders

 * <span class="badge builder"></span> [BuiltinRoleRefBuilder](./builder-BuiltinRoleRefBuilder.md)
 * <span class="badge builder"></span> [CustomRoleRefBuilder](./builder-CustomRoleRefBuilder.md)
 * <span class="badge builder"></span> [RoleBindingBuilder](./builder-RoleBindingBuilder.md)
 * <span class="badge builder"></span> [RoleBindingSubjectBuilder](./builder-RoleBindingSubjectBuilder.md)
## Functions

### <span class="badge function"></span> NewRoleBinding

NewRoleBinding creates a new RoleBinding object.

```go
func NewRoleBinding() *RoleBinding
```

### <span class="badge function"></span> NewBuiltinRoleRef

NewBuiltinRoleRef creates a new BuiltinRoleRef object.

```go
func NewBuiltinRoleRef() *BuiltinRoleRef
```

### <span class="badge function"></span> NewCustomRoleRef

NewCustomRoleRef creates a new CustomRoleRef object.

```go
func NewCustomRoleRef() *CustomRoleRef
```

### <span class="badge function"></span> NewRoleBindingSubject

NewRoleBindingSubject creates a new RoleBindingSubject object.

```go
func NewRoleBindingSubject() *RoleBindingSubject
```

### <span class="badge function"></span> NewBuiltinRoleRefOrCustomRoleRef

NewBuiltinRoleRefOrCustomRoleRef creates a new BuiltinRoleRefOrCustomRoleRef object.

```go
func NewBuiltinRoleRefOrCustomRoleRef() *BuiltinRoleRefOrCustomRoleRef
```

### <span class="badge function"></span> RoleBindingConverter

RoleBindingConverter accepts a `RoleBinding` object and generates the Go code to build this object using builders.

```go
func RoleBindingConverter(input RoleBinding) string
```

### <span class="badge function"></span> BuiltinRoleRefConverter

BuiltinRoleRefConverter accepts a `BuiltinRoleRef` object and generates the Go code to build this object using builders.

```go
func BuiltinRoleRefConverter(input BuiltinRoleRef) string
```

### <span class="badge function"></span> CustomRoleRefConverter

CustomRoleRefConverter accepts a `CustomRoleRef` object and generates the Go code to build this object using builders.

```go
func CustomRoleRefConverter(input CustomRoleRef) string
```

### <span class="badge function"></span> RoleBindingSubjectConverter

RoleBindingSubjectConverter accepts a `RoleBindingSubject` object and generates the Go code to build this object using builders.

```go
func RoleBindingSubjectConverter(input RoleBindingSubject) string
```

