---
title: <span class="badge object-type-enum"></span> RuleExecErrState
---
# <span class="badge object-type-enum"></span> RuleExecErrState

## Definition

```php
final class RuleExecErrState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, RuleExecErrState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function oK(): self
    {
        if (!isset(self::$instances["OK"])) {
            self::$instances["OK"] = new self("OK");
        }

        return self::$instances["OK"];
    }

    public static function alerting(): self
    {
        if (!isset(self::$instances["Alerting"])) {
            self::$instances["Alerting"] = new self("Alerting");
        }

        return self::$instances["Alerting"];
    }

    public static function error(): self
    {
        if (!isset(self::$instances["Error"])) {
            self::$instances["Error"] = new self("Error");
        }

        return self::$instances["Error"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "OK") {
            return self::oK();
        }

        if ($value === "Alerting") {
            return self::alerting();
        }

        if ($value === "Error") {
            return self::error();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum RuleExecErrState");
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
