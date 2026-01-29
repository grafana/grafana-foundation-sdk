<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class FieldColorConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\FieldColor $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\FieldColorBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue("'.$input->mode.'")';
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
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode::fromValue("'.$input->seriesBy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

