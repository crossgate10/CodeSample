namespace RealWorldCode
{
    abstract class Element
    {
        public abstract void Accept(IVisitor visitor);
    }
}