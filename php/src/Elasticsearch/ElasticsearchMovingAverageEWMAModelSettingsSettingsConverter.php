<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchMovingAverageEWMAModelSettingsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder())',
        ];
            if ($input->alpha !== null && $input->alpha !== "") {
    
        
    $buffer = 'alpha(';
        $arg0 =\var_export($input->alpha, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

