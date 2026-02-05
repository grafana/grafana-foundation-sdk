import * as common from '../common';
export interface Options {
    timezone?: common.TimeZone[];
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    orientation?: common.VizOrientation;
}
export declare const defaultOptions: () => Options;
export type FieldConfig = common.GraphFieldConfig;
export declare const defaultFieldConfig: () => FieldConfig;
