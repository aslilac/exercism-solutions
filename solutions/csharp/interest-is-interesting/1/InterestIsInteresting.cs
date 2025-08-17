using System;

static class SavingsAccount
{
    public static float InterestRate(decimal balance)
    {
        switch (balance) {
        case < 0:
            return 3.213F;
        case < 1000:
            return 0.5F;
        case < 5000:
            return 1.621F;
        default:
            return 2.475F;
        }
    }

    public static decimal Interest(decimal balance) => balance * Convert.ToDecimal(InterestRate(balance) / 100);

    public static decimal AnnualBalanceUpdate(decimal balance) => balance + Interest(balance);

    public static int YearsBeforeDesiredBalance(decimal balance, decimal targetBalance)
    {
        var years = 0;
        while (balance < targetBalance) {
            years++;
            balance = AnnualBalanceUpdate(balance);
        }
        return years;
    }
}
