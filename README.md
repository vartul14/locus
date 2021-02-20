# locus

1. Install golang on the system using the link: https://golang.org/doc/install
2. run command 'go run main.go' after navigating inside the project in terminal or an supporting IDE
3. Postman collection
{
	"info": {
		"_postman_id": "893008b0-dc42-4fa1-8682-430c71de1022",
		"name": "Transactions",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "UpdateTransaction",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"request-id\": \"sdgg-gdffd-serrvds-sdcsds\",\n    \"transaction-id\": 2,\n    \"amount\": 1.3,\n    \"type\": \"payment\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8070/transactionservice/transaction/",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8070",
					"path": [
						"transactionservice",
						"transaction",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "GetTransaction",
			"request": {
				"method": "GET",
				"header": [],
				"url": null
			},
			"response": []
		},
		{
			"name": "GetTransactionByType",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8070/transactionservice/types/?type=payment&request-id=dbkgj-gbldvk",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8070",
					"path": [
						"transactionservice",
						"types",
						""
					],
					"query": [
						{
							"key": "type",
							"value": "payment"
						},
						{
							"key": "request-id",
							"value": "dbkgj-gbldvk"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "GetTransactionSum",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8070/transactionservice/sum/?request-id=dbkgj-gbldvk&transaction-id=20",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8070",
					"path": [
						"transactionservice",
						"sum",
						""
					],
					"query": [
						{
							"key": "request-id",
							"value": "dbkgj-gbldvk"
						},
						{
							"key": "transaction-id",
							"value": "20"
						}
					]
				}
			},
			"response": []
		}
	]
}
