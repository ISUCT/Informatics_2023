# Работа с проектом

## Настройка окружения

### Golang

Установите `golang` для Вашей операционной системы в соответствии с инструкциями оффициальной документации. 

Запустите выкачивание зависимостей:

```shell
go mod tidy
```

Для того, чтобы проверить код на соответствие правилам форматирования, следует установить линтер. В случае с ОС `Windows` установка выполняется через терминал `Git Bash`:

```shell
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
```

Для запуска линтера следует выполнить следующие команды:
```shell
$(go env GOPATH)/bin/golangci-lint run -vvv --out-format tab --print-resources-usage --timeout=180s --config=.golangci.yml
```

### Запуск тестов

```shell
go test -count=1 -p 1 -v ./...
```