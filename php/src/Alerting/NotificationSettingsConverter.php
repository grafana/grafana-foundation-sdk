<?php

namespace Grafana\Foundation\Alerting;

final class NotificationSettingsConverter
{
    public static function convert(\Grafana\Foundation\Alerting\NotificationSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\NotificationSettingsBuilder())',
        ];
            if ($input->groupBy !== null && count($input->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->groupBy as $arg1) {
        $tmpgroup_byarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroup_byarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupInterval !== null && $input->groupInterval !== "") {
    
        
    $buffer = 'groupInterval(';
        $arg0 =\var_export($input->groupInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupWait !== null && $input->groupWait !== "") {
    
        
    $buffer = 'groupWait(';
        $arg0 =\var_export($input->groupWait, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->muteTimeIntervals !== null && count($input->muteTimeIntervals) >= 1) {
    
        
    $buffer = 'muteTimeIntervals(';
        $tmparg0 = [];
        foreach ($input->muteTimeIntervals as $arg1) {
        $tmpmute_time_intervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpmute_time_intervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->receiver !== "") {
    
        
    $buffer = 'receiver(';
        $arg0 =\var_export($input->receiver, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeatInterval !== null && $input->repeatInterval !== "") {
    
        
    $buffer = 'repeatInterval(';
        $arg0 =\var_export($input->repeatInterval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

