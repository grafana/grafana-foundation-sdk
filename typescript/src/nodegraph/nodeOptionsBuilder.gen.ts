// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as nodegraph from '../nodegraph';

export class NodeOptionsBuilder implements cog.Builder<nodegraph.NodeOptions> {
    protected readonly internal: nodegraph.NodeOptions;

    constructor() {
        this.internal = nodegraph.defaultNodeOptions();
    }

    /**
     * Builds the object.
     */
    build(): nodegraph.NodeOptions {
        return this.internal;
    }

    // Unit for the main stat to override what ever is set in the data frame.
    mainStatUnit(mainStatUnit: string): this {
        this.internal.mainStatUnit = mainStatUnit;
        return this;
    }

    // Unit for the secondary stat to override what ever is set in the data frame.
    secondaryStatUnit(secondaryStatUnit: string): this {
        this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }

    // Define which fields are shown as part of the node arc (colored circle around the node).
    arcs(arcs: cog.Builder<nodegraph.ArcOption>[]): this {
        const arcsResources = arcs.map(builder1 => builder1.build());
        this.internal.arcs = arcsResources;
        return this;
    }
}
