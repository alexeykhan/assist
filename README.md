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

## Функционал 

### Decompose

<b>Подход: от желаемого результата.</b>

Используйте команду `decompose`, чтобы декомпозировать
финансовую цель, например, узнать минимальные необходимые
условия для достижения вашей цели к конкретному сроку.

<b>Пример использования:</b>

```
$ ./bin/assist decompose savings --goal=1234567.89 --years=10 --interest=6.5
$ ./bin/assist decompose savings -g=1234567.89 -y=10 -i=6.5
$ ./bin/assist decompose savings --help
```

<b>Параметры и опции команды:</b>

```
-g, --goal float32      Ваша финансовая цель, которую     
                        нужно достгнуть за заданный период
-h, --help bool         Документация по команде
-i, --interest float32  Доходность вашего инвестиционного
                        портфеля в процентах годовых
-y, --years uint8       Количество лет, за которое        
                        необходимо накопить нужную сумму
```

<b>Пример выполнения команды:</b>

![Assist Decompose Savings](./.github/decompose_savings.png?raw=true)

[comment]: <> (### Calculate)

[comment]: <> (<b>Подход: от текущей ситуации.</b>)

[comment]: <> (Используйте команду `calculate`, чтобы посмотреть, каких)

[comment]: <> (результатов можно достигнуть за указанный период, если)

[comment]: <> (соблюдать конкретные условия. Например:)

[comment]: <> (`./bin/assist calculate savings --payment 35000 --years 10 --interest=6.5`)

## Локальная сборка

`make linter && make assist && ./bin/assist -h`
