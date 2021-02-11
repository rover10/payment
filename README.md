# payment
This is a Go service which implements an api for getting transaction history of a typical upi app. It returns a list of 10 recent transactions made from or into the user's account. It also implements a basic pagination which let the caller to define mininum & maximum transactions to be returned in the response. This also prevent stress on database. Default value of the  limit is 10. The max limit is 100 and min is 1.

There is also an ``offset`` query param which can be set to get next transactions. The offset is the id of the last transaction in the list.
For e.g. ``v1/payment/history/1?limit=2&offset=19`` would get next 2 transaction below transaction id 19.

The api takes userId path parameter ``v1/payment/history/{userId}?limit=2&offset=19``

```
[
  {
    "id": 18,
    "utr": "UTR166884422267896YYZZ",
    "amount": 3500,
    "paymentTime": "2021-02-11T19:50:47.275203Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "1230*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "Housing Development Finance Corporation Limited Bank",
      "account": "1220*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/hdfc.png"
    }
  },
  {
    "id": 17,
    "utr": "UTR1668844222678904DZZ",
    "amount": 100,
    "paymentTime": "2021-02-11T19:50:47.275203Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "1234*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "Housing Development Finance Corporation Limited Bank",
      "account": "1220*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/hdfc.png"
    }
  }
]
```

To get next 3 transactions ``v1/payment/history/1?limit=5&offset=17``

```
[
  {
    "id": 15,
    "utr": "UTR1668844222678904DSE",
    "amount": 100,
    "paymentTime": "2021-02-11T19:50:03.817003Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "1230*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "State Bank of India",
      "account": "1234*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    }
  },
  {
    "id": 14,
    "utr": "UTR1888844446789045DSE",
    "amount": 500,
    "paymentTime": "2021-02-11T19:49:54.236802Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "1234*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "State Bank of India",
      "account": "1230*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    }
  },
  {
    "id": 13,
    "utr": "UTR1200123456789045DSE",
    "amount": 1000,
    "paymentTime": "2021-02-11T19:49:50.662226Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "1234*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "State Bank of India",
      "account": "1230*****",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    }
  }
]
```

## Schema
The schema SQL script can be found at https://github.com/rover10/payment/blob/master/lib/database/schema.sql. 


## ER diagram
The ER diagram for this is   https://github.com/rover10/payment/blob/master/ER.png
 
