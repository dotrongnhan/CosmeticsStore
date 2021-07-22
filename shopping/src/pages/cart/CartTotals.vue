<template>
  <div
    class="
      bo9
      w-size18
      p-l-40 p-r-40 p-t-30 p-b-38
      m-t-30 m-r-0 m-l-auto
      p-lr-15-sm
    "
  >
    <h5 class="m-text20 p-b-24">Cart Totals</h5>

    <div class="flex-w flex-sb bo10 p-t-15 p-b-20">
      <span class="s-text18 w-size19 w-full-sm"> Shipping: </span>

      <div class="w-size20 w-full-sm">
        <p class="s-text8 p-b-23">
          There are no shipping methods available. Please double check your
          address, or contact us if you need any help.
        </p>
      </div>
    </div>

    <div class="flex-w flex-sb-m p-t-26 p-b-30">
      <span class="m-text22 w-size19 w-full-sm"> Total: </span>

      <span class="m-text21 w-size20 w-full-sm">
        {{ currency(subTotal) }}
      </span>
    </div>

    <div class="size15 trans-0-4">
      <!-- Button -->
      <button class="flex-c-m sizefull bg1 bo-rad-23 hov1 s-text1 trans-0-4">
        Proceed to Checkout
      </button>
    </div>
  </div>
</template>

<script>
import { currency } from "@/utils/currency";


export default {
  name: "CartTotals",

  props: {
    subTotal: Number,
  },
  data() {
    return {
      loaded: false,
      paidFor: false,
      product: {
        price: 100,
        description: "gi gi do",
        img: "https://scontent-hkg4-1.xx.fbcdn.net/v/t1.6435-9/196437483_838057373517144_5524917560022717507_n.jpg?_nc_cat=107&ccb=1-3&_nc_sid=09cbfe&_nc_ohc=3ifrj73xS0kAX_11YwU&_nc_ht=scontent-hkg4-1.xx&oh=b4a8c153c5e31a3de79c0cc7eaae4e5b&oe=60FCD11D",
      }
    }
  },
  mounted() {
    const script = document.createElement("script")
    script.src =
        "https://www.paypal.com/sdk/js?client-id=AbTnXKHy0QyzQb-I6J9JxIeVYJGmcnc8HblkL5ycQjAuo6epSNptQtxUjF0bRhmfhboY0Z0--nubE1RH&currency=USD"
    script.addEventListener("load", this.setLoaded)
    document.body.appendChild(script)
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
                    description: this.product.description,
                    amount: {
                      currency_code: "USD",
                      value: this.product.price
                    }
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
  },
};
</script>