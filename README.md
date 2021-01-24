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

<b>ПЕРСОНАЛЬНЫЙ АССИСТЕНТ ДЛЯ ПЛАНИРОВАНИЯ ЛИЧНЫХ ФИНАНСОВ</b>

Используйте команду `decompose`, чтобы декомпозировать
финансовую цель, например, узнать минимальные необходимые
условия для достижения вашей цели к конкретному сроку.
Подход: от желаемого результата.

![Assist Decompose Savings](./.github/decompose_savings.png?raw=true)

Используйте команду `calculate`, чтобы посмотреть, каких
результатов можно достигнуть за указанный период, если
соблюдать конкретные условия. Подход: от текущей ситуации.

## Примеры использования

```
$ ./bin/assist decompose --help
$ ./bin/assist calculate --help
$ ./bin/assist --help
```

## Локальная сборка

`make linter && make assist && ./bin/assist -h`
