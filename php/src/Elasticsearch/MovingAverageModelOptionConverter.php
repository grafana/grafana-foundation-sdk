<?php

namespace Grafana\Foundation\Elasticsearch;

final class MovingAverageModelOptionConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\MovingAverageModelOption $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\MovingAverageModelOptionBuilder())',
        ];
            if ($input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'value(';
        $arg0 ='\Grafana\Foundation\Elasticsearch\MovingAverageModel::fromValue("'.$input->value.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

