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
  methods: {
    currency,
  },
};
</script>