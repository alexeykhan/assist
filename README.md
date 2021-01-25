<pre>
  ______    ______    ______   ______   ______   ________  
 /      \  /      \  /      \ /      | /      \ /        |
/$$$$$$  |/$$$$$$  |/$$$$$$  |$$$$$$/ /$$$$$$  |$$$$$$$$/  
$$ |__$$ |$$ \__$$/ $$ \__$$/   $$ |  $$ \__$$/    $$ |    
$$    $$ |$$      \ $$      \   $$ |  $$      \    $$ |    
$$$$$$$$ | $$$$$$  | $$$$$$  |  $$ |   $$$$$$  |   $$ |    
$$ |  $$ |/  \__$$ |/  \__$$ | _$$ |_ /  \__$$ |   $$ |    
$$ |  $$ |$$    $$/ $$    $$/ / $$   |$$    $$/    $$ |    
$$/   $$/  $$$$$$/   $$$$$$/  $$$$$$/  $$$$$$/     $$/
</pre>

# 💸 assist
    
Персональный ассистент для планирования личных финансов.

<p>
    <a href="https://pkg.go.dev/github.com/alexeykhan/assist">
        <img src="https://img.shields.io/badge/pkg.go.dev-reference-00ADD8?logo=go&logoColor=white" alt="GoDoc Reference">
    </a>
    <a href="https://pkg.go.dev/github.com/alexeykhan/assist">
        <img src="https://img.shields.io/badge/version-0.1.0-00ADD8&logoColor=white" alt="Version">
    </a>
    <a href="https://github.com/alexeykhan/amocrm">
        <img src="https://img.shields.io/badge/build-passes-success" alt="Build Status">
    </a>
    <a href="https://github.com/alexeykhan/assist/blob/master/LICENSE.md">
        <img src="https://img.shields.io/badge/licence-MIT-success" alt="License">
    </a>
</p>

## Юзкейсы

Представим, что вы решили обеспечить себе достойную старость, создав собственный 
пенсионный счет. Для этого решили посчитать, сколько денег нужно инвестировать 
на постоянной основе, чтобы при определенном уровне доходности портфеля можно было 
выйти на пенсию и до конца жизни поддерживать комфортный уровень жизни.

Пример решения данной задачи с помощью assist:

```
1. Сколько денег нужно будет через 35 лет, чтобы иметь такую же покупательскую 
   способность, как сегодня, имея 150,000 рублей, если среднегодовая инфляция 
   не будет превышать отметку 5%?

$ ./bin/assist calculate inflation --current=150000 --years=35 --inflation=5
  > Исходная сумма увеличится до 827,402.30
  
2. Сколько денег нужно накопить к пенсии, чтобы с доходностью портфеля 6.5% годовых
   можно было на протяжении следующих 25 лет тратить по 827,402.30 руб в месяц и к 
   концу срока потратить все накопления?
  
$ ./bin/assist decompose retirement --expenses=827402.30 --years=25 --interest=6.5
  > Минимальная сумма накоплений составит: 123,204,271.23

3. Сколько денег нужно откладывать каждый месяц на протяжении следующих 35 лет при
   доходности портфеля 6.5% годовых, чтобы к концу срока накопить 123,204,271.23 руб?
    
$ ./bin/assist decompose savings --goal=123204271.23 --years=35 --interest=6.5
  > Сумма ежемесячных инвестиций составит: 76,572.69
  
Итогo: чтобы через 35 лет выйти на пенсию и поддерживать комфортный уровень жизни, 
сопоставимый c сегодняшними 150,000 рублей в месяц, необходимо на протяжении всего
этого срока ежемесячно инвестировать не менее 76,572.69 рублей и поддерживать
доходность инвестиций на уровне не ниже 6.5% годовых при условии, что инфляция
не будет превышать 5% в год.
```

## Функционал

### decompose

Подход: от желаемого результата. Используйте группу команд `decompose`, чтобы 
декомпозировать финансовую цель — выяснить обязательные условия для ее достижения.

<b>Доступные команды:</b>

```
$ ./bin/assist decompose retirement --help
$ ./bin/assist decompose savings --help
$ ./bin/assist decompose --help
```

#### decompose retirement

Используйте команду `decompose retirement`, чтобы узнать минимальную сумму накоплений
с ежегодной доходностью R%, которая позволит вам при выходе на пенсию на протяжении
следующих N лет ежемесячно тратить X руб в месяц.

<b>Пример использования:</b>

```
$ ./bin/assist decompose retirement --expenses=150000.00 --years=25 --interest=6.5 --detailed=M
$ ./bin/assist decompose retirement -e=150000.00 -y=25 -i=6.5 -d=M
$ ./bin/assist decompose retirement --help
```

<b>Параметры и опции команды:</b>

```
-d, --detailed bool   Выводить детализированный ответ. M — по
                      месяцам, Y — по годам.
-e, --expenses float  Сумма ежемесячных расходов в течение    
                      пенсионного периода
-h, --help bool       help for retirement
-i, --interest float  Доходность вашего инвестиционного       
                      портфеля в процентах годовых
-y, --years uint8     Количество лет, на протяжении которых   
                      будут ежемесячные траты
```

