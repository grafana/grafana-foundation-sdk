# Contributing to the SDK

Foundation-SDK uses [cog](https://github.com/grafana/cog) to generate the SDK.

You can contribute to the SDK in several ways:
* Adding new schemas
* Improving existing schemas
* Fixing bugs
* Improving the documentation

## Adding New Schemas

To add a new schema, you must add it to the [configuration file](../.cog/config.yaml).
Valid schema types are CUE, JSON Schema, and OpenAPI.

> [!NOTE]
> We only accept schemas owned by Grafana, preferably ones that are currently supported (e.g., all new k8s schemas).
> Third-party libraries are not accepted. For those, check the [documentation](https://grafana.github.io/grafana-foundation-sdk/) to learn how to use them.

Configuration for each kind of input is documented in the [input cog documentation](https://grafana.github.io/cog/pipelines/creating_pipeline/#inputs).
It's important to fill in the metadata section with the correct information.
Metadata includes `package`, `kind`, and `variant`.

* **Package**: Used as the package name of the generated files and the folder for the generated code.
* **Kind**: Can be `core` or `composable`. `composable` it is used only for panels and datasources. All other schemas should be `core`.
* **Variant**: Mandatory when the kind is `composable`. Possible values are `panelcfg` and `dataquery`.

In general, this is sufficient to get your new schema working.

### What if We Want to Add a Schema from a Grafana Resource Without a Schema?

We have some schemas under [the schemas](../.cog/schemas) folder. This method is used as a last resort, and we try to avoid it as much as possible.
These schemas could become outdated at any time since they don't have active support, and we update them on demand.

The ideal approach is to find a replacement with a schema that is actively updated.

## Modifying Schemas

Sometimes, the version of the schema could change, or we may have a different source of truth.
In this case, you need to modify the schema. If this schema is used in Grafana versions prior to 12, it's recommended to use an `if` clause in the schema configuration to generate the schema under a version not supported.

> For example, `test datasource` doesn't exist in kind-registry from Grafana v11 onwards, so we use the `if` condition to generate the new schema only for Grafana v11+.

## Compiler Passes

Compiler passes are used to modify the generated code and are defined in the [compiler directory](../.cog/compiler). Documentation is available [here](https://grafana.github.io/cog/reference/schema_transformations/).

In general, you don't need to add them, but there are some cases where you might want to:
* **Missing fields**: This usually happens when the schema is not updated, and it's complex to do so. For example, legacy schemas in Grafana versions prior to 12. It's preferable to update the schema instead of adding the fields directly in the documentation, but sometimes it's complex.
* **Omit fields**: Fields that exist in the schema but don't make sense to generate.
* **Add references**: Add references from a field to another schema. For example, datasources use a `retype_field` rule to use the same object in the `datasource` field.
* **Cog bugs**: Bugs in cog that aren't resolved yet.
* And more...

## Veneers

Veneers are used to modify the generated builders.

They are mainly used to "prettify" the output, and it's desirable to add some to the schemas to improve the user experience.

**Use cases:**
* Avoid creating a builder for a specific object. For example, if the model only has one field, it's not necessary to create a builder.
* When you only have one field in your definition, and you want to ask for the field in the option instead of the whole object.
* Create new options. Useful when we have disjunctions and want to create different options for each value.
* Promote fields to have them as constructor arguments.
* And more...

## Generation and Publishing

To generate the code with your changes, run:
```bash
make generate
```

Verify that:
- The new package folder is generated (in the case of a new schema)
- New changes are applied (if you modified an existing schema)

After that, you can create a PR with the new changes.

After the PR is reviewed and merged, you only need to create and push a new tag to generate the new release.
The tag should follow the [semver](https://semver.org/) format. (e.g., v0.0.3)

> [!WARNING]
> Compiler passes and veneers are executed in the order they are defined in the configuration. If you add multiple modifications against the same field/object, the first action in the configuration will affect all subsequent ones.

---

# Advanced Topics

## Working with Templates

Sometimes you need to add extra functionality to your code. To deal with that, you can create a Go template that will be included in the generated code.
These templates are located in the [templates](../.cog/templates) folder, and inside each language folder, you will find an `overrides` folder.

If you need to work with them, it is desirable to do it for every language, but not all languages need them. It can be complex because you need to know the syntax of each language.
You can find the kinds of accepted templates [here](https://github.com/grafana/cog/blob/main/internal/jennies/template/blocks.go). They have a specific structure depending on the result you want to achieve.

### Template Types

* **object_\<package\>_\<object\>_custom_unmarshal**: Override the default unmarshaling of an object.
* **object_\<package\>_\<object\>_custom_strict_unmarshal**: Same as above but for strict unmarshaling.
* **api_reference_object_\<package\>_\<object\>_extra**: Add an extra information block for an object in documentation.
* **api_reference_builder_\<package\>_\<object\>_extra**: Add an extra information block for a builder in documentation.
* **api_reference_package_\<package\>_extra**: Add an extra information block for a package in documentation.
* **object_variant_\<variant\>**: Add extra methods to a variant.
* **object_\<package\>_\<object\>_custom_methods**: Add extra methods to an object.
* **schema_variant_\<variant\>**: Add extra code into an object to make the variants work properly.
* **variant_\<variant\>_field_unmarshal**: Define variant unmarshal functions.
* **dynamic_files**: Generate extra files.
* **assignment_\<package\>_\<object\>**: This one isn't defined in the template blocks, but it's used to generate custom assignment functions for builders. It can be combined with the `properties` compiler pass.

  Example: [WithPanel function](https://github.com/grafana/grafana-foundation-sdk/blob/next%2Bcog-v0.0.x/go/dashboard/dashboard_builder_gen.go#L185) + [properties](https://github.com/grafana/grafana-foundation-sdk/blob/next%2Bcog-v0.0.x/go/dashboard/dashboard_builder_gen.go#L14-L16).

## Missing Compiler Passes or Veneer Actions

If you need to add a new compiler pass or veneer action, you can add them in cog, or you can create an issue in the [cog repository](https://github.com/grafana/cog/issues).

**Examples:**
* Compiler pass: https://github.com/grafana/cog/pull/729
* Veneer: https://github.com/grafana/cog/pull/701

## Missing Flags in the Generate Command

Cog accepts a list of parameters that can be modified in the configuration file. These parameters are dynamic, and you can add more if you need to add checks in the configuration.

For example, imagine that you add a new schema that requires a minimum version of the SLO service. You can add `slo_version` as a parameter and update the Makefile. You can combine it in the `if` clause with the parameters you need.

More specific parameters that may affect language generation require cog updates.
