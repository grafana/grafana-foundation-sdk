<?php

namespace Grafana\Foundation\Athena;

final class ConnectionArgsConverter
{
    public static function convert(\Grafana\Foundation\Athena\ConnectionArgs $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Athena\ConnectionArgsBuilder())',
        ];
            if ($input->region !== null && $input->region !== "" && $input->region !== "__default") {
    
        
    $buffer = 'region(';
        $arg0 =\var_export($input->region, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->catalog !== null && $input->catalog !== "" && $input->catalog !== "__default") {
    
        
    $buffer = 'catalog(';
        $arg0 =\var_export($input->catalog, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->database !== null && $input->database !== "" && $input->database !== "__default") {
    
        
    $buffer = 'database(';
        $arg0 =\var_export($input->database, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resultReuseEnabled !== null && $input->resultReuseEnabled !== false) {
    
        
    $buffer = 'resultReuseEnabled(';
        $arg0 =\var_export($input->resultReuseEnabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resultReuseMaxAgeInMinutes !== null && $input->resultReuseMaxAgeInMinutes !== (float) 60) {
    
        
    $buffer = 'resultReuseMaxAgeInMinutes(';
        $arg0 =\var_export($input->resultReuseMaxAgeInMinutes, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

