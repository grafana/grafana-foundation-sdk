import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';

export interface CustomPanelOptions {
  makeBeautiful?: boolean;
}

export const defaultCustomPanelOptions = (): CustomPanelOptions => ({});

export class CustomPanelBuilder extends dashboard.PanelBuilder {
  constructor() {
    super();

    this.internal.type = "custom-panel"; // panel plugin ID
  }

  makeBeautiful(): this {
    if (!this.internal.options) {
      this.internal.options = defaultCustomPanelOptions();
    }
    this.internal.options.makeBeautiful = true;
    return this;
  }
}
