// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class LineStyleBuilder implements cog.Builder<common.LineStyle> {
    protected readonly internal: common.LineStyle;

    constructor() {
        this.internal = common.defaultLineStyle();
    }

    /**
     * Builds the object.
     */
    build(): common.LineStyle {
        return this.internal;
    }

    fill(fill: "solid" | "dash" | "dot" | "square"): this {
        this.internal.fill = fill;
        return this;
    }

    dash(dash: number[]): this {
        this.internal.dash = dash;
        return this;
    }
}
