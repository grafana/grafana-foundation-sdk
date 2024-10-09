<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchRateSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchRateSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchRateSettingsBuilder())',
        ];
            if ($input->unit !== null && $input->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->mode !== null && $input->mode !== "") {
    
        
    $buffer = 'mode(';
        $arg0 =\var_export($input->mode, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

