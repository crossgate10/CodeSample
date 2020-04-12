using System;
using System.Collections.Generic;

namespace RealWorldCode
{
    abstract class Document
    {
        private List<Page> pages = new List<Page>();
        public Document()
        {
            this.CreatePages();
        }

        public abstract void CreatePages();

        public List<Page> Pages { get { return this.pages; } }
    }
}