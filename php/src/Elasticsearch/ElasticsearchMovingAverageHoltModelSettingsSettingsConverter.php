<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchMovingAverageHoltModelSettingsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder())',
        ];
            if ($input->alpha !== null && $input->alpha !== "") {
    
        
    $buffer = 'alpha(';
        $arg0 =\var_export($input->alpha, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->beta !== null && $input->beta !== "") {
    
        
    $buffer = 'beta(';
        $arg0 =\var_export($input->beta, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

