// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as statushistory from '../statushistory';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<statushistory.Options> {
    protected readonly internal: statushistory.Options;

    constructor() {
        this.internal = statushistory.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): statushistory.Options {
        return this.internal;
    }

    // Set the height of the rows
    rowHeight(rowHeight: number): this {
        if (!(rowHeight >= 0)) {
            throw new Error("rowHeight must be >= 0");
        }
        if (!(rowHeight <= 1)) {
            throw new Error("rowHeight must be <= 1");
        }
        this.internal.rowHeight = rowHeight;
        return this;
    }

    // Show values on the columns
    showValue(showValue: common.VisibilityMode): this {
        this.internal.showValue = showValue;
        return this;
    }

    // Controls the column width
    colWidth(colWidth: number): this {
        if (!(colWidth <= 1)) {
            throw new Error("colWidth must be <= 1");
        }
        this.internal.colWidth = colWidth;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }

    timezone(timezone: common.TimeZone[]): this {
        this.internal.timezone = timezone;
        return this;
    }

    // Enables pagination when > 0
    perPage(perPage: number): this {
        if (!(perPage >= 1)) {
            throw new Error("perPage must be >= 1");
        }
        this.internal.perPage = perPage;
        return this;
    }
}

