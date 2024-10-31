---
title: <span class="badge object-type-enum"></span> TextMode
---
# <span class="badge object-type-enum"></span> TextMode

## Definition

```php
final class TextMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TextMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function hTML(): self
    {
        if (!isset(self::$instances["HTML"])) {
            self::$instances["HTML"] = new self("html");
        }

        return self::$instances["HTML"];
    }

    public static function markdown(): self
    {
        if (!isset(self::$instances["Markdown"])) {
            self::$instances["Markdown"] = new self("markdown");
        }

        return self::$instances["Markdown"];
    }

    public static function code(): self
    {
        if (!isset(self::$instances["Code"])) {
            self::$instances["Code"] = new self("code");
        }

        return self::$instances["Code"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "html") {
            return self::hTML();
        }

        if ($value === "markdown") {
            return self::markdown();
        }

        if ($value === "code") {
            return self::code();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TextMode");
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
