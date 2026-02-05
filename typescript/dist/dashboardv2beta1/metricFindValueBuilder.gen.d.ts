import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class MetricFindValueBuilder implements cog.Builder<dashboardv2beta1.MetricFindValue> {
    protected readonly internal: dashboardv2beta1.MetricFindValue;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.MetricFindValue;
    text(text: string): this;
    value(value: string | number): this;
    group(group: string): this;
    expandable(expandable: boolean): this;
}