<b>Пример выполнения команды:</b>

```
$ ./bin/assist decompose retirement -e 150000 -y 25 -i 6.5 -d Y

 ДЕКОМПОЗИЦИЯ ПЕНСИИ

 Задача: рассчитать сумму, которую необходимо накопить, чтобы
 при доходности портфеля 6.50% годовых можно было на протяжении
 25 лет тратить 150000 рублей в месяц без дополнительного
 дохода, потратив к концу срока все сбережения.

 ┏━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┓
 ┃   ГОД   ┃    ПРОЦЕНТЫ    ┃    РАСХОДЫ     ┃   НАКОПЛЕНИЯ  ┃
 ┣━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃    1    ┃   1431214.80   ┃   1800000.00   ┃  21966952.43  ┃
 ┃    2    ┃   2837731.38   ┃   3600000.00   ┃  21573469.01  ┃
 ┃    3    ┃   4217895.65   ┃   5400000.00   ┃  21153633.27  ┃
 ┃    4    ┃   5569942.73   ┃   7200000.00   ┃  20705680.36  ┃
 ┃    5    ┃   6891989.58   ┃   9000000.00   ┃  20227727.21  ┃
 ┃    6    ┃   8182027.03   ┃  10800000.00   ┃  19717764.65  ┃
 ┃    7    ┃   9437911.33   ┃  12600000.00   ┃  19173648.96  ┃
 ┃    8    ┃  10657355.20   ┃  14400000.00   ┃  18593092.83  ┃
 ┃    9    ┃  11837918.15   ┃  16200000.00   ┃  17973655.78  ┃
 ┃   10    ┃  12976996.26   ┃  18000000.00   ┃  17312733.88  ┃
 ┃   11    ┃  14071811.20   ┃  19800000.00   ┃  16607548.83  ┃
 ┃   12    ┃  15119398.59   ┃  21600000.00   ┃  15855136.22  ┃
 ┃   13    ┃  16116595.52   ┃  23400000.00   ┃  15052333.14  ┃
 ┃   14    ┃  17060027.23   ┃  25200000.00   ┃  14195764.86  ┃
 ┃   15    ┃  17946092.99   ┃  27000000.00   ┃  13281830.61  ┃
 ┃   16    ┃  18770950.87   ┃  28800000.00   ┃  12306688.50  ┃
 ┃   17    ┃  19530501.68   ┃  30600000.00   ┃  11266239.31  ┃
 ┃   18    ┃  20220371.68   ┃  32400000.00   ┃  10156109.31  ┃
 ┃   19    ┃  20835894.22   ┃  34200000.00   ┃   8971631.85  ┃
 ┃   20    ┃  21372090.11   ┃  36000000.00   ┃   7707827.74  ┃
 ┃   21    ┃  21823646.70   ┃  37800000.00   ┃   6359384.33  ┃
 ┃   22    ┃  22184895.54   ┃  39600000.00   ┃   4920633.16  ┃
 ┃   23    ┃  22449788.54   ┃  41400000.00   ┃   3385526.17  ┃
 ┃   24    ┃  22611872.59   ┃  43200000.00   ┃   1747610.22  ┃
 ┃   25    ┃  22664262.37   ┃  45000000.00   ┃      0.00     ┃
 ┣━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃  ИТОГО  ┃  22664262.37   ┃  45000000.00   ┃      0.00     ┃
 ┗━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┛

 > Минимальная сумма накоплений составит: 22335737.63
 > Сумма начисленных процентов за период: 22664262.37
```

#### decompose savings

Используйте команду `decompose savings`, чтобы узнать минимальную сумму ежемесячных
инвестиций с заданной среднегодовой доходностью и ежемесячной капиализацией процентов,
чтобы к концу заданного срока в годах накопить нужную сумму.

<b>Пример использования:</b>

```
$ ./bin/assist decompose savings --goal=1234567.89 --years=10 --interest=6.5 --detailed=M
$ ./bin/assist decompose savings -g=1234567.89 -y=10 -i=6.5 -d=M
$ ./bin/assist decompose savings --help
```

<b>Параметры и опции команды:</b>

```
-d, --detailed bool     Выводить детализированный ответ. M — по
                        месяцам, Y — по годам.
-g, --goal float32      Ваша финансовая цель, которую     
                        нужно достгнуть за заданный период
-h, --help bool         Документация по команде
-i, --interest float32  Доходность вашего инвестиционного
                        портфеля в процентах годовых
-y, --years uint8       Количество лет, за которое        
                        необходимо накопить нужную сумму
```

<b>Пример выполнения команды:</b>

