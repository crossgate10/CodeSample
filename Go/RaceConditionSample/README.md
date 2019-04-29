#### Refeference: 
[用一個小例子談談 Golang 中的 Race Condition](https://larrylu.blog/race-condition-in-golang-c49a6e242259)

[互斥鎖(Mutex)](https://zh.wikipedia.org/wiki/%E4%BA%92%E6%96%A5%E9%94%81)

##### ＠Cases
_May not see race condition_
- When a is small (a < 5)

_You'll see race condition_

- When a is Big (a > 1000000)

_Solutions:_

- Mutex
