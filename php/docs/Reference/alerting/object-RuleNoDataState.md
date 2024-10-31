---
title: <span class="badge object-type-enum"></span> RuleNoDataState
---
# <span class="badge object-type-enum"></span> RuleNoDataState

## Definition

```php
final class RuleNoDataState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, RuleNoDataState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function alerting(): self
    {
        if (!isset(self::$instances["Alerting"])) {
            self::$instances["Alerting"] = new self("Alerting");
        }

        return self::$instances["Alerting"];
    }

    public static function noData(): self
    {
        if (!isset(self::$instances["NoData"])) {
            self::$instances["NoData"] = new self("NoData");
        }

        return self::$instances["NoData"];
    }

    public static function oK(): self
    {
        if (!isset(self::$instances["OK"])) {
            self::$instances["OK"] = new self("OK");
        }

        return self::$instances["OK"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Alerting") {
            return self::alerting();
        }

        if ($value === "NoData") {
            return self::noData();
        }

        if ($value === "OK") {
            return self::oK();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum RuleNoDataState");
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
