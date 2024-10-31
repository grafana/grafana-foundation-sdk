---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes.

## Definition

```php
final class QueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function tIMESERIESLIST(): self
    {
        if (!isset(self::$instances["TIME_SERIES_LIST"])) {
            self::$instances["TIME_SERIES_LIST"] = new self("timeSeriesList");
        }

        return self::$instances["TIME_SERIES_LIST"];
    }

    public static function tIMESERIESQUERY(): self
    {
        if (!isset(self::$instances["TIME_SERIES_QUERY"])) {
            self::$instances["TIME_SERIES_QUERY"] = new self("timeSeriesQuery");
        }

        return self::$instances["TIME_SERIES_QUERY"];
    }

    public static function sLO(): self
    {
        if (!isset(self::$instances["SLO"])) {
            self::$instances["SLO"] = new self("slo");
        }

        return self::$instances["SLO"];
    }

    public static function aNNOTATION(): self
    {
        if (!isset(self::$instances["ANNOTATION"])) {
            self::$instances["ANNOTATION"] = new self("annotation");
        }

        return self::$instances["ANNOTATION"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "timeSeriesList") {
            return self::tIMESERIESLIST();
        }

        if ($value === "timeSeriesQuery") {
            return self::tIMESERIESQUERY();
        }

        if ($value === "slo") {
            return self::sLO();
        }

        if ($value === "annotation") {
            return self::aNNOTATION();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryType");
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
