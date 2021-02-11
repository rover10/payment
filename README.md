# payment
This service implements an api for getting transaction history of a typical upi app. It returns a list of 10 recent transactions made from or into the user's account. It also implements a basic pagination which let the caller to define mininum & maximum transactions to be returned in the response. The max limit is 100 and min is 1. The default value of the reponse limit is 10. 

There is also offset which can be set to get next transactions. The offset is the id of the last transaction in the list.
For e.g. ``v1/payment/history/1?limit=2&offset=19`` would get next 2 transaction below transaction id 19.

```
[
  {
    "id": 18,
    "utr": "UTR166884422267896YYZZ",
    "amount": 3500,
    "fromAccountId": 2,
    "toAccountId": 3,
    "paymentTime": "2021-02-11T19:50:47.275203Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "123000009",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "Housing Development Finance Corporation Limited Bank",
      "account": "123000009",
      "bankIcon": "http://icon.mobile.matchmove/mobile/hdfc.png"
    }
  },
  {
    "id": 17,
    "utr": "UTR1668844222678904DZZ",
    "amount": 100,
    "fromAccountId": 1,
    "toAccountId": 3,
    "paymentTime": "2021-02-11T19:50:47.275203Z",
    "status": "success",
    "payment": "sent",
    "fromBank": {
      "name": "State Bank of India",
      "account": "123456789",
      "bankIcon": "http://icon.mobile.matchmove/mobile/sbi.png"
    },
    "toBank": {
      "name": "Housing Development Finance Corporation Limited Bank",
      "account": "123456789",
      "bankIcon": "http://icon.mobile.matchmove/mobile/hdfc.png"
    }
  }
]
```



