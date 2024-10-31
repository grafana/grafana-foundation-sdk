// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Support for legacy graph panel.
// @deprecated this a deprecated panel type
export class GraphPanelBuilder implements cog.Builder<dashboard.GraphPanel> {
    protected readonly internal: dashboard.GraphPanel;

    constructor() {
        this.internal = dashboard.defaultGraphPanel();
        this.internal.type = "graph";
    }

    /**
     * Builds the object.
     */
    build(): dashboard.GraphPanel {
        return this.internal;
    }

    // @deprecated this is part of deprecated graph panel
    legend(legend: {
	show: boolean;
	sort?: string;
	sortDesc?: boolean;
}): this {
        this.internal.legend = legend;
        return this;
    }
}
