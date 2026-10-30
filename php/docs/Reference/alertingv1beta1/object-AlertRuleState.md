---
title: <span class="badge object-type-enum"></span> AlertRuleState
---
# <span class="badge object-type-enum"></span> AlertRuleState

## Definition

```php
final class AlertRuleState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AlertRuleState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function inactive(): self
    {
        if (!isset(self::$instances["Inactive"])) {
            self::$instances["Inactive"] = new self("Inactive");
        }

        return self::$instances["Inactive"];
    }

    public static function healthy(): self
    {
        if (!isset(self::$instances["Healthy"])) {
            self::$instances["Healthy"] = new self("Healthy");
        }

        return self::$instances["Healthy"];
    }

    public static function firing(): self
    {
        if (!isset(self::$instances["Firing"])) {
            self::$instances["Firing"] = new self("Firing");
        }

        return self::$instances["Firing"];
    }

    public static function pending(): self
    {
        if (!isset(self::$instances["Pending"])) {
            self::$instances["Pending"] = new self("Pending");
        }

        return self::$instances["Pending"];
    }

    public static function recovering(): self
    {
        if (!isset(self::$instances["Recovering"])) {
            self::$instances["Recovering"] = new self("Recovering");
        }

        return self::$instances["Recovering"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Inactive") {
            return self::inactive();
        }

        if ($value === "Healthy") {
            return self::healthy();
        }

        if ($value === "Firing") {
            return self::firing();
        }

        if ($value === "Pending") {
            return self::pending();
        }

        if ($value === "Recovering") {
            return self::recovering();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AlertRuleState");
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
