---
title: <span class="badge object-type-enum"></span> VariableType
---
# <span class="badge object-type-enum"></span> VariableType

Dashboard variable type

`query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.

`adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).

`constant`: 	Define a hidden constant.

`datasource`: Quickly change the data source for an entire dashboard.

`interval`: Interval variables represent time spans.

`textbox`: Display a free text input field with an optional default value.

`custom`: Define the variable options manually using a comma-separated list.

`system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables

## Definition

```php
final class VariableType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function query(): self
    {
        if (!isset(self::$instances["query"])) {
            self::$instances["query"] = new self("query");
        }

        return self::$instances["query"];
    }

    public static function adhoc(): self
    {
        if (!isset(self::$instances["adhoc"])) {
            self::$instances["adhoc"] = new self("adhoc");
        }

        return self::$instances["adhoc"];
    }

    public static function constant(): self
    {
        if (!isset(self::$instances["constant"])) {
            self::$instances["constant"] = new self("constant");
        }

        return self::$instances["constant"];
    }

    public static function datasource(): self
    {
        if (!isset(self::$instances["datasource"])) {
            self::$instances["datasource"] = new self("datasource");
        }

        return self::$instances["datasource"];
    }

    public static function interval(): self
    {
        if (!isset(self::$instances["interval"])) {
            self::$instances["interval"] = new self("interval");
        }

        return self::$instances["interval"];
    }

    public static function textbox(): self
    {
        if (!isset(self::$instances["textbox"])) {
            self::$instances["textbox"] = new self("textbox");
        }

        return self::$instances["textbox"];
    }

    public static function custom(): self
    {
        if (!isset(self::$instances["custom"])) {
            self::$instances["custom"] = new self("custom");
        }

        return self::$instances["custom"];
    }

    public static function system(): self
    {
        if (!isset(self::$instances["system"])) {
            self::$instances["system"] = new self("system");
        }

        return self::$instances["system"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "query") {
            return self::query();
        }

        if ($value === "adhoc") {
            return self::adhoc();
        }

        if ($value === "constant") {
            return self::constant();
        }

        if ($value === "datasource") {
            return self::datasource();
        }

        if ($value === "interval") {
            return self::interval();
        }

        if ($value === "textbox") {
            return self::textbox();
        }

        if ($value === "custom") {
            return self::custom();
        }

        if ($value === "system") {
            return self::system();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableType");
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
