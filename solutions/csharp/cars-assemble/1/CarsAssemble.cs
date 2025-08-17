using System;
using System.Diagnostics;

static class AssemblyLine
{
    public static double SuccessRate(int speed)
    {
        switch (speed) {
        case 0:
            return 0.0;
        case <= 4:
            return 1.0;
        case <= 8:
            return 0.9;
        case <= 9:
            return 0.8;
        case <= 10:
            return 0.77;
        default:
            throw new UnreachableException("unreachable");
        }
    }
    
    public static double ProductionRatePerHour(int speed) => speed * SuccessRate(speed) * 221;

    public static int WorkingItemsPerMinute(int speed) {
        var perHour = ProductionRatePerHour(speed);
        return Convert.ToInt32(Math.Floor(perHour / 60));
    }
}
