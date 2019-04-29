using System.Collections.Generic;

namespace Visitor
{
    public class ObjectStructure
    {
        private List<Element> elements = new List<Element>();

        public void Attach(Element element)
        {
            this.elements.Add(element);
        }

        public void Detach(Element element)
        {
            this.elements.Remove(element);
        }

        public void Accept(Visitor visitor)
        {
            foreach (Element element  in this.elements)
            {
                element.Accept(visitor);
            }
        }
    }
}
