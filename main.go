package main

import (
	"encoding/json"
	"fmt"
)

/*
*
main: Cli tool, Make sure working as binary, compatible in MacOS and WinOS
1. read the export config
2. eliminate the part that causing error
3. print it back
*/
func main() {

	str := "{\n  \"rabbit_version\": \"3.11.10\",\n  \"rabbitmq_version\": \"3.11.10\",\n  \"product_name\": \"RabbitMQ\",\n  \"product_version\": \"3.11.10\",\n  \"users\": [\n    {\n      \"name\": \"admin\",\n      \"password_hash\": \"BLzW0EASGs4hWzUdBkVxu5X0K+24Ef/ro9UqSpHFwBl6SzmQA\",\n      \"hashing_algorithm\": \"rabbit_password_hashing_sha256\",\n      \"tags\": [\n        \"administrator\"\n      ],\n      \"limits\": {}\n    },\n    {\n      \"name\": \"user\",\n      \"password_hash\": \"qadi61ibCLcq169rbZjd8t01JUlWTFq0IIju7ybKd7N4xtmR\",\n      \"hashing_algorithm\": \"rabbit_password_hashing_sha256\",\n      \"tags\": [\n        \"administrator\"\n      ],\n      \"limits\": {}\n    },\n    {\n      \"name\": \"guest\",\n      \"password_hash\": \"FDnEIy9fQ/PLiTLarDw66yxAxa8E9UiVCAHgFRddZdR/z80S\",\n      \"hashing_algorithm\": \"rabbit_password_hashing_sha256\",\n      \"tags\": [\n        \"administrator\"\n      ],\n      \"limits\": {}\n    }\n  ],\n  \"vhosts\": [\n    {\n      \"name\": \"/\"\n    }\n  ],\n  \"permissions\": [\n    {\n      \"user\": \"admin\",\n      \"vhost\": \"/\",\n      \"configure\": \".*\",\n      \"write\": \".*\",\n      \"read\": \".*\"\n    },\n    {\n      \"user\": \"guest\",\n      \"vhost\": \"/\",\n      \"configure\": \".*\",\n      \"write\": \".*\",\n      \"read\": \".*\"\n    },\n    {\n      \"user\": \"user\",\n      \"vhost\": \"/\",\n      \"configure\": \".*\",\n      \"write\": \".*\",\n      \"read\": \".*\"\n    }\n  ],\n  \"topic_permissions\": [],\n  \"parameters\": [],\n  \"global_parameters\": [\n    {\n      \"name\": \"internal_cluster_id\",\n      \"value\": \"rabbitmq-cluster-id-whJGT2qC0hzkoUyl2_5ZAA\"\n    }\n  ],\n  \"policies\": [],\n  \"queues\": [\n    {\n      \"name\": \"FILE_UPLOAD\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"digital_id_notification_ttl_queue\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {\n        \"x-dead-letter-exchange\": \"digital_id_notification_dead_letter_exchange\",\n        \"x-dead-letter-routing-key\": \"digital_id_notification_dead_letter_routing_key\"\n      }\n    },\n    {\n      \"name\": \"INVOICE_SUBMIT\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"POST_TOKEN\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"transfer_payment_ttl_queue\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {\n        \"x-dead-letter-exchange\": \"transfer_dead_letter_exchage\",\n        \"x-dead-letter-routing-key\": \"transfer_dead_letter_routing_key\"\n      }\n    },\n    {\n      \"name\": \"httpsBearTokenAuth\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"kenWi_json_request\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"Crypto_POST\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"OffQueue\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"KKK_POST_URL_ENCODE_FORM\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    },\n    {\n      \"name\": \"AGG_POST\",\n      \"vhost\": \"/\",\n      \"durable\": true,\n      \"auto_delete\": false,\n      \"arguments\": {}\n    }\n  ]\n}"

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &dataMap)
	if err != nil {
		panic(err)
	}

	delete(dataMap, "users")

	prettyStr, err := json.MarshalIndent(dataMap, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyStr))

}
