<?php

namespace Grafana\Foundation\Elasticsearch;

final class ExtendedStatConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ExtendedStat $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ExtendedStatBuilder())',
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
        $arg0 ='\Grafana\Foundation\Elasticsearch\ExtendedStatMetaType::fromValue("'.$input->value.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

