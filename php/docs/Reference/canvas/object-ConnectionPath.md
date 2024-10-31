---
title: <span class="badge object-type-enum"></span> ConnectionPath
---
# <span class="badge object-type-enum"></span> ConnectionPath

## Definition

```php
final class ConnectionPath implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ConnectionPath>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function straight(): self
    {
        if (!isset(self::$instances["Straight"])) {
            self::$instances["Straight"] = new self("straight");
        }

        return self::$instances["Straight"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "straight") {
            return self::straight();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ConnectionPath");
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
