[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# Go биндинги для Translations Plugify плагина

Этот репозиторий содержит Go биндинги для плагина [plugify-plugin-translations](https://github.com/untrustedmodders/plugify-plugin-translations). Биндинги генерируются автоматически и синхронизируются с оригинальным плагином.

## Документация по API

Полную документацию по API можно найти здесь на сайте [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-translations/refs/heads/main/plugify-plugin-translations.pplugin.in)

## Установка

```bash
go get github.com/fr0nch/go-plugify-translations
```

## Обновление

Этот репозиторий использует GitHub Actions для автоматической проверки обновлений в [plugify-plugin-translations](https://github.com/untrustedmodders/plugify-plugin-translations) в каждые 2:00 по UTC. При обнаружении новой версии, биндинги обновляются, и создается новый релиз.

## Лицензия

Этот модуль Go для Plugify лицензирован на условиях [MIT License](LICENSE).