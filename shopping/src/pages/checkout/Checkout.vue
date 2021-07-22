<template>
  <div class="container">
    <div class="py-5 text-center"></div>
    <div class="row">
      <div class="col-md-4 order-md-2 mb-4">
        <h4 class="d-flex justify-content-between align-items-center mb-3">
          <span class="text-muted">Your cart</span>
          <span class="badge badge-secondary badge-pill">{{totalItems}}</span>
        </h4>
        <ul class="list-group mb-3">
          <li v-for="product in products" :key="product.id" class="list-group-item d-flex justify-content-between lh-condensed">
            <div class="row">
              <div class="col-3">
                <img :src="product.product.image" style="width: 100%" alt="">
              </div>
              <div class="col-6">
                <p>{{`${product.product.product_name} x ${product.quantity}`}}</p>
              </div>
              <div class="col-3">
                <p>{{currency(product.product.price * product.quantity)}}</p>
              </div>
            </div>
          </li>
          <li class="list-group-item d-flex justify-content-between">
            <span>Total (USD)</span>
            <strong>{{currency(subTotal)}}</strong>
          </li>
        </ul>
      </div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Billing address</h4>
        <Form class="needs-validation" :validation-schema="schema">
          <div class="mb-3">
            <label for="username">Full name</label>
            <div class="input-group">
              <Field v-model="infor.fullName" type="text" class="form-control" name="fullname" id="username" placeholder="Full name" />
            </div>
            <ErrorMessage name="fullname" class="text-danger d-block" />
          </div>
            <div class="mb-1">
              <label for="phone">Phone</label>
              <div class="input-group">
                <Field v-model="infor.phone" type="text" class="form-control" name="phone" id="phone" placeholder="Phone" />
              </div>
              <ErrorMessage name="phone" class="text-danger d-block" />
            </div>
          <div class="mb-1">
            <label for="email">Email</label>
            <div class="input-group">
              <Field v-model="infor.email" type="email" class="form-control" name="email" id="email" placeholder="you@example.com"/>
            </div>
            <ErrorMessage name="email" class="text-danger d-block" />
          </div>

          <div class="mb-1">
            <label for="address">Address</label>
            <Field v-model="infor.address" type="text" class="form-control" name="address" id="address" placeholder="1234 Main St" />
            <ErrorMessage name="address" class="text-danger d-block" />
          </div>
          <h4 class="mb-1">Payment</h4>
          <div v-show="infor.fullName !=='' && infor.address !== '' && infor.email!=='' && infor.phone !=='' ">
            <div class="text-center s-text20 font-weight-bold mb-3">Please pay to complete checkout</div>
            <div class="d-flex justify-content-center" id="paypal-button-container"></div>
          </div>
          <div v-if="infor.fullName ==='' || infor.address === '' || infor.email==='' || infor.phone ==='' " class="text-center text-warning">Please fill in the recipient's information</div>
        </Form>
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters, mapState} from "vuex";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import {currency} from "../../utils/currency";

export default {
  name: "Checkout",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  mounted() {
    const script = document.createElement("script")
    script.src =
        "https://www.paypal.com/sdk/js?client-id=AbTnXKHy0QyzQb-I6J9JxIeVYJGmcnc8HblkL5ycQjAuo6epSNptQtxUjF0bRhmfhboY0Z0--nubE1RH&currency=USD"
    script.addEventListener("load", this.setLoaded)
    document.body.appendChild(script)
  },
  data() {
    const schema = yup.object().shape({
      email: yup
          .string()
          .required("Email is required!")
          .email("Email is invalid!"),
      fullname: yup
          .string()
          .required("Full name is required!"),
      phone: yup
          .string()
          .required("Phone is required"),
      address: yup
          .string()
          .required("Address is required")
    });

    return {
      schema,
      loaded: false,
      paidFor: false,
      infor: {
        fullName: "",
        phone: "",
        address: "",
        email: ""
      }
    };
  },
  computed: {
    ...mapState("cart", ["products"]),
    ...mapGetters("cart", ["totalItems","subTotal"])
  },
  methods: {
    currency,
    setLoaded: function() {
      this.loaded = true;
      window.paypal
          .Buttons({
            createOrder: (data, actions) => {
              return actions.order.create({
                purchase_units: [
                  {
                    description: "this is description",
                    amount: {
                      currency_code: "USD",
                      value: this.subTotal
                    },
                    style: {
                      size: 'large',
                      color: 'gold',
                      shape: 'pill',
                    },
                  }
                ]
              });
            },
            onApprove: async (data, actions) => {
              const order = await actions.order.capture();
              this.paidFor = true;
              console.log(order);
            },
            onError: err => {
              console.log(err);
            }
          })
          .render('#paypal-button-container');
    }
  }
}
</script>

<style scoped>
.bd-placeholder-img {
  font-size: 1.125rem;
  text-anchor: middle;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.form-control {
  border: 1px solid #ced4da !important;
}

@media (min-width: 768px) {
  .bd-placeholder-img-lg {
    font-size: 3.5rem;
  }
}

</style>