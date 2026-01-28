<?php

namespace Grafana\Foundation\Xychart;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Xychart\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mapping(';
        $arg0 ='\Grafana\Foundation\Xychart\SeriesMapping::fromValue("'.$input->mapping.'")';
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
    
    
            if (count($input->series) >= 1) {
    
        
    $buffer = 'series(';
        $tmparg0 = [];
        foreach ($input->series as $arg1) {
        $tmpseriesarg1 = \Grafana\Foundation\Xychart\XYSeriesConfigConverter::convert($arg1);
        $tmparg0[] = $tmpseriesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

