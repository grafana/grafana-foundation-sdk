---
title: <span class="badge object-type-enum"></span> FormatOptions
---
# <span class="badge object-type-enum"></span> FormatOptions

## Definition

```php
final class FormatOptions implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, FormatOptions>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function timeSeries(): self
    {
        if (!isset(self::$instances["TimeSeries"])) {
            self::$instances["TimeSeries"] = new self(0);
        }

        return self::$instances["TimeSeries"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self(1);
        }

        return self::$instances["Table"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["Logs"])) {
            self::$instances["Logs"] = new self(2);
        }

        return self::$instances["Logs"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::timeSeries();
        }

        if ($value === 1) {
            return self::table();
        }

        if ($value === 2) {
            return self::logs();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FormatOptions");
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
