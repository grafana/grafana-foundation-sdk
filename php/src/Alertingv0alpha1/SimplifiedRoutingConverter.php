<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class SimplifiedRoutingConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\SimplifiedRoutingBuilder())',
        ];
            if ($input->receiver !== "") {
    
        
    $buffer = 'receiver(';
        $arg0 =\var_export($input->receiver, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBy !== null && count($input->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->groupBy as $arg1) {
        $tmpgroupByarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroupByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupWait !== "") {
    
        
    $buffer = 'groupWait(';
        $arg0 =\var_export($input->groupWait, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupInterval !== "") {
    
        
    $buffer = 'groupInterval(';
        $arg0 =\var_export($input->groupInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeatInterval !== "") {
    
        
    $buffer = 'repeatInterval(';
        $arg0 =\var_export($input->repeatInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->muteTimeIntervals !== null && count($input->muteTimeIntervals) >= 1) {
    
        
    $buffer = 'muteTimeIntervals(';
        $tmparg0 = [];
        foreach ($input->muteTimeIntervals as $arg1) {
        $tmpmuteTimeIntervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpmuteTimeIntervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->activeTimeIntervals !== null && count($input->activeTimeIntervals) >= 1) {
    
        
    $buffer = 'activeTimeIntervals(';
        $tmparg0 = [];
        foreach ($input->activeTimeIntervals as $arg1) {
        $tmpactiveTimeIntervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpactiveTimeIntervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

