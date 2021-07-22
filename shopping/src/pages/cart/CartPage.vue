<template>
  <div>
    <section class="cart bgwhite p-t-30 p-b-100">
      <div class="container">
        <!-- Cart item -->
            <div class="container-table-cart pos-relative">
              <div class="wrap-table-shopping-cart bgwhite">
                <table class="table-shopping-cart">
                  <tr class="table-head">
                    <th class="column-1"></th>
                    <th class="column-2">Product</th>
                    <th class="column-3">Price</th>
                    <th class="column-4 t-center">Quantity</th>
                    <th class="column-4 t-center">Total</th>
                    <th  class="column-1"></th>
                  </tr>

                  <tr
                      class="table-row"
                      v-for="product in products"
                      :key="product.id"
                  >
                    <td class="column-1">
                      <div class="cart-img-product b-rad-4 o-f-hidden">
                        <img :src="product.product.image" alt="IMG-PRODUCT" />
                      </div>
                    </td>
                    <td class="column-2">{{ product.product.product_name }}</td>
                    <td class="column-3">{{ currency(product.product.price) }}</td>
                    <td class="column-4">
                      <div class="flex-w bo5 of-hidden m-l-r-auto w-size17">
                        <button
                            class="
                        btn-num-product-down
                        color1
                        flex-c-m
                        size7
                        bg8
                        eff2
                      "
                            @click="
                        changeCart(product, -1)
                      "
                        >
                          <i class="fs-12 fa fa-minus" aria-hidden="true"></i>
                        </button>

                        <input
                            class="size8 m-text18 t-center num-product"
                            type="number"
                            name="num-product1"
                            :value="product.quantity"
                            disabled
                        />

                        <button
                            class="btn-num-product-up color1 flex-c-m size7 bg8 eff2"
                            @click="changeCart(product, 1)"
                        >
                          <i class="fs-12 fa fa-plus" aria-hidden="true"></i>
                        </button>
                      </div>
                    </td>
                    <td class="column-5 t-center">{{ currency(product.quantity * Number(product.product.price)) }}</td>
                    <td class="column-1">
                      <button @click="deleteOrderItem(product.id)" class="flex-c-m w-50 bg1 bo-rad-8 hov1 s-text1 trans-0-4">
                        X
                      </button>
                    </td>
                  </tr>

                  <tr class="table-row">
                    <td class="column-1" colspan="4">
                      <div class="size15 trans-0-4">
                      <button @click="this.$router.push('/payment')" class="flex-c-m sizefull bg1 bo-rad-23 hov1 s-text1 trans-0-4 w-50">
                        Proceed to Checkout
                      </button>
                    </div></td>
                    <td class="text-right p-r-10"><b>Subtotal:</b></td>
                    <td class="column-5">{{ currency(subTotal) }}</td>
                  </tr>
                </table>
              </div>
            </div>
      </div>
    </section>
  </div>
</template>

<script>
import {mapState, mapMutations, mapGetters, mapActions} from "vuex";
import { currency } from "@/utils/currency";

export default {
  name: "CartPage",

  computed: {
    ...mapState("cart", ["products", "isLoading"]),
    ...mapGetters("cart", ["subTotal"]),
    ...mapState("users", ["user"])
  },
  methods: {
    currency,
    go(a) {
      console.log(a)
    },
    ...mapMutations("cart", ["updateProductQuantity"]),
    ...mapActions("cart", ["updateCart", "deleteOrderItem"]),
    changeCart(data, quantity) {
      this.updateCart({userId: this.user.id, product: data.product, quantity: quantity})

    }
  },
};
</script>