```
$ ./bin/assist decompose savings -g 1234567.89 -y 10 -i 6.5 -d Y

 ДЕКОМПОЗИЦИЯ НАКОПЛЕНИЯ СУММЫ

 Задача: рассчитать сумму, которую необходимо инвестировать
 каждый месяц на протяжении 10 лет, чтобы при средней
 доходности портфеля 6.50% годовых и ежемесячной
 капитализации процентов накопить 1234567.88 руб.

 ┏━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┓
 ┃  ГОД  ┃    ВЛОЖЕНИЯ    ┃    ПРОЦЕНТЫ    ┃   НАКОПЛЕНИЯ  ┃
 ┣━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃   1   ┃    87499.12    ┃    2654.39     ┃    90153.50   ┃
 ┃   2   ┃   174998.23    ┃    11346.53    ┃   186344.75   ┃
 ┃   3   ┃   262497.34    ┃    26480.77    ┃   288978.09   ┃
 ┃   4   ┃   349996.47    ┃    48488.56    ┃   398484.94   ┃
 ┃   5   ┃   437495.59    ┃    77830.23    ┃   515325.75   ┃
 ┃   6   ┃   524994.69    ┃   114996.95    ┃   639991.62   ┃
 ┃   7   ┃   612493.44    ┃   160512.75    ┃   773006.50   ┃
 ┃   8   ┃   699992.19    ┃   214936.83    ┃   914929.62   ┃
 ┃   9   ┃   787490.94    ┃   278865.75    ┃   1066357.62  ┃
 ┃  10   ┃   874989.69    ┃   352936.03    ┃   1227927.25  ┃
 ┣━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃ ИТОГО ┃   874989.69    ┃   359587.31    ┃   1234578.50  ┃
 ┗━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┛

 > Сумма ежемесячных инвестиций составит: 7291.59
 > Сумма собственных вложений за период: 874989.69
 > Сумма начисленных процентов за период: 359587.31
```

### calculate

Подход: от текущей ситуации. Используйте группу команд `calculate`, чтобы
прогнозировать финансовые результаты в зависимости от исходных данных.

<b>Доступные команды:</b>

```
$ ./bin/assist calculate retirement --help
$ ./bin/assist calculate savings --help
$ ./bin/assist calculate --help
```

#### calculate savings

Используйте команду `calculate savings`, чтобы узнать, какую сумму сможете накопить 
с учетом сложного процента, если на протяжении следующих N лет будете ежемесячно 
инвестировать X рублей под R% годовых с ежемесячной капитализацией процентов.

<b>Пример использования:</b>

```
$ ./bin/assist calculate savings --payment=20000.00 --years=25 --interest=6.5 --detailed=M
$ ./bin/assist calculate savings -p=20000.00 -y=25 -i=6.5 -d=M
$ ./bin/assist calculate savings --help
```

<b>Параметры и опции команды:</b>

```
-d, --detailed bool   Выводить детализированный ответ. M — по
                      месяцам, Y — по годам.
-h, --help bool       Документация по команде
-i, --interest float  Доходность вашего инвестиционного       
                      портфеля в процентах годовых
-p, --payment float   Размер ежемесячного пополнения          
                      инвестиционного портфеля
-y, --years uint8     Количество лет, на протяжении которых   
                      будут производиться накопления
```

<b>Пример выполнения команды:</b>

```
$ ./bin/assist calculate savings -e 150000 -y 25 -i 6.5 -d Y

 РАСЧЕТ БУДУЩИХ НАКОПЛЕНИЙ

 Задача: рассчитать сумму, которую можно накопить с учетом
 сложного процента, если на протяжении следующих 10 лет
 ежемесячно инвестировать 20000.00 рублей под 6.50% годовых с
 ежемесячной капитализацией процентов.

 ┏━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┓
 ┃   ГОД   ┃    ВЛОЖЕНИЯ    ┃    ПРОЦЕНТЫ    ┃   НАКОПЛЕНИЯ  ┃
 ┣━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃    1    ┃   240000.00    ┃    7280.68     ┃   247280.68   ┃
 ┃    2    ┃   480000.00    ┃    31122.21    ┃   511122.21   ┃
 ┃    3    ┃   720000.00    ┃    72633.70    ┃   792633.70   ┃
 ┃    4    ┃   960000.00    ┃   132998.53    ┃   1092998.53  ┃
 ┃    5    ┃   1200000.00   ┃   213479.35    ┃   1413479.35  ┃
 ┃    6    ┃   1440000.00   ┃   315423.37    ┃   1755423.37  ┃
 ┃    7    ┃   1680000.00   ┃   440268.00    ┃   2120268.00  ┃
 ┃    8    ┃   1920000.00   ┃   589546.96    ┃   2509546.96  ┃
 ┃    9    ┃   2160000.00   ┃   764896.65    ┃   2924896.65  ┃
 ┃   10    ┃   2400000.00   ┃   968063.08    ┃   3368063.08  ┃
 ┣━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
 ┃  ИТОГО  ┃   2400000.00   ┃   986306.76    ┃   3386306.76  ┃
 ┗━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┛

 > Накопленная сумма составит: 3386306.76
 > Сумма собственных вложений за период: 2400000.00
 > Сумма начисленных процентов за период: 986306.76
```

## Локальная сборка

`make linter && make assist && ./bin/assist -h`
