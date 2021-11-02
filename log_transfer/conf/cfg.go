package conf

// LogTransferCfg 嵌套结构体
type LogTransferCfg struct {
	KafkaCfg  `ini:"kafka"`
	EsCfg	`ini:"es"`
}

// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

// EsCfg ...
type EsCfg struct {
	Address string `ini:"address"`
}

