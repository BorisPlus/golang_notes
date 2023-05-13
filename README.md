# Golang. Заметки

## Составитель отчетов

* [Составитель](./templator/README.md) формирует отчет по шаблону из файлов кода и Go-переменных.

> Есть идеи по шаблонизаторам?

## Слайсы

* [Трюки](https://ueokande.github.io/go-slice-tricks/)

## Стек

* Решение задачи [стека](./stack/README.md) из вебинара (с тестированием производительности вариантов реализаций).

## Интерфейсы

* Решение задачи [интерфейса](./interfaces/README.md) из вебинара.

* Решение расширенной подзадачи [двусвязного списка](./dlist/README.md) из курса.

* Решение [задач](./mathan/README.md) математического анализа данных.

  * Алгоритм [классификации точки пространства](https://github.com/BorisPlus/golang_notes/blob/master/mathan/README.md#%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D1%8F) (соотнесения с тем или иным классом точек, определяемым его центром - центроидом) заранее не знает метрику (функцию подсчета расстояния, размерность пространства), имеется только **интерфейс**.

  * Алгоритм [иерархической кластеризации](https://github.com/BorisPlus/golang_notes/blob/master/mathan/README.md#%D0%B4%D0%B5%D0%BC%D0%BE%D0%BD%D1%81%D1%82%D1%80%D0%B0%D1%86%D0%B8%D1%8F-%D0%BA%D0%BB%D0%B0%D1%81%D1%82%D0%B5%D1%80%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D0%B8). Субъективно конечно, но вариант использования моего пакета конечным потребителем кажется проще, чем этого [https://github.com/knightjdr/hclust].

## Бенчмарки

* Решение задачи на [Бенчмарки]([./interfaces/README.md](https://github.com/BorisPlus/OTUS-Go-2023-03/blob/master/hw04_lru_cache/REPORT.md#benchmark-%D0%B8%D0%BB%D0%B8-%D0%BA%D0%B0%D0%BA-%D1%8F-01-%D1%81%D0%BB%D0%BE%D0%B6%D0%BD%D0%BE%D1%81%D1%82%D1%8C-%D0%BF%D1%80%D0%B5%D0%B4%D1%8A%D1%8F%D0%B2%D0%BB%D1%8F%D0%BB)) расскрывает возможность ведения своих собственных вычислений в рамках тестирования. В документе представлен вариант реализации функции [метрики внутри бенчмарка](https://github.com/BorisPlus/OTUS-Go-2023-03/blob/master/hw04_lru_cache/REPORT.md#%D0%BF%D0%BE%D0%B4%D1%85%D0%BE%D0%B4-benchmark-%D1%81-%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%D0%BC-benchmarkreportmetric) (это действительно хорошо, становится возможным собирать данные, заведомо отсутствуюшие в интерфейсе бенчмарка, в примере - это дисперсия), а также вариант возможности вычисления [показателей, агрегирующих итоговые значения нагрузочного тестирования](https://github.com/BorisPlus/OTUS-Go-2023-03/blob/master/hw04_lru_cache/REPORT.md#%D0%B7%D0%B0%D0%BF%D1%83%D1%81%D0%BA-benchmark-%D1%81-%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%D0%BC-benchmarkresult), в примере - это усредненные значения.
* 

## Указатели

* Нужно понять, что это за [указатель](./pointers/README.md).
