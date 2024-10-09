<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchCumulativeSumSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettingsBuilder())',
        ];
            if ($input->format !== null && $input->format !== "") {
    
        
    $buffer = 'format(';
        $arg0 =\var_export($input->format, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

