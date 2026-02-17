# Schema transformations

`cog` supports modifying its [intermediate representation](https://grafana.github.io/cog/reference/glossary/#intermediate-representation)
of a schema to add missing elements, fix incorrect ones, …

This feature can be useful when the original schemas can not be fixed directly. It should be used with parsimony and fixing the original schema should always be preferred.

Schema transformations should be declared in the [codegen unit](https://grafana.github.io/cog/reference/glossary/#unit)
defining a schema:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/unit.json
# .cog/resources/expr/config.yaml

inputs:
  - jsonschema:
      url: 'https://raw.githubusercontent.com/grafana/grafana/main/pkg/expr/query.request.schema.json'
      package: expr
      metadata:
        kind: composable
        variant: dataquery
        identifier: __expr__

      # List of files containing transformations to apply to the current schema.
      # These transformations will only be applied to the current input.
      transformations:
        - '%__config_dir%/schema.transforms.yaml'
```

And the transformations themselves might look like:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/compiler_passes.json
# .cog/resources/expr/schema.transforms.yaml

passes:
  # changes the type of the expr.typeMath.datasource field
  # to `common.DataSourceRef`
  - retype_field:
      field: expr.typeMath.datasource
      as: &common_datasourceRef
        kind: ref
        ref: { referred_pkg: common, referred_type: DataSourceRef }
  - retype_field:
      field: expr.typeReduce.datasource
      as: *common_datasourceRef
  - retype_field:
      field: expr.typeResample.datasource
      as: *common_datasourceRef
  - retype_field:
      field: expr.typeClassicConditions.datasource
      as: *common_datasourceRef
  - retype_field:
      field: expr.typeThreshold.datasource
      as: *common_datasourceRef
  - retype_field:
      field: expr.typeSql.datasource
      as: *common_datasourceRef

  # Changes the type of an entire object
  - retype_object:
      object: expr.Expr
      as:
        kind: disjunction
        hints: { implements_variant: dataquery }
        disjunction:
          branches:
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeMath } }
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeReduce } }
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeResample } }
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeClassicConditions } }
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeThreshold } }
            - { kind: ref, ref: { referred_pkg: expr, referred_type: TypeSql } }
```

> [!NOTE]
> See cog's documentation for a complete [reference of the available schema transformations](https://grafana.github.io/cog/reference/schema_transformations/).

## See also

* [Adding new resources](./adding-resources.md)
* [Builder transformations](./builder-transformations.md)
