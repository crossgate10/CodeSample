using System;

namespace Factory_Strategy.Protocol
{
    public interface ICarStrategy
    {
        ICar CreateCar(Type type);
    }
}
