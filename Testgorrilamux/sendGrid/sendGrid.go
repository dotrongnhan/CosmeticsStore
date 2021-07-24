package sendGrid

import (
	"Testgorillamux/models"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"strconv"
)

func SendEmailBySendGrid(data models.Data) {

	Subtotal := strconv.FormatFloat(float64(data.Total), 'E', -1, 32)

	str := ""
	for i := 0; i < len(data.Orders); i++ {
		a := fmt.Sprintf("%f", data.Orders[i].Product.Price)
		quantity := fmt.Sprint(data.Orders[i].Quantity)
		str  += `<tr>
		<td>` + strconv.Itoa(i + 1) + `</td>
		<td>` + data.Orders[i].Product.ProductName + `</td>
		<td>` + a + `</td>
		<td>` + quantity + `</td>
		</tr>`

	}

	from := mail.NewEmail("CosmeticsStore", "nhandt@reactplus.com")
	subject := "Thanks"
	to := mail.NewEmail(data.AddressShipping.FullName, data.AddressShipping.Email)
	plainTextContent := "Inform buyer"
	htmlContent := `<h2>Inform receiver</h2>
	<h4>` + data.AddressShipping.FullName + `</h4>
	<h4>` + data.AddressShipping.Phone + `</h4>
	<h4>` + data.AddressShipping.Email + `</h4>
	<h4>` + data.AddressShipping.Address + `</h4>
	<h2>Inform bill</h2>
	<table>
		<thead>
			<tr>
				<th>STT</th>
				<th>Product Name</th>
				<th>Price</th>
				<th>Quantity</th>
			</tr>
		</thead>
		<tbody>
` + str + `
		</tbody>
	</table>
	<h2>Sumary</h2>
	<h4>Subtotal: $` + Subtotal + ` </h4>`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}