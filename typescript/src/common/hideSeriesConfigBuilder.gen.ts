// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class HideSeriesConfigBuilder implements cog.Builder<common.HideSeriesConfig> {
    protected readonly internal: common.HideSeriesConfig;

    constructor() {
        this.internal = common.defaultHideSeriesConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.HideSeriesConfig {
        return this.internal;
    }

    tooltip(tooltip: boolean): this {
        this.internal.tooltip = tooltip;
        return this;
    }

    legend(legend: boolean): this {
        this.internal.legend = legend;
        return this;
    }

    viz(viz: boolean): this {
        this.internal.viz = viz;
        return this;
    }
}
