---
title: <span class="badge object-type-enum"></span> SupportingQueryType
---
# <span class="badge object-type-enum"></span> SupportingQueryType

## Definition

```php
final class SupportingQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SupportingQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function logsVolume(): self
    {
        if (!isset(self::$instances["LogsVolume"])) {
            self::$instances["LogsVolume"] = new self("logsVolume");
        }

        return self::$instances["LogsVolume"];
    }

    public static function logsSample(): self
    {
        if (!isset(self::$instances["LogsSample"])) {
            self::$instances["LogsSample"] = new self("logsSample");
        }

        return self::$instances["LogsSample"];
    }

    public static function dataSample(): self
    {
        if (!isset(self::$instances["DataSample"])) {
            self::$instances["DataSample"] = new self("dataSample");
        }

        return self::$instances["DataSample"];
    }

    public static function infiniteScroll(): self
    {
        if (!isset(self::$instances["InfiniteScroll"])) {
            self::$instances["InfiniteScroll"] = new self("infiniteScroll");
        }

        return self::$instances["InfiniteScroll"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "logsVolume") {
            return self::logsVolume();
        }

        if ($value === "logsSample") {
            return self::logsSample();
        }

        if ($value === "dataSample") {
            return self::dataSample();
        }

        if ($value === "infiniteScroll") {
            return self::infiniteScroll();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SupportingQueryType");
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
