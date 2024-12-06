// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as nodegraph from '../nodegraph';

export class EdgeOptionsBuilder implements cog.Builder<nodegraph.EdgeOptions> {
    protected readonly internal: nodegraph.EdgeOptions;

    constructor() {
        this.internal = nodegraph.defaultEdgeOptions();
    }

    /**
     * Builds the object.
     */
    build(): nodegraph.EdgeOptions {
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
}
