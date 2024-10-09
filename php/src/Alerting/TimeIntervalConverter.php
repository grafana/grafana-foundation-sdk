<?php

namespace Grafana\Foundation\Alerting;

final class TimeIntervalConverter
{
    public static function convert(\Grafana\Foundation\Alerting\TimeInterval $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\TimeIntervalBuilder())',
        ];
            if ($input->daysOfMonth !== null && count($input->daysOfMonth) >= 1) {
    
        
    $buffer = 'daysOfMonth(';
        $tmparg0 = [];
        foreach ($input->daysOfMonth as $arg1) {
        $tmpdays_of_montharg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdays_of_montharg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->location !== null && $input->location !== "") {
    
        
    $buffer = 'location(';
        $arg0 =\var_export($input->location, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->months !== null && count($input->months) >= 1) {
    
        
    $buffer = 'months(';
        $tmparg0 = [];
        foreach ($input->months as $arg1) {
        $tmpmonthsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpmonthsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->times !== null && count($input->times) >= 1) {
    
        
    $buffer = 'times(';
        $tmparg0 = [];
        foreach ($input->times as $arg1) {
        $tmptimesarg1 = \Grafana\Foundation\Alerting\TimeRangeConverter::convert($arg1);
        $tmparg0[] = $tmptimesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->weekdays !== null && count($input->weekdays) >= 1) {
    
        
    $buffer = 'weekdays(';
        $tmparg0 = [];
        foreach ($input->weekdays as $arg1) {
        $tmpweekdaysarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpweekdaysarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->years !== null && count($input->years) >= 1) {
    
        
    $buffer = 'years(';
        $tmparg0 = [];
        foreach ($input->years as $arg1) {
        $tmpyearsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpyearsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

