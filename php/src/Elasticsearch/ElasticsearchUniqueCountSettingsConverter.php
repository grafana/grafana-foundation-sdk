<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchUniqueCountSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettingsBuilder())',
        ];
            if ($input->precisionThreshold !== null && $input->precisionThreshold !== "") {
    
        
    $buffer = 'precisionThreshold(';
        $arg0 =\var_export($input->precisionThreshold, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->missing !== null && $input->missing !== "") {
    
        
    $buffer = 'missing(';
        $arg0 =\var_export($input->missing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

