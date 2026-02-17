# Navigating the repository

## Directory layout

* `./.cog/`: Codegen configuration files.
* `./.github/`: GitHub-specific configuration files.
* `./.mkdocs/`: Configuration and non language-specific files used to generate the public website.
* `./examples/`: Contains examples scripts showcasing how to build some resources (dashboards, ...) in various languages. Used both to serve as inspiration as well as an experimentation playground for library features.
* `./scripts/`: Various scripts related to the codegen or release processes.
* `./(go|java|jsonschema|openapi|php|python|typescript)/`: Generated SDK codebase, for each supported language.

## Codegen pipeline configuration

Codegen configuration files live in the `./.cog/` folder. It contains the following:

* [`./config.yaml`](../.cog/config.yaml): entrypoint of the codegen configuration. It describes the entire [codegen pipeline](https://grafana.github.io/cog/pipelines/creating_pipeline/).
* [`./resources/*`](../.cog/resources/): each folder within the `./resources/` directory contains a [codegen unit](https://grafana.github.io/cog/pipelines/splitting_pipeline/) describing one or several resources. E.g: `timeseries` panel, `prometheus` query, …
* [`./templates/*`](../.cog/templates/): contains language-specific [templates](https://grafana.github.io/cog/reference/template_blocks/) used by cog to customize the generated code.

## See also

* [Adding new resources](./adding-resources.md)
* [Schema transformations](./schema-transformations.md)
* [Builder transformations](./builder-transformations.md)
