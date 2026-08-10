package common

import "errors"

// todo: Посмотреть библиотеки, которые позволяют удобно работать с ошибками

var DublicateIdError = errors.New("У составного объекта не может быть два одинаковых ключа")