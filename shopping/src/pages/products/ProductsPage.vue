<template>
  <div>
    <!-- Title Page -->
    <section
      class="bg-title-page p-t-50 p-b-40 flex-col-c-m"
      :style="{
        backgroundImage:
          'url(' + 'https://cdn.shopify.com/s/files/1/0373/4889/collections/millennial-hero-image_7bbfb2e3-0ffd-4ff0-84f0-5c5e29be850a_1600x.jpg?v=1592395128' + ')',
      }"

    >
      <h2 v-show="value.label !== 'Default category'" class="l-text2 t-center">{{value.label}}</h2>

    </section>

    <!-- Content page -->
    <section class="bgwhite p-t-55 p-b-65">
      <div class="container">
        <div class="row">
          <div class="col-sm-6 col-md-8 col-lg-12 p-b-50">
            <div class="flex-sb-m flex-w p-b-35">
              <div class="flex-w">
                <div class="bo4 w-size12 m-t-5 m-b-5 m-r-10">
                  <Select2
                    :options="[
                      { value: '', label: 'Default Sorting' },
                      { value: -1, label: 'Price: low to high' },
                      { value: 1, label: 'Price: high to low' },
                    ]"
                    @change="sort"
                    :value="direction"
                  />
                </div>

                <div class="bo4 w-size12 m-t-5 m-b-5 m-r-10">
                  <Select2
                    :options="[
                      { value: '', label: 'Default category' },
                      { value: 4, label: 'Toner' },
                      { value: 2, label: 'Mask' },
                      { value: 5, label: 'Lipstick' },
                      { value: 3, label: 'Face Scream' },
                      { value: 6, label: 'Serum' },
                    ]"
                    @change="category"
                    :value="currentCategory"
                  />
                </div>
              </div>
              <Pagination
                  :length="products.length"
                  :pageSize="limit"
                  :pageIndex="pageIndex"
                  @change="changePage"
              />
            </div>

            <transition name="fade">
              <div>
                <div class="row">
                  <div
                    class="col-sm-12 col-md-6 col-lg-3 p-b-50"
                    v-for="product in listProduct"
                    :key="product.id"
                  >
                    <div class="block2">
                      <div
                        class="block2-img wrap-pic-w of-hidden pos-relative"
                      >
                        <img :src="product.image" alt="IMG-PRODUCT" />

                        <div class="block2-overlay trans-0-4">
                          <a
                            href="#"
                            class="block2-btn-addwishlist hov-pointer trans-0-4"
                          >
                            <i
                              class="icon-wishlist icon_heart_alt"
                              aria-hidden="true"
                            ></i>
                            <i
                              class="icon-wishlist icon_heart dis-none"
                              aria-hidden="true"
                            ></i>
                          </a>

                          <div class="block2-btn-addcart w-size1 trans-0-4">
                            <button
                                @click="addToCart(product)"
                              class="
                                flex-c-m
                                size1
                                bg4
                                bo-rad-23
                                hov1
                                s-text1
                                trans-0-4
                              "
                            >
                              Add to Cart
                            </button>
                          </div>
                        </div>
                      </div>

                      <div class="block2-txt p-t-20">
                        <router-link
                          :to="'/products/' + product.id"
                          @click="this.$store.dispatch(products.products/getProductById(id))"
                          class="block2-name dis-block s-text3 p-b-5"
                        >
                          {{ product.product_name }}
                        </router-link>

                        <span class="block2-price m-text6 p-r-5">
                          {{ currency(product.price) }}
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </transition>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import {mapState} from "vuex";
import { currency } from "@/utils/currency";
import Select2 from "@/components/Select2.vue";
import Pagination from "@/components/Pagination.vue";

export default {
  name: "ProductsPage",
  components: {
    Select2,
    Pagination,
  },
  data() {
    return {
      value: {},
      direction: {},
      pageIndex: 1,
      limit: 8,
    }
  },

  created() {
    this.$store.dispatch("products/getProducts", {});
    console.log(this.$store.state.products.category)
  },
  methods: {
    currency,
    sort(direction) {
      this.$store.commit("products/SORT_PRODUCT", direction.value)
      this.direction = direction
    },
    category(value) {
      this.$store.dispatch("products/getProductByCategory", value.value)
      this.direction = ""
      this.value = this.currentCategory
    },
    addToCart(product) {
      this.$store.commit('cart/addProductToCart', {product: product, quantity: 1})
    },
    changePage(value) {
      this.pageIndex = value
    }
  },
  computed: {
    ...mapState("products", ["products"]),
    listProduct() {
      const index = this.limit * (this.pageIndex - 1)
      return this.products.slice(index, index + this.limit)
    },
    currentCategory() {
      return this.$store.state.products.category
    }
  },
};
</script>