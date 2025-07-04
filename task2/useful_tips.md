## Лучшие практики и частые ошибки

1. Всегда используйте экспортируемые поля. Поля структур, начинающиеся с маленькой буквы, не будут видны пакету encoding/json и будут проигнорированы. Это самая частая ошибка новичков.

2. Всегда проверяйте ошибку от json.Marshal, json.Unmarshal, encoder.Encode, decoder.Decode. Ошибки могут возникнуть из-за невалидного JSON, несовместимости типов и т.д.

3. Используйте struct для известной структуры. Это даёт вам безопасность типов, автодополнение в IDE и более читаемый код. Прибегайте к map[string]interface{} только тогда, когда структура данных действительно динамическая.

4. Используйте Encoder/Decoder для I/O операций. Это экономит память и часто работает быстрее, так как избегает лишних аллокаций.

5. Помните про float64. При анмаршалинге в interface{}, все JSON-числа (целые и с плавающей точкой) становятся float64. Будьте готовы к этому при приведении типов.

6. Используйте json.Valid для быстрой проверки. Если вам нужно только проверить, является ли []byte валидным JSON, без его разбора, используйте json.Valid(data). Это быстрее, чем json.Unmarshal.

7. Сторонние библиотеки. Для большинства задач стандартной библиотеки более чем достаточно. Однако в сценариях, где производительность критична (high-load сервисы, парсинг гигабайт логов), обратите внимание на сторонние, более быстрые реализации, например, json-iterator/go. Но начинать всегда стоит со стандартной библиотеки.
