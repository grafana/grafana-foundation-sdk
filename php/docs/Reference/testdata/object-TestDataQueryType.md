---
title: <span class="badge object-type-enum"></span> TestDataQueryType
---
# <span class="badge object-type-enum"></span> TestDataQueryType

## Definition

```php
final class TestDataQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TestDataQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function randomWalk(): self
    {
        if (!isset(self::$instances["RandomWalk"])) {
            self::$instances["RandomWalk"] = new self("random_walk");
        }

        return self::$instances["RandomWalk"];
    }

    public static function slowQuery(): self
    {
        if (!isset(self::$instances["SlowQuery"])) {
            self::$instances["SlowQuery"] = new self("slow_query");
        }

        return self::$instances["SlowQuery"];
    }

    public static function randomWalkWithError(): self
    {
        if (!isset(self::$instances["RandomWalkWithError"])) {
            self::$instances["RandomWalkWithError"] = new self("random_walk_with_error");
        }

        return self::$instances["RandomWalkWithError"];
    }

    public static function randomWalkTable(): self
    {
        if (!isset(self::$instances["RandomWalkTable"])) {
            self::$instances["RandomWalkTable"] = new self("random_walk_table");
        }

        return self::$instances["RandomWalkTable"];
    }

    public static function exponentialHeatmapBucketData(): self
    {
        if (!isset(self::$instances["ExponentialHeatmapBucketData"])) {
            self::$instances["ExponentialHeatmapBucketData"] = new self("exponential_heatmap_bucket_data");
        }

        return self::$instances["ExponentialHeatmapBucketData"];
    }

    public static function linearHeatmapBucketData(): self
    {
        if (!isset(self::$instances["LinearHeatmapBucketData"])) {
            self::$instances["LinearHeatmapBucketData"] = new self("linear_heatmap_bucket_data");
        }

        return self::$instances["LinearHeatmapBucketData"];
    }

    public static function noDataPoints(): self
    {
        if (!isset(self::$instances["NoDataPoints"])) {
            self::$instances["NoDataPoints"] = new self("no_data_points");
        }

        return self::$instances["NoDataPoints"];
    }

    public static function dataPointsOutsideRange(): self
    {
        if (!isset(self::$instances["DataPointsOutsideRange"])) {
            self::$instances["DataPointsOutsideRange"] = new self("datapoints_outside_range");
        }

        return self::$instances["DataPointsOutsideRange"];
    }

    public static function cSVMetricValues(): self
    {
        if (!isset(self::$instances["CSVMetricValues"])) {
            self::$instances["CSVMetricValues"] = new self("csv_metric_values");
        }

        return self::$instances["CSVMetricValues"];
    }

    public static function predictablePulse(): self
    {
        if (!isset(self::$instances["PredictablePulse"])) {
            self::$instances["PredictablePulse"] = new self("predictable_pulse");
        }

        return self::$instances["PredictablePulse"];
    }

    public static function predictableCSVWave(): self
    {
        if (!isset(self::$instances["PredictableCSVWave"])) {
            self::$instances["PredictableCSVWave"] = new self("predictable_csv_wave");
        }

        return self::$instances["PredictableCSVWave"];
    }

    public static function streamingClient(): self
    {
        if (!isset(self::$instances["StreamingClient"])) {
            self::$instances["StreamingClient"] = new self("streaming_client");
        }

        return self::$instances["StreamingClient"];
    }

    public static function simulation(): self
    {
        if (!isset(self::$instances["Simulation"])) {
            self::$instances["Simulation"] = new self("simulation");
        }

        return self::$instances["Simulation"];
    }

    public static function uSA(): self
    {
        if (!isset(self::$instances["USA"])) {
            self::$instances["USA"] = new self("usa");
        }

        return self::$instances["USA"];
    }

    public static function live(): self
    {
        if (!isset(self::$instances["Live"])) {
            self::$instances["Live"] = new self("live");
        }

        return self::$instances["Live"];
    }

    public static function grafanaAPI(): self
    {
        if (!isset(self::$instances["GrafanaAPI"])) {
            self::$instances["GrafanaAPI"] = new self("grafana_api");
        }

        return self::$instances["GrafanaAPI"];
    }

    public static function arrow(): self
    {
        if (!isset(self::$instances["Arrow"])) {
            self::$instances["Arrow"] = new self("arrow");
        }

        return self::$instances["Arrow"];
    }

    public static function annotations(): self
    {
        if (!isset(self::$instances["Annotations"])) {
            self::$instances["Annotations"] = new self("annotations");
        }

        return self::$instances["Annotations"];
    }

    public static function tableStatic(): self
    {
        if (!isset(self::$instances["TableStatic"])) {
            self::$instances["TableStatic"] = new self("table_static");
        }

        return self::$instances["TableStatic"];
    }

    public static function serverError500(): self
    {
        if (!isset(self::$instances["ServerError500"])) {
            self::$instances["ServerError500"] = new self("server_error_500");
        }

        return self::$instances["ServerError500"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["Logs"])) {
            self::$instances["Logs"] = new self("logs");
        }

        return self::$instances["Logs"];
    }

    public static function nodeGraph(): self
    {
        if (!isset(self::$instances["NodeGraph"])) {
            self::$instances["NodeGraph"] = new self("node_graph");
        }

        return self::$instances["NodeGraph"];
    }

    public static function flameGraph(): self
    {
        if (!isset(self::$instances["FlameGraph"])) {
            self::$instances["FlameGraph"] = new self("flame_graph");
        }

        return self::$instances["FlameGraph"];
    }

    public static function rawFrame(): self
    {
        if (!isset(self::$instances["RawFrame"])) {
            self::$instances["RawFrame"] = new self("raw_frame");
        }

        return self::$instances["RawFrame"];
    }

    public static function cSVFile(): self
    {
        if (!isset(self::$instances["CSVFile"])) {
            self::$instances["CSVFile"] = new self("csv_file");
        }

        return self::$instances["CSVFile"];
    }

    public static function cSVContent(): self
    {
        if (!isset(self::$instances["CSVContent"])) {
            self::$instances["CSVContent"] = new self("csv_content");
        }

        return self::$instances["CSVContent"];
    }

    public static function trace(): self
    {
        if (!isset(self::$instances["Trace"])) {
            self::$instances["Trace"] = new self("trace");
        }

        return self::$instances["Trace"];
    }

    public static function manualEntry(): self
    {
        if (!isset(self::$instances["ManualEntry"])) {
            self::$instances["ManualEntry"] = new self("manual_entry");
        }

        return self::$instances["ManualEntry"];
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
        if ($value === "random_walk") {
            return self::randomWalk();
        }

        if ($value === "slow_query") {
            return self::slowQuery();
        }

        if ($value === "random_walk_with_error") {
            return self::randomWalkWithError();
        }

        if ($value === "random_walk_table") {
            return self::randomWalkTable();
        }

        if ($value === "exponential_heatmap_bucket_data") {
            return self::exponentialHeatmapBucketData();
        }

        if ($value === "linear_heatmap_bucket_data") {
            return self::linearHeatmapBucketData();
        }

        if ($value === "no_data_points") {
            return self::noDataPoints();
        }

        if ($value === "datapoints_outside_range") {
            return self::dataPointsOutsideRange();
        }

        if ($value === "csv_metric_values") {
            return self::cSVMetricValues();
        }

        if ($value === "predictable_pulse") {
            return self::predictablePulse();
        }

        if ($value === "predictable_csv_wave") {
            return self::predictableCSVWave();
        }

        if ($value === "streaming_client") {
            return self::streamingClient();
        }

        if ($value === "simulation") {
            return self::simulation();
        }

        if ($value === "usa") {
            return self::uSA();
        }

        if ($value === "live") {
            return self::live();
        }

        if ($value === "grafana_api") {
            return self::grafanaAPI();
        }

        if ($value === "arrow") {
            return self::arrow();
        }

        if ($value === "annotations") {
            return self::annotations();
        }

        if ($value === "table_static") {
            return self::tableStatic();
        }

        if ($value === "server_error_500") {
            return self::serverError500();
        }

        if ($value === "logs") {
            return self::logs();
        }

        if ($value === "node_graph") {
            return self::nodeGraph();
        }

        if ($value === "flame_graph") {
            return self::flameGraph();
        }

        if ($value === "raw_frame") {
            return self::rawFrame();
        }

        if ($value === "csv_file") {
            return self::cSVFile();
        }

        if ($value === "csv_content") {
            return self::cSVContent();
        }

        if ($value === "trace") {
            return self::trace();
        }

        if ($value === "manual_entry") {
            return self::manualEntry();
        }

        if ($value === "variables-query") {
            return self::variablesQuery();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TestDataQueryType");
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
