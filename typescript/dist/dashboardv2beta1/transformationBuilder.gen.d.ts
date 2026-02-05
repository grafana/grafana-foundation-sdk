import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TransformationBuilder implements cog.Builder<dashboardv2beta1.TransformationKind> {
    protected readonly internal: dashboardv2beta1.TransformationKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TransformationKind;
    kind(kind: string): this;
    id(id: string): this;
    disabled(disabled: boolean): this;
    filter(filter: dashboardv2beta1.MatcherConfig): this;
    topic(topic: dashboardv2beta1.DataTopic): this;
    options(options: any): this;
}
