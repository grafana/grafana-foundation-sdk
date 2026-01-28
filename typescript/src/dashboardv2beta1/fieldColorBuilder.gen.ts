// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Map a field to a color.
export class FieldColorBuilder implements cog.Builder<dashboardv2beta1.FieldColor> {
    protected readonly internal: dashboardv2beta1.FieldColor;

    constructor() {
        this.internal = dashboardv2beta1.defaultFieldColor();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.FieldColor {
        return this.internal;
    }

    // The main color scheme mode.
    mode(mode: dashboardv2beta1.FieldColorModeId): this {
        this.internal.mode = mode;
        return this;
    }

    // The fixed color value for fixed or shades color modes.
    fixedColor(fixedColor: string): this {
        this.internal.fixedColor = fixedColor;
        return this;
    }

    // Some visualizations need to know how to assign a series color from by value color schemes.
    seriesBy(seriesBy: dashboardv2beta1.FieldColorSeriesByMode): this {
        this.internal.seriesBy = seriesBy;
        return this;
    }
}

