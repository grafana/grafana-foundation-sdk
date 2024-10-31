---
title: <span class="badge object-type-enum"></span> CloudWatchQueryMode
---
# <span class="badge object-type-enum"></span> CloudWatchQueryMode

## Definition

```php
final class CloudWatchQueryMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, CloudWatchQueryMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function metrics(): self
    {
        if (!isset(self::$instances["Metrics"])) {
            self::$instances["Metrics"] = new self("Metrics");
        }

        return self::$instances["Metrics"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["Logs"])) {
            self::$instances["Logs"] = new self("Logs");
        }

        return self::$instances["Logs"];
    }

    public static function annotations(): self
    {
        if (!isset(self::$instances["Annotations"])) {
            self::$instances["Annotations"] = new self("Annotations");
        }

        return self::$instances["Annotations"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Metrics") {
            return self::metrics();
        }

        if ($value === "Logs") {
            return self::logs();
        }

        if ($value === "Annotations") {
            return self::annotations();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum CloudWatchQueryMode");
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
