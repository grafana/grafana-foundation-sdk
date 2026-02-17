# Injecting custom logic into

Cog supports the injection of custom logic into types and builders.

For example, to automatically position panels on the grid whenever they're added to a dashboard or to customize how a type should be marshalled to JSON.

To achieve that, cog looks for [template blocks](https://pkg.go.dev/text/template#example-Template-Block) following
a particular naming convention, and renders them.

A [list of every supported template block](https://grafana.github.io/cog/reference/template_blocks/) is available in cog's documentation.


The Foundation SDK stores its own templates in the [`.cog/templates`](../.cog/templates/) directory. Inside each language folder within that directory, you will find an `overrides` folder with the relevant templates.

## See also

* [Adding new resources](./adding-resources.md)
* [Schema transformations](./schema-transformations.md)
* [Builder transformations](./builder-transformations.md)
