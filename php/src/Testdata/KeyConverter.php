<?php

namespace Grafana\Foundation\Testdata;

final class KeyConverter
{
    public static function convert(\Grafana\Foundation\Testdata\Key $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\KeyBuilder())',
        ];
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'tick(';
        $arg0 =\var_export($input->tick, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->uid !== null && $input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

