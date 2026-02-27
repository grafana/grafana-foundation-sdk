# Adding new resources

> [!NOTE]
> The `make add-resource` command starts a CLI tool that will guide you through
> the process of adding a new resource to the SDK.

## Requirements

As described in [cog's documentation](https://grafana.github.io/cog/pipelines/creating_pipeline/#inputs),
a schema for the resources to add must exist in either CUE, jsonschema or OpenAPI.

Furthermore, only schemas owned by Grafana will be accepted and *versioned schemas*
are strongly preferred (e.g. the [v2beta1 dashboard schema](https://github.com/grafana/grafana/blob/main/apps/dashboard/kinds/v2beta1/dashboard_spec.cue)).

## Adding a schema to the codegen pipeline

New resources are expected to be defined as [*codegen units*](https://grafana.github.io/cog/pipelines/splitting_pipeline/) within the `.cog/resources` directory.

Taking the [`Playist.v0alpha1`](https://github.com/grafana/grafana/blob/2a14494b2d6ab60f860d8b27603d0ccb264336f6/apps/playlist/kinds/playlist.cue#L3)
resource as an example, the `.cog/resources/playlistv0alpha1/config.yaml` file
would be created with the following content:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/unit.json
# .cog/resources/playlistv0alpha1/config.yaml

inputs:
  - cue:
      # Where to find the schema.
      # Required.
      url: 'https://raw.githubusercontent.com/grafana/grafana/v12.3.2/apps/playlist/kinds/playlist.cue'

      # Name of the package in which code for this resource will be generated.
      # Including the resource name + its version is a good practice.
      # Optional, but explicitly setting this value is recommended.
      package: playlistv0alpha1

      # The actual schema for a playlist isn't at the root of the file, so we
      # have to tell cog where to find it.
      # Optional.
      subpath: playlistv0alpha1.schema.spec

      # The playlist schema is defined under the `spec` name, which cog will use
      # unless told otherwise. Playlist is a better name here.
      # Optional.
      forced_envelope: Playlist
```

> [!NOTE]
> See [cog's documentation](https://grafana.github.io/cog/pipelines/creating_pipeline/#inputs)
> to learn more about how to configure inputs.

Next, run the `make preview` command to verify that what is generated matches
your expectations.

If everything looks good, edit the [`.github/CODEOWNERS`](../.github/CODEOWNERS)
file to set your team as owner of the `.cog/resources/[new-resource]/` directory
and submit a PR.

> [!NOTE]
> Only the codegen configuration is needed to add a resource to the SDK. No need
> to commit and push the generated code as there is automation that handles it
> for us.

The following PRs can be used as examples of how to add new resources to the SDK:

* [Playlist.v0alpha1](https://github.com/grafana/grafana-foundation-sdk/pull/1107)

## Adding a panel or dataquery schema

To handle panel and dataqueries correctly, schemas need to explicitly be identified as such.

Example:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/grafana/cog/main/schemas/unit.json
# .cog/resources/expr/config.yaml

inputs:
  - jsonschema:
      url: 'https://raw.githubusercontent.com/grafana/grafana/v12.3.2/pkg/expr/query.request.schema.json'
      package: expr
      # This option lets cog know more about the schema being described.
      metadata:
        # Both panels and dataqueries are composable kinds.
        kind: composable
        # This particular schema is a dataquery, as opposed to "panelcfg" which would describe a panel.
        variant: dataquery
        # `identifier` is the identifier of this particular dataquery type (or panel) within Grafana itself.
        # e.g.: `prometheus`, `loki`, `timeseries`, `heatmap`, …
        identifier: __expr__
```

## There is no schema for the resource I want to add :(

Generally speaking, resources without a schema should not be a part of the Foundation SDK.

However, if a resource must be included and no schema for it can be defined elsewhere, its schema can be defined and maintained within the SDK itself. See the [`units`](../.cog/resources/units/) resource definition for an example.

This is method used as a last resort, and we try to avoid it as much as possible as these schemas would be at risk of quickly becoming outdated – being far removed from the application actually defining and using them.

## See also

* [Schema transformations](./schema-transformations.md)
* [Builder transformations](./builder-transformations.md)
