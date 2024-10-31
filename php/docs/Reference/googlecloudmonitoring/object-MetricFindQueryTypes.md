---
title: <span class="badge object-type-enum"></span> MetricFindQueryTypes
---
# <span class="badge object-type-enum"></span> MetricFindQueryTypes

## Definition

```php
final class MetricFindQueryTypes implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MetricFindQueryTypes>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function projects(): self
    {
        if (!isset(self::$instances["Projects"])) {
            self::$instances["Projects"] = new self("projects");
        }

        return self::$instances["Projects"];
    }

    public static function services(): self
    {
        if (!isset(self::$instances["Services"])) {
            self::$instances["Services"] = new self("services");
        }

        return self::$instances["Services"];
    }

    public static function defaultProject(): self
    {
        if (!isset(self::$instances["DefaultProject"])) {
            self::$instances["DefaultProject"] = new self("defaultProject");
        }

        return self::$instances["DefaultProject"];
    }

    public static function metricTypes(): self
    {
        if (!isset(self::$instances["MetricTypes"])) {
            self::$instances["MetricTypes"] = new self("metricTypes");
        }

        return self::$instances["MetricTypes"];
    }

    public static function labelKeys(): self
    {
        if (!isset(self::$instances["LabelKeys"])) {
            self::$instances["LabelKeys"] = new self("labelKeys");
        }

        return self::$instances["LabelKeys"];
    }

    public static function labelValues(): self
    {
        if (!isset(self::$instances["LabelValues"])) {
            self::$instances["LabelValues"] = new self("labelValues");
        }

        return self::$instances["LabelValues"];
    }

    public static function resourceTypes(): self
    {
        if (!isset(self::$instances["ResourceTypes"])) {
            self::$instances["ResourceTypes"] = new self("resourceTypes");
        }

        return self::$instances["ResourceTypes"];
    }

    public static function aggregations(): self
    {
        if (!isset(self::$instances["Aggregations"])) {
            self::$instances["Aggregations"] = new self("aggregations");
        }

        return self::$instances["Aggregations"];
    }

    public static function aligners(): self
    {
        if (!isset(self::$instances["Aligners"])) {
            self::$instances["Aligners"] = new self("aligners");
        }

        return self::$instances["Aligners"];
    }

    public static function alignmentPeriods(): self
    {
        if (!isset(self::$instances["AlignmentPeriods"])) {
            self::$instances["AlignmentPeriods"] = new self("alignmentPeriods");
        }

        return self::$instances["AlignmentPeriods"];
    }

    public static function selectors(): self
    {
        if (!isset(self::$instances["Selectors"])) {
            self::$instances["Selectors"] = new self("selectors");
        }

        return self::$instances["Selectors"];
    }

    public static function sLOServices(): self
    {
        if (!isset(self::$instances["SLOServices"])) {
            self::$instances["SLOServices"] = new self("sloServices");
        }

        return self::$instances["SLOServices"];
    }

    public static function sLO(): self
    {
        if (!isset(self::$instances["SLO"])) {
            self::$instances["SLO"] = new self("slo");
        }

        return self::$instances["SLO"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "projects") {
            return self::projects();
        }

        if ($value === "services") {
            return self::services();
        }

        if ($value === "defaultProject") {
            return self::defaultProject();
        }

        if ($value === "metricTypes") {
            return self::metricTypes();
        }

        if ($value === "labelKeys") {
            return self::labelKeys();
        }

        if ($value === "labelValues") {
            return self::labelValues();
        }

        if ($value === "resourceTypes") {
            return self::resourceTypes();
        }

        if ($value === "aggregations") {
            return self::aggregations();
        }

        if ($value === "aligners") {
            return self::aligners();
        }

        if ($value === "alignmentPeriods") {
            return self::alignmentPeriods();
        }

        if ($value === "selectors") {
            return self::selectors();
        }

        if ($value === "sloServices") {
            return self::sLOServices();
        }

        if ($value === "slo") {
            return self::sLO();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricFindQueryTypes");
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
