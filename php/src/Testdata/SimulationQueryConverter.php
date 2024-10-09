<?php

namespace Grafana\Foundation\Testdata;

final class SimulationQueryConverter
{
    public static function convert(\Grafana\Foundation\Testdata\SimulationQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\SimulationQueryBuilder())',
        ];
            
    
        {
    $buffer = 'key(';
        $arg0 = \Grafana\Foundation\Testdata\KeyConverter::convert($input->key);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->config !== null) {
    
        
    $buffer = 'config(';
        $arg0 = "[";
        foreach ($input->config as $key => $arg1) {
            $tmpconfigarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpconfigarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stream !== null) {
    
        
    $buffer = 'stream(';
        $arg0 =\var_export($input->stream, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->last !== null) {
    
        
    $buffer = 'last(';
        $arg0 =\var_export($input->last, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

