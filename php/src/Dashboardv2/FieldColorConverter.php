<?php

namespace Grafana\Foundation\Dashboardv2;

final class FieldColorConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\FieldColor $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\FieldColorBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2\FieldColorModeId::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->fixedColor !== null && $input->fixedColor !== "") {
    
        
    $buffer = 'fixedColor(';
        $arg0 =\var_export($input->fixedColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->seriesBy !== null) {
    
        
    $buffer = 'seriesBy(';
        $arg0 ='\Grafana\Foundation\Dashboardv2\FieldColorSeriesByMode::fromValue("'.$input->seriesBy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

