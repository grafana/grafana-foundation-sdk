---
title: <span class="badge object-type-enum"></span> DataqueryScenarioId
---
# <span class="badge object-type-enum"></span> DataqueryScenarioId

## Definition

```php
final class DataqueryScenarioId implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DataqueryScenarioId>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function annotations(): self
    {
        if (!isset(self::$instances["Annotations"])) {
            self::$instances["Annotations"] = new self("annotations");
        }

        return self::$instances["Annotations"];
    }

    public static function arrow(): self
    {
        if (!isset(self::$instances["Arrow"])) {
            self::$instances["Arrow"] = new self("arrow");
        }

        return self::$instances["Arrow"];
    }

    public static function csvContent(): self
    {
        if (!isset(self::$instances["CsvContent"])) {
            self::$instances["CsvContent"] = new self("csv_content");
        }

        return self::$instances["CsvContent"];
    }

    public static function csvFile(): self
    {
        if (!isset(self::$instances["CsvFile"])) {
            self::$instances["CsvFile"] = new self("csv_file");
        }

        return self::$instances["CsvFile"];
    }

    public static function csvMetricValues(): self
    {
        if (!isset(self::$instances["CsvMetricValues"])) {
            self::$instances["CsvMetricValues"] = new self("csv_metric_values");
        }

        return self::$instances["CsvMetricValues"];
    }

    public static function datapointsOutsideRange(): self
    {
        if (!isset(self::$instances["DatapointsOutsideRange"])) {
            self::$instances["DatapointsOutsideRange"] = new self("datapoints_outside_range");
        }

        return self::$instances["DatapointsOutsideRange"];
    }

    public static function errorWithSource(): self
    {
        if (!isset(self::$instances["ErrorWithSource"])) {
            self::$instances["ErrorWithSource"] = new self("error_with_source");
        }

        return self::$instances["ErrorWithSource"];
    }

    public static function exponentialHeatmapBucketData(): self
    {
        if (!isset(self::$instances["ExponentialHeatmapBucketData"])) {
            self::$instances["ExponentialHeatmapBucketData"] = new self("exponential_heatmap_bucket_data");
        }

        return self::$instances["ExponentialHeatmapBucketData"];
    }

    public static function flameGraph(): self
    {
        if (!isset(self::$instances["FlameGraph"])) {
            self::$instances["FlameGraph"] = new self("flame_graph");
        }

        return self::$instances["FlameGraph"];
    }

    public static function grafanaApi(): self
    {
        if (!isset(self::$instances["GrafanaApi"])) {
            self::$instances["GrafanaApi"] = new self("grafana_api");
        }

        return self::$instances["GrafanaApi"];
    }

    public static function linearHeatmapBucketData(): self
    {
        if (!isset(self::$instances["LinearHeatmapBucketData"])) {
            self::$instances["LinearHeatmapBucketData"] = new self("linear_heatmap_bucket_data");
        }

        return self::$instances["LinearHeatmapBucketData"];
    }

    public static function live(): self
    {
        if (!isset(self::$instances["Live"])) {
            self::$instances["Live"] = new self("live");
        }

        return self::$instances["Live"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["Logs"])) {
            self::$instances["Logs"] = new self("logs");
        }

        return self::$instances["Logs"];
    }

    public static function manualEntry(): self
    {
        if (!isset(self::$instances["ManualEntry"])) {
            self::$instances["ManualEntry"] = new self("manual_entry");
        }

        return self::$instances["ManualEntry"];
    }

    public static function noDataPoints(): self
    {
        if (!isset(self::$instances["NoDataPoints"])) {
            self::$instances["NoDataPoints"] = new self("no_data_points");
        }

        return self::$instances["NoDataPoints"];
    }

    public static function nodeGraph(): self
    {
        if (!isset(self::$instances["NodeGraph"])) {
            self::$instances["NodeGraph"] = new self("node_graph");
        }

        return self::$instances["NodeGraph"];
    }

    public static function predictableCsvWave(): self
    {
        if (!isset(self::$instances["PredictableCsvWave"])) {
            self::$instances["PredictableCsvWave"] = new self("predictable_csv_wave");
        }

        return self::$instances["PredictableCsvWave"];
    }

    public static function predictablePulse(): self
    {
        if (!isset(self::$instances["PredictablePulse"])) {
            self::$instances["PredictablePulse"] = new self("predictable_pulse");
        }

        return self::$instances["PredictablePulse"];
    }

    public static function randomWalk(): self
    {
        if (!isset(self::$instances["RandomWalk"])) {
            self::$instances["RandomWalk"] = new self("random_walk");
        }

        return self::$instances["RandomWalk"];
    }

    public static function randomWalkTable(): self
    {
        if (!isset(self::$instances["RandomWalkTable"])) {
            self::$instances["RandomWalkTable"] = new self("random_walk_table");
        }

        return self::$instances["RandomWalkTable"];
    }

    public static function randomWalkWithError(): self
    {
        if (!isset(self::$instances["RandomWalkWithError"])) {
            self::$instances["RandomWalkWithError"] = new self("random_walk_with_error");
        }

        return self::$instances["RandomWalkWithError"];
    }

    public static function rawFrame(): self
    {
        if (!isset(self::$instances["RawFrame"])) {
            self::$instances["RawFrame"] = new self("raw_frame");
        }

        return self::$instances["RawFrame"];
    }

    public static function serverError500(): self
    {
        if (!isset(self::$instances["ServerError500"])) {
            self::$instances["ServerError500"] = new self("server_error_500");
        }

        return self::$instances["ServerError500"];
    }

    public static function simulation(): self
    {
        if (!isset(self::$instances["Simulation"])) {
            self::$instances["Simulation"] = new self("simulation");
        }

        return self::$instances["Simulation"];
    }

    public static function slowQuery(): self
    {
        if (!isset(self::$instances["SlowQuery"])) {
            self::$instances["SlowQuery"] = new self("slow_query");
        }

        return self::$instances["SlowQuery"];
    }

    public static function streamingClient(): self
    {
        if (!isset(self::$instances["StreamingClient"])) {
            self::$instances["StreamingClient"] = new self("streaming_client");
        }

        return self::$instances["StreamingClient"];
    }

    public static function tableStatic(): self
    {
        if (!isset(self::$instances["TableStatic"])) {
            self::$instances["TableStatic"] = new self("table_static");
        }

        return self::$instances["TableStatic"];
    }

    public static function trace(): self
    {
        if (!isset(self::$instances["Trace"])) {
            self::$instances["Trace"] = new self("trace");
        }

        return self::$instances["Trace"];
    }

    public static function usa(): self
    {
        if (!isset(self::$instances["Usa"])) {
            self::$instances["Usa"] = new self("usa");
        }

        return self::$instances["Usa"];
    }

    public static function variablesQuery(): self
    {
        if (!isset(self::$instances["VariablesQuery"])) {
            self::$instances["VariablesQuery"] = new self("variables-query");
        }

        return self::$instances["VariablesQuery"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "annotations") {
            return self::annotations();
        }

        if ($value === "arrow") {
            return self::arrow();
        }

        if ($value === "csv_content") {
            return self::csvContent();
        }

        if ($value === "csv_file") {
            return self::csvFile();
        }

        if ($value === "csv_metric_values") {
            return self::csvMetricValues();
        }

        if ($value === "datapoints_outside_range") {
            return self::datapointsOutsideRange();
        }

        if ($value === "error_with_source") {
            return self::errorWithSource();
        }

        if ($value === "exponential_heatmap_bucket_data") {
            return self::exponentialHeatmapBucketData();
        }

        if ($value === "flame_graph") {
            return self::flameGraph();
        }

        if ($value === "grafana_api") {
            return self::grafanaApi();
        }

        if ($value === "linear_heatmap_bucket_data") {
            return self::linearHeatmapBucketData();
        }

        if ($value === "live") {
            return self::live();
        }

        if ($value === "logs") {
            return self::logs();
        }

        if ($value === "manual_entry") {
            return self::manualEntry();
        }

        if ($value === "no_data_points") {
            return self::noDataPoints();
        }

        if ($value === "node_graph") {
            return self::nodeGraph();
        }

        if ($value === "predictable_csv_wave") {
            return self::predictableCsvWave();
        }

        if ($value === "predictable_pulse") {
            return self::predictablePulse();
        }

        if ($value === "random_walk") {
            return self::randomWalk();
        }

        if ($value === "random_walk_table") {
            return self::randomWalkTable();
        }

        if ($value === "random_walk_with_error") {
            return self::randomWalkWithError();
        }

        if ($value === "raw_frame") {
            return self::rawFrame();
        }

        if ($value === "server_error_500") {
            return self::serverError500();
        }

        if ($value === "simulation") {
            return self::simulation();
        }

        if ($value === "slow_query") {
            return self::slowQuery();
        }

        if ($value === "streaming_client") {
            return self::streamingClient();
        }

        if ($value === "table_static") {
            return self::tableStatic();
        }

        if ($value === "trace") {
            return self::trace();
        }

        if ($value === "usa") {
            return self::usa();
        }

        if ($value === "variables-query") {
            return self::variablesQuery();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DataqueryScenarioId");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

```
