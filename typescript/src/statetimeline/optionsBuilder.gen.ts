// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as statetimeline from '../statetimeline';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<statetimeline.Options> {
    protected readonly internal: statetimeline.Options;

    constructor() {
        this.internal = statetimeline.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): statetimeline.Options {
        return this.internal;
    }

    // Show timeline values on chart
    showValue(showValue: common.VisibilityMode): this {
        this.internal.showValue = showValue;
        return this;
    }

    // Controls the row height
    rowHeight(rowHeight: number): this {
        if (!(rowHeight <= 1)) {
            throw new Error("rowHeight must be <= 1");
        }
        this.internal.rowHeight = rowHeight;
        return this;
    }

    // Merge equal consecutive values
    mergeValues(mergeValues: boolean): this {
        this.internal.mergeValues = mergeValues;
        return this;
    }

    // Controls value alignment on the timelines
    alignValue(alignValue: common.TimelineValueAlignment): this {
        this.internal.alignValue = alignValue;
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

