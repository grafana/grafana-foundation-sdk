// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as nodegraph from '../nodegraph';

export class ArcOptionBuilder implements cog.Builder<nodegraph.ArcOption> {
    protected readonly internal: nodegraph.ArcOption;

    constructor() {
        this.internal = nodegraph.defaultArcOption();
    }

    /**
     * Builds the object.
     */
    build(): nodegraph.ArcOption {
        return this.internal;
    }

    // Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    // The color of the arc.
    color(color: string): this {
        this.internal.color = color;
        return this;
    }
}
