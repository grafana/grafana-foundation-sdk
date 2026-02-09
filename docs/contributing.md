# Contributing to the SDK

Foundation-SDK uses [cog](https://github.com/grafana/cog) to generate the SDK.

You can contribute to the SDK in several ways:
* Adding new schemas
* Improving existing schemas
* Fixing bugs
* Improving the documentation

## Adding New Schemas

To add a new schema, you must add it to the [configuration file](../.cog/post-12/config.yaml).
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

In general, this is enough to get your new schema working.

### What if We Want to Add a Schema from a Grafana Resource Without a Schema?

We have some schemas under [the schemas](../.cog/post-12/schemas) folder. This method is used as a last resort, and we try to avoid it as much as possible.
These schemas could become outdated at any time since they don't have active support, and we update them on demand.

The ideal approach is to find a replacement with a schema that is actively updated.

## Compiler Passes

Compiler passes are used to modify the generated code and are defined in the [compiler directory](../.cog/post-12/compiler).

In general, you have to fix the schema directly, but there are some cases where you might want to:

* **Omit fields**: Fields that exist in the schema but don't make sense to generate in Foundation SDK.
* **Add references**: Add references from a field to another schema. For example, datasources use a `retype_field` rule to use the same object in the `datasource` field.
* **Cog bugs**: Bugs in cog that aren't resolved yet.

Documentation is available [here](https://grafana.github.io/cog/reference/schema_transformations/).

## Veneers

Veneers are used to modify the generated builders.

They are mainly used to "prettify" the output, and it's desirable to add some to the schemas to improve the user experience.

**Use cases:**
* Avoid creating a builder for a specific object. For example, if the model only has one field, it's not necessary to create a builder.
* When you only have one field in your definition, and you want to ask for the field in the option instead of the whole object.
* Create new options. Useful when we have disjunctions and want to create different options for each value.
* Promote fields to have them as constructor arguments.
* And more...

Documentation is available [here](https://grafana.github.io/cog/reference/builders_transformations/).

## Generation

You can test your changes locally by running the following command:
```bash
make generate
```

Verify that:
- The new package folder is generated (in the case of a new schema)
- New changes are applied (if you modified an existing schema)

You only have to do a PR with the configuration changes, ask for a review, and merge it.

If you want to do a new release, you can follow the [release documentation](releasing.md).

> [!WARNING]
> Compiler passes and veneers are executed in the order they are defined in the configuration. If you add multiple modifications against the same field/object, the first action in the configuration will affect all later ones.

---

# Advanced Topics

## Working with Templates

Sometimes you need to add extra functionality to your code. To deal with that, you can create a Go template that will be included in the generated code.
These templates are located in the [templates](../.cog/post-12/templates) folder, and inside each language folder, you will find an `overrides` folder.

If you need to work with them, it is desirable to do it for every language, but not all languages need them. It can be complex because you need to know the syntax of each language.
You can find the kinds of accepted templates [here](https://github.com/grafana/cog/blob/main/internal/jennies/template/blocks.go). They have a specific structure depending on the result you want to achieve.

You can find documentation about the templates [here](https://grafana.github.io/cog/reference/template_blocks/).

Example: [WithPanel function](https://github.com/grafana/grafana-foundation-sdk/blob/next%2Bcog-v0.0.x/go/dashboard/dashboard_builder_gen.go#L185) + [properties](https://github.com/grafana/grafana-foundation-sdk/blob/next%2Bcog-v0.0.x/go/dashboard/dashboard_builder_gen.go#L14-L16).

## Missing Compiler Passes or Veneer Actions

If you need to add a new compiler pass or veneer action, you can add them in cog, or you can create an issue in the [cog repository](https://github.com/grafana/cog/issues).

**Examples:**
* Compiler pass: https://github.com/grafana/cog/pull/729
* Veneer: https://github.com/grafana/cog/pull/701
