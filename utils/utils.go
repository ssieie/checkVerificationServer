package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func ParsePostData(data io.ReadCloser) (rData map[string]interface{}, err error) {

	res, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	var r interface{}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	rData, ok := r.(map[string]interface{})
	if !ok {
		return nil, errors.New("parse data err 断言失败")
	}

	return rData, nil
}

type Z map[string]interface{}

func JSON(res Z) (data []byte) {
	data, _ = json.Marshal(res)
	return
}

func decryptData(encryptedData []byte, key []byte) (map[string]interface{}, error) {
	// 将密钥转换成 AES 密钥类型
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 将加密数据和初始向量拆分出来
	iv := encryptedData[:aes.BlockSize]
	cipherData := encryptedData[aes.BlockSize:]

	// 使用 CBC 模式解密数据
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherData, cipherData)

	// 将解密后的字符串转换成 JSON 对象
	var data map[string]interface{}
	err = json.Unmarshal(cipherData, &data)
	if err != nil {
		return nil, err
	}

	// 验证请求参数是否完整
	if _, ok := data["data"]; !ok {
		return nil, errors.New("请求参数不完整")
	}

	// 返回解密后的数据
	return data, nil
}

type ConfigType struct {
	Host  string `json:"host"`
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
