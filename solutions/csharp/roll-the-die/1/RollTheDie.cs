using System;

public class Player
{
    public int RollDie()
    {
        var r = new Random();
        return r.Next() % 18 + 1;
    }

    public double GenerateSpellStrength()
    {
        var r = new Random();
        return r.NextDouble() * 100;
    }
}
