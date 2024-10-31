---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`0`: Never refresh the variable

`1`: Queries the data source every time the dashboard loads.

`2`: Queries the data source when the dashboard time range changes.

## Definition

```php
final class VariableRefresh implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, VariableRefresh>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function never(): self
    {
        if (!isset(self::$instances["never"])) {
            self::$instances["never"] = new self(0);
        }

        return self::$instances["never"];
    }

    public static function onDashboardLoad(): self
    {
        if (!isset(self::$instances["onDashboardLoad"])) {
            self::$instances["onDashboardLoad"] = new self(1);
        }

        return self::$instances["onDashboardLoad"];
    }

    public static function onTimeRangeChanged(): self
    {
        if (!isset(self::$instances["onTimeRangeChanged"])) {
            self::$instances["onTimeRangeChanged"] = new self(2);
        }

        return self::$instances["onTimeRangeChanged"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::never();
        }

        if ($value === 1) {
            return self::onDashboardLoad();
        }

        if ($value === 2) {
            return self::onTimeRangeChanged();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableRefresh");
    }

    public function jsonSerialize(): int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

```
