// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as nodegraph from '../nodegraph';

export class OptionsBuilder implements cog.Builder<nodegraph.Options> {
    protected readonly internal: nodegraph.Options;

    constructor() {
        this.internal = nodegraph.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): nodegraph.Options {
        return this.internal;
    }

    nodes(nodes: cog.Builder<nodegraph.NodeOptions>): this {
        const nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }

    edges(edges: cog.Builder<nodegraph.EdgeOptions>): this {
        const edgesResource = edges.build();
        this.internal.edges = edgesResource;
        return this;
    }
}

