<?php

namespace Grafana\Foundation\Team;

final class TeamConverter
{
    public static function convert(\Grafana\Foundation\Team\Team $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Team\TeamBuilder('.\var_export($input->name, true).'))',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->email !== null && $input->email !== "") {
    
        
    $buffer = 'email(';
        $arg0 =\var_export($input->email, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

