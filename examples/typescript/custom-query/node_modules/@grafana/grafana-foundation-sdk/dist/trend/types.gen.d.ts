import * as common from '../common';
export interface Options {
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    xField?: string;
}
export declare const defaultOptions: () => Options;
export type FieldConfig = common.GraphFieldConfig;
export declare const defaultFieldConfig: () => FieldConfig;
