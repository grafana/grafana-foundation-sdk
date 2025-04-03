<?php

namespace Grafana\Foundation\Debug;

final class UpdateConfigConverter
{
    public static function convert(\Grafana\Foundation\Debug\UpdateConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Debug\UpdateConfigBuilder())',
        ];
            
    
        {
    $buffer = 'render(';
        $arg0 =\var_export($input->render, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'dataChanged(';
        $arg0 =\var_export($input->dataChanged, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'schemaChanged(';
        $arg0 =\var_export($input->schemaChanged, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

