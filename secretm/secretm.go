package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/marcoswlrich/ecombrutus/awsgo"
	"github.com/marcoswlrich/ecombrutus/models"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var datasSecret models.SecretRDSJson
	fmt.Println(" > Pedido Secreto " + nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datasSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &datasSecret)
	fmt.Println(" > Leitura Secret Ok" + nameSecret)
	return datasSecret, nil
}
