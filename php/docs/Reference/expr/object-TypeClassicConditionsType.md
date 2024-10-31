---
title: <span class="badge object-type-enum"></span> TypeClassicConditionsType
---
# <span class="badge object-type-enum"></span> TypeClassicConditionsType

## Definition

```php
final class TypeClassicConditionsType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeClassicConditionsType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("");
        }

        return self::$instances["None"];
    }

    public static function timeseriesWide(): self
    {
        if (!isset(self::$instances["TimeseriesWide"])) {
            self::$instances["TimeseriesWide"] = new self("timeseries-wide");
        }

        return self::$instances["TimeseriesWide"];
    }

    public static function timeseriesLong(): self
    {
        if (!isset(self::$instances["TimeseriesLong"])) {
            self::$instances["TimeseriesLong"] = new self("timeseries-long");
        }

        return self::$instances["TimeseriesLong"];
    }

    public static function timeseriesMany(): self
    {
        if (!isset(self::$instances["TimeseriesMany"])) {
            self::$instances["TimeseriesMany"] = new self("timeseries-many");
        }

        return self::$instances["TimeseriesMany"];
    }

    public static function timeseriesMulti(): self
    {
        if (!isset(self::$instances["TimeseriesMulti"])) {
            self::$instances["TimeseriesMulti"] = new self("timeseries-multi");
        }

        return self::$instances["TimeseriesMulti"];
    }

    public static function directoryListing(): self
    {
        if (!isset(self::$instances["DirectoryListing"])) {
            self::$instances["DirectoryListing"] = new self("directory-listing");
        }

        return self::$instances["DirectoryListing"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self("table");
        }

        return self::$instances["Table"];
    }

    public static function numericWide(): self
    {
        if (!isset(self::$instances["NumericWide"])) {
            self::$instances["NumericWide"] = new self("numeric-wide");
        }

        return self::$instances["NumericWide"];
    }

    public static function numericMulti(): self
    {
        if (!isset(self::$instances["NumericMulti"])) {
            self::$instances["NumericMulti"] = new self("numeric-multi");
        }

        return self::$instances["NumericMulti"];
    }

    public static function numericLong(): self
    {
        if (!isset(self::$instances["NumericLong"])) {
            self::$instances["NumericLong"] = new self("numeric-long");
        }

        return self::$instances["NumericLong"];
    }

    public static function logLines(): self
    {
        if (!isset(self::$instances["LogLines"])) {
            self::$instances["LogLines"] = new self("log-lines");
        }

        return self::$instances["LogLines"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "") {
            return self::none();
        }

        if ($value === "timeseries-wide") {
            return self::timeseriesWide();
        }

        if ($value === "timeseries-long") {
            return self::timeseriesLong();
        }

        if ($value === "timeseries-many") {
            return self::timeseriesMany();
        }

        if ($value === "timeseries-multi") {
            return self::timeseriesMulti();
        }

        if ($value === "directory-listing") {
            return self::directoryListing();
        }

        if ($value === "table") {
            return self::table();
        }

        if ($value === "numeric-wide") {
            return self::numericWide();
        }

        if ($value === "numeric-multi") {
            return self::numericMulti();
        }

        if ($value === "numeric-long") {
            return self::numericLong();
        }

        if ($value === "log-lines") {
            return self::logLines();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeClassicConditionsType");
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
