// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class GraphThresholdsStyleConfigBuilder implements cog.OptionsBuilder<common.GraphThresholdsStyleConfig> {
    private readonly internal: common.GraphThresholdsStyleConfig;

    constructor() {
        this.internal = common.defaultGraphThresholdsStyleConfig();
    }

    build(): common.GraphThresholdsStyleConfig {
        return this.internal;
    }

    mode(mode: common.GraphTresholdsStyleMode): this {
        this.internal.mode = mode;
        return this;
    }
}
