package biz

import (
	"data4perf/models"
	"encoding/json"
	"fmt"
)

func GetProductName(id string) (name string, err error) {
	var names []string
	models.Orm.Table("product").Where("id = ?", id).Pluck("product", &names)

	if len(names) == 0 {
		err = fmt.Errorf("未找到[%v]产品信息，请核对", id)
		Logger.Error("%s", err)
		return
	}

	name = names[0]

	return
}

func GetEnvTypeByName(product string) (envType int, err error) {
	var envTypes []int
	models.Orm.Table("product").Where("product = ?", product).Pluck("env_type", &envTypes)

	if len(envTypes) == 0 {
		err = fmt.Errorf("未找到[%v]产品信息，请核对", product)
		Logger.Warning("%s", err)
		return
	}

	envType = envTypes[0]

	return
}

func GetProductEnv(name string) (envModel ProductEnvModel, err error) {
	var product Product
	models.Orm.Table("product").Where("product = ?", name).Find(&product)

	if len(product.Name) == 0 {
		err = fmt.Errorf("未找到[%v]产品信息，请核对", name)
		Logger.Error("%s", err)
		return
	}
	envModel.Protocol = product.Protocol
	envModel.Ip = product.Ip
	envModel.Name = product.Name
	auth := make(map[string]string)

	if len(product.Auth) == 0 {
		err = fmt.Errorf("未配置鉴权信息，请先配置")
		Logger.Error("%s", err)
		return
	}
	json.Unmarshal([]byte(product.Auth), &auth)
	for k, v := range auth {
		var varDataModel VarDataModel
		varDataModel.Name = k
		varDataModel.TestValue = append(varDataModel.TestValue, v)
		envModel.Auth = append(envModel.Auth, varDataModel)
	}

	return
}

func GetProductApps(id string) (name string, err error) {
	var names []string
	models.Orm.Table("product").Where("id = ?", id).Pluck("apps", &names)

	if len(names) == 0 {
		err = fmt.Errorf("未找到[%v]产品信息，请核对", id)
		Logger.Error("%s", err)
		return
	}

	name = names[0]

	return
}

func (dbProduct DbProduct) GetPrivateParameter() (privateParameter map[string]interface{}) {
	privateParameter = make(map[string]interface{})
	if len(dbProduct.PrivateParameter) > 2 {
		err := json.Unmarshal([]byte(dbProduct.PrivateParameter), &privateParameter)
		if err != nil {
			Logger.Error("%s", err)
		}
	}
	return
}
