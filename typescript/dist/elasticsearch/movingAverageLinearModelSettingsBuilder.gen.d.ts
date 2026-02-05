import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class MovingAverageLinearModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageLinearModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageLinearModelSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageLinearModelSettings;
    window(window: string): this;
    predict(predict: string): this;
}
