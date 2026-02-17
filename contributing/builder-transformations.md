# Builder transformations

`cog` supports modifying its [intermediate representation](https://grafana.github.io/cog/reference/glossary/#intermediate-representation)
of [builders](https://grafana.github.io/cog/reference/glossary/#builder).

These transformations can be useful to ensure that [builders](https://grafana.github.io/cog/reference/glossary/#builder)
and [options](https://grafana.github.io/cog/reference/glossary/#builder-option)
expose meaningful names, make their APIs easier to understand and use, …

Builder transformations are defined within their own configuration files and
are then referenced in [codegen units](https://grafana.github.io/cog/reference/glossary/#unit).

**Common use-cases:**

* Avoid creating a builder for a specific object. For example, if the object only has one field, it's not necessary to create a builder.
* [Promote options to constructor arguments](https://github.com/grafana/grafana-foundation-sdk/blob/775cf20675c978d42f51cc3832c4cd20749b9ee2/go/dashboardv2beta1/dashboard_builder_gen.go#L16).
* Rename builders or options to expose friendlier names.
* And more...

> [!NOTE]
> See cog's documentation for a complete reference of the available [builder](https://grafana.github.io/cog/reference/builders_transformations/)
> and [option](https://grafana.github.io/cog/reference/options_transformations/) transformations.

Builder transformations should be declared in the [codegen unit](https://grafana.github.io/cog/reference/glossary/#unit)
defining a schema:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/unit.json
# .cog/resources/prometheus/config.yaml

inputs:
  - kindsys_composable:
      entrypoint: '%__config_dir%/'
      metadata:
        kind: composable
        variant: dataquery
        identifier: prometheus
      cue_imports:
        - '%kind_registry_path%/grafana/%kind_registry_version%/common:github.com/grafana/grafana/packages/grafana-schema/src/common'
      transformations:
        - '%__config_dir%/schema.transforms.yaml'

# List of files or directories containing builder/option transformations.
builder_transformations:
  - '%__config_dir%/common.builder.transforms.yaml'
```

And the transformations themselves might look like:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/veneers.json
# .cog/resources/prometheus/common.builder.transforms.yaml

# The transformations defined here will apply to every language supported by the SDK.
language: all

# The object names used in this file imply the following package.
package: prometheus

# These transformations are applied to builders.
builders:
  # Create a new option that sets two fields at once.
  # Like the UI, support a "both instant and range" query
  - add_option:
      by_object: Dataquery
      option:
        name: rangeAndInstant
        assignments:
          - path: range
            method: direct
            value: { constant: true }
          - path: instant
            method: direct
            value: { constant: true }

# These transformations are applied to options within builders.
options:
  # Ensure that enabling "range query mode" disables other modes
  - unfold_boolean:
      by_name: Dataquery.range
      true_as: range
      false_as: notRange
  - omit: { by_name: Dataquery.notRange }
  - add_assignment:
      by_name: Dataquery.range
      assignment:
        path: instant
        method: direct
        value: { constant: false }

  # Ensure that enabling "instant query mode" disables other modes
  - unfold_boolean:
      by_name: Dataquery.instant
      true_as: instant
      false_as: notInstant
  - omit: { by_name: Dataquery.notInstant }
  - add_assignment:
      by_name: Dataquery.instant
      assignment:
        path: range
        method: direct
        value: { constant: false }
```

## See also

* [Adding new resources](./adding-resources.md)
* [Schema transformations](./schema-transformations.md)
