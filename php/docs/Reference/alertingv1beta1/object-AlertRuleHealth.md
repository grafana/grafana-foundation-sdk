---
title: <span class="badge object-type-enum"></span> AlertRuleHealth
---
# <span class="badge object-type-enum"></span> AlertRuleHealth

## Definition

```php
final class AlertRuleHealth implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AlertRuleHealth>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function unknown(): self
    {
        if (!isset(self::$instances["Unknown"])) {
            self::$instances["Unknown"] = new self("Unknown");
        }

        return self::$instances["Unknown"];
    }

    public static function oK(): self
    {
        if (!isset(self::$instances["OK"])) {
            self::$instances["OK"] = new self("OK");
        }

        return self::$instances["OK"];
    }

    public static function paused(): self
    {
        if (!isset(self::$instances["Paused"])) {
            self::$instances["Paused"] = new self("Paused");
        }

        return self::$instances["Paused"];
    }

    public static function error(): self
    {
        if (!isset(self::$instances["Error"])) {
            self::$instances["Error"] = new self("Error");
        }

        return self::$instances["Error"];
    }

    public static function noData(): self
    {
        if (!isset(self::$instances["NoData"])) {
            self::$instances["NoData"] = new self("NoData");
        }

        return self::$instances["NoData"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Unknown") {
            return self::unknown();
        }

        if ($value === "OK") {
            return self::oK();
        }

        if ($value === "Paused") {
            return self::paused();
        }

        if ($value === "Error") {
            return self::error();
        }

        if ($value === "NoData") {
            return self::noData();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AlertRuleHealth");
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
