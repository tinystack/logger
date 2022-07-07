# logger
Go logger Package

## 示例

```go
import "github.com/tinystack/logger"

// 创建新实例
newLogger := NewLogger(
    logger.WithWriter(os.Stdout),
    logger.WithEncoder(EncoderJSON),
    logger.WithCaller(true),
    logger.WithLevel(DebugLevel),
)

// 打印日志
newLogger.Info("print logger message")
newLogger.Infof("print logger message: %s", "string message")
newLogger.Infot("print logger message", logger.T{"key1": "value1", "key2": T{"key3": "value3"}})

// 更新全局默认实例
logger.UpdateDefaultLogger(newLogger)
```