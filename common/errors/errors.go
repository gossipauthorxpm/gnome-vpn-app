package errors

// todo: Посмотреть библиотеки, которые позволяют удобно работать с ошибками

var (
	Throwable = CreateErrorType("Общее исключение", nil)
	RuntimeException = CreateErrorType("Исключение времени выполнения", Throwable)
	LogicException = CreateErrorType("Пользовательское исключение", Throwable)
	InvalidArgumentException = CreateErrorType("Неверный аргумент", RuntimeException)
	IOException = CreateErrorType("Ошибка ввода\\вывода", RuntimeException)
)