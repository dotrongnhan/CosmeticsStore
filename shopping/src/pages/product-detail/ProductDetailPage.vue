<template>
  <div>
    <div class="bread-crumb bgwhite flex-w p-l-52 p-r-15 p-t-30 p-l-15-sm">
      <router-link to="/" class="s-text16">
        Home
        <i class="fa fa-angle-right m-l-8 m-r-9" aria-hidden="true"></i>
      </router-link>

      <router-link to="/products" class="s-text16">
        Products
        <i class="fa fa-angle-right m-l-8 m-r-9" aria-hidden="true"></i>
      </router-link>

      <span class="s-text17"> {{ product.name }} </span>
    </div>

    <div class="container bgwhite p-t-35 p-b-80">
      <div v-if="isLoading" data-loader="ball-scale"></div>

      <div v-else class="flex-w flex-sb">
        <div class="w-size13 p-t-30 respon5">
          <div class="wrap-slick3 flex-sb flex-w">
              <div
                data-thumb="@/assets/images/thumb-item-01.jpg"
              >
                <div class="wrap-pic-w">
                  <img
                    :src="product.image"
                    alt="IMG-PRODUCT"
                  />
                </div>
              </div>
          </div>
        </div>

        <div class="w-size14 p-t-30 respon5">
          <h4 class="product-detail-name m-text16 p-b-13">
            {{ product.product_name}}
          </h4>

          <span class="m-text17"> {{ currency(product.price) }} </span>

          <div class="p-t-33 p-b-60">

              <div class="w-size16 flex-m flex-w">
                <div class="flex-w bo5 of-hidden m-r-22 m-t-10 m-b-10">
                  <button
                    class="btn-num-product-down color1 flex-c-m size7 bg8 eff2"
                    @click="updateQuantityProduct(-1)"
                  >
                    <i class="fs-12 fa fa-minus" aria-hidden="true"></i>
                  </button>

                  <input
                    class="size8 m-text18 t-center num-product"
                    type="number"
                    name="num-product"
                    :value="quantity"
                    @keyup="setQuantityProduct($event.target.value)"
                  />

                  <button
                    class="btn-num-product-up color1 flex-c-m size7 bg8 eff2"
                    @click="updateQuantityProduct(1)"
                  >
                    <i class="fs-12 fa fa-plus" aria-hidden="true"></i>
                  </button>
                </div>

                <div
                  class="
                    btn-addcart-product-detail
                    size9
                    trans-0-4
                    m-t-10 m-b-10
                  "
                >
                  <button
                    class="
                      flex-c-m
                      sizefull
                      bg1
                      bo-rad-23
                      hov1
                      s-text1
                      trans-0-4
                    "
                    @click="submitNewProduct(product)">
                    Add to Cart
                  </button>
                </div>
              </div>
          </div>

          <div class="wrap-dropdown-content bo6 p-t-15 p-b-14 active-dropdown-content">
            <h5
                class="
        js-toggle-dropdown-content
        flex-sb-m
        cs-pointer
        m-text19
        color0-hov
        trans-0-4
      "
            >
              Description
              <i
                  class="down-mark fs-12 color1 fa fa-minus dis-none"
                  aria-hidden="true"
              ></i>
              <i class="up-mark fs-12 color1 fa fa-plus" aria-hidden="true"></i>
            </h5>

            <div class="p-t-15 p-b-23">
              <p class="s-text8">
                {{product.description}}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <section class="relateproduct bgwhite p-t-45 p-b-138">
      <div class="container">
        <div class="sec-title p-b-60">
          <h3 class="m-text5 t-center">Related Products</h3>
        </div>

        <ProductsCarousel :products="products" />
      </div>
    </section>
  </div>
</template>

<script>
import {mapMutations, mapState} from "vuex";
import { currency } from "@/utils/currency";
import ProductsCarousel from "@/components/ProductsCarousel.vue";
import router from "../../router";

export default {
  name: "ProductDetailPage",
  components: {
    ProductsCarousel,
  },

  data() {
    return {
      quantity: 1
    };
  },
  computed: mapState("products", ["products", "product"]),
  created() {
    this.$store.dispatch("products/getProductById", this.$route.params.id);
  },
  methods: {
    currency,
    ...mapMutations("cart", ["addProductToCart"]),
    submitNewProduct(product) {
      this.addProductToCart({
        product: product,
        quantity: this.quantity
      })
      router.push('/products')
    },
    updateQuantityProduct(number) {
      if(this.quantity === 0 && number === -1) {
        this.quantity = 0
      }else if (this.quantity <= -number || this.quantity < 0) {
        this.quantity = 0
      } else {
        this.quantity = this.quantity + number
      }
    },
    setQuantityProduct(number) {
      number = parseInt(number)
      if(number < 0) {
        this.quantity = 0
      } else {
        this.quantity = number
      }
    }
  },
};
</script>