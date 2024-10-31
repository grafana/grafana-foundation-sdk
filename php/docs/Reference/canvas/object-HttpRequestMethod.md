---
title: <span class="badge object-type-enum"></span> HttpRequestMethod
---
# <span class="badge object-type-enum"></span> HttpRequestMethod

## Definition

```php
final class HttpRequestMethod implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HttpRequestMethod>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function gET(): self
    {
        if (!isset(self::$instances["GET"])) {
            self::$instances["GET"] = new self("GET");
        }

        return self::$instances["GET"];
    }

    public static function pOST(): self
    {
        if (!isset(self::$instances["POST"])) {
            self::$instances["POST"] = new self("POST");
        }

        return self::$instances["POST"];
    }

    public static function pUT(): self
    {
        if (!isset(self::$instances["PUT"])) {
            self::$instances["PUT"] = new self("PUT");
        }

        return self::$instances["PUT"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "GET") {
            return self::gET();
        }

        if ($value === "POST") {
            return self::pOST();
        }

        if ($value === "PUT") {
            return self::pUT();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HttpRequestMethod");
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
