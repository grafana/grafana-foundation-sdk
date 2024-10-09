<?php

namespace Grafana\Foundation\Elasticsearch;

final class GeoHashGridSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\GeoHashGridSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\GeoHashGridSettingsBuilder())',
        ];
            if ($input->precision !== null && $input->precision !== "") {
    
        
    $buffer = 'precision(';
        $arg0 =\var_export($input->precision, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

