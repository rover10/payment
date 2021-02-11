# payment
This service implements an api for getting transaction history of a typical upi app. It returns a list of 10 recent transactions made from or into the user's account. It also implements a basic pagination which let the caller to define mininum & maximum transactions to be returned in the response. The max limit is 100 and min is 1. The default value of the reponse limit is 10. 

There is also offset which can be set to get next transactions. The offset is the id of the last transaction in the list.
For e.g. ``v1/payment/history/1?limit=10&offset=15`` would get next 10 transaction below transaction id 15.




