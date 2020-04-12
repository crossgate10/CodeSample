using System;

namespace Factory_Strategy.Protocol
{
    public interface ICarFactory
    {
        ICar CreateCar();
        bool AppliesTo(Type type);
    }
}
