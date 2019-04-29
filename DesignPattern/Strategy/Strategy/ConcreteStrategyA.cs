namespace StrategySample
{
    using System;

    class ConcreteStrategyA : Strategy
    {
        public override void AlgorithmInterface()
        {
            Console.WriteLine(
              "Called ConcreteStrategyA.AlgorithmInterface()");
        }
    }
}