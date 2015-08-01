# towed

Проверка за вдигнат автомобил по регистрационнен номер чрез https://www.sofiatraffic.bg/bg/tow-away/

## Инсталация

      $ go get github.com/mytrile/wunderlist

## Пример

``` go
package main

import (
  "fmt"
  "github.com/mytrile/towed"
)

func main() {
    towed.Lifted("CA1234AB")
}
```


### Мета

* Author  - Dimitar Kostov
* Email   - mitko.kostov@gmail.com
* Blog    - <http://mytrile.github.io>
* Twitter - [mytrile](https://twitter.com/mytrile)

## License

(The MIT License)

Copyright (c) 2013 Dimitar Kostov <mitko.kostov@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.