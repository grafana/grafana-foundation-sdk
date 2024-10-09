<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchDerivativeSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettingsBuilder())',
        ];
            if ($input->unit !== null && $input->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

