class Lasagna
{
    public int ExpectedMinutesInOven() => 40;

    public int RemainingMinutesInOven(int soFar) => ExpectedMinutesInOven() - soFar;

    public int PreparationTimeInMinutes(int layers) => layers * 2;

    public int ElapsedTimeInMinutes(int layers, int soFar) => PreparationTimeInMinutes(layers) + soFar;
}
