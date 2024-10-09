<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchRawDataSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettingsBuilder())',
        ];
            if ($input->size !== null && $input->size !== "") {
    
        
    $buffer = 'size(';
        $arg0 =\var_export($input->size, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

