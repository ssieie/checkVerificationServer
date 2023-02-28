package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Z map[string]interface{}

func JSON(res Z) (data []byte) {
	data, _ = json.Marshal(res)
	return
}

var sysKey = []byte("wc666wc666wc6666")

func decrypt(cipherText string, key []byte) (map[string]interface{}, error) {
	// 解析加密后的数据
	cipherData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, errors.New("base64 decode error:" + err.Error())
	}
	// 提取随机向量和密文
	iv := cipherData[:aes.BlockSize]
	ciphertext := cipherData[aes.BlockSize:]

	// 创建解密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("cipher error:" + err.Error())
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充
	padding := int(plaintext[len(plaintext)-1])
	plaintext = plaintext[:len(plaintext)-padding]

	// 将解密后的数据反序列化为原始数据类型
	var data map[string]interface{}
	if err := json.Unmarshal(plaintext, &data); err != nil {
		return nil, errors.New("json decode error:" + err.Error())
	}
	return data, nil
}

type reqData struct {
	InputData string `json:"input_data"`
}

func Parse(r io.ReadCloser) (rData map[string]interface{}, err error) {

	res, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var data reqData

	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	if data.InputData == "" {
		return nil, errors.New("InputData is empty")
	}

	parseData, err := decrypt(data.InputData, sysKey)
	if err != nil {
		return nil, err
	}

	return parseData, nil
}

type ConfigType struct {
	Port  string `json:"port"`
	Mysql string `json:"mysql"`
	Redis string `json:"redis"`
}

var SysConfig ConfigType

func GetConfig() ConfigType {
	return SysConfig
}

func ParsConfig() error {
	conf, err := os.Open("./config.json")

	if err != nil {
		return errors.New(fmt.Sprintf("打开配置文件错误:%s", err))
	}
	data, err := io.ReadAll(conf)
	if err != nil {
		return errors.New(fmt.Sprintf("读取配置文件错误:%s", err))
	}
	err = json.Unmarshal(data, &SysConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("序列化配置文件错误:%s", err))
	}

	return nil
}

const TimeFormat = "2006-01-02 15:04:05"
