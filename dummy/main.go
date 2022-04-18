package main


import (
	"fmt"
	razorpay "github.com/razorpay/razorpay-go"
)

func main() {
	client := razorpay.NewClient("rzp_test_k6uL897VPBz20q", "EnLs21M47BllR3X8PSFtjtbd")

	// data := map[string]interface{}{
	// 	"name": "sunil Kumar",
	// 	"contact": 9123456780,
	// 	"email": "sunil.kumar@example.com",
	// 	"fail_existing": 0,
	// 	"notes": map[string]interface{}{
	// 		"notes_key_1": "Tea, Earl Grey, Hot",
	// 		"notes_key_2": "Tea, Earl Grey… decaf.",
	// 	  },
	// }

	// customerId := "inv_JG1LJqepIOI47l"

// data := map[string]interface{}{
//   "name": "sunila Kumar",
//   "email": "sunila.kumar@example.com",
//   "contact": 9000000555,
// }

// data := map[string]interface{}{
// 	"type": "upi_qr",
// 	"name": "Store_1",
// 	"usage": "single_use",
// 	"fixed_amount": true,
// 	"payment_amount": 300,
// 	"description": "For Store 1",
// 	"customer_id": "cust_JFz35u2L3c6KJl",
// }

// data := map[string]interface{}{
// 	"notes": map[string]interface{}{
// 		"notes_key_1": "Tea, Earl Grey, Hot",
// 		"notes_key_2": "Tea, Earl Grey… decaf.",
// 		},
// }

// virtualId := "va_J3gpXyhnwtTQ3B"

// para_attr := map[string]interface{}{
//   "type": "bank_account",
//   "bank_account": map[string]interface{}{
//     "ifsc": "UTIB0000013",
//     "account_number": 914411112235679,
//   },
// }

// elements := make(map[string]interface{})
// elements["0"] = "vpa"

// para_attr := map[string]interface{}{
// 	"types": elements,
// 	"vpa": map[string]interface{}{
// 	  "descriptor": "gaubikumar",
// 	},
//   }

// para_attr := map[string]interface{}{
// 	"customer_id" : "cust_J3oSNi1KgzqVEX",
// }
// line_item := make(map[string]interface{})
// line_item["0"] = "bank_account"

// para_attr := map[string]interface{}{
//   "receivers": map[string]interface{}{
//     "types": line_item,
//   },
//   "description": "Virtual Account created for Raftar Soft",
//   "customer_id": "cust_IOyIY3JvbVny9o",
//   "close_by": 1650335857,
//   "notes": map[string]interface{}{
//     "project_name": "Banking Software",
//   },
// }

// types := make(map[string]interface{})
// types["0"] = "qr_code"

// data:= map[string]interface{}{
//   "receivers": map[string]interface{}{
//     "types": types,
//   },
//   "qr_code": map[string]interface{}{
// 	"type": "upi_qr",
// 	"name": "Store_1",
// 	"usage": "single_use",
// 	"fixed_amount": true,
// 	"payment_amount": 300,
// 	"description": "For Store 1",
// 	"customer_id": "cust_IOyIY3JvbVny9o",
// 	"close_by": 1650335857,
// 	"notes": map[string]interface{}{
// 	  "purpose": "Test UPI QR code notes",
// 	},
//   },
//   "description": "First QR code",
//   "amount_expected": 100,
//   "notes": map[string]interface{}{
//     "receiver_key": "receiver_value",
//   },
// }
data := map[string]interface{}{
	"notes": map[string]interface{}{
    "key_1": "value1",
    "key_2": "value2",
  },
}
body, err := client.Payment.Refund("pay_JKpA1zAwaOi0xH",1200, data, nil)
//body, err := client.VirtualAccount.AllowedPayer("va_JGrZBme0FqgsVh", para_attr, nil)
//body, err := client.VirtualAccount.AllowedPayer("va_JGrZBme0FqgsVh", para_attr, nil)
//body, err := client.VirtualAccount.DeleteAllowedPayer("va_JGrZBme0FqgsVh", "ba_JKkEXLugMXpfi5", nil, nil)  
// body, err := client.Invoice.Delete("inv_JKjgIXXBxkHSG9",nil, nil)
// para_attr := map[string]interface{}{
// 	"customer_id" : "cust_J3oSNi1KgzqVEX",
// }

// para_attr := map[string]interface{}{
// 	"amount": 100,
// 	"currency": "INR",
// 	"order_id": "order_EAkbvXiCJlwhHR",
// 	"email": "gaurav.kumar@example.com",
// 	"contact": 9090909090,
// 	"method": "upi",
//   }

//body, err := client.Payment.CreatePaymentJson(para_attr, nil)

// para_attr := map[string]interface{}{
// 	"email": "gaurav.kumar@example.com",
// 	"contact": "9876543567",
// 	"amount": 10000,
// 	"currency": "INR",
// 	"order_id": "order_JIPONOfVuBVB4B",
// 	"customer_id": "cust_DzYEzfJLV03rkp",
// 	"token": "token_JIOj6ss6Ug43Gg",
// 	"recurring": "1",
// 	"notes": map[string]interface{}{
// 	  "note_key_1": "Tea. Earl grey. Hot.",
// 	  "note_key_2": "Tea. Earl grey. Decaf.",
// 	},
// 	"description": "Creating recurring payment for Gaurav Kumar",
//   }
 
// data := map[string]interface{}{
// 	"notes": map[string]interface{}{
// 		"key1": "value1",
// 		"key2": "value2",
// 	},
// }

// data := map[string]interface{}{
// 	"amount": 50000,
// 	"currency": "INR",
// 	"receipt": "some_receipt_id",
// 	"payment": map[string]interface{}{
// 	  "capture ": "automatic",
// 	  "capture_options ": map[string]interface{}{
// 		"automatic_expiry_period ": 12,
// 		"manual_expiry_period ": 7200,
// 		"refund_speed": "optimum",
// 	  },  
// 	},
//   }
//body, err := client.Payment.FetchPaymentDowntime(nil, nil)
//body, err := client.Customer.Create(data, nil)
	//   data := map[string]interface{}{
	// 	"amount":          1234,
	// 	"currency":        "INR",
	// 	"receipt":         "some_receipt_id",
	// 	"payment_capture": 1,
	// }
	// body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err)
		return
	  }
    
	  fmt.Println(body)  
}
