// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Map a field to a color.
export class FieldColorBuilder implements cog.Builder<dashboard.FieldColor> {
    protected readonly internal: dashboard.FieldColor;

    constructor() {
        this.internal = dashboard.defaultFieldColor();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.FieldColor {
        return this.internal;
    }

    // The main color scheme mode.
    mode(mode: dashboard.FieldColorModeId): this {
        this.internal.mode = mode;
        return this;
    }

    // The fixed color value for fixed or shades color modes.
    fixedColor(fixedColor: string): this {
        this.internal.fixedColor = fixedColor;
        return this;
    }

    // Some visualizations need to know how to assign a series color from by value color schemes.
    seriesBy(seriesBy: dashboard.FieldColorSeriesByMode): this {
        this.internal.seriesBy = seriesBy;
        return this;
    }
}
