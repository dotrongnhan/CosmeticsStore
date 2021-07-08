<template>
  <div>
    <section class="cart bgwhite p-t-70 p-b-100">
      <div class="container">
        <!-- Cart item -->
        <div class="container-table-cart pos-relative">
          <div class="wrap-table-shopping-cart bgwhite">
            <table class="table-shopping-cart">
              <tr class="table-head">
                <th class="column-1"></th>
                <th class="column-2">Product</th>
                <th class="column-3">Price</th>
                <th class="column-4 p-l-70">Quantity</th>
                <th class="column-5">Total</th>
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
                  <div class="flex-w bo5 of-hidden w-size17">
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
                        updateProductQuantity({
                          productId: product.id,
                          value: product.quantity - 1,
                        })
                      "
                    >
                      <i class="fs-12 fa fa-minus" aria-hidden="true"></i>
                    </button>

                    <input
                      class="size8 m-text18 t-center num-product"
                      type="number"
                      name="num-product1"
                      :value="product.quantity"
                      @input="
                        updateProductQuantity({
                          productId: product.id,
                          value: $event.target.value,
                        })
                      "
                    />

                    <button
                      class="btn-num-product-up color1 flex-c-m size7 bg8 eff2"
                      @click="
                        updateProductQuantity({
                          productId: product.id,
                          value: product.quantity + 1,
                        })
                      "
                    >
                      <i class="fs-12 fa fa-plus" aria-hidden="true"></i>
                    </button>
                  </div>
                </td>
                <td class="column-5">{{ currency(product.quantity * Number(product.product.price)) }}</td>
              </tr>

              <tr class="table-row">
                <td class="column-1" colspan="3"/>
                <td class="text-right p-r-10"><b>Subtotal:</b></td>
                <td class="column-5">{{ currency(subTotal) }}</td>
              </tr>
            </table>
          </div>
        </div>

        <!-- Total -->
        <CartTotals :subTotal="subTotal" />
      </div>
    </section>
  </div>
</template>

<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import { currency } from "@/utils/currency";
import CartTotals from "./CartTotals.vue";

export default {
  name: "CartPage",

  components: {
    CartTotals,
  },

  computed: {
    ...mapState("cart", ["products", "isLoading"]),
    ...mapGetters("cart", ["subTotal"]),
  },

  methods: {
    currency,
    ...mapMutations("cart", ["updateProductQuantity"]),
  },
};
</script>