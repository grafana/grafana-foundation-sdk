<?php

namespace Grafana\Foundation\Elasticsearch;

final class BaseBucketAggregationConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\BaseBucketAggregation $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\BaseBucketAggregationBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Elasticsearch\BucketAggregationType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 =\var_export($input->settings, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

