# Template veneers

The templates contained in this folder are used to customize the types and builders generated by cog.
For example, to automatically position panels on the grid whenever they're added to a dashboard.

To achieve that, cog looks for [template blocks](https://pkg.go.dev/text/template#example-Template-Block) following
a particular naming convention, and renders them.

For types, the following blocks are supported:

* `object_{{ Package }}_{{ ObjectName }}_custom_unmarshal`: allows the definition of a custom unmarshal function for the object.
* `object_{{ Package }}_{{ ObjectName }}_field_{{ FieldName}}_custom_strict_unmarshal`: allows the definition of a custom — strict — unmarshal logic for a field.

For builders, the following blocks are supported:

* `pre_assignment_{{ BuilderName }}_{{ OptionName }}`: will be rendered within an option, before the assignment.
* `post_assignment_{{ BuilderName }}_{{ OptionName }}`: will be rendered within an option, after the assignment.

For API reference documentation, the following blocks are supported:

* `api_reference_package_{{ Package }}_extra`: additional content that will be rendered on the documentation page for a package.
* `api_reference_object_{{ Package }}_{{ ObjectName }}_extra`: additional content that will be rendered on the documentation page for an object.
* `api_reference_builder_{{ Package }}_{{ BuilderName }}_extra`: additional content that will be rendered on the documentation page for a builder.

> [!NOTE]
> The variables components of block names are **case-sensitive**.
