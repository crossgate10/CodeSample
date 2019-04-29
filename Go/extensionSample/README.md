#### Reference: [POST](https://stackoverflow.com/questions/22153309/extension-methods-in-golang)
  
    type MyString string
    
    func (m *MyString) Method() {}
    
    func main() {
      var s MyString = ""
      s.Method()
    }
