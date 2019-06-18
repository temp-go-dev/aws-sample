package s3mightiness

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func S3Access() *aws.Config {
	// /*aws.Configかんたんまとめ
	// s3Config := &aws.Config{
	// 	//認証情報オブジェクト
	// 	Credentials: credentials.NewStaticCredentials("hogehoge", "hogehoge", ""),
	// 	//エンドポイントURL
	// 	Endpoint: aws.String("http://127.0.0.1:9001"),
	// 	//リージョンはクライアントごとに設定
	// 	Region: aws.String("us-east-1"),
	// 	//SSL無効にする。デフォルトはfalse
	// 	DisableSSL: aws.Bool(true),
	// 	//trueに設定すると、リクエストはパススタイルになる
	// 	//デフォルトはS3クライアントは可能であれば仮想ホスト型バケットアドレッシングを使用
	// 	S3ForcePathStyle: aws.Bool(true),
	// 	//最大リトライ回数
	// 	MaxRetries: aws.Int(2),
	// 	//デフォルトのログレベルは0（ログがない）。ロギングを有効にするには、LogLevel値に設定する
	// 	DisableParamValidation: aws.Bool(false),
	// 	//要求と応答のチェックサムの計算を無効にする
	// 	DisableComputeChecksums: aws.Bool(false),
	// 	/* S3高速化機能を有効にするには、これを `true`に設定する。
	// 	S3と互換性のあるすべての操作に対して、Accelerateは要求に対して高速化エンドポイントを使用する。
	// 	互換性がない要求は通常のS3要求にフォールバックする。
	// 	高速化を有効にしてS3クライアントで高速化を使用するには、バケットを有効にする必要がある。
	// 	バケットが高速化に対して有効になっていない場合は、エラーが返される。
	// 	バケット名は、Accelerateでも機能するようにDNS互換である必要がある*/
	// 	S3UseAccelerate: aws.Bool(true),
	// 	//リクエストされたエンドポイントホストにモデル化された情報をプレフィックスするSDKの動作を無効にする
	// 	DisableEndpointHostPrefix: aws.Bool(false),
	// 	//そのモデルに定義がある操作でのエンドポイント検出を可能にする。デフォルトではfalse
	// 	EnableEndpointDiscovery: aws.Bool(false),
	// 	//残りのプロトコル要求を行うときにURLパスを消去する
	// 	DisableRestProtocolURICleaning: aws.Bool(false),
	// }*/

	return &aws.Config{
		Credentials:      credentials.NewStaticCredentials("hogehoge", "hogehoge", ""),
		Endpoint:         aws.String("http://127.0.0.1:9001"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
}
