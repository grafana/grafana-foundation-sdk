<?php

namespace Grafana\Foundation\Timeseries;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Timeseries\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Timeseries\OptionsBuilder())',
        ];
            if ($input->timezone !== null && count($input->timezone) >= 1) {
    
        
    $buffer = 'timezone(';
        $tmparg0 = [];
        foreach ($input->timezone as $arg1) {
        $tmptimezonearg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptimezonearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->orientation !== null) {
    
        
    $buffer = 'orientation(';
        $arg0 ='\Grafana\Foundation\Common\VizOrientation::fromValue("'.$input->orientation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

