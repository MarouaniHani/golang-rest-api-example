##REST API

####CreateAccount
* ( POST ): `/accounts`.

1. Request
* Url:  `/accounts`

* Body :
```
{
	"customer_id":"011671897080a38fa57f3c190a310d49",
	"username":"hany",
	"password":"test"
}
```

1. Responses

* status: 200

```
{
  "data": {
    "customer_id": "fd10113c54f4f4bfcaacbd1c2d188a77",
    "username": "hany",
    "password": "test"
  },
  "success": true
}
```
####ListOrdersByCustomerID
* ( GET ): `/customers/fd10113c54f4f4bfcaacbd1c2d188a77/orders`.

1. Request
* Url:  `/customers/:id/orders`

1. Responses

* status: 200

```
{
  "data": [
    {
      "id": "011f9dff2545a2cf8ac1809faed3ec88",
      "status": "delivered",
      "purchased_at": "2017-08-01T19:00:07Z",
      "pickedup_at": "2017-08-03T12:41:48Z",
      "delivered_at": "2017-08-09T21:04:42Z",
      "items": [
        {
          "price": 109.73,
          "freight": 16.53
        }
      ],
      "payment": 
        {
          "method": "credit_card",
          "amount": 126.26
        }
    }
  ],
  "success": true
}
```
####GetOrderByCustomerID
* ( GET ): `/orders/011f9dff2545a2cf8ac1809faed3ec88/customers/fd10113c54f4f4bfcaacbd1c2d188a77`.

1. Request
* Url:  `/orders/:order_id/customers/:customer_id`

1. Responses

* status: 200

```
{
  "data": {
    "id": "011f9dff2545a2cf8ac1809faed3ec88",
    "status": "delivered",
    "purchased_at": "2017-08-01T19:00:07Z",
    "pickedup_at": "2017-08-03T12:41:48Z",
    "delivered_at": "2017-08-09T21:04:42Z",
    "items": [
      {
        "price": 109.73,
        "freight": 16.53
      }
    ],
    "payment": 
      {
        "method": "credit_card",
        "amount": 126.26
      }
  },
  "success": true
}
```
####CreateOrder
* ( POST ): `/customers/fd10113c54f4f4bfcaacbd1c2d188a77/orders`.

1. Request
* Url:  `/customers/:customer_id/orders`

* Body :
```
{
	"product_ids":["016f3b29107ed03252e477b08445cec4","126abb319a746b23254ef06ddf10181d"]
}
```

1. Responses

* status: 200

```
{
  "data": {
    "id": "4957652351158172624",
    "status": "submitted",
    "purchased_at": "2017-08-01T19:00:07Z",
    "pickedup_at": "2017-08-01T19:00:07Z",
    "delivered_at": "2017-08-01T19:00:07Z",
    "items": [
      {
        "price": 0,
        "freight": 0
      },
      {
        "price": 0,
        "freight": 0
      }
    ],
    "payment": {
      "method": "",
      "amount": 0
    }
  },
  "success": true
}
```
####PayOrder
* ( POST ): `/customers/fd10113c54f4f4bfcaacbd1c2d188a77/orders/4957652351158172624/payments`.

1. Request
* Url:  `/customers/:customer_id/orders/:order_id/payments`

* Body :
```
{
	"method":"credit card",
	"amount":100.0
}
```

1. Responses

* status: 200

```
{
  "data": {
    "id": "4957652351158172624",
    "status": "submitted",
    "purchased_at": "2017-08-01T19:00:07Z",
    "pickedup_at": "2017-08-01T19:00:07Z",
    "delivered_at": "2017-08-01T19:00:07Z",
    "items": [
      {
        "price": 0,
        "freight": 0
      },
      {
        "price": 0,
        "freight": 0
      }
    ],
    "payment": {
      "method": "credit card",
      "amount": 100
    }
  },
  "success": true
}
```