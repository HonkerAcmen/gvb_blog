package config

type ConfigUpload struct {
	PhotoSize int    `yaml:"photoSize" json:"photoSize"`
	PhotoPath string `yaml:"photoPath" json:"photoPath"`
	VideoSize int    `yaml:"videoSize" json:"videoSize"`
	VideoPath string `yaml:"videoPath" json:"videoPath"`
	FileSize  int    `yaml:"fileSize" json:"fileSize"`
	FilePath  string `yaml:"filePath" json:"filePath"`
}
