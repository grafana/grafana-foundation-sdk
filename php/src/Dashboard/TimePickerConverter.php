<?php

namespace Grafana\Foundation\Dashboard;

final class TimePickerConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\TimePickerConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\TimePickerBuilder())',
        ];
            if ($input->hidden !== null && $input->hidden !== false) {
    
        
    $buffer = 'hidden(';
        $arg0 =\var_export($input->hidden, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refreshIntervals !== null && count($input->refreshIntervals) >= 1) {
    
        
    $buffer = 'refreshIntervals(';
        $tmparg0 = [];
        foreach ($input->refreshIntervals as $arg1) {
        $tmprefresh_intervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmprefresh_intervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeOptions !== null && count($input->timeOptions) >= 1) {
    
        
    $buffer = 'timeOptions(';
        $tmparg0 = [];
        foreach ($input->timeOptions as $arg1) {
        $tmptime_optionsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptime_optionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->nowDelay !== null && $input->nowDelay !== "") {
    
        
    $buffer = 'nowDelay(';
        $arg0 =\var_export($input->nowDelay, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

