---
title: <span class="badge object-type-enum"></span> LokiQueryDirection
---
# <span class="badge object-type-enum"></span> LokiQueryDirection

## Definition

```php
final class LokiQueryDirection implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LokiQueryDirection>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function forward(): self
    {
        if (!isset(self::$instances["forward"])) {
            self::$instances["forward"] = new self("forward");
        }

        return self::$instances["forward"];
    }

    public static function backward(): self
    {
        if (!isset(self::$instances["backward"])) {
            self::$instances["backward"] = new self("backward");
        }

        return self::$instances["backward"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "forward") {
            return self::forward();
        }

        if ($value === "backward") {
            return self::backward();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LokiQueryDirection");
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
