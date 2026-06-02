<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class AlignmentTypes implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AlignmentTypes>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function aLIGNDELTA(): self
    {
        if (!isset(self::$instances["ALIGNDELTA"])) {
            self::$instances["ALIGNDELTA"] = new self("ALIGN_DELTA");
        }

        return self::$instances["ALIGNDELTA"];
    }

    public static function aLIGNRATE(): self
    {
        if (!isset(self::$instances["ALIGNRATE"])) {
            self::$instances["ALIGNRATE"] = new self("ALIGN_RATE");
        }

        return self::$instances["ALIGNRATE"];
    }

    public static function aLIGNINTERPOLATE(): self
    {
        if (!isset(self::$instances["ALIGNINTERPOLATE"])) {
            self::$instances["ALIGNINTERPOLATE"] = new self("ALIGN_INTERPOLATE");
        }

        return self::$instances["ALIGNINTERPOLATE"];
    }

    public static function aLIGNNEXTOLDER(): self
    {
        if (!isset(self::$instances["ALIGNNEXTOLDER"])) {
            self::$instances["ALIGNNEXTOLDER"] = new self("ALIGN_NEXT_OLDER");
        }

        return self::$instances["ALIGNNEXTOLDER"];
    }

    public static function aLIGNMIN(): self
    {
        if (!isset(self::$instances["ALIGNMIN"])) {
            self::$instances["ALIGNMIN"] = new self("ALIGN_MIN");
        }

        return self::$instances["ALIGNMIN"];
    }

    public static function aLIGNMAX(): self
    {
        if (!isset(self::$instances["ALIGNMAX"])) {
            self::$instances["ALIGNMAX"] = new self("ALIGN_MAX");
        }

        return self::$instances["ALIGNMAX"];
    }

    public static function aLIGNMEAN(): self
    {
        if (!isset(self::$instances["ALIGNMEAN"])) {
            self::$instances["ALIGNMEAN"] = new self("ALIGN_MEAN");
        }

        return self::$instances["ALIGNMEAN"];
    }

    public static function aLIGNCOUNT(): self
    {
        if (!isset(self::$instances["ALIGNCOUNT"])) {
            self::$instances["ALIGNCOUNT"] = new self("ALIGN_COUNT");
        }

        return self::$instances["ALIGNCOUNT"];
    }

    public static function aLIGNSUM(): self
    {
        if (!isset(self::$instances["ALIGNSUM"])) {
            self::$instances["ALIGNSUM"] = new self("ALIGN_SUM");
        }

        return self::$instances["ALIGNSUM"];
    }

    public static function aLIGNSTDDEV(): self
    {
        if (!isset(self::$instances["ALIGNSTDDEV"])) {
            self::$instances["ALIGNSTDDEV"] = new self("ALIGN_STDDEV");
        }

        return self::$instances["ALIGNSTDDEV"];
    }

    public static function aLIGNCOUNTTRUE(): self
    {
        if (!isset(self::$instances["ALIGNCOUNTTRUE"])) {
            self::$instances["ALIGNCOUNTTRUE"] = new self("ALIGN_COUNT_TRUE");
        }

        return self::$instances["ALIGNCOUNTTRUE"];
    }

    public static function aLIGNCOUNTFALSE(): self
    {
        if (!isset(self::$instances["ALIGNCOUNTFALSE"])) {
            self::$instances["ALIGNCOUNTFALSE"] = new self("ALIGN_COUNT_FALSE");
        }

        return self::$instances["ALIGNCOUNTFALSE"];
    }

    public static function aLIGNFRACTIONTRUE(): self
    {
        if (!isset(self::$instances["ALIGNFRACTIONTRUE"])) {
            self::$instances["ALIGNFRACTIONTRUE"] = new self("ALIGN_FRACTION_TRUE");
        }

        return self::$instances["ALIGNFRACTIONTRUE"];
    }

    public static function aLIGNPERCENTILE99(): self
    {
        if (!isset(self::$instances["ALIGNPERCENTILE99"])) {
            self::$instances["ALIGNPERCENTILE99"] = new self("ALIGN_PERCENTILE_99");
        }

        return self::$instances["ALIGNPERCENTILE99"];
    }

    public static function aLIGNPERCENTILE95(): self
    {
        if (!isset(self::$instances["ALIGNPERCENTILE95"])) {
            self::$instances["ALIGNPERCENTILE95"] = new self("ALIGN_PERCENTILE_95");
        }

        return self::$instances["ALIGNPERCENTILE95"];
    }

    public static function aLIGNPERCENTILE50(): self
    {
        if (!isset(self::$instances["ALIGNPERCENTILE50"])) {
            self::$instances["ALIGNPERCENTILE50"] = new self("ALIGN_PERCENTILE_50");
        }

        return self::$instances["ALIGNPERCENTILE50"];
    }

    public static function aLIGNPERCENTILE05(): self
    {
        if (!isset(self::$instances["ALIGNPERCENTILE05"])) {
            self::$instances["ALIGNPERCENTILE05"] = new self("ALIGN_PERCENTILE_05");
        }

        return self::$instances["ALIGNPERCENTILE05"];
    }

    public static function aLIGNPERCENTCHANGE(): self
    {
        if (!isset(self::$instances["ALIGNPERCENTCHANGE"])) {
            self::$instances["ALIGNPERCENTCHANGE"] = new self("ALIGN_PERCENT_CHANGE");
        }

        return self::$instances["ALIGNPERCENTCHANGE"];
    }

    public static function aLIGNNONE(): self
    {
        if (!isset(self::$instances["ALIGNNONE"])) {
            self::$instances["ALIGNNONE"] = new self("ALIGN_NONE");
        }

        return self::$instances["ALIGNNONE"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "ALIGN_DELTA") {
            return self::aLIGNDELTA();
        }

        if ($value === "ALIGN_RATE") {
            return self::aLIGNRATE();
        }

        if ($value === "ALIGN_INTERPOLATE") {
            return self::aLIGNINTERPOLATE();
        }

        if ($value === "ALIGN_NEXT_OLDER") {
            return self::aLIGNNEXTOLDER();
        }

        if ($value === "ALIGN_MIN") {
            return self::aLIGNMIN();
        }

        if ($value === "ALIGN_MAX") {
            return self::aLIGNMAX();
        }

        if ($value === "ALIGN_MEAN") {
            return self::aLIGNMEAN();
        }

        if ($value === "ALIGN_COUNT") {
            return self::aLIGNCOUNT();
        }

        if ($value === "ALIGN_SUM") {
            return self::aLIGNSUM();
        }

        if ($value === "ALIGN_STDDEV") {
            return self::aLIGNSTDDEV();
        }

        if ($value === "ALIGN_COUNT_TRUE") {
            return self::aLIGNCOUNTTRUE();
        }

        if ($value === "ALIGN_COUNT_FALSE") {
            return self::aLIGNCOUNTFALSE();
        }

        if ($value === "ALIGN_FRACTION_TRUE") {
            return self::aLIGNFRACTIONTRUE();
        }

        if ($value === "ALIGN_PERCENTILE_99") {
            return self::aLIGNPERCENTILE99();
        }

        if ($value === "ALIGN_PERCENTILE_95") {
            return self::aLIGNPERCENTILE95();
        }

        if ($value === "ALIGN_PERCENTILE_50") {
            return self::aLIGNPERCENTILE50();
        }

        if ($value === "ALIGN_PERCENTILE_05") {
            return self::aLIGNPERCENTILE05();
        }

        if ($value === "ALIGN_PERCENT_CHANGE") {
            return self::aLIGNPERCENTCHANGE();
        }

        if ($value === "ALIGN_NONE") {
            return self::aLIGNNONE();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AlignmentTypes");
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